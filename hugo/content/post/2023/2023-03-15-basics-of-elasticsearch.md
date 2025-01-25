---
date: "2023-03-15T00:00:00Z"
tags:
- elasticsearch
title: The Basics of Elasticsearch
---

Following videos are helpful to understand the overview of Elasticsearch more.

Elasticsearch architecture
<iframe width="560" height="315" src="https://www.youtube.com/embed/2WJFMYAri_8" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

About search relevance
<iframe width="560" height="315" src="https://www.youtube.com/embed/CCTgroOcyfM" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>


## Documents and index

![documents](/assets/images/posts/2023/03/15/basics-of-elasticsearch/documents.jpg)

- Document: JSON object, equivalent to a row in a table of RDBMS
- Index: The set of documents collected by the same type of data. For exaple, one index is for a user, second one is for a product in an e-commerce service.


## Elasticsearch Cluster
![cluster architecture](/assets/images/posts/2023/03/15/basics-of-elasticsearch/cluster.jpg)

Elasticsearch is a distributed system for search.

- Shard: the part of index. On multiple nodes, all documents in an index is split by multiple shards.
    - Primary shard: The original shard
    - Replica shard: The copy of its primary shard. It can be used to increase the throughput

## Search

Basic terms for search:
- Term is a word to search documents

Elastic search searches a term by looking up an **inverse index** and find documents

To create an inverse index, first, a tokenization is required to split tokens from a document.
In most cases, this splits the document into each word

Then from the tokens, create inverse indices to look up a document by each token


### Relevance score

There are a few algorithms to score a search result.
The related factor is

- Term frequency: Frequency of a term in a document
- Document frequency: Frequency of a query in all documents

### Trade off of relevance: Precision vs Recall

![precision vs recall](/assets/images/posts/2023/03/15/basics-of-elasticsearch/precision vs recall.jpg)

- Precision: Accuracy of positive results
    - True Positive / (True Positive + False Positive)
- Recall: How much the correct data is retrieved
    - True Positive / (True Positive + False Negative)


## Questions

- How to search the data on Elasticsearch? Is inverted index included in each shard to look up a word?
    - Yes. See [this article](https://www.devinline.com/2018/09/elasticsearch-inverted-index-and-its-storage.html#:~:text=Elasticsearch%20uses%20a%20special%20data,documents%20in%20which%20it%20appears.)
