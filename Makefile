.PHONY: prototool gwy
filename:=$(shell echo '$(name)' | perl -pe 's/([A-Z])/_\L\1/g' | sed 's/^_//')
SOURCES :=	$(shell find . -name "*.proto" -not -path ./vendor/\*)
TARGETS_TMPL :=	$(foreach source, $(SOURCES), $(source)_tmpl)


init: deps prototool $(TARGETS_TMPL) ## bootstrap project
	go mod vendor
	go install
	grpclab serve

deps: ## download dependencies and tls certificates
	go mod vendor
	git clone https://github.com/googleapis/googleapis && mv googleapis
	brew install protobuf
	brew install prototool
	go get -u \
		google.golang.org/grpc \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/ckaznocha/protoc-gen-lint \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
		moul.io/protoc-gen-gotemplate \
		github.com/gogo/protobuf/...

sol2proto: ## make sol2proto name=token_service pkg=token_service
	mkdir -p contracts/$(pkg)
	sol2proto --pkg $(pkg) --abi $(name).abi > contracts/$(pkg)/$(pkg).proto

abigen: ## make abigen name=token_service pkg=token_service
	abigen --type $(name) --abi $(name).abi --pkg $(pkg) --out ./contracts/$(pkg)/$(pkg).go --bin $(name).bin

contract: ## make contract name=token_service pkg=token_service
	grpc-contract --types $(pkg) --path ./contracts/$(pkg) --pb-path ./contracts/$(pkg) > ./contracts/$(pkg)/$(pkg)_server.go

gen:
	cd contracts/token_service; docker run -v `pwd`:/defs colemanword/gen-grpc-gateway:1.17_0  -f token_service.proto -s Token

build: ## build docker image
	bash build.sh

push: ## push docker image
	bash push.sh


prototool:
	prototool all contracts
	go mod vendor

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort