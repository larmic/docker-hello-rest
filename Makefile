CONTAINER_NAME=docker-hello-rest
IMAGE_NAME=larmic/docker-hello-rest

docker-build:
	@echo "Remove docker image if already exists"
	docker rmi -f ${IMAGE_NAME}:LATEST
	@echo "Build go docker image"
	DOCKER_BUILDKIT=1 docker build -t ${IMAGE_NAME}:LATEST .
	@echo "Prune intermediate images"
	docker image prune --filter label=stage=intermediate -f

docker-run:
	docker run -ti --rm -p 8080:8080 --name ${CONTAINER_NAME} ${IMAGE_NAME}:LATEST