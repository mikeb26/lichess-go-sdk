export GO111MODULE=on
export GOFLAGS=-mod=vendor

OPENAPIGENGO=java -jar ./openapi-generator-cli.jar generate -g go -o ./openapi --git-host github.com --git-user-id mikeb26 --git-repo-id lichess-go-sdk

.PHONY: build
build: vendor openapi example FORCE

vendor: go.mod openapi
	go mod download
	go mod vendor

.PHONY: deps
deps: openapi
	rm -rf go.mod go.sum vendor
	go mod init github.com/mikeb26/lichess-go-sdk
	GOPROXY=direct go mod tidy

openapi-generator-cli.jar:
	curl -L https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.2.0/openapi-generator-cli-6.2.0.jar --output openapi-generator-cli.jar

# openapi.json can be obtained by visiting https://lichess.org/api and clicking 'Download'
openapi: openapi-generator-cli.jar openapi.json
	$(OPENAPIGENGO) -i ./openapi.json
	rm openapi/go.mod

example: openapi FORCE
	go build github.com/mikeb26/lichess-go-sdk/cmd/example

.PHONY: clean
clean:
	rm -f ./example

FORCE:
