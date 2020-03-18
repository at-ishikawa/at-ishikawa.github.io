---
title: Terraform for GCP
---
Wrote in November, 2019.

Troubleshootings
===

GKE
---

### To use kubernetes provider
Got `Error: Post https://[ip_address]/api/v1/namespaces/[namespace]/secrets: x509: certificate signed by unknown authority`.
This is related [GitHub issue](https://github.com/terraform-providers/terraform-provider-helm/issues/37).

**google_container_cluster** resource outputs *cluster_ca_certificate*, *client_key*, or *client_certificate* and they are base64 encoded.


### Secrets cannot be created
Got `Error: secrets is forbidden: User "system:anonymous" cannot create resource "secrets" in API group "" in the namespace "[namespace]"`
This is unresolved.

If you run the terraform localy and you have kubernetes credentials, it may be because you do not have cluster admin roles.
But if you try to create a cluster at the same time with adding secrets by terraform on CI, it seems no way yet.
There is an [issue](https://github.com/terraform-providers/terraform-provider-kubernetes/issues/176) on GitHub related with this error, but no good solution yet.


Cloud SQL
---
### Cloud SQL instance cannot be created
Got `Error: googleapi: Error 409: The instance or operation is not in an appropriate state to handle the request., invalidState`
This is related [GitHub issue](https://github.com/hashicorp/terraform/issues/20972) and [PR](https://github.com/GoogleCloudPlatform/magic-modules/pull/1634).

If Cloud SQL instance is tried to be created but another one with the same name was created and deleted within a week, then this error may happen.
This may be fixed in recent versions of google providers.


Cloud Storage
---
### Bucket cannot be created
Got `Error: googleapi: Error 403: The bucket you tried to create is a domain name owned by another user., forbidden`.
The domain name of GCS has to be verified by following the steps of [this page](https://cloud.google.com/storage/docs/domain-name-verification#verification).
