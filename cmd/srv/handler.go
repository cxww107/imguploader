package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/cxww107/imguploader/db"

	img "github.com/cxww107/imguploader/imgupload/image"

	"google.golang.org/grpc/credentials"

	"github.com/globalsign/mgo/bson"
)

type imageFileServer struct {
	DBName string
}

func getCreds() (*credentials.TransportCredentials, error) {
	p, err := filepath.Abs("./cmd/srv/certs")
	if err != nil {
		return nil, err
	}

	cert, err := tls.LoadX509KeyPair(filepath.Join(p, "cert.pem"),
		filepath.Join(p, "key.pem"))
	if err != nil {
		return nil, fmt.Errorf("key or cert pem: %v", err)
	}

	ca, err := ioutil.ReadFile(filepath.Join(p, "ca.pem"))
	if err != nil {
		return nil, fmt.Errorf("failed to read CA pem: %v", err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, fmt.Errorf("failed to append CA cert")
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	})

	return &creds, nil
}

func (i imageFileServer) MarkProcessed(ctx context.Context, in *img.ID) (*img.Void, error) {
	return nil, fmt.Errorf("method MarkProcessed is not implemented yet")
}

func (i imageFileServer) GetNewImagesIDs(ctx context.Context, v *img.Void) (*img.IDs, error) {
	m, err := db.NewMongo(os.Getenv("IMGUPL_MONGO_ADDRS"), i.DBName, mongoSSLMode, os.Getenv("DOCKER_SECRET_LOC"))
	defer m.Session.Close()

	if err != nil {
		return nil, err
	}

	ids := make([]string, 0)
	iter := m.DB(m.DBName).GridFS("images").Find(nil).Iter()

	var r struct {
		ID bson.ObjectId `bson:"_id,omitempty"`
	}

	for iter.Next(&r) {
		ids = append(ids, r.ID.Hex())
	}

	return &img.IDs{Ids: ids}, nil
}

func (i imageFileServer) GetCount(context.Context, *img.Void) (*img.Count, error) {
	return nil, fmt.Errorf("not implemented")
}

func (i imageFileServer) PostImage(ctx context.Context, file *img.File) (*img.Void, error) {
	log.Println("Startet Post Image")
	m, err := db.NewMongo(os.Getenv("IMGUPL_MONGO_ADDR"), i.DBName, mongoSSLMode, os.Getenv("DOCKER_SECRET_LOC"))
	if err != nil {
		return &img.Void{}, err
	}
	defer m.Session.Close()

	d, err := json.Marshal(file)
	if err != nil {
		return &img.Void{}, fmt.Errorf("Post image failed: %v", err)
	}

	log.Println("I start insert one file")
	err = m.InsertOneFile(ctx, "images", d)
	if err != nil {
		return &img.Void{}, err
	}

	return &img.Void{Msg: "1 file inserted"}, nil
}

func (i imageFileServer) GetImage(ctx context.Context, in *img.ID) (*img.File, error) {
	m, err := db.NewMongo(os.Getenv("IMGUPL_MONGO_ADDR"), i.DBName, mongoSSLMode, os.Getenv("DOCKER_SECRET_LOC"))
	defer m.Session.Close()
	if err != nil {
		return nil, err
	}

	id := bson.ObjectIdHex(in.ID)

	resc := make(chan interface{})
	defer close(resc)

	// Read in goroutine
	go func() {
		gf, err := m.DB(m.DBName).GridFS("images").OpenId(id)
		if err != nil {
			resc <- err
		}
		defer gf.Close()

		d, err := ioutil.ReadAll(gf)
		if err != nil {
			resc <- err
		}

		resc <- d
	}()

	// Wait for result or cancellation signal
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout reached in GetImage")
	case res := <-resc:
		switch val := res.(type) {
		case error:
			return nil, err
		case []byte:
			var f img.File
			err := json.Unmarshal(val, &f)
			if err != nil {
				return nil, err
			}
			return &f, nil
		}

	}
	return nil, fmt.Errorf("internal error")
}
