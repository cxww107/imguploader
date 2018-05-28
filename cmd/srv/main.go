package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cxww107/imguploader/imgupload"
	img "github.com/cxww107/imguploader/imgupload/image"

	"google.golang.org/grpc"
)

var mongoSSLMode bool

var envs = [2]string{
	"IMGUPL_MONGO_ADDR",
	"DOCKER_SECRET_LOC",
}

func main() {
	if err := checkEnvs(); err != nil {
		log.Fatal(err)
	}

	if os.Getenv("MONGO_AUTH_MODE") == "SSL" {
		mongoSSLMode = true
	}

	creds, err := getCreds()
	if err != nil {
		log.Fatalf("Failed to create credentials: %v", err)
	}

	srv := grpc.NewServer(grpc.MaxMsgSize(15*1024*1024),
		grpc.MaxRecvMsgSize(imgupload.MaxImgSizeMb*1024*1024),
		grpc.Creds(*creds))

	handler := imageFileServer{DBName: "images"}

	img.RegisterFileHandlerServer(srv, handler)
	l, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatalf("Cound not listen to port %s: %v", "8888", err)
	}

	log.Fatal(srv.Serve(l))
}

func checkEnvs() error {
	for _, env := range envs {
		if os.Getenv(env) == "" {
			return fmt.Errorf("Enviroment variable %v is not set", env)
		}
	}

	return nil
}
