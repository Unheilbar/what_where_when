# deploy
    docker build --tag=grpc-api .
    docker run -it -p 8080:8080 -p 8081:8081 grpc-api
