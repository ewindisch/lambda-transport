package main
// typedef void (*voidFunc) (unsigned char *context, unsigned char *event);
// void callCfuncPointer(voidFunc ptr, unsigned char *context, unsigned char *event);
import "C"

import (
  "fmt"
  "log"
  "os"
  "net/rpc"
  "net"
  "unsafe"
  "context"
  "time"
  //"github.com/aws/aws-lambda-go/lambdacontext"
  "github.com/aws/aws-lambda-go/lambda/messages"
)

func makeHandler(cFunction unsafe.Pointer) func(ctx context.Context, payload []byte) {
  return func(ctx context.Context, payload []byte) {
    C.callCfuncPointer(cFunction, ctx, payload);
  }
}

type Function struct {
	/* C function pointer */
	handler func(ctx context.Context, payload []byte)
  //unsafe.Pointer
}

func (fn *Function) Invoke(req *messages.InvokeRequest, response *messages.InvokeResponse) error {
	/*defer func() {
		if err := recover(); err != nil {
			panicInfo := getPanicInfo(err)
			response.Error = &messages.InvokeResponse_Error{
				Message:    panicInfo.Message,
				Type:       getErrorType(err),
				StackTrace: panicInfo.StackTrace,
				ShouldExit: true,
			}
		}
	}()*/

	deadline := time.Unix(req.Deadline.Seconds, req.Deadline.Nanos).UTC()
	invokeContext, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	payload, err := fn.handler(req.ClientContext, req.Payload)
	if err != nil {
		response.Error = &messages.InvokeResponse_Error{
      Message: err.Error(),
      Type:    "GoError",
    }
    return nil
	}
	response.Payload = payload
	return nil
}

//export Start
func Start(userHandler unsafe.Pointer) {
  port := os.Getenv("_LAMBDA_SERVER_PORT")
  lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
  if err != nil {
    log.Fatal(err)
  }

  /*
  function := new(lambda.Function)
  function.handler = handler
  rs := reflect.ValueOf(&function).Elem()
  rf := rs.Field(0)
  rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
  ri := reflect.ValueOf(&handler).Elem()
  rf.Set(ri)
  */

  function := new(Function)
  function.handler = makeHandler(userHandler)

  rpc.Register(function)
  rpc.Accept(lis)
  log.Fatal("accept should not have returned")
}

func main() { }
