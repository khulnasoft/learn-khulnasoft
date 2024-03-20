FROM golang:1.14-alpine3.11 as build-img

RUN apk update && apk add --no-cache git libc-dev gcc pkgconf && mkdir /home/meshplay
COPY ${PWD} /go/src/github.com/khulnasoft/learn-khulnasoft/smi-conformance/
WORKDIR /go/src/github.com/khulnasoft/learn-khulnasoft/smi-conformance/
# RUN git rev-parse HEAD > /home/meshplay/version
# RUN git describe --tags `git rev-list --tags --max-count=1` >> /home/com/version
RUN go mod vendor && go build -a -ldflags "-s -w" -o /home/meshplay/smi_conformance main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates && mkdir /home/test-yamls && mkdir /home/test-yamls/traffic-access && mkdir /home/test-yamls/traffic-spec && mkdir /home/test-yamls/traffic-split
COPY --from=build-img /home/meshplay/** /home/
COPY --from=build-img /go/src/github.com/khulnasoft/learn-khulnasoft/smi-conformance/test-gen/test-yamls/traffic-access/** /home/test-yamls/traffic-access/
COPY --from=build-img /go/src/github.com/khulnasoft/learn-khulnasoft/smi-conformance/test-gen/test-yamls/traffic-split/** /home/test-yamls/traffic-split/
COPY --from=build-img /go/src/github.com/khulnasoft/learn-khulnasoft/smi-conformance/test-gen/test-yamls/traffic-spec/** /home/test-yamls/traffic-spec/
WORKDIR /home/
EXPOSE 10008
CMD ["sh","-c","./smi_conformance"]