test:
	@go test -v ./...

table: docker-build docker-run

table-halftime: docker-build docker-run-halftime

docker-build:
	@docker build --tag nifstable .


docker-run:
	@docker run nifstable

docker-run-halftime:
	@docker run nifstable -halftime

docker-remove:
	@docker rmi nifstable --force