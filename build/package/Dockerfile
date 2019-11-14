FROM golang:1.13.4

ARG VERSION
ARG GIT_COMMIT

WORKDIR /go/src/github.com/tehcyx/go-cli-workshop

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/cmd.Version=${VERSION} -X github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/cmd.GitCommit=${GIT_COMMIT} -s" -a -installsuffix cgo -i -o bin/app ./cmd/cli-workshop

FROM scratch

ARG VERSION
ARG GIT_COMMIT

COPY --from=0 /go/src/github.com/tehcyx/go-cli-workshop/bin/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]