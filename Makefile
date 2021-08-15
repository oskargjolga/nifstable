test:
	@go clean -testcache ./...
	@go test ./...

docker-build:
	@docker build --tag nifstable .

docker-run:
	@docker run nifstable

docker-remove:
	@docker rmi nifstable --force