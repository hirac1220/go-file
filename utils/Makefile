# precommit check
precommit: goimports fmt errcheck staticcheck gosec

goimports:
	find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/hirac1220/go-file/utils"

fmt:
	cd file && go fmt ./...

vet:
	cd file && go vet ./...

errcheck:
	cd file && errcheck -ignore 'fmt:Fprint*' ./...

staticcheck:
	cd file && staticcheck ./...

golangci:
	cd file && golangci-lint run ./...

gosec:
	cd file && gosec ./...


# unit test
test_all: test_file

test_file:
	cd file && go test -v -coverpkg=. ./...
