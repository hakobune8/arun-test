# Deployment Guide

## Prerequisites
- `kubectl` configured for your target cluster
- `helm` v3+
- Container registry access for pushing images

## Building the Image
```bash
docker build -t arun-test:latest .
docker push arun-test:latest
```

## Deploying with Helm
```bash
helm upgrade --install arun-test ./charts/arun-test \
  --set image.repository=<your-registry>/arun-test \
  --set image.tag=latest \
  --set image.pullPolicy=IfNotPresent
```

## Verifying Deployment
```bash
kubectl get pods -l app.kubernetes.io/name=arun-test
kubectl logs -l app.kubernetes.io/name=arun-test
kubectl get svc arun-test
```

## Troubleshooting
- **Pod CrashLoopBackOff**: Check logs for missing environment variables or port conflicts.
- **Service Unreachable**: Verify that the Service selector matches the Deployment labels.
- **Resource Limits**: Adjust `resources.requests` and `resources.limits` in `values.yaml` if the game server requires more CPU/memory.