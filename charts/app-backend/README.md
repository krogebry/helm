# <CHARTNAME>

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

Manging infrastructure for the <CHARTNAME> service

## Source Code

* <source1>

## Maintainers
| Name | Email | Url |
| ---- | ------ | --- |
| Bryan Kroger | bryan.kroger@gmail.com | https://github.com/bryan-kroger |

## Local Installation

[Optional]: first get a the current values of the chart so we can replicate what is currently set.

```shell
helm-deploy get values <CHARTNAME> > /tmp/<CHARTNAME>.yaml
```

Skip the variables and just set the image.tag value:

```
helm upgrade --install <CHARTNAME> ./ \
  --set image.tag=<YOUR_TAG>
```

Where:
* YOUR_TAG is any tag you're using to test.

Use the values, this is handy for testing a new version of the helm chart in place:

```shell
helm-deploy upgrade --install <CHARTNAME> ./ --values /tmp/<CHARTNAME>.yaml
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `100` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `80` |  |
| database.hostname | string | `"fqdn.amazonaws.com"` | FQDN that lives in MTF wired up to the RDS instance via a private link. |
| database.port | string | `"5432"` | Port number |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"Always"` |  |
| image.repository | string | `"docker.com/<CHARTNAME>"` |  |
| image.tag | string | `"dev"` |  |
| imagePullSecrets | list | `[]` |  |
| ingress.annotations | object | `{}` |  |
| ingress.className | string | `""` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hosts[0].host | string | `"chart-example.local"` |  |
| ingress.hosts[0].paths[0].path | string | `"/"` |  |
| ingress.hosts[0].paths[0].pathType | string | `"ImplementationSpecific"` |  |
| ingress.tls | list | `[]` |  |
| shared_clusteri.app_name | string | `"my-app"` | MTF application name |
| shared_clusteri.cloud | string | `"aws"` | MTF cluster: [aws, gke] |
| shared_clusteri.cluster_family_id | int | `0` | MTF cluster family ID: [0, 1] |
| shared_clusteri.cluster_family_name | string | `"gc"` | MTF cluster family: ["gc", "gpu"] |
| shared_clusteri.domain | string | `"infra.mydomain.com"` | MTF domain |
| shared_clusteri.env_name | string | `"apps"` | MTF prod or non-prod: [apps-prod, apps] |
| shared_clusteri.namespace_name | string | `"application-name-dev"` | MTF namespace name |
| shared_clusteri.region | string | `"ap-northeast-1"` | MTF region: [us-east-1, ap-northeast-1, us-central1] |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources.limits.cpu | string | `"400m"` |  |
| resources.limits.memory | string | `"4000Mi"` |  |
| resources.requests.cpu | string | `"200m"` |  |
| resources.requests.memory | string | `"2000Mi"` |  |
| securityContext.allowPrivilegeEscalation | bool | `false` |  |
| securityContext.capabilities.drop[0] | string | `"all"` |  |
| securityContext.privileged | bool | `false` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsGroup | int | `1000` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `1000` |  |
| securityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| service.port | int | `8080` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `false` |  |
| serviceAccount.name | string | `""` |  |
| tolerations | list | `[]` |  |
