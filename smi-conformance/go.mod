module github.com/khulnasoft/learn-khulnasoft/smi-conformance

go 1.13

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/khulnasoft/service-mesh-performance v0.0.0-20240115121425-f0a0a76e1ed5
	github.com/kudobuilder/kuttl v0.0.0-00010101000000-000000000000
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/net v0.23.0 // indirect
	google.golang.org/grpc v1.56.3
	google.golang.org/protobuf v1.33.0
	k8s.io/api v0.17.16
	k8s.io/client-go v0.17.16
	sigs.k8s.io/controller-runtime v0.5.1
)

replace github.com/kudobuilder/kuttl => github.com/khulnasoft/kuttl v0.4.1-0.20200806180306-b7e46afd657f
