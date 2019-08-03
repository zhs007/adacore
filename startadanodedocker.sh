docker container stop adanode
docker container rm adanode
docker run -d \
    --name adanode \
    -p 7201:7201 \
    -v $PWD/cfg:/app/adanode/cfg \
    -v $PWD/logs:/app/adanode/logs \
    -v $PWD/output:/app/adanode/output \
    adanode