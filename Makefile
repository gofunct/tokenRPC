filename:=$(shell echo '$(name)' | perl -pe 's/([A-Z])/_\L\1/g' | sed 's/^_//')

sol2proto: ## make sol2proto name=token_service pkg=token_service
	mkdir -p contracts/$(pkg)
	sol2proto --pkg $(pkg) --abi $(name).abi > contracts/$(pkg)/$(pkg).proto

abigen: ## make abigen name=token_service pkg=token_service
	abigen --type $(name) --abi $(name).abi --pkg $(pkg) --out ./contracts/$(pkg)/$(pkg).go --bin $(name).bin

contract: ## make contract name=token_service pkg=token_service
	grpc-contract -type $(pkg) -path ./contracts/$(pkg) > ./contracts/$(pkg)/$(pkg)_server.go

server:
	go build -v -o ./build/bin/server ./cmd/server

client:
	go build -v -o ./build/bin/client ./cmd/client

gen:
	cd contracts/token_service; docker run -v `pwd`:/defs colemanword/protoc-all -f token_service.proto -l go && mv gen/pb-go/* .

build: ## build docker image
	bash build.sh

push: ## push docker image
	bash push.sh

grapi:
	docker run -v `pwd`:/defs colemanword/grapi:1.17_0

proto:
	cd contracts/name_service; docker run -v `pwd`:/defs colemanword/protoc-all:1.17_0 -f name_service.proto -l go

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort