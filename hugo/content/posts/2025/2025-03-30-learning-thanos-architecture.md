---
title: Learning Grafana Mimir architecture
date: 2025-03-30
---

The overview of Thanos is quite concisely explained in the following video.

<iframe width="560" height="315" src="https://www.youtube.com/embed/SR2sG4wno-s?si=-8uyUNQPTmAUb-Tg" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>


In this article, I collected some information, mainly from components pages in the official documents like [this compactor page](https://thanos.io/tip/components/compact.md/).


## Components
### Query Frontend

This component is stateless and follow a few features.

- Query splitting
    - Split long query into multiple short queries, based on `--query-range.split.interval` flag value. The default is 24 hours.
- Retry and caching


### Querier/Query

This is stateless and horizontally scalable.
This component is to aggregate and deduplicate multiple metrics.

Thanos Query sends queries from HA prometheus hair and deduplicate when there is a gap based on a replica label set by `--query.replica-label` flag.
The gap between Prometheus HA group happens when, for instance, there is an outage of one instance.
There are 2 algorithms for deduplication

- penalty: It's mainly used from Prometheus HA group
- chain: 1;1 deduplication for samples, and useful for receivers.


### Sidecar

This is a component running along with a Prometheus instance.
This components

- Runs queries to Prometheus data
- Upload TSDB blocks to an object storage
    - It seems it's not recommended to upload compacted blocks from Prometheus


### Store

This is a component to access an object storage bucket.
Memcached and Redis can be used for index, chunks of TSDB blocks, and metadata caches.


### Compactor

This is a component to run compaction and downsampling.

- Compaction: This is to reduce the number of blocks and the size of index indices
    - blocks from single source is recognized by external labels. External labels have to be unique and persistent.
- Downsampling: reduce overall resolution without losing accuracy
    - Creating 5m downsampling for blocks older than 40 hours (2 days)
    - Creating 1h downsampling for blocks older than 10 days (2 weeks)
- Vertical compaction
    - Overlapping TSDB blocks into one using external labels.
    - There are a few use cases for this feature with risks, races between multiple compactions, backfilling, or offline deduplication
- Deleting blocks after retention periods

For scalability, use label sharding for horizontal scaling.


### Ruler

> NOTE: It is recommended to keep deploying rules inside the relevant Prometheus servers locally. Use ruler only on specific cases.


### Receiver

This component is to receive remote write requests from Prometheus, retain blocks, and uploads TSDB blocks for every 2 hours by default.

This supports

- Replication by gRPC or Cap'N Proto
- Multi-tenancy using external labels
- Active Series Limiting
    - This is to prevent ingesting too high number of active series per tenant


## Some features
### Sharding

For sharding, there are 2 options, one of them is only applicable to store gateway.

- relabelling: to select values of external labels for TSDB blocks
- time partitioning: to select blocks by `--min-time` and `--max-time` flags. This is only available for store gateway
