package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cxww107/imguploader/imgupload"

	img "github.com/cxww107/imguploader/imgupload/image"

	"google.golang.org/grpc"
)

const (
	// URL is server address
	URL = "0.0.0.0:8888"

	// ServerCN is common name from server certificate
	ServerCN = "localhost"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Command missed. Variants: upload, download")
		os.Exit(1)
	}

	creds, err := getCreds(ServerCN)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(URL,
		grpc.WithTransportCredentials(*creds),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(imgupload.MaxImgSizeMb*1024*1024),
			grpc.MaxCallSendMsgSize(imgupload.MaxImgSizeMb*1024*1024)))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := img.NewFileHandlerClient(conn)

	switch f := flag.Arg(0); f {
	case "upload":
		if flag.NArg() < 2 {
			err = fmt.Errorf("include filename or folder in command")
			break
		}
		err = upload(ctx, client, flag.Arg(1))
	case "download":
		err = download(ctx, client)
		if err != nil {
			break
		}
	default:
		err = fmt.Errorf("unknown command, %v", f)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
