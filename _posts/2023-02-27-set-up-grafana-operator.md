---
title: Set up a Grafana operator
date: 2023-02-27T00:00:00
tags:
  - kubernetes
  - grafana
  - grafana-operator
---

Follow [this document](https://github.com/grafana-operator/grafana-operator/blob/master/documentation/deploy_grafana.md) mainly.

# Set up a Grafana

First, clone a repository and apply a change.

```bash
kubectl kustomize deploy/manifests -o grafana-operator.yml
kubectl apply -f grafana-operator.yml
```

Then create a Grafana resource in the same namespace.
I tried to create it in the different namespace, but didn't work for some reasons.

```yml
apiVersion: integreatly.org/v1alpha1
kind: Grafana
metadata:
  name: example-grafana
  namespace: grafana-operator-system
spec:
  client:
    preferService: true
  ingress:
    enabled: True
    pathType: Prefix
    path: "/"
  config:
    log:
      mode: "console"
      level: "error"
    log.frontend:
      enabled: true
    auth:
      disable_login_form: False
      disable_signout_menu: True
    auth.anonymous:
      enabled: True
  service:
    name: "grafana-service"
    labels:
      app: "grafana"
      type: "grafana-service"
  dashboardLabelSelector:
    - matchExpressions:
        - { key: app, operator: In, values: [grafana] }
  resources:
    # Optionally specify container resources
    limits:
      cpu: 200m
      memory: 200Mi
    requests:
      cpu: 100m
      memory: 100Mi
```

After a grafana-deployment pod should be deployed, port-forward and confirm you can access the grafana without auth.

## Set up a Prometheus datasource

Then create a datasource from [an example](https://github.com/grafana-operator/grafana-operator/blob/master/deploy/examples/datasources/Prometheus.yaml)

```yml
apiVersion: integreatly.org/v1alpha1
kind: GrafanaDataSource
metadata:
  name: example-grafanadatasource
  namespace: grafana-operator-system
spec:
  name: middleware.yaml
  datasources:
    - name: Prometheus
      type: prometheus
      access: proxy
      url: http://prometheus-operated.prometheus:9090
      isDefault: true
      version: 1
      editable: true
      jsonData:
        tlsSkipVerify: true
        timeInterval: "5s"
```

Also, update the above Grafana resource to login by an admin user.

```diff
    auth.anonymous:
-      enabled: True
+      enabled: False
```

Then confirm an admin password from the secret

```bash
kubectl view-secret grafana-admin-credentials GF_SECURITY_ADMIN_PASSWORD
```

Now you can see a Prometheus from a datasource.
