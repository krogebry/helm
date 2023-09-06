FROM grafana/k6 as k6

FROM releases-docker.jfrog.io/jfrog/jfrog-cli-v2-jf:2.45.0 AS jfrog

FROM jnorwood/helm-docs:latest AS helm-docs

FROM golangci/golangci-lint:v1.54 AS golangci-lint

FROM docker:24.0-dind AS dind

FROM golang:1.21.0-bookworm as go

FROM debian:bookworm-slim

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates gnupg \
    curl jq wget \
    make gcc gettext-base build-essential \
    git gh \
    kubernetes-client \
    python3 python3-pip \
    unzip watch openssh-client \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Terraform Version manager
ENV TF_VERSION=1.5.6
ENV PATH=${PATH}:/opt/tfenv/bin
RUN git clone --depth=1 https://github.com/tfutils/tfenv.git /opt/tfenv && \
    tfenv install ${TF_VERSION} && \
    tfenv use ${TF_VERSION}

# Install Terraform docs utility
RUN curl -fsSL -o ./terraform-docs.tar.gz https://terraform-docs.io/dl/v0.16.0/terraform-docs-v0.16.0-linux-amd64.tar.gz && \
    tar -xzf terraform-docs.tar.gz && \
    chmod +x terraform-docs && \
    mv terraform-docs /usr/local/bin/ && \
    rm -rf terraform-docs.tar.gz

# Install Terraform lint utility
RUN curl -fsSL https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash

# Install AWS cli
RUN curl -fsSL -o awscliv2.zip https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf awscliv2.zip

# Install helm and helm-diff
RUN curl -fsSL https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash && \
    helm plugin install https://github.com/databus23/helm-diff

# Install yq utility
RUN curl -fsSL -o yq_linux_amd64.tar.gz https://github.com/mikefarah/yq/releases/download/v4.34.1/yq_linux_amd64.tar.gz && \
    tar -xzf yq_linux_amd64.tar.gz && \
    chmod +x yq_linux_amd64 && \
    mv yq_linux_amd64 /usr/local/bin/yq && \
    rm -rf yq_linux_amd64.tar.gz

# Install `just`
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin/

# Install correct docker and containerd
RUN install -m 0755 -d /etc/apt/keyrings && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg && \
    chmod a+r /etc/apt/keyrings/docker.gpg && \
    echo "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
    "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

RUN apt-get update && apt-get install --no-install-recommends -y \
    containerd.io docker-buildx-plugin docker-ce docker-ce-cli docker-compose-plugin  \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=k6 /usr/bin/k6 /usr/bin/k6

RUN groupadd --gid 10000 actions
RUN useradd -m actions -u 10000 -g 10000

# Install kube-score
RUN wget -O /usr/local/bin/kube-score https://github.com/zegl/kube-score/releases/download/v1.16.1/kube-score_1.16.1_linux_amd64 \
  && chmod +x /usr/local/bin/kube-score

# Docker in docker
COPY --from=dind /usr/local/bin/docker /usr/local/bin/

# golangci-lint
COPY --from=golangci-lint /usr/bin/golangci-lint /usr/local/bin/

# Helm docs
COPY --from=helm-docs /usr/bin/helm-docs /usr/local/bin/helm-docs

# JFrog CLI
COPY --from=jfrog /usr/local/bin/jf /usr/local/bin/jf

COPY --from=go /usr/local/go/ /usr/local/go/
ENV PATH=${PATH}:/usr/local/go/bin/

# Go test format
RUN go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest && \
    mv /root/go/bin/gotestfmt /usr/local/bin/

WORKDIR /home/actions/runner

ENV RUNNER_VERSION=2.308.0

RUN wget -q https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz \
  && tar xzf ./actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz \
  && rm ./actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz

RUN chown -R actions:actions ~actions && ./bin/installdependencies.sh
COPY entrypoint.sh ./

USER actions
ENV COMPlus_EnableDiagnostics=0

ENTRYPOINT ["/usr/bin/dumb-init", "--", "./entrypoint.sh"]

# RUN git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"
