generate:
	oapi-codegen -generate "types" -package oapi ./doc/reference/lab.yaml > ./pkg/oapi/types.gen.go
	oapi-codegen -generate "chi-server" -package oapi ./doc/reference/lab.yaml > ./pkg/oapi/server.gen.go
	oapi-codegen -generate "spec" -package oapi ./doc/reference/lab.yaml > ./pkg/oapi/spec.gen.go