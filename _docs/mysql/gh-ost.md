---
title: gh-ost
---

gh-ost
===

[gh-ost](https://github.com/github/gh-ost) is an online migration tool for MySQL developed by GitHub.

There is a video to describe what issue gh-ost is used to solve.

<iframe width="560" height="315" src="https://www.youtube.com/embed/2zksJnRSgv0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>


Getting Started
---

### Build a binary from source code

We need go which is version > 1.15

```
> wget https://github.com/github/gh-ost/releases/download/v1.1.2/gh-ost-binary-linux-20210617134741.tar.gz
--2021-07-28 15:40:14--  https://github.com/github/gh-ost/releases/download/v1.1.2/gh-ost-binary-linux-20210617134741.tar.gz
Resolving github.com (github.com)... 192.30.255.112
Connecting to github.com (github.com)|192.30.255.112|:443... connected.
HTTP request sent, awaiting response... 302 Found
Location: https://github-releases.githubusercontent.com/54378638/0a93f200-cf85-11eb-8870-035762e21f3a?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20210728%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210728T063905Z&X-Amz-Expires=300&X-Amz-Signature=993eae9ee5e3b05bc105d62feb28ffe1d76ee2dd3c1ef7c2a77d7e16c72b3ccb&X-Amz-SignedHeaders=host&actor_id=0&key_id=0&repo_id=54378638&response-content-disposition=attachment%3B%20filename%3Dgh-ost-binary-linux-20210617134741.tar.gz&response-content-type=application%2Foctet-stream [following]
--2021-07-28 15:40:14--  https://github-releases.githubusercontent.com/54378638/0a93f200-cf85-11eb-8870-035762e21f3a?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20210728%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210728T063905Z&X-Amz-Expires=300&X-Amz-Signature=993eae9ee5e3b05bc105d62feb28ffe1d76ee2dd3c1ef7c2a77d7e16c72b3ccb&X-Amz-SignedHeaders=host&actor_id=0&key_id=0&repo_id=54378638&response-content-disposition=attachment%3B%20filename%3Dgh-ost-binary-linux-20210617134741.tar.gz&response-content-type=application%2Foctet-stream
Resolving github-releases.githubusercontent.com (github-releases.githubusercontent.com)... 185.199.108.154, 185.199.110.154, 185.199.111.154, ...
Connecting to github-releases.githubusercontent.com (github-releases.githubusercontent.com)|185.199.108.154|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 4710185 (4.5M) [application/octet-stream]
Saving to: ‘gh-ost-binary-linux-20210617134741.tar.gz’

gh-ost-binary-linux-202106171 100%[=================================================>]   4.49M  4.18MB/s    in 1.1s

2021-07-28 15:40:15 (4.18 MB/s) - ‘gh-ost-binary-linux-20210617134741.tar.gz’ saved [4710185/4710185]

> tar -xzvf gh-ost-binary-linux-20210617134741.tar.gz
gh-ost
> ./gh-ost --version
1.1.2
> mv gh-ost /path/to/bin/
```

### Test to run a gh-ost cli
