---
title: Getting Started with Grafana Tempo
date: 2023-03-12
tags:
  - grafana
  - grafana tempo
---

# Grafana Tempo
Follow [this example](https://grafana.com/docs/tempo/latest/getting-started/example-demo-app/#helm).

- Install grafana tempo with [microservices-tempo-values.yaml](https://github.com/grafana/tempo/blob/main/example/helm/microservices-tempo-values.yaml).
    - I commented out `storage` values because when I tried to use GCS, it failed due to the permission issue to access GCS
- Added the grafana tempo as the data source of Grafana and see if it works
    - The service is `-gateway` created by the helm chart

# Update Grafana Operator

I deployed the Grafana by grafana-operator on [this post](/2023/02/27/set-up-grafana-operator/).
So, it was required to update Grafana by following [this example](https://github.com/grafana-operator/grafana-operator/blob/a33ef1d58c85196c7cae7158e9c796c4df9da084/deploy/examples/datasources/Tempo.yaml#L4)

```yml
@@ -4,6 +4,7 @@ metadata:
   name: example-grafana
   namespace: grafana-operator-system
 spec:
+  baseImage: grafana/grafana:9.4.3
   client:
     preferService: true
   ingress:
@@ -11,6 +12,8 @@ spec:
     pathType: Prefix
     path: "/"
   config:
+    feature_toggles:
+      enable: "tempoSearch,tempoServiceGraph,tempoApmTable,traceqlEditor"
     log:
       mode: "console"
       level: "error"
```

Note that according to [this comment](https://github.com/grafana/helm-charts/issues/813#issuecomment-967048718), the Grafana version has to be at least 8.2.


# Deploy load test

Deploy a k6 load test from [this document](https://github.com/grafana/tempo/tree/main/example/helm)
Note that `endpoint` might need to be updated depends on the configuration.
(This didn't work for some reasons).

{% comment %}
Instead, I developed a simple implementation to generate a trace by a golang.
{% endcomment %}
