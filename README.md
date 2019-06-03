### Math Op service
MathOp service is a microservice implemented in go-lang which can be used by its clients to perform basic math operations like addition and square-root. 

#### Project Structure
`config/config.yaml` contains the port number on which the service will run. 

`service/service.proto` This contains the mathOp gRPC service definitions.

`main.go` contains the service implementation.

##### How to run it?
1. Clone this repo and place it in the `{GOPATH}/src` directory
2. `cd mathOp_service`
3. Execute `go run main.go` or run `docker run -d -p 8080:8080 bigrewal/mathop-service:v1.0`

##### Future Improvements

1. Add a logger interceptor (centralised logging) instead of having log statements in the service implementation
2. Add a request interceptor which will make sure that the incoming request is valid before forwarding it to the service.
3. Need to check if the mathOp gRPC service definition needs to be put in a seperate repository?

##### Kubernetes

Deployment and Service objects exist in mathop-service.yaml
