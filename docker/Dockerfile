FROM golang:1.15-alpine AS build

WORKDIR $GOPATH/src/github.com/JokeTrue/my-arts

COPY . ./
RUN go mod download
RUN go build -o /go/bin/my-arts ./cmd/main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/my-arts /bin/

ENTRYPOINT ["my-arts"]