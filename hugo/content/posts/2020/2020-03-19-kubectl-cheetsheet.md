---
date: "2020-03-19T00:00:00Z"
tags:
- kubernetes
- kubectl
title: kubectl cheetsheet
---

# Collect recent error logs
If the logs are outputted by [zap](https://github.com/uber-go/zap), error messages are aggregated by checking level = error.
This log does not work very well if the field `error` contains some unique values like id.

```
$ kubectl logs -l key=value --since=5m | jq -r 'select(.level=="error") | .error' | sort | uniq -c | sort -bgr
```

Instead of `error` field, `msg` or `errorVerbose` could be used.
