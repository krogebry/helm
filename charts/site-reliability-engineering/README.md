# Site Reliability Engineering

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

A Helm chart for Kubernetes

## Source Code

## Maintainers
| Name | Email | Url |
| ---- | ------ | --- |

## Local testing

[docs](https://github.com/prometheus/blackbox_exporter/blob/master/CONFIGURATION.md#grpc_probe)

The `blackbox.yaml` file is used for testing local changes and exploring with the idea.

First start the container:
```shell
docker run \
  --rm \
  -p 9115:9115 \
  -v ./blackbox.yaml:/etc/blackbox.yaml \
  -v $(pwd):/config \
  --name blackbox_exporter \
  quay.io/prometheus/blackbox-exporter:latest --config.file=/etc/blackbox.yaml --log.level=debug
```

Now curl the local endpoint to invoke a module test:

```shell
curl "http://localhost:9115/probe?module=http_2xx_proxy&target=https://visualizer.dev.apc.arene.com"
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `100` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `80` |  |
| config.modules.dns_mtfuji.dns.preferred_ip_protocol | string | `"ip4"` |  |
| config.modules.dns_mtfuji.dns.query_name | string | `"check"` |  |
| config.modules.dns_mtfuji.dns.query_type | string | `"CNAME"` |  |
| config.modules.dns_mtfuji.dns.transport_protocol | string | `"udp"` |  |
| config.modules.dns_mtfuji.dns.validate_answer_rrs.fail_if_not_matches_regexp[0] | string | `"aws.infra.mtfuji.w3n.io"` |  |
| config.modules.dns_mtfuji.prober | string | `"dns"` |  |
| config.modules.dns_mtfuji.timeout | string | `"5s"` |  |
| config.modules.grpc.grpc.preferred_ip_protocol | string | `"ip4"` |  |
| config.modules.grpc.grpc.tls | bool | `true` |  |
| config.modules.grpc.prober | string | `"grpc"` |  |
| config.modules.grpc_plain.grpc.tls | bool | `false` |  |
| config.modules.grpc_plain.prober | string | `"grpc"` |  |
| config.modules.http_2xx_proxy.http.follow_redirects | bool | `false` |  |
| config.modules.http_2xx_proxy.http.preferred_ip_protocol | string | `"ip4"` |  |
| config.modules.http_2xx_proxy.http.proxy_url | string | `"http://tmc-bridge-63f406f71d3cf22a.elb.ap-northeast-1.amazonaws.com:3128"` |  |
| config.modules.http_2xx_proxy.http.valid_http_versions[0] | string | `"HTTP/1.1"` |  |
| config.modules.http_2xx_proxy.http.valid_http_versions[1] | string | `"HTTP/2.0"` |  |
| config.modules.http_2xx_proxy.prober | string | `"http"` |  |
| config.modules.http_2xx_proxy.timeout | string | `"5s"` |  |
| config.modules.https_oauth_blocked.http.fail_if_header_not_matches[0].header | string | `"location"` |  |
| config.modules.https_oauth_blocked.http.fail_if_header_not_matches[0].regexp | string | `"login.windows.net"` |  |
| config.modules.https_oauth_blocked.http.fail_if_not_ssl | bool | `true` |  |
| config.modules.https_oauth_blocked.http.fail_if_ssl | bool | `false` |  |
| config.modules.https_oauth_blocked.http.ip_protocol_fallback | bool | `false` |  |
| config.modules.https_oauth_blocked.http.method | string | `"GET"` |  |
| config.modules.https_oauth_blocked.http.no_follow_redirects | bool | `true` |  |
| config.modules.https_oauth_blocked.http.preferred_ip_protocol | string | `"ip4"` |  |
| config.modules.https_oauth_blocked.http.tls_config.insecure_skip_verify | bool | `false` |  |
| config.modules.https_oauth_blocked.http.valid_http_versions[0] | string | `"HTTP/1.1"` |  |
| config.modules.https_oauth_blocked.http.valid_http_versions[1] | string | `"HTTP/2.0"` |  |
| config.modules.https_oauth_blocked.http.valid_status_codes[0] | int | `302` |  |
| config.modules.https_oauth_blocked.prober | string | `"http"` |  |
| config.modules.https_oauth_blocked.timeout | string | `"5s"` |  |
| config.modules.tls_connect.prober | string | `"tcp"` |  |
| config.modules.tls_connect.tcp.tls | bool | `true` |  |
| config.modules.tls_connect.timeout | string | `"5s"` |  |
| configExistingSecretName | string | `""` |  |
| containerPort | int | `9115` |  |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"docker.artifactory-ha.tmc-stargate.com/arene/plane/core/sre/blackbox-exporter"` |  |
| image.tag | string | `"v0.23.0"` |  |
| imagePullSecrets[0].name | string | `"regcred"` |  |
| infra.env_name | string | `"dev"` |  |
| ingress.annotations | object | `{}` |  |
| ingress.className | string | `""` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hosts[0].host | string | `"chart-example.local"` |  |
| ingress.hosts[0].paths[0].path | string | `"/"` |  |
| ingress.hosts[0].paths[0].pathType | string | `"ImplementationSpecific"` |  |
| ingress.tls | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations."prometheus.io/port" | string | `"8125"` |  |
| podAnnotations."prometheus.io/prefix" | string | `"app.woven.arene.apc.control-plane."` |  |
| podAnnotations."prometheus.io/scrape" | string | `"true"` |  |
| podAnnotations."prometheus.istio.io/merge-metrics" | string | `"true"` |  |
| podAnnotations."seccomp.security.alpha.kubernetes.io/pod" | string | `"runtime/default"` |  |
| podAnnotations."sidecar.istio.io/inject" | string | `"false"` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources.limits.cpu | string | `"200m"` |  |
| resources.limits.memory | string | `"2Gi"` |  |
| resources.requests.cpu | string | `"100m"` |  |
| resources.requests.memory | string | `"1Gi"` |  |
| secretConfig | bool | `false` |  |
| securityContext.allowPrivilegeEscalation | bool | `false` |  |
| securityContext.capabilities.drop[0] | string | `"all"` |  |
| securityContext.privileged | bool | `false` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsGroup | int | `1000` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `1000` |  |
| securityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| service.port | int | `8125` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `""` |  |
| tolerations | list | `[]` |  |
