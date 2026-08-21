# trainspotter

One simple trainspotting application to keep track which ICE trains I already saw and when I saw them.
I want to collect them all :)

## Deployment

### Local (Docker Compose)

```bash
docker compose up
```

- Frontend: http://localhost:3000
- Backend: http://localhost:8080

### Kubernetes

**Prerequisites**

1. Build and load images into the cluster:
```bash
make build
```

2. Create secrets:
```bash
kubectl create secret generic postgres-secret \
  --from-literal=POSTGRES_USER=trainspotter \
  --from-literal=POSTGRES_PASSWORD=<password> \
  --from-literal=POSTGRES_DB=trainspotter

kubectl create secret generic backend-secret \
  --from-literal=DATABASE_URL=postgresql://trainspotter:<password>@postgres:5432/trainspotter
```

3. Apply manifests:
```bash
kubectl apply -f deploy/k8s/ --recursive
```

- Frontend: http://localhost:30000
