python -m grpc_tools.protoc -I. --python_out=./stubs --pyi_out=. --grpc_python_out=./stubs *.proto