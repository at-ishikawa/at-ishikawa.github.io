CONTAINER_IMAGE=at-ishikawa.github.io/examples/golang/opentelemetry

docker/build:
	docker build --target build -t $(CONTAINER_IMAGE):builder .
	docker build -t $(CONTAINER_IMAGE):latest .
