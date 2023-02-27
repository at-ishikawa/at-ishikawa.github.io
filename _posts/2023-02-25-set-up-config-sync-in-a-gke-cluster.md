---
title: Set up a config sync in a GKE cluster
date: 2023-02-25T00:00:00
tags:
  - gke
  - kubernetes
  - config sync
---

Basically, follow [this document](https://cloud.google.com/anthos-config-management/docs/tutorials/config-sync-multi-repo)

# Enable a Config Sync
1. Select **Install Config Sync > your cluster"
1. Clear the Enable Policy Controller checkbox and click Next.
1. Leave the Enable Config sync checkbox enabled.
1. In the Repository list, select "Custom".
1. Set your repository on your URL and click "SHOW ADVANCED SETTINGS"
1. Choose how to read a GitHub repository by one of the ways described in [this document](https://cloud.google.com/anthos-config-management/docs/how-to/installing-kubectl).
    1. I used a GitHub Token which will be expired shortly. And set it out on our cluster.
1. Set the Configuration directory as the root directory on the repository. I set `config-sync`
1. Change the Source format to hierarchy. You can see the details about it in [here](https://cloud.google.com/anthos-config-management/docs/concepts/hierarchical-repo).

# Set up a GitHub repository

Create files under the config sync directory set above and see if the config sync works.
The guestbook is the manifest set no the other getting started in [this post](/2023/02/09/kubebuilder/).

```bash
> tree config-sync
config-sync
├── namespaces
│   └── kubebuilder
│       ├── guestbook.yml
│       └── namespace.yml
└── system
    └── repo.yml

3 directories, 3 files
```

I added its namespace and guestbook manifests on each yml file.
The system/repo.yml file looks like next.

```yml
# system/repo.yaml
kind: Repo
apiVersion: configmanagement.gke.io/v1
metadata:
  name: repo
spec:
  version: "0.1.0"
```

# Install CLI to manage a config sync: nomos

By following [this section](https://cloud.google.com/anthos-config-management/docs/tutorials/config-sync-multi-repo#sync-status), we can install the CLI by

```shell
gcloud components install nomos
```

Then you can check the status like next

```bash
> nomos status
Connecting to clusters...

*personal
  --------------------
  <root>:root-sync                         https://github.com/path/to/repository/config-sync@main
  SYNCED @ 2023-02-25 21:10:34 -0800 PST   53444e744f2d4f7e07a6d602400a2d9d05e63620
  Managed resources:
     NAMESPACE     NAME                                                      STATUS    SOURCEHASH
                   namespace/kubebuilder                                     Current   53444e7
     kubebuilder   guestbook.webapp.at-ishikawa.github.io/guestbook-sample   Current   53444e7

```

# Configure a monitoring on Google Cloud Monitoring

As default, there are errors on permissions on an open telemetry collector installed by a config-system namespaces.

```shell
otel-collector-67d9f55576-xfkmd otel-collector 2023-02-26T06:13:31.097Z warn    batchprocessor/batch_processor.go:178Sender failed    {"kind": "processor", "name": "batch", "pipeline": "metrics/kubernetes", "error": "failed to export time series to GCM: rpc error: code = PermissionDenied desc = Permission monitoring.timeSeries.create denied (or the resource may not exist)."}
```

In this case, grant a permission the service account a write permission for Google Monitoring, by following by [this document](https://cloud.google.com/anthos-config-management/docs/how-to/monitoring-config-sync#custom-monitoring).

- Create a service account on the GCP
- Grant the service account workload identity user for the default SA in the config-management-monitoring namespace
- Grant the `roles/monitoring.metricWriter` role to the GCP service account
- Add an annotation `iam.gke.io/gcp-service-account: $GSA_NAME@$PROJECT_ID.iam.gserviceaccount.com` on the default SA on the config-management-monitoring namespace

Then rollout the otel-collector deployment in the namespace.

# Other configurations

## CRD
Put CRDs under `/cluster` directory.
See [this document](https://cloud.google.com/anthos-config-management/docs/how-to/cluster-scoped-objects#configure_customresourcedefinitions) for more details.

# Trouble shooting

1. If `system/repo` is missing under the directory of a config sync, then we'll get next error. In that case, add a file `repo.yml` described in the above.

```
KNV1017: The system/ directory must declare a Repo Resource. path: system/ For more information, see https://g.co/cloud/acm-errors#knv1017
```

1. If there is no namespace file but there is another resource, it shows an error and create a resource for the namespace.

```
KNV1044: The directory "kubebuilder" has configs, but is missing a Namespace config. All bottom level subdirectories MUST have a Namespace config.

path: namespaces/kubebuilder/

For more information, see https://g.co/cloud/acm-errors#knv1044
```
