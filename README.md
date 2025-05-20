# showmethatoken
> **B'cause why not?**

[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](LICENSE)

`showmethatoken` is a minimal HTTP service that displays the Bearer token found in the `Authorization` header of incoming requests. It is especially useful in Single Sign-On (SSO) environments where retrieving the access token can be non-trivial.

## 📌 Use Case

This project is designed to assist Kubernetes and OpenShift administrators or users in enterprise environments.
Combined with tools such as [oauth2-proxy](https://github.com/oauth2-proxy/oauth2-proxy), it provides a simple way to inspect tokens passed by reverse proxies or ingress controllers.

## 🔧 How It Works

The service listens for HTTP requests. When a request includes an `Authorization: Bearer <token>` header, it extracts and returns the token in a JSON response.

## 🛠 Deployment

`showmethatoken` is intended to be deployed inside Kubernetes/OpenShift clusters. It is typically used behind a reverse proxy such as oauth2-proxy. A typical setup includes:

1. Deploying `showmethatoken` as a service.
2. Using oauth2-proxy to perform authentication and forward requests to `showmethatoken`.
3. Accessing the service via an Ingress or Route.

## ⚠️ Security Warning

This service exposes the raw Bearer token in the response. **Do not expose this service without proper access control.** It is recommended to use it only in isolated or development environments.

## 📦 Build & Run

### Locally

```bash
go build -o showmethatoken .
./showmethatoken
```

### Docker

```bash
docker build -t showmethatoken .
docker run -p 8080:8080 showmethatoken
```

## 🚀 Releasing with GoReleaser

This project uses [GoReleaser](https://goreleaser.com) to automate builds and Docker image publication.

### Environment Variables

To publish Docker images, make sure to define the following environment variables:

* `DOCKER_REGISTRY`: the target container registry (e.g. `ghcr.io`, `docker.io`, or your internal registry)
* `DOCKER_REGISTRY_NAMESPACE`: the namespace or organization in the registry (e.g. `mycompany`, `username`)
* `GITHUB_TOKEN`: your GITHUB Token

### Examples

```bash
export DOCKER_REGISTRY=docker.io
export DOCKER_REGISTRY_NAMESPACE=myorg
export GITHUB_TOKEN=mytoken
goreleaser release
```

GoReleaser will tag, build, and push the image as:

```bash
$DOCKER_REGISTRY/$DOCKER_REGISTRY_NAMESPACE/showmethatoken:<version>
```

## 📝 License

Licensed under the [Apache License 2.0](LICENSE).

---

© 2025 `showmethatoken` contributors
