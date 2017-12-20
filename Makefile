
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o build/rapyd-darwin-amd64 rapyd.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/rapyd-linux-amd64 rapyd.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o build/rapyd-windows-amd64 rapyd.go

build: build-macos build-linux build-windows

clean:
	rm -rf build
