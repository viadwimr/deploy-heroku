lint:
	go fmt ./...

test:
	go test -v ./...

build:
	go build -v .

run:
	./deploy-heroku