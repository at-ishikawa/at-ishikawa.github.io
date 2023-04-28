# https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#use-multi-stage-builds
# syntax=docker/dockerfile:1
FROM golang:1.20-alpine AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
WORKDIR /go/src/opentelemetry-go-client-fab/
# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/opentelemetry-go-client-fab/
RUN go build -o /bin/opentelemetry-go-client-fab

# This results in a single layer image
FROM scratch
COPY --from=build /bin/opentelemetry-go-client-fab /bin/opentelemetry-go-client-fab
ENTRYPOINT ["/bin/opentelemetry-go-client-fab"]
