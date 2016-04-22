package main

import (
  "log"
  "net"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "crypto/sha512"
  "fmt"
  pb "./proto"
)

const (
  port = ":50051"
)

type server struct{}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
  var err string
  var session string
  if in.Username == "user" && in.Password == "pass" {
    sha_512 := sha512.New()
    sha_512.Write([]byte(in.Username + ":" + in.Password))
    session = fmt.Sprintf("%x", sha_512.Sum(nil))
  } else {
    err = "Authentication failed"
  }
  return &pb.LoginResponse{Error: err, Session: session}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterAuthenticationServer(s, &server{})
  s.Serve(lis)
}
