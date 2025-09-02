ARG BASE_IMAGE=docker.io/library/alpine:3.22.0

FROM ${BASE_IMAGE} AS builder

ARG TARGETARCH

ARG KUBECTL_VERSION=1.33.4

WORKDIR /workspace

RUN apk add --update --no-cache curl \
    && curl -LO https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/${TARGETARCH}/kubectl \
    && curl -LO https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/${TARGETARCH}/kubectl.sha256 \
    && echo "$(cat kubectl.sha256) kubectl" | sha256sum -cs \
    && chmod 0755 kubectl

FROM ${BASE_IMAGE}

COPY --from=builder /workspace/kubectl /usr/bin/kubectl

COPY deliver_kubectl.sh .

ENTRYPOINT ["./deliver_kubectl.sh"]
