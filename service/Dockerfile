FROM golang:1.13.7 as bd
WORKDIR /github.com/khulnasoft/sample-app-service
ADD . .
RUN go build -a -o ./main .

FROM golang:1.13.7
COPY --from=bd /github.com/khulnasoft/sample-app-service/main /home/main
WORKDIR /home/
EXPOSE 9091
CMD ["./main"]