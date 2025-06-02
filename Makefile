
mkdir:
	mkdir -p bin/linux_amd64 && mkdir -p bin/windows_amd64

build-linux: mkdir
	env GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/dep-comparer cmd/app/main.go

build-windows:
	env	 GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64/dep-comparer.exe cmd/app/main.go

zip: mkdir build-linux build-windows
	zip linux_amd64.zip bin/linux_amd64/dep-comparer && zip windows_amd64.zip bin/windows_amd64/dep-comparer.exe

test:
	go test ./...