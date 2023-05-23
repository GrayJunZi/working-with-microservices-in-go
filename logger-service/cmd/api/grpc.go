package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
	"logger/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		resp := &logs.LogResponse{Result: "failed"}
		return resp, err
	}

	resp := &logs.LogResponse{Result: "logged!"}
	return resp, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
		return
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC Serverr started on port %s", gRrpcPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
