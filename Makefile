PATH := ${PWD}/bin:${PATH}
export PATH

.DEFAULT_GOAL := help

REVISION ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date +'%Y-%m-%dT%H:%M:%SZ')

GO_BUILD_FLAGS := -v
GO_TEST_FLAGS := -v -timeout 2m
GO_COVER_FLAGS := -coverprofile coverage.txt -covermode atomic
SRC_FILES := $(shell go list -f '{{range .GoFiles}}{{printf "%s/%s\n" $$.Dir .}}{{end}}' ./...)

XC_ARCH := 386 amd64
XC_OS := darwin linux windows


#  App
#----------------------------------------------------------------
BIN_DIR := ./bin
OUT_DIR := ./dist
GENERATED_BINS :=
PACKAGES :=

define cmd-tmpl

$(eval NAME := $(notdir $(1)))
$(eval OUT := $(addprefix $(BIN_DIR)/,$(NAME)))
$(eval LDFLAGS := -ldflags "-X main.revision=$(REVISION) -X main.buildDate=$(BUILD_DATE)")

$(OUT): $(SRC_FILES)
	go build $(GO_BUILD_FLAGS) $(LDFLAGS) -o $(OUT) $(1)

.PHONY: $(NAME)
$(NAME): $(OUT)

.PHONY: $(NAME)-package
$(NAME)-package: $(NAME) $(BIN_DIR)/gox
	gox \
		$(LDFLAGS) \
		-os="$(XC_OS)" \
		-arch="$(XC_ARCH)" \
		-output="$(OUT_DIR)/$(NAME)_{{.OS}}_{{.Arch}}" \
		$(1)

$(eval GENERATED_BINS += $(OUT))
$(eval PACKAGES += $(NAME)-package)

endef

$(foreach src,$(wildcard ./cmd/*),$(eval $(call cmd-tmpl,$(src))))


#  Commands
#----------------------------------------------------------------
.PHONY: setup
setup: ## clean all binaries in bin/
ifdef CI
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	dep ensure -v -vendor-only
	dep ensure -add github.com/golang/mock/mockgen
	go install ./cmd/tools
	tools --build --verbose

.PHONY: clean
clean: ## clean all binaries in bin/
	rm -rf $(BIN_DIR)/*

.PHONY: gen
gen: ## go generate all files
	go generate ./...

.PHONY: lint
lint: ## use tools reviewdog -diff="git diff master" ifdef CI tools reviewdog -reporter=github-pr-review

ifdef CI
	tools reviewdog -reporter=github-pr-review
else
	tools reviewdog -diff="git diff master"
endif

.PHONY: test
test: ## run all tests
	go test $(GO_TEST_FLAGS) ./...

.PHONY: cover
cover: ## run test coverage
	go test $(GO_TEST_FLAGS) $(GO_COVER_FLAGS) ./...

.PHONY: test-e2e
test-e2e-hack: ## run /e2e/hack/run_test.sh
	@./e2e/hack/run_test.sh

.PHONY: all ## compile all binaries to bin/
all: $(GENERATED_BINS)

.PHONY: packages
packages: $(PACKAGES)

run:
	mkdir -p contracts/$(pkg)
	sol2proto --pkg $(pkg) --abi $(name).abi > contracts/$(pkg)/$(pkg).proto
	protoc --go_out=plugins=grpc:. contracts/$(pkg)/$(pkg).proto
	abigen --type $(name) --abi $(name).abi --pkg $(pkg) --out ./contracts/$(pkg)/$(pkg).go --bin $(name).bin
	grpc-contract -type $(pkg) -path ./contracts/$(pkg) > ./contracts/$(pkg)/$(pkg)_server.go

server:
	go build -v -o ./build/bin/server ./cmd/server
	@echo "Done building."
	@echo "Run \"$(GOBIN)/server\" to launch server."

client:
	go build -v -o ./build/bin/client ./cmd/client
	@echo "Done building."
	@echo "Run \"$(GOBIN)/client\" to launch client."


build-alpine: ## Compile optimized for alpine linux.
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME}

package: ## Build final docker image with just the go binary inside
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):local .

tag: ## Tag image created by package with latest, git commit and version'
	@echo "Tagging: latest ${VERSION} $(GIT_COMMIT)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag ## Push tagged images to registry'
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort