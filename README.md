<h1 align="center">Kubernetes Hello World API</h1>
<p align="center">A "Hello world" API designed to be deployed on Kubernetes. Built using Go Gin.</p>
<br />

## Table of contents
- [Table of contents](#table-of-contents)
- [Prerequisites](#prerequisites)
- [Getting started](#getting-started)
  - [1. Install the projects dependencies](#1-install-the-projects-dependencies)
  - [2. Development](#2-development)
  - [3. Build and run](#3-build-and-run)
  - [4. Containerize the project (optional)](#4-containerize-the-project-optional)

## Prerequisites

 - [Go](https://golang.org/doc/install)
 - [Docker](https://docs.docker.com/engine/install/)

## Getting started

### 1. Install the projects dependencies

Before we begin, lets install the projects dependencies,

```bash
go install
```

### 2. Development

We can run and watch for any code changes using a package called [Air](https://github.com/cosmtrek/air),

```bash
make dev
```
### 3. Build and run

Now lets build an executable binary and run the project,

```bash
make build && \
./kubernetes-hello-world-api
```

### 4. Containerize the project (optional)

To build a Docker image of the project,

```bash
make docker-build
```

This next step is intended for use with the [bakeruk/kubernetes-hello-world-kustomization](https://github.com/bakeruk/kubernetes-hello-world-kustomization) project. This will tag and push the newly built Docker image to the the `bakeruk/kubernetes-hello-world-kustomization` local registry,

```bash
make docker-push-local
```