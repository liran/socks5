VERSION=1.0.0


lint:
	staticcheck ./...

build: lint
	GOOS=linux GOARCH=amd64 go build -o dist/socks5 cmd/server/main.go

test:
	go test -timeout 45s ./... -race -tags all -v -cover

build-docker: build
	docker build --no-cache --platform linux/amd64 -t ranqiwu/socks5:${VERSION} -f cmd/server/Dockerfile .
	docker tag ranqiwu/socks5:${VERSION} ranqiwu/socks5:latest
	docker push ranqiwu/socks5:${VERSION}
	docker push ranqiwu/socks5:latest
	docker rmi -f ranqiwu/socks5:${VERSION}
	docker rmi -f ranqiwu/socks5:latest