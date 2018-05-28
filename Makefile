start-server:
	docker-compose up -d
compile-client:
	go get -u github.com/gobuffalo/packr/...
	packr install ./cmd/imgupclient/