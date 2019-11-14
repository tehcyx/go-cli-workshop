VERSION=$(shell cat VERSION)
GIT_COMMIT=$(shell git rev-list -1 HEAD)

.PHONY: build
build: 
	./before-commit.sh ci
	go build -ldflags "-X github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/cmd.Version=${VERSION} -X github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/cmd.GitCommit=${GIT_COMMIT} -s" -a -installsuffix cgo -i -o bin/workshop ./cmd/cli-workshop

.PHONY: ci-pr
ci-pr: build

.PHONY: ci-master
ci-master: build

.PHONY: ci-release
ci-release: build

.PHONY: test
test:
	go test -coverprofile=cover.out ./...
	@echo "Total test coverage: $$(go tool cover -func=cover.out | grep total | awk '{print $$3}')"
	@rm cover.out

.PHONY: clean
clean:
	rm -rf bin

.PHONY: docker
docker:
	docker build -t workshop:$(VERSION) -f build/package/Dockerfile --build-arg VERSION="${VERSION}" --build-arg GIT_COMMIT="${GIT_COMMIT}" ./
