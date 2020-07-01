package protos

//go:generate protoc -I ./ ./raft.proto --go_out=plugins=grpc:.
