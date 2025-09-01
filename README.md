# DevOps Go App

A comprehensive DevOps demonstration project featuring a simple Go web application with complete CI/CD pipeline implementation using Docker, Kubernetes, Helm, and GitOps practices with ArgoCD.

## Table of Contents

- [Project Overview](#project-overview)
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Development](#development)
- [Docker](#docker)
- [Kubernetes Deployment](#kubernetes-deployment)
- [Helm Charts](#helm-charts)
- [GitOps with ArgoCD](#gitops-with-argocd)
- [CI/CD Pipeline](#cicd-pipeline)
- [Contributing](#contributing)

## Project Overview

This project demonstrates modern DevOps practices through a simple Go web application. It includes:

- **Go Web Application**: Simple HTTP server serving static content
- **Containerization**: Docker support with multi-stage builds
- **Orchestration**: Kubernetes manifests and Helm charts
- **Local Development**: Kind cluster configuration
- **GitOps**: ArgoCD for continuous deployment
- **Infrastructure as Code**: Complete automation scripts

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.19 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [Helm](https://helm.sh/docs/intro/install/)
- [Git](https://git-scm.com/downloads)

## Project Structure

```
devops-GO-APP/
├── cleanup.sh              # Cleanup script for resources
├── docker-build.sh         # Docker build automation script
├── Dockerfile              # Multi-stage Docker build
├── gitops/
│   └── argocd/
│       ├── 01-install.md    # ArgoCD installation guide
│       ├── argocd-install.sh # ArgoCD setup script
│       └── go-app.yaml      # ArgoCD application manifest
├── go-app/
│   ├── go.mod              # Go module definition
│   ├── main.go             # Main application code
│   ├── main_test.go        # Application tests
│   └── static/
│       └── abcd.html       # Static web content
├── helm/
│   └── go-app-chart/       # Helm chart for the application
│       ├── Chart.yaml      # Chart metadata
│       ├── templates/      # Kubernetes templates
│       │   ├── deployment.yaml
│       │   ├── ingress.yaml
│       │   └── service.yaml
│       └── values.yaml     # Chart values
├── k8s/
│   └── manifests/          # Raw Kubernetes manifests
│       ├── deployment.yaml
│       ├── ingress.yaml
│       └── service.yaml
├── helm-install.sh         # Helm deployment script
├── k8s-install.sh          # Kubernetes deployment script
├── kind-config.yaml        # Kind cluster configuration
├── kind-setup.sh           # Kind cluster setup script
└── README_RAW.md          # Raw commands reference
```

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/abcofdevops/devops-GO-APP.git
cd devops-GO-APP
```

### 2. Set Up Local Development Environment

```bash
# Create Kind cluster using automation script
./kind-setup.sh

# Or manually:
kind create cluster --config kind-config.yaml
```

**Script Details:**
- `kind-setup.sh` - Automates Kind cluster creation with custom configuration

## 💻 Development

### Running the Go Application Locally

```bash
# Navigate to the application directory
cd go-app

# Initialize Go module (if not already done)
go mod init go-app

# Run the application
go run main.go
```

The application will be available at `http://localhost:8080`

### Building the Application

```bash
# Build the binary
go build

# Run the built binary
./go-app
```

### Running Tests

```bash
# Run tests
go test -v
```

## 🐳 Docker

### Building the Docker Image

```bash
# Build using the automation script
./docker-build.sh

# Or build manually
docker build -t go-app:latest .
```

**Script Details:**
- `docker-build.sh` - Automates Docker image building with proper tagging and optimization

### Running with Docker

```bash
# Run the container
docker run -p 8080:8080 go-app:latest
```

### Local Registry Setup

For local development with Kind, you can set up a local Docker registry:

```bash
# Run local registry
docker run -d -p 5000:5000 --restart=always --name registry registry:2

# Tag and push to local registry
docker tag go-app:v1 localhost:5000/go-app:v1
docker push localhost:5000/go-app:v1
```

### Loading Image to Kind

```bash
# Load image directly to Kind
kind load docker-image go-app:latest --name kind

# Verify images in Kind cluster
docker exec -it kind-control-plane crictl images
```

## Kubernetes Deployment

### Using Raw Manifests

```bash
# Deploy using the automation script
./k8s-install.sh

# Or deploy manually
kubectl apply -f k8s/manifests/deployment.yaml
kubectl apply -f k8s/manifests/service.yaml
kubectl apply -f k8s/manifests/ingress.yaml
```

**Script Details:**
- `k8s-install.sh` - Automates Kubernetes deployment with proper sequencing and validation

### Setting Up Ingress Controller

```bash
# Install NGINX Ingress Controller for Kind
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for ingress controller to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

### Accessing the Application

```bash
# Port forward to access the application
kubectl port-forward deployment/go-app 8080:8080

# Or port forward the service
kubectl port-forward svc/go-app 8080:80
```

Visit `http://localhost:8080` or `http://go-app.local` (if ingress is configured)

### Useful Commands

```bash
# Check deployment status
kubectl get deployments
kubectl get pods
kubectl get services

# Restart deployment
kubectl rollout restart deployment go-app

# Wait for deployment to be ready
kubectl wait --for=condition=Ready pod -l app=go-app
```

## Helm Charts

### Installing with Helm

```bash
# Install using the automation script
./helm-install.sh

# Or install manually
helm install go-app ./helm/go-app-chart
```

**Script Details:**
- `helm-install.sh` - Automates Helm chart installation with dependency management and validation

### Managing Helm Releases

```bash
# List releases
helm list

# Upgrade release
helm upgrade go-app ./helm/go-app-chart

# Uninstall release
helm uninstall go-app
```

### Customizing Values

Edit `helm/go-app-chart/values.yaml` to customize the deployment:

```yaml
# Example values
image:
  repository: go-app
  tag: latest
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 80

ingress:
  enabled: true
  host: go-app.local
```

## GitOps with ArgoCD

### Installing ArgoCD

```bash
# Run the installation automation script
./gitops/argocd/argocd-install.sh

# Or install manually
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

**Script Details:**
- `gitops/argocd/argocd-install.sh` - Automates ArgoCD installation, configuration, and initial setup

### Accessing ArgoCD

```bash
# Port forward ArgoCD server
kubectl port-forward svc/argocd-server -n argocd 8000:443

# Get initial admin password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

Access ArgoCD at `https://localhost:8000` with username `admin` and the retrieved password.

### Deploying Application via ArgoCD

```bash
# Apply the ArgoCD application manifest
kubectl apply -f gitops/argocd/go-app.yaml -n argocd
```

## CI/CD Pipeline

This project is configured for GitHub Actions. The workflow includes:

1. **Build**: Compile the Go application
2. **Test**: Run unit tests
3. **Docker**: Build and push Docker images
4. **Deploy**: Update Kubernetes manifests

For more information on setting up GitHub Actions, refer to the [GitHub Marketplace](https://github.com/marketplace).

## Automation Scripts Overview

This project includes several shell scripts to automate the entire DevOps workflow:

| Script | Purpose | Location |
|--------|---------|----------|
| `kind-setup.sh` | Sets up local Kind Kubernetes cluster | Root directory |
| `docker-build.sh` | Builds and tags Docker images | Root directory |
| `k8s-install.sh` | Deploys application using raw Kubernetes manifests | Root directory |
| `helm-install.sh` | Deploys application using Helm charts | Root directory |
| `argocd-install.sh` | Installs and configures ArgoCD | `gitops/argocd/` |
| `cleanup.sh` | Cleans up all resources and environments | Root directory |

### Script Execution Order

For a complete deployment workflow:

```bash
# 1. Set up local development environment
./kind-setup.sh

# 2. Build Docker image
./docker-build.sh

# 3. Choose deployment method:
# Option A: Raw Kubernetes manifests
./k8s-install.sh

# Option B: Helm charts
./helm-install.sh

# Option C: GitOps with ArgoCD
./gitops/argocd/argocd-install.sh

# 4. Clean up when done
./cleanup.sh
```

To clean up all resources:

```bash
# Run comprehensive cleanup script
./cleanup.sh

# Or manually clean up
helm uninstall go-app 2>/dev/null || true
kubectl delete -f k8s/manifests/ 2>/dev/null || true
kubectl delete namespace argocd 2>/dev/null || true
kind delete cluster
docker rm -f registry 2>/dev/null || true
```

**Script Details:**
- `cleanup.sh` - Comprehensive cleanup script that removes all resources, containers, clusters, and registries created during the project lifecycle

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 📞 Support

If you have any questions or need help with this project, please:

1. Check the existing [Issues](https://github.com/abcofdevops/devops-GO-APP/issues)
2. Create a new issue with detailed information
3. Join the discussion in [Discussions](https://github.com/abcofdevops/devops-GO-APP/discussions)

---

**Happy DevOps-ing!**