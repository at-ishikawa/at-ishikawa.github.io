GO := @go
DOCKER_COMPOSE := docker-compose
MYSQL_TEST_ROOT_PASSWORD := password

setup:
	$(DOCKER_COMPOSE) up -d
	sleep 10
	$(DOCKER_COMPOSE) exec mysql bash -c "mysql -uroot -p$(MYSQL_TEST_ROOT_PASSWORD) test < /workspace/internal/user/table.sql"
	$(GO) run ./cmd/seed/user

clean:
	$(DOCKER_COMPOSE) rm -s -f

run/helloworld/%:
	$(GO) run ./cmd/helloworld/$*

proto:
	protoc --go_out=plugins=grpc:. protos/*.proto

benchmark:
	go test -v -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out ./internal/user
