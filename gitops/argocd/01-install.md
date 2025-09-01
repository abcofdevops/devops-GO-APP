# GitOps with ArgoCD

## Installing ArgoCD

```bash
# Run the installation automation script
./gitops/argocd/argocd-install.sh

# Or install manually
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

**Script Details:**
- `gitops/argocd/argocd-install.sh` - Automates ArgoCD installation, configuration, and initial setup

## Accessing ArgoCD

```bash
# Port forward ArgoCD server
kubectl port-forward svc/argocd-server -n argocd 8000:443

# Get initial admin password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

Access ArgoCD at `https://localhost:8000` with username `admin` and the retrieved password.

## Deploying Application via ArgoCD

```bash
# Apply the ArgoCD application manifest
kubectl apply -f gitops/argocd/go-app.yaml -n argocd
```