package user

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/protos"
)

type Server struct {
	pb.UserServiceServer
	db *sqlx.DB
}

type Record struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const limit = 1000

func (s *Server) countUserRecord(ctx context.Context, keyword string) (int, error) {
	rows, err := s.db.NamedQueryContext(ctx, "SELECT count(*) AS total FROM users WHERE name LIKE :keyword", map[string]interface{}{
		"keyword": "%" + keyword + "%",
	})
	if err != nil {
		return 0, fmt.Errorf("failed to NamedQuery: %w", err)
	}
	rows.Next()
	var count int
	if err := rows.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to Scan: %w", err)
	}
	return count, nil
}

func (s *Server) searchUserRecord(ctx context.Context, keyword string, offset int) ([]Record, error) {
	rows, err := s.db.NamedQueryContext(ctx, "SELECT * FROM users WHERE name LIKE :keyword ORDER BY id DESC LIMIT :limit OFFSET :offset", map[string]interface{}{
		"keyword": "%" + keyword + "%",
		"offset":  offset,
		"limit":   limit,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to NamedQuery: %w", err)
	}
	var records []Record
	for rows.Next() {
		var record Record
		if err := rows.StructScan(&record); err != nil {
			return nil, fmt.Errorf("failed to StructScan: %w", err)
		}
		records = append(records, record)
	}
	return records, nil
}

func (s *Server) convertDBRecordsToProto(records []Record) []*pb.User {
	protos := make([]*pb.User, len(records))
	for i := 0; i < len(records); i++ {
		protos[i] = &pb.User{
			Id:   int64(records[i].ID),
			Name: records[i].Name,
		}
	}
	return protos
}

func (s *Server) PaginateUsers(ctx context.Context, req *pb.PaginateUsersRequest) (*pb.PaginateUsersResponse, error) {
	startTime := time.Now()

	var totalCount int
	var offset int
	var err error
	if req.PageToken != "" {
		offset, err = strconv.Atoi(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid PageToken: %s", req.PageToken)
		}
	} else {
		totalCount, err = s.countUserRecord(ctx, req.Keyword)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to countUserRecord: %+v", err)
		}
	}

	records, err := s.searchUserRecord(ctx, req.Keyword, offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to searchUserRecord: %+v", err)
	}
	duration := time.Now().Sub(startTime)
	metadata.AppendToOutgoingContext(ctx, "duration", duration.String())
	var nextToken string
	if len(records) == limit {
		nextToken = strconv.Itoa(offset + limit)
	}
	return &pb.PaginateUsersResponse{
		Users:      s.convertDBRecordsToProto(records),
		NextToken:  nextToken,
		TotalCount: uint32(totalCount),
	}, nil
}

func (s *Server) StreamUsers(req *pb.StreamUsersRequest, stream pb.UserService_StreamUsersServer) error {
	startTime := time.Now()
	if err := stream.SetHeader(metadata.New(map[string]string{
		"start": startTime.Format(time.RFC3339),
	})); err != nil {
		return status.Errorf(codes.Internal, "failed to SetHeader: %+v", err)
	}

	ctx := stream.Context()
	totalCount, err := s.countUserRecord(ctx, req.Keyword)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to countUserRecord: %+v", err)
	}

	var offset int
	for {
		if offset >= totalCount {
			break
		}

		eg, egctx := errgroup.WithContext(ctx)
		for i := 0; i < int(req.Concurrency); i++ {
			nextOffset := offset + i*limit
			if nextOffset >= totalCount {
				break
			}

			eg.Go(func() error {
				records, err := s.searchUserRecord(egctx, req.Keyword, nextOffset)
				if err != nil {
					return status.Errorf(codes.Internal, "failed to searchUserRecord: %+v", err)
				}

				if err := stream.Send(&pb.StreamUsersResponse{
					Users:      s.convertDBRecordsToProto(records),
					TotalCount: uint32(totalCount),
				}); err != nil {
					return fmt.Errorf("failed to Send: %+v", err)
				}
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return status.Errorf(codes.Internal, "failed to Wait: %+v", err)
		}

		offset = offset + limit*int(req.Concurrency)
	}

	duration := time.Now().Sub(startTime)
	stream.SetTrailer(metadata.New(map[string]string{
		"duration": duration.String(),
	}))
	return nil
}

func RunServer(address string, mysqlDataSource string, logger *log.Logger) (*grpc.Server, error) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: (%w)", err)
	}

	db, err := sqlx.Connect("mysql", mysqlDataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql: (%w)", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &Server{
		db: db,
	})
	reflection.Register(grpcServer)
	logger.Printf("run a server on port: %s", address)
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			logger.Printf("stopped serving: %v", err)
		}
	}()
	return grpcServer, nil
}
