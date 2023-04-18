build : getDep clean amd64 amd64-v2 amd64-v3

getDep:
	go get -v -t -d ./...

clean :
	rm -rf target; go clean; go clean --cache

amd64 :
	go clean; env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w " -o target/linux-amd64/vm-manager ./cmd/vm-manager

amd64-v2 :
	go clean; env GOOS=linux GOARCH=amd64 GOAMD64=v2 go build -ldflags "-s -w " -o target/linux-amd64-v2/vm-manager ./cmd/vm-manager

amd64-v3 :
	go clean; env GOOS=linux GOARCH=amd64 GOAMD64=v3 go build -ldflags "-s -w " -o target/linux-amd64-v3/vm-manager ./cmd/vm-manager
