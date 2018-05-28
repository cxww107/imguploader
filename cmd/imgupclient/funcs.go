package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"sync"

	"github.com/gobuffalo/packr"

	img "github.com/cxww107/imguploader/imgupload/image"

	"google.golang.org/grpc/credentials"
)

const numberWorkers = 10

type imgWorkerFunc func(context.Context, interface{}, img.FileHandlerClient) chan interface{}

func merge(ctx context.Context, workers []chan interface{}) chan interface{} {
	out := make(chan interface{}, len(workers))

	var wg sync.WaitGroup

	output := func(worker chan interface{}) {
		defer wg.Done()
		// until worker is working
		for res := range worker {
			select {
			case <-ctx.Done():
				return
			case out <- res:
				// got result
			}
		}
	}

	wg.Add(len(workers))
	for i := range workers {
		worker := workers[i]
		go output(worker)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func dispatchTasks(ctx context.Context, tasks []interface{}, worker imgWorkerFunc, client img.FileHandlerClient) chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		var wa []interface{}

		if len(tasks) >= numberWorkers {
			wa = tasks[:numberWorkers]
		} else {
			wa = tasks[:len(tasks)]
		}

		for {
			workers := make([]chan interface{}, 0)
			for _, id := range wa {
				worker := worker(ctx, id, client)

				workers = append(workers, worker)
			}

			res := merge(ctx, workers)
			for r := range res {
				out <- r
			}

			if cap(wa) <= len(wa) {
				break
			}

			wa = wa[len(wa):cap(wa)]
			if len(wa) >= numberWorkers {
				wa = wa[:numberWorkers]
			} else {
				wa = wa[:len(wa)]
			}

		}
	}()

	return out
}

func getCreds(addr string) (*credentials.TransportCredentials, error) {

	box := packr.NewBox("./certs")

	certPEM, err := box.MustBytes("client.cert.pem")
	if err != nil {
		return nil, err
	}

	keyPEM, err := box.MustBytes("client.key.pem")
	if err != nil {
		return nil, err
	}

	pair, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	ca, err := box.MustBytes("ca.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, fmt.Errorf("failed append a ca cert")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   addr,
		Certificates: []tls.Certificate{pair},
		RootCAs:      certPool,
	})

	return &creds, nil
}
