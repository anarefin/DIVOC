PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
source := $(wildcard *.go **/*.go */*/*.go)

all: divoc-server certificate_processor analytics_feed

deps:
	@GO111MODULE=off go get -u github.com/myitcv/gobin
	@$(GOPATH)/bin/gobin github.com/go-swagger/go-swagger/cmd/swagger@v0.25.0
	@$(GOPATH)/bin/gobin github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24.0

swagger_gen/restapi/server.go: ../interfaces/vaccination-api.yaml
	@mkdir $(PWD)/swagger_gen || echo 'exists'
	@$(GOPATH)/bin/swagger generate server --exclude-main -t ./swagger_gen -f $(PWD)/../interfaces/vaccination-api.yaml --principal models.JWTClaimBody

api_docs:
	@echo "Installing swagger-merger" && npm install swagger-merger -g
	@swagger-merger -i $(PWD)/../interfaces/vaccination-api.yaml -o $(PWD)/bundle.yaml

divoc-server: $(source) swagger_gen/restapi/server.go
	@echo "Building Server to $(PWD)/divoc-server ..."
	@CGO_ENABLED=1 GO111MODULE=on go build github.com/divoc/api/cmd/divoc-server

certificate_processor: $(source)
	@echo "Building certificate processor $(PWD)/certificate_processor"
	@CGO_ENABLED=1 GO111MODULE=on go build github.com/divoc/api/cmd/certificate_processor

analytics_feed: $(source)
	@echo "Building analytics feed $(PWD)/analytics_feed"
	@CGO_ENABLED=1 GO111MODULE=on go build github.com/divoc/api/cmd/analytics_feed

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.test' | xargs rm -fv
	@rm -rf build
	@rm -rf release
	#@rm `find swagger_gen|grep -v configure_divoc`
	@rm  `grep -lr '// Code generated by go-swagger; DO NOT EDIT.' swagger_gen`

run: ./divoc-server
	@echo "Running api server"
	./divoc-server --scheme http --port 8000