package main

import (
  "C"
  "fmt"
  "log"
  "os"
  "net/rpc"
  "net"
  "reflect"
  "unsafe"
  "github.com/aws/aws-lambda-go/lambda"
)

//export Start
func Start(handler interface{}) {
  port := os.Getenv("_LAMBDA_SERVER_PORT")
  lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
  if err != nil {
    log.Fatal(err)
  }
  function := new(lambda.Function)
  //function.handler = handler

  //fuh := reflect.ValueOf(&function).Elem()
  rs := reflect.ValueOf(&function).Elem()
  rf := rs.Field(0)
  rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()

  /* handler must become a Go function... */
  ri := reflect.ValueOf(&handler).Elem()

  //rf := rs.Field(n) 
  rf.Set(ri)

  rpc.Register(function)
  rpc.Accept(lis)
  log.Fatal("accept should not have returned")
}
