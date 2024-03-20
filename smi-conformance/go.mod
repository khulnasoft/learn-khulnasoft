module github.com/khulnasoft/learn-khulnasoft/smi-conformance

go 1.13

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/khulnasoft/service-mesh-performance v0.0.0-20240115121425-f0a0a76e1ed5
	github.com/kudobuilder/kuttl v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/grpc v1.56.3
	google.golang.org/protobuf v1.30.0
	k8s.io/api v0.28.3
	k8s.io/client-go v0.28.3
	sigs.k8s.io/controller-runtime v0.16.3
)

replace github.com/kudobuilder/kuttl => github.com/khulnasoft/kuttl v0.0.0-20231129091724-5ec7d24cd2ca
