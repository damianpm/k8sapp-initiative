FROM golang:1.13.7-alpine3.11 as stage1

COPY go.mod go.sum /codebase/
COPY pkg /codebase/pkg
RUN cd /codebase && go build -v -o /codebase/bin/server ./pkg/main.go

FROM alpine:3.11 as stage2
COPY --from=stage1 /codebase/bin/server /server

USER root
RUN apk add --update curl

ENV PORT=8080

CMD ["sh", "-c", "/server"]