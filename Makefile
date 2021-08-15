test:
	@go test -v ./...

make docker-build-and-run: docker-build docker-run

docker-build:
	@docker build --tag nifstable .

docker-run:
	@docker run nifstable

docker-remove:
	@docker rmi nifstable --force