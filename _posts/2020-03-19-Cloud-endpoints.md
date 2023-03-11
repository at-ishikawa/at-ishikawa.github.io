---
date: 2020-03-19
title: Cloud endpoints
---
Written in Nov, 2019.

The Cloud endpoint is actually the NGINX proxy which offers the following features on GCP.
1. Authentication and validation
1. Logging and monitoring in GCP

The overall architecture for this is described in [official page](https://cloud.google.com/endpoints/docs/openapi/architecture-overview).

# Supported protocol
They support
1. OpenAPI. The spec of OpenAPI is described [here](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md).
1. gRPC
1. Rest APIs using Cloud Endpoints Framework

# Supported environments
1. Endpoint on Cloud Run
1. Endpoint on GKE as a sidecar
1. Endpoint on App Engine

For Cloud Run or GKE, docker images are available in GCP registry. You can use the image like `gcr.io/endpoints-release/endpoints-runtime:1`. See released and secure image versions in [here](https://github.com/cloudendpoints/esp#released-esp-docker-images).

# Reference
* [Official page](https://cloud.google.com/endpoints/)
* [Quickstart](https://cloud.google.com/endpoints/docs/quickstart-endpoints)
