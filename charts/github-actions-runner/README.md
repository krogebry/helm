<!--
STOP! README.md is automatically generated using helm-docs
Run `helm-docs .` to generate.
If you're looking at README.md.gotmpl, then you're in the right place.
-->
# github-actions-runner

![Version: 0.10.3-0.3.7](https://img.shields.io/badge/Version-0.10.3--0.3.7-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.291.1-20.04-1](https://img.shields.io/badge/AppVersion-2.291.1--20.04--1-informational?style=flat-square)

Deploys a StatefulSet of [GitHub Actions self-hosted runners](https://github.com/actions/runner)—
allowing you to run your [GitHub Actions](https://github.com/features/actions) workflows on your
own infrastructure.

## Requirements

- Kubernetes >1.14
- Helm ~3

## Runner Registration (Authentication)

This chart intentionally avoids using a GitHub Personal Access Token with `admin:org` privileges. As such, you will need to provide a runner registration token upon chart installation and when increasing the replica count.

**⚠️ Important: Follow these steps to setup your runners correctly**.

#### Registration Tokens

An Actions Runner must _register_ before obtaining the credentials necessary to perform actions. Registration is accomplished by calling the GitHub-provided `config.sh` script with a _registration token_. Once a runner has registered, it no longer needs the registration token. Please be sure to provide a registration token to this chart via values (`runner.registrationToken`) _before_ the token expires.

#### Obtaining a registration token

[**ℹ️ Note: Registration tokens expire in 1 hour**](https://github.community/t/api-to-generate-runners-token/16963/8)

A registration token is provided when selecting "Add new" -> "New Runner" in the GitHub UI as described in [the "Adding self-hosted runners" documentation](https://docs.github.com/en/free-pro-team@latest/actions/hosting-your-own-runners/adding-self-hosted-runners).

#### Using a registration token with this chart

A registration token can be used to register multiple runners before its expiry. Therefore, this chart utilizes one registration token to register as many runners as the highest of these values: `replicaCount` or `autoscaling.maxReplicas`.

When autoscaling is disabled, the registration token will be used as many as `replicaCount` times to register runners (each pod). After registration, the long-lived runner credentials for each pod are stored on its persistent volume so that the runner pod can be rescheduled or restarted and resume working again without registering again.

When autoscaling is enabled, it is necessary to use the registration token immediately to obtain `autoscaling.maxReplicas` long-lived runner credentials so that the registration token may expire and as many as `autoscaling.maxReplicas` long-lived runner credentials have been saved to persistent volumes and can be started as needed.

### Why a StatefulSet?

Runners are stateful! That is, they need a short-lived registration token in order to register themselves, but after that point, they store their own credentials on disk. To preserve the runner credentials on disk for each pod, we use a persistent volume claim per pod, so each pod can be rescheduled or restarted and when started, can authenticate by itself (without the registration token) and continue working.

The token provided when using "add new runner" in the UI is a temporary registration token that can be used to register an actions runner. Once the runner is registered (using config.sh), it stores its own credentials locally. The registration token is short-lived (1 hour) and cannot be renewed. But after registration is complete, the runner should be able to continue to stay authenticated using its own credentials.

So where does the runner store credentials? ./.runner, ./.credentials, etc, local to the run.sh and config.sh script. Because they're local to the scripts, we can't mount a volume on top of the scripts, so we copy the scripts to a volume directory and allow the registration credentials to be written alongside the scripts inside the volume. Thus, when the runner restarts, it doesn't need to re-register, it simply reconnects with existing credentials.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `4` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `50` |  |
| dind.enabled | bool | `false` | If you'd like your runner to be docker-enabled and you understand and accept the security risks associated with a priveleged container running actions, enable here |
| dind.image.pullPolicy | string | `"IfNotPresent"` |  |
| dind.image.repository | string | `"docker"` |  |
| dind.image.tag | string | `"19.03.13-dind"` |  |
| dind.resources | object | `{}` |  |
| dind.securityContext.privileged | bool | `true` |  |
| fullnameOverride | string | `""` |  |
| hostNetwork.enabled | bool | `false` |  |
| imagePullSecrets | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| persistence.accessModes[0] | string | `"ReadWriteOnce"` |  |
| persistence.annotations | object | `{}` |  |
| persistence.enabled | bool | `true` | Highly recommended so runners can retain their registration and renew tokens automatically |
| persistence.size | string | `"8Gi"` |  |
| podAnnotations | object | `{}` |  |
| podManagementPolicy | string | `"Parallel"` |  |
| podSecurityContext.runAsGroup | int | `1000` |  |
| podSecurityContext.runAsNonRoot | bool | `true` |  |
| podSecurityContext.runAsUser | int | `1000` |  |
| podSecurityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| replicaCount | int | `1` | If not using autoscaling, set the desired runner count |
| resources.limits.cpu | string | `"500m"` |  |
| resources.limits.memory | string | `"1028Mi"` |  |
| resources.requests.cpu | string | `"500m"` |  |
| resources.requests.memory | string | `"1028Mi"` |  |
| runner.extraEnv | list | `[]` |  |
| runner.ghr_pat_token | string | `""` | The registration token obtained from GitHub's "Add new runner" registrationToken: "" |
| runner.image.pullPolicy | string | `"IfNotPresent"` |  |
| runner.image.repository | string | `"docker.artifactory-ha.tri-ad.tech/mtfuji/devops/github-actions-runner"` |  |
| runner.image.tag | string | `"0.3.1"` | Default is the chart appVersion. |
| runner.labels | string | `"mtf-runner,dev"` | Optional labels for the runner to match against in your workflows |
| runner.name | string | `""` | Set if you want custom runner names. Defaults to the pod name |
| runner.resources.limits.cpu | string | `"1000m"` |  |
| runner.resources.limits.memory | string | `"1028Mi"` |  |
| runner.resources.requests.cpu | string | `"1000m"` |  |
| runner.resources.requests.memory | string | `"1028Mi"` |  |
| runner.scope | string | `""` | Either `myorg` or `myorg/myrepo` for an organization-scoped runner or repo-scoped runner, respectively. |
| runner.securityContext.allowPrivilegeEscalation | bool | `false` |  |
| runner.securityContext.capabilities.drop[0] | string | `"all"` |  |
| runner.securityContext.privileged | bool | `false` |  |
| runner.securityContext.readOnlyRootFilesystem | bool | `true` |  |
| runner.securityContext.runAsGroup | int | `1000` |  |
| runner.securityContext.runAsNonRoot | bool | `true` |  |
| runner.securityContext.runAsUser | int | `1000` |  |
| runner.securityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| runner.url | string | `"https://github.tri-ad.tech/TRI-AD/avmetricsproxy"` |  |
| securityContext.allowPrivilegeEscalation | bool | `false` |  |
| securityContext.capabilities.drop[0] | string | `"all"` |  |
| securityContext.privileged | bool | `false` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsGroup | int | `1000` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `1000` |  |
| securityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `""` | If not set and create is true, a name is generated using the fullname template |
| tolerations | list | `[]` |  |
