package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	img "github.com/cxww107/imguploader/imgupload/image"
)

func upload(ctx context.Context, client img.FileHandlerClient, filename string) error {
	root, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("error during creation file absolute path: %v", err)
	}

	paths := make([]string, 0)

	if filename[len(filename)-1:] == "/" {
		// check if it is a folder
		var dir []os.FileInfo
		dir, err = ioutil.ReadDir(root)
		if err != nil {
			return err
		}
		for _, d := range dir {
			paths = append(paths, filepath.Join(root, d.Name()))
		}

	} else {
		// check if it is a file
		_, err = ioutil.ReadFile(root)
		if err != nil {
			return err
		}
		paths = append(paths, root)
	}

	tasks := make([]interface{}, len(paths))
	for i, v := range paths {
		tasks[i] = v
	}

	res := dispatchTasks(ctx, tasks, uploader, client)
	for r := range res {
		switch val := r.(type) {
		case error:
			log.Println(val)
		case *img.Void:
			fmt.Println(val.Msg)
		}
	}

	return nil
}

func uploader(ctx context.Context, path interface{}, client img.FileHandlerClient) chan interface{} {

	out := make(chan interface{})

	go func() {
		defer close(out)
		res := make(chan interface{})
		go func() {
			defer close(res)
			p, ok := path.(string)
			if !ok {
				res <- fmt.Errorf("uploader received wrong type: should be string")
				return
			}

			d, err := ioutil.ReadFile(p)
			if err != nil {
				res <- fmt.Errorf("uploader failed to read a file: %v", err)
				return
			}
			file := img.File{
				UnixCreatedAt: time.Now().Unix(),
				Data:          d,
				Processed:     false,
				Extension:     filepath.Ext(p)}

			c, cancel := context.WithTimeout(ctx, time.Second*60)
			defer cancel()

			v, err := client.PostImage(c, &file)
			if err != nil {
				res <- err
			}
			res <- v
		}()

		// Wait for result or leave
		select {
		case <-ctx.Done():
			return
		case val := <-res:
			if v, ok := val.(*img.Void); ok {
				out <- &img.Void{Msg: fmt.Sprintf("%v: %v", v.Msg, path)}
				return
			}
			out <- val
		}
	}()

	return out
}
