run:
    docker build -t majula . && \
    docker rm majula_backend
    docker run \
        --name majula_backend \
        -p 8000:8000 \
        majula
