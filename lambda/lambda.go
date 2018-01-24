package lambda

import (
  "C"
  "fmt"
  "log"
  "os"
  "net/rpc"
  "net"
  "reflect"
  //"github.com/aws/aws-lambda-go/lambda"
)

// exports Start()
func Start(handler interface{}) {
  port := os.Getenv("_LAMBDA_SERVER_PORT")
  lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
  if err != nil {
    log.Fatal(err)
  }
  function := new(Function)
  //fuh := reflect.ValueOf(&function).Elem()
  //function.handler = handler
  //fuh.handler = handler
  rpc.Register(function)
  rpc.Accept(lis)
  log.Fatal("accept should not have returned")
}
