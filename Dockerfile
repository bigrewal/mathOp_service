FROM golang
WORKDIR $GOPATH/src/mathOp_service
COPY . .
RUN go get -d -v ./...
EXPOSE 8080
ENTRYPOINT go run main.go
