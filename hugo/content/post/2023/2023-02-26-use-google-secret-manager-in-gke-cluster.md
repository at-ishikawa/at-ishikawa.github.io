---
date: "2023-02-26T19:00:00Z"
tags:
- kubernetes
- google secret manager
- secret store csi driver
title: Use Google Secret Manager in a GKE cluster
---

There are an [awesome article](https://medium.com/google-cloud/consuming-google-secret-manager-secrets-in-gke-911523207a79) about the options to use the Google Secret Manager and their pros and cons.
In this article, use Secrets Store CSI Driver by following [this page](https://secrets-store-csi-driver.sigs.k8s.io/getting-started/installation.html).

# Set up a secrets store CSI driver

First, install the secrets-store-csi-driver in kube-system namespace

```bash
helm repo add secrets-store-csi-driver https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts
helm install csi-secrets-store secrets-store-csi-driver/secrets-store-csi-driver --namespace kube-system
```

Next, install the GCP provider for the secrets store CSI driver from [this repository](https://github.com/GoogleCloudPlatform/secrets-store-csi-driver-provider-gcp).
Unfortunately, there is no helm chart as the time of this post is written, according to [this GitHub issue](https://github.com/GoogleCloudPlatform/secrets-store-csi-driver-provider-gcp/issues/131).

Just download `deploy/provider-gcp-plugin.yaml` and apply it to the cluster.

```yml
kubectl apply -f deploy/provider-gcp-plugin.yaml
```

# Use the secrets from a kubernetes

Just follow [the usage](https://github.com/GoogleCloudPlatform/secrets-store-csi-driver-provider-gcp#usage) described in the secrets-store-csi-driver-provider-gcp repository.

These are overview, though I changed names of namespaces:

* The secret name is `secrets_store_csi_driver_test`
* The pod uses a secrets store CSI driver needs a permission to access Google Secret Manager, like with a workload identity
* GCP Service account name is `gke-secrets-store-csi-test` with the role `roles/secretmanager.secretAccessor`
* The pod is in `test-secrets-store-csi` namespace and the service account name is `default`

Then I create following k8s resources.

```yml
apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: app-secrets
spec:
  provider: gcp
  parameters:
    secrets: |
      - resourceName: "projects/$PROJECT_ID/secrets/secrets_store_csi_driver_test/versions/latest"
        path: "good1.txt"
      - resourceName: "projects/$PROJECT_ID/secrets/secrets_store_csi_driver_test/versions/latest"
        path: "good2.txt"
```

```yml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: mypodserviceaccount
  namespace: default
  annotations:
    iam.gke.io/gcp-service-account: gke-workload@$PROJECT_ID.iam.gserviceaccount.com
---
apiVersion: v1
kind: Pod
metadata:
  name: mypod
  namespace: default
spec:
  serviceAccountName: mypodserviceaccount
  containers:
  - image: gcr.io/google.com/cloudsdktool/cloud-sdk:slim
    imagePullPolicy: IfNotPresent
    name: mypod
    resources:
      requests:
        cpu: 100m
    stdin: true
    stdinOnce: true
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    tty: true
    volumeMounts:
      - mountPath: "/var/secrets"
        name: mysecret
  volumes:
  - name: mysecret
    csi:
      driver: secrets-store.csi.k8s.io
      readOnly: true
      volumeAttributes:
        secretProviderClass: "app-secrets"
```

After I deployed k8s resources and configure GCP resources, I was able to see the secrets are mounted:

```bash
[personal|test-secrets-store-csi] > kubectl exec -it mypod /bin/bash
root@mypod:/# ls /var/secrets
good1.txt  good2.txt
root@mypod:/# cat /var/secrets/good1.txt
foo
root@mypod:/# cat /var/secrets/good2.txt
foo
```


## Sync as Kubernetes Secret

This is an alpha feature, but it's possible to sync and generate the data as a k8s secret, described in [this page](https://secrets-store-csi-driver.sigs.k8s.io/topics/sync-as-kubernetes-secret.html).
**Note that the volume mount is required for the Sync With Kubernetes Secrets**

To use this feature, at first, enable the feature on the helm chart.

```yml
syncSecret:
  enabled: true
```

Then change k8s resources like next.
I checked how to write the configuration for secret objects from the [test code of the provider](https://github.com/GoogleCloudPlatform/secrets-store-csi-driver-provider-gcp/blob/12450d1ac7098fa3fea1593d666950c0c2c1d008/test/e2e/templates/test-sync.yaml.tmpl#L24-L29).
```diff
@@ -11,6 +11,12 @@ spec:
         path: "good1.txt"
       - resourceName: "projects/$PROJECT_ID/secrets/secrets_store_csi_driver_test/versions/latest"
         path: "good2.txt"
+  secretObjects:
+  - data:
+    - key: test-secret-contents
+      objectName: "good1.txt"
+    type: Opaque
+    secretName: test-synced-secret
 ---
 apiVersion: v1
 kind: ServiceAccount
@@ -31,6 +37,12 @@ spec:
   - image: gcr.io/google.com/cloudsdktool/cloud-sdk:slim
     imagePullPolicy: IfNotPresent
     name: mypod
+    env:
+    - name: SECRET_CONTENTS
+      valueFrom:
+        secretKeyRef:
+            name: test-synced-secret
+            key: test-secret-contents
     resources:
       requests:
         cpu: 100m
```

If you don't mount it on a volume mount, you'll get an error like this

```yml
status:
  conditions:
    containerStatuses:
    - waiting:
        message: secret "test-synced-secret" not found
        reason: CreateContainerConfigError
```

When you succeed to deploy, you can see the secrets like

```fish
[personal|test-secrets-store-csi] > kubectl get secrets
NAME                  TYPE                                  DATA   AGE
default-token-qbwgz   kubernetes.io/service-account-token   3      10d
test-synced-secret    Opaque                                1      1s
[personal|test-secrets-store-csi] > kubectl view-secret test-synced-secret
Choosing key: test-secret-contents
foo
[personal|test-secrets-store-csi] > kubectl exec -it mypod /bin/bash
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
root@mypod2:/# echo $SECRET_CONTENTS
foo
root@mypod2:/# exit
exit
```


## Additional features

There are a a few features that are still alpha:

* [Auto rotation](https://secrets-store-csi-driver.sigs.k8s.io/topics/secret-auto-rotation.html)
    * The secrets mounted on a pod will be automatically updated, but this doesn't support rotating secrets on an application side. Application has to implement a way to update secrets by detecting the updates
    * `SecretProviderClassPodStatus` resource stores the binding of a secret and a pod, and it also contains the version of the secret.


