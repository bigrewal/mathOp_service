#### Math Op service.

##### How to run it?

1. `cd mathOp_service`
2. Execute `go run main.go` or run `docker run -d -p 8080:8080 bigrewal/mathop-service:v1.0`

##### Future Improvements

1. Add a logger interceptor (centralised logging) instead of having log statements in the service implementation

2. Add a request interceptor which will make sure that the incoming request is valid before forwarding it to the service.

###### Kubernetes

Deployment and Service objects exist in mathop-service.yaml
