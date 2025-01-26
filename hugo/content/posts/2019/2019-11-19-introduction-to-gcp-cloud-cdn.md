---
date: "2019-11-19T00:00:00Z"
tags:
- gcp
- cloud cdn
title: Introduction to GCP Cloud CDN
---

# Target upstream services
Cloud CDN can have only GCP load balancer as the upstream services.
And GCP load balancer can configure one of followings for backends.

1. Backend services
    1. GCE instance groups
    1. GCE Network endpoint groups. These are the groups of VM instances
1. Backend buckets

See [official document](https://cloud.google.com/cdn/docs/overview) for the architecture using this.


# Configurations
Cache can be controlled by response headers of origin servers, like cache expiration times, like `Cache-Control: max-age` header.
The details are described in [here](https://cloud.google.com/cdn/docs/caching#expiration).


# Reference
[Cloud CDN](https://cloud.google.com/cdn/docs/overview)
