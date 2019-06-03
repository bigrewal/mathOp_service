package main

import (
	"context"
	"io/ioutil"
	"log"
	"math"
	"mathOp_service/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
)

type server struct{}

type conf struct {
	Port string `yaml:"port"`
}

func main() {
	var c conf
	c.getConf()

	listener, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	service.RegisterMathOpServiceServer(srv, &server{})
	reflection.Register(srv)

	log.Println("Running on port:", c.Port)

	if e := srv.Serve(listener); e != nil {
		log.Fatalf("failed to serve: %v", e)
	}

}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (s *server) Add(ctx context.Context, request *service.AddRequest) (*service.AddResponse, error) {
	log.Println("Addition request received: ", request)
	a, b := request.GetX(), request.GetY()
	result := a + b
	log.Println("Addition request response:", result)
	return &service.AddResponse{Answer: result}, nil
}

func (s *server) Sqrt(ctx context.Context, request *service.SqrtRequest) (*service.SqrtResponse, error) {
	log.Println("Sqrt request received: ", request)
	a := request.GetX()

	result := math.Sqrt(a)
	log.Println("Sqrt request response:", result)
	return &service.SqrtResponse{Answer: result}, nil
}
