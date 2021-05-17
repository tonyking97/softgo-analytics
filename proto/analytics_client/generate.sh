#!/usr/bin/env bash
NAME="*.proto"
#FOLDERNAME="python"
#mkdir -p $FOLDERNAME
protoc $NAME --go_out=. --go-grpc_out=.
#python -m grpc_tools.protoc -I. --python_out=./$FOLDERNAME/. --grpc_python_out=./$FOLDERNAME/. $NAME

#for Google's implementation
#protoc -I=. $NAME --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
