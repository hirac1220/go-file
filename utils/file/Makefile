# install packages
install:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/kisielk/errcheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	brew install golangci-lint
	brew upgrade golangci-lint
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(GOPATH)/bin latest


# precommit check
precommit: goimports fmt errcheck staticcheck gosec

goimports:
	find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/hirac1220/go-file/utils"

fmt:
	go fmt ./...

vet:
	go vet ./...

errcheck:
	errcheck

staticcheck:
	staticcheck ./...

golangci:
	golangci-lint run ./...

gosec:
	gosec ./...


# unit test
test_all: test_file

test_file:
	go test -v -coverpkg=. ./...

test_cover_file:
	go test -v -coverprofile=cover_file.out  -coverpkg=./...
	go tool cover -html=cover_file.out -o cover_file.html
