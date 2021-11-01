FROM golang:1.17

WORKDIR /go/src/gitlab.com/tokend/subgroup/project

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/project gitlab.com/tokend/subgroup/project


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/project /usr/local/bin/project
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["project"]
