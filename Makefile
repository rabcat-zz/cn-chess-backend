build:
	@echo "Building BE"
	@go build -o be/be_server be/main/main.go

vendor: go.mod go.sum ## pull the vendor pkgs for deps
	@GO111MODULE=on go mod vendor

run: vendor ## run the server from source code
	@go run be/main/main.go