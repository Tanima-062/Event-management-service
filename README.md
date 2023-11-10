## Dockerize event-management-service app

## Configure Dockerfile

## To build docker and create image

1. docker build -t eventdocker .

## To see docker image 

2. docker ps 

## To run docker 

3. docker run --network="host" -p 8888:8888 -t eventdocker