goose-create name:
	docker run --rm \
	  -v ./internal/infrastructure/db/migrations:/migrations \
	  -e GOOSE_COMMAND="create" \
	  -e GOOSE_COMMAND_ARG="{{name}} sql" \
	  ghcr.io/kukymbr/goose-docker:latest-amd64

run:
    docker build -t majula . && \
    docker rm majula_backend || true
    docker run \
        --name majula_backend \
        -p 8000:8000 \
        majula
