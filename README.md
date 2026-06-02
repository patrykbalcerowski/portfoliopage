# Kubernetes Homelab

Personal homelab environment built on K3s and Proxmox.

The repository contains Kubernetes manifests, infrastructure configuration and a small Go application used for testing GitOps workflows, monitoring and deployment automation.

## Stack

* K3s
* ArgoCD
* Helm
* Terraform
* Ansible
* Prometheus
* Grafana
* Cloudflare Tunnel
* Nginx Gateway API
* Go
* Proxmox

## Implemented

* [x] GitOps deployments with ArgoCD
* [x] Helm-based application management
* [x] Monitoring with Prometheus and Grafana
* [x] Cloudflare Tunnel integration
* [x] Gateway API routing
* [x] Custom application written in Go
* [x] Infrastructure automation
* [x] Kubernetes homelab running on Proxmox

## Environment Overview

The cluster is hosted on Proxmox and runs K3s as the Kubernetes distribution.

Applications are deployed using Helm and managed through ArgoCD. Changes pushed to the repository are automatically synchronized to the cluster following GitOps principles.

External access is provided through Cloudflare Tunnel and routed internally using Nginx Gateway API. Monitoring is handled by Prometheus and Grafana.

The repository also contains a custom Go application used for exposing metrics and serving a simple portfolio page.

## Deployment Flow

```text
Git Repository
      │
      ▼
   ArgoCD
      │
      ▼
 Kubernetes
      │
      ├── Helm Releases
      ├── Monitoring Stack
      ├── Cloudflare Tunnel
      └── Custom Go Application
```


## Repository Structure

```text
argocd/         ArgoCD applications
deployment/     Infrastructure deployments
infra/          Gateway API, monitoring and cluster configuration
go-application/ Custom Go application
node-exporter/  Kubernetes manifests
```

## Deployment Notes

This repository contains configuration specific to my environment.

Before deployment:

* Update all Route, Gateway and hostname definitions.
* Replace domain names with your own.
* Configure Cloudflare Tunnel credentials.
* Update DNS records in your domain.
* Review ArgoCD application paths and repository references.
* Review any hardcoded IP addresses or network ranges.

## Deployment

### 1. Install ArgoCD

```bash
kubectl create namespace argocd

kubectl apply -n argocd \
-f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

### 2. Configure Cloudflare Tunnel

Update `cloudflared.yaml` with your tunnel token or create the appropriate Kubernetes Secret.

```bash
kubectl apply -f cloudflared.yaml
```

Complete the tunnel configuration in the Cloudflare Dashboard:

https://developers.cloudflare.com/cloudflare-one/networks/connectors/cloudflare-tunnel/get-started/create-remote-tunnel/

When creating public hostnames, routes should point to the Kubernetes Gateway service, for example:

```text
http://main-gateway-nginx.default.svc.cluster.local
```

### 3. Configure DNS

Create DNS entries for the services you want to expose.

Example:

```text
argocd.example.com
grafana.example.com
app.example.com
```

### 4. Bootstrap Applications

```bash
kubectl apply -f argocd/
```

ArgoCD will deploy and maintain the remaining applications and infrastructure components defined in this repository.


## Example DNS Records

```text
argocd.example.com     -> Cloudflare Tunnel
grafana.example.com    -> Cloudflare Tunnel
app.example.com        -> Cloudflare Tunnel
```

## Purpose

This repository exists primarily as a portfolio project and learning environment used to explore:

* Kubernetes administration
* GitOps workflows
* Infrastructure as Code
* Monitoring and observability
* Linux systems administration
* Go development
