---
title: Set up a prometheus by a prometheus operator in a kubernetes cluster
date: 2023-02-26
last_modified_at: 2023-03-12
tags:
  - kubernetes
  - prometheus
  - prometheus-operator
---

This document is written by following [this document](https://prometheus-operator.dev/docs/user-guides/getting-started/).

# Install a prometheus operator

First, create CRDs and resources in the `default` namespace.

```bash
set LATEST (curl -s https://api.github.com/repos/prometheus-operator/prometheus-operator/releases/latest | jq -cr .tag_name)
curl -sLO https://github.com/prometheus-operator/prometheus-operator/releases/download/$LATEST/bundle.yaml
kubectl apply -f bundle.yaml
```

Note that I failed to create a prometheus resource because the CRD definition was too big like next.

```bash
The CustomResourceDefinition "prometheuses.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes
```

Hence, I created the resource by the next command.

```
kubectl apply --server-side=true -f https://raw.githubusercontent.com/prometheus-operator/kube-prometheus/main/manifests/setup/0prometheusCustomResourceDefinition.yaml
```

## Create a Prometheus resource

Create a prometheus server with exposing its admin API
```yml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: prometheus
spec:
  serviceAccountName: prometheus
  podMonitorSelector:
    matchLabels:
      team: frontend
  resources:
    requests:
      memory: 400Mi
  enableAdminAPI: true
```

```yml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
```

```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus
rules:
- apiGroups: [""]
  resources:
  - nodes
  - nodes/metrics
  - services
  - endpoints
  - pods
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources:
  - configmaps
  verbs: ["get"]
- apiGroups:
  - networking.k8s.io
  resources:
  - ingresses
  verbs: ["get", "list", "watch"]
- nonResourceURLs: ["/metrics"]
  verbs: ["get"]
```

```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
- kind: ServiceAccount
  name: prometheus
  namespace: default
```

## An application monitor

On an application to be monitored, create a ServiceMonitor CRD.
For example, if there is the next example app,

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: example-app
  template:
    metadata:
      labels:
        app: example-app
    spec:
      containers:
      - name: example-app
        image: fabxc/instrumented_app
        ports:
        - name: web
          containerPort: 8080
---
kind: Service
apiVersion: v1
metadata:
  name: example-app
  labels:
    app: example-app
spec:
  selector:
    app: example-app
  ports:
  - name: web
    port: 8080
```

Then deploy the ServiceMonitor like next

```yml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: example-app
  labels:
    team: frontend
spec:
  selector:
    matchLabels:
      app: example-app
  endpoints:
  - port: web
```

Confirm if this configuration works.
At first, confirm what kind of metrics are output by example-app.
```bash
kubectl port-forward service/example-app 18080:8080
curl localhost:18080/metrics
```

Then access to the prometheus and see if the query runs

```bash
kubectl port-forward svc/prometheus-operated 9091:9090
```

Then open the browser and see if prometheus query shows the correct data.

If you can't, see [the troubleshooting page](https://prometheus-operator.dev/docs/operator/troubleshooting/) for more details.

# Install Node Exporter

To utilize some metrics of nodes, follow [this article](https://www.civo.com/learn/kubernetes-node-monitoring-with-prometheus-and-grafana).

At first, deploy

```yml
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: node-exporter
  name: node-exporter
  namespace: prometheus
spec:
  selector:
    matchLabels:
      app: node-exporter
  template:
    metadata:
      annotations:
        cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
      labels:
        app: node-exporter
    spec:
      containers:
      - args:
        - --web.listen-address=0.0.0.0:9100
        - --path.procfs=/host/proc
        - --path.sysfs=/host/sys
        image: quay.io/prometheus/node-exporter:v0.18.1
        imagePullPolicy: IfNotPresent
        name: node-exporter
        ports:
        - containerPort: 9100
          hostPort: 9100
          name: metrics
          protocol: TCP
        resources:
          limits:
            cpu: 200m
            memory: 50Mi
          requests:
            cpu: 100m
            memory: 30Mi
        volumeMounts:
        - mountPath: /host/proc
          name: proc
          readOnly: true
        - mountPath: /host/sys
          name: sys
          readOnly: true
      hostNetwork: true
      hostPID: true
      restartPolicy: Always
      tolerations:
      - effect: NoSchedule
        operator: Exists
      - effect: NoExecute
        operator: Exists
      volumes:
      - hostPath:
          path: /proc
          type: ""
        name: proc
      - hostPath:
          path: /sys
          type: ""
        name: sys
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: node-exporter
  name: node-exporter
  namespace: prometheus
spec:
  ports:
  - name: node-exporter
    port: 9100
    protocol: TCP
    targetPort: 9100
  selector:
    app: node-exporter
  sessionAffinity: None
  type: ClusterIP
---
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  labels:
    app: node-exporter
    serviceMonitorSelector: prometheus
  name: node-exporter
  namespace: prometheus
spec:
  endpoints:
  - honorLabels: true
    interval: 30s
    path: /metrics
    targetPort: 9100
  jobLabel: node-exporter
  namespaceSelector:
    matchNames:
    - prometheus
  selector:
    matchLabels:
      app: node-exporter
```

Then change the Prometheus resources

```diff
-  serviceMonitorSelector:
-    matchLabels:
-      team: frontend
+  serviceMonitorSelector: {}
```

# Install Kube State Metrics

Install from helm charts following [the official document](https://github.com/kubernetes/kube-state-metrics#helm-chart).

At first, install kube-state-metrics into a namespace, let's say `kube-state-metrics`.
Then set up a ServiceMonitor
```yml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  labels:
    app: kube-state-metrics
    serviceMonitorSelector: prometheus
  name: kube-state-metrics
  namespace: prometheus
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: kube-state-metrics
  jobLabel: kube-state-metrics
  namespaceSelector:
    matchNames:
    - kube-state-metrics
  endpoints:
  - port: http
```

Then I was able to see some metrics exported by kube-state-metrics.
