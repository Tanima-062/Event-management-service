FROM golang:1.20.5
WORKDIR /app
COPY . /app
RUN go build /app
EXPOSE 8888
ENTRYPOINT ["./event-management-service"]
