---
title: Securing Swift
description: Securing Swift
menu:
  product_swift_0.8.1:
    identifier: guides-security
    name: Securing Swift
    parent: guides
    weight: 15
product_name: swift
menu_name: product_swift_0.8.1
section_menu_id: guides
---

# Securing Swift

There are 3 aspects to securing swift connections:

- User <-> Swift Server
- Swift Server <-> Tiller Server
- Swift Server <-> Chart Repo


## Serve Swift api over SSL

To serve Swift api over SSL connections, you provide a server certificate pair via `--tls-cert-file` and `--tls-private-key-file` flags. If you use a self-signed certificate pair, pass the CA certificate via `--tls-ca-file` flag.

```
  --tls-ca-file string                      File containing CA certificate
  --tls-cert-file string                    File container server TLS certificate
  --tls-private-key-file string             File containing server TLS private key
```

You can generate certificate pair using tools like [onessl](https://github.com/kubepack/onessl), openssl, cfssl. Below are instructions for generating certificate pairs using onessl:

- Download onessl binary from project's Github release page:

```console
# Linux amd64:
curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.3.0/onessl-linux-amd64 \
  && chmod +x onessl \
  && sudo mv onessl /usr/local/bin/

# Linux arm64:
curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.3.0/onessl-linux-arm64 \
  && chmod +x onessl \
  && sudo mv onessl /usr/local/bin/

# Mac OSX amd64:
curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.3.0/onessl-darwin-amd64 \
  && chmod +x onessl \
  && sudo mv onessl /usr/local/bin/
```

- Now create a ca certificate pair using onessl.

```console
$ mkdir swift
$ cd swift
$ onessl create ca-cert

$ ls -b1
ca.crt
ca.key
```

- Now create a server certificate pair using the ca-certificate from previous step. Pass the domain names used to connect to Swift server using `--domains` flag

```console
$ onessl create server-cert --domains=swift.kube-system.svc
Wrote server certificates in  /home/tamal/Desktop/swift

$ ls -b1
ca.crt
ca.key
server.crt
server.key
```

- Now create a secret to upload certificates to a Kubernetes cluster and mount the secret is Swfift deployment.

```console
$ kubectl create secret generic swift-ssl -n kube-system \
  --from-file=./ca.crt \
  --from-file=./server.crt \
  --from-file=./server.key

secret "swift-ssl" created
```


## Using authentication with Chart Repository

Swift can download charts from Chart repository using basic auth, bearer auth and/or client cert auth. `InstallRelease` and `UpdateRelease` api support the following input parameters with their respectie api calls:

| Parameter              |            | Description                                                                |
|------------------------|------------| ---------------------------------------------------------------------------|
| `chart_url`            | `Required` | URL to download chart archive.                                              |
| `ca_bundle`            | `Optional` | PEM encoded CA bundle used to sign server certificate of chart repository. |
| `username`             | `Optional` | Username for basic authentication to the chart repository.                 |
| `password`             | `Optional` | Password for basic authentication to the chart repository.                 |
| `token`                | `Optional` | Bearer token for authentication to the chart repository.                   |
| `client_certificate`   | `Optional` | PEM-encoded data passed as a client cert to chart repository.              |
| `client_key`           | `Optional` | PEM-encoded data passed as a client key to chart repository.               |
| `insecure_skip_verify` | `Optional` | Skip certificate verification for chart repository.                        |


## Securely connecting to Tiller Server

You can run Swift server in the same pod as the Tiller server and connect over localhost. This will ensure that traffic between Tiller server and Swift proxy is not visible to outside parties.

If you are using [SSL with your Tiller Server](https://github.com/kubernetes/helm/blob/master/docs/tiller_ssl.md), you can pass the ca certificate and/or client certificate pair to Swift server to secure connect to Tiller server over TLS.

```
  --tiller-ca-file string                   File containing CA certificate for Tiller server
  --tiller-client-cert-file string          File container client TLS certificate for Tiller server
  --tiller-client-private-key-file string   File containing client TLS private key for Tiller server
  --tiller-endpoint string                  Endpoint of Tiller server, eg, [scheme://]host:port
  --tiller-insecure-skip-verify             Skip certificate verification for Tiller server
```

The server certificate used with Tiller should have the following Subject Alternative Names (SANS) to ensure that Swift can connect to it whether running in the same namespace or in different namespaces.

| Subject Alternative Name (SANS) | Usage                                                    |
|---------------------------------|----------------------------------------------------------|
| tiller-svc                      | When tiller and swift both running in same namespace     |
| tiller-svc.tiller-namespace     |                                                          |
| tiller-svc.tiller-namespace.svc | When tiller and swift are running in different namespace |
