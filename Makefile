test:
	@go test -v ./...

docker-build:
	@docker build --tag nifstable .

docker-run:
	@docker run nifstable

docker-remove:
	@docker rmi nifstable --force