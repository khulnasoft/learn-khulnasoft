VER=$(shell git rev-parse --short HEAD)

build-service:
	cd service && go build -a -o ./main .

run-service-a:
	SERVICE_NAME="service-a" \
	PORT=9091 \
	./service/main

run-service-b:
	SERVICE_NAME="service-b" \
	PORT=9092 \
	./service/main

build-img-service:
	cd service && docker build -t khulnasoft/learn-khulnasoft:latest -t khulnasoft/learn-khulnasoft:$(VER) .

image-push:
	docker push khulnasoft/learn-khulnasoft
