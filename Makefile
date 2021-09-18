build-prod-api-image:
	docker build -f docker/api/Dockerfile.api.prod -t api-prod-image:latest .