package main

import (
	"context"
	"cxroot/laundry-app-upload-service/db"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	img "github.com/cxww107/imguploader/imgupload/image"

	"github.com/rs/xid"
)

func readEnvs(env string) (string, error) {
	v := os.Getenv(env)
	if v == "" {
		return "", fmt.Errorf("not found")
	}

	return v, nil
}

func BenchmarkPostImage(b *testing.B) {
	os.Setenv("MONGO_ADDRS", "localhost:27018")
	dbName := xid.New().String()
	handler := imageFileServer{DBName: dbName}
	ctx := context.Background()

	d, err := ioutil.ReadFile("./test_files/test.png")
	if err != nil {
		b.Fatal(err)
	}

	file := &img.File{
		UnixCreatedAt: time.Now().Unix(),
		Processed:     false,
		Extension:     "png",
		Data:          d,
	}

	for i := 0; i < b.N; i++ {
		if _, err := handler.PostImage(ctx, file); err != nil {
			b.Fatal(err)
		}

	}

	db, err := db.NewMongo("localhost:27018", handler.DBName, false)
	if err != nil {
		b.Fatal(err)
	}

	if err = db.DB(handler.DBName).DropDatabase(); err != nil {
		b.Fatal(err)
	}
}
