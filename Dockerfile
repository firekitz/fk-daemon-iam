FROM golang:1.18.2 as builder

COPY . /go/fk-daemon-iam/

WORKDIR /go/fk-daemon-iam

RUN CGO_ENABLED=0 go build ./cmd/fk-daemon-iam/main.go

FROM alpine:3.14
ARG mode
#RUN echo $mode
ENV GO_ENV $mode

RUN addgroup -S golang && adduser -S golang -G golang
RUN mkdir -p /usr/src/golang-app && chown -R golang:golang /usr/src/golang-app
USER golang

WORKDIR /usr/src/golang-app
COPY --chown=golang:golang --from=builder /go/fk-daemon-iam/main .

EXPOSE 9090
EXPOSE 8080
CMD ["./main"]
