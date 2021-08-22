run:
	go run ./cmd/web

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/url-parser ./cmd/web

test:
	go test -v ./...

build-image:
	docker build -t url-parser:latest .

run-container:
	docker run -d --rm -p 8080:8080 --name url-parser-container url-parser