package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	img "github.com/cxww107/imguploader/imgupload/image"

	"github.com/rs/xid"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func download(ctx context.Context, client img.FileHandlerClient) error {
	// Get IDs of new images
	ids, err := client.GetNewImagesIDs(ctx, &img.Void{})
	if err != nil {
		return err
	}

	tasks := make([]interface{}, len(ids.Ids))
	for i, id := range ids.Ids {
		tasks[i] = &img.ID{ID: id}
	}

	imgc := dispatchTasks(ctx, tasks, downloader, client)

	for f := range imgc {
		switch v := f.(type) {
		case *img.File:
			go createFile(v)
		case error:
			log.Println(v)
		default:
			log.Println("failed to save a file: received wrong type from mongodb")
		}
	}

	return nil
}

func downloader(ctx context.Context, fileID interface{}, client img.FileHandlerClient) chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		res := make(chan interface{})
		go func() {
			defer close(res)
			id, ok := fileID.(*img.ID)
			if !ok {
				res <- fmt.Errorf("downloader received wrong type for ID. It shoud be *img.ID")
				return
			}
			f, err := client.GetImage(ctx, id)
			if err != nil {
				res <- err
				return
			}
			res <- f
		}()
		// Wait for result or leave
		select {
		case <-ctx.Done():
			return
		case v := <-res:
			out <- v
		}
	}()

	return out
}

func createFile(file *img.File) {

	fmt.Println("start writing file")

	if file == nil {
		log.Println("Error writing file. Received nil file")
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Saving of file is failed: %v \n", err)
		return
	}

	filename := filepath.Join(cwd, xid.New().String())
	fmt.Println(filename)
	nf, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("Saving of file is failed: %v \n", err)
		return
	}
	defer nf.Close()

	_, err = nf.Write(file.Data)
	if err != nil {
		log.Printf("Saving of file is failed: %v \n", err)
		return
	}

	log.Printf("File saved: %v\n", filename)
}
