#! /bin/sh

protoc --proto_path=. \
       --micro_out=. \
       --micro_opt=paths=source_relative \
       --go_out=. \
       --go_opt=paths=source_relative \
       ./GreeterServer/api/*.proto