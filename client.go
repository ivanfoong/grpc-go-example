package main

import (
  "log"
  "os"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "./proto"
)

const (
  address = "localhost:50051"
  defaultUsername = "user"
  defaultPassword = "pass"
)

func main() {
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect:: %v", err)
  }
  defer conn.Close()
  c := pb.NewAuthenticationClient(conn)
  
  username := defaultUsername
  password := defaultPassword
  if len(os.Args) > 2 {
    username = os.Args[1]
    password = os.Args[2]
  }
  r, err := c.Login(context.Background(), &pb.LoginRequest{Username: username, Password: password})
  if err != nil {
    log.Fatalf("could not login: %v", err)
  }
  if r.Error != "" {
    log.Panicf("could not login: %s", r.Error)
  } else {
    log.Printf("Session: %s", r.Session)
  }
}
