FROM golang:1.10 as build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]
