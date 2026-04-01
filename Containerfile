# syntax=docker/dockerfile:1
FROM debian:bookworm-slim

# renovate: suite=bookworm depName=curl
ARG CURL_VERSION="7.88.1-10+deb12u8"
# renovate: suite=bookworm depName=git
ARG GIT_VERSION="1:2.39.5-0+deb12u2"

# renovate: suite=bookworm depName=file
ARG FILE="5.46-5"
    
# renovate: suite=bookworm depName=ca-certificates
ARG CA_CERTIFICATES_VERSION="20230311"

ENV DEBIAN_FRONTEND=noninteractive \
    CURL_VERSION=${CURL_VERSION} \
    GIT_VERSION=${GIT_VERSION} \
    CA_CERTIFICATES_VERSION=${CA_CERTIFICATES_VERSION}

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    curl=${CURL_VERSION} \
    git=${GIT_VERSION} \
    ca-certificates=${CA_CERTIFICATES_VERSION} \
 && rm -rf /var/lib/apt/lists/*

CMD ["bash", "-lc", "echo 'mock container with pinned Debian package variables'"]
