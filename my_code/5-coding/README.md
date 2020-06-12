# Extra: Golang code generator

## Installation of Go gRPC plugin

~~~bash
# Install the module
$ export GO111MODULE=on # Enable module-aware mode
$ go get google.golang.org/grpc@v1.28.1

# Install the protobuf CLI
$ brew install protobuf
$ protoc --version
libprotoc 3.11.4

# Install the protoc plugin for go
$ go get github.com/golang/protobuf/protoc-gen-go
~~~

## Add mandatory go_option

Add the following option to the existing `proto` files; replace the final package name for the module.

~~~go
option go_package = "rdc.pt/golang/example_simple";
~~~

## Generate code

~~~bash
# Generate all
$ protoc -I=proto/ --go_out=plugins=grpc:golang proto/*.proto
~~~

...this will generate under `golang/` the following directory structure:

~~~txt
.
└── rdc.pt
    └── golang
        ├── example_complex
        │   └── complex.pb.go
        ├── example_enumerations
        │   └── enum_example.pb.go
        └── example_simple
            └── simple.pb.go
~~~
