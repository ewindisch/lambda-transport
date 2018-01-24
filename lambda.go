package main
// typedef void (*voidFunc) (char *context, char *event);
// void callCfuncPointer(voidFunc ptr, char *context, char *event);
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
  "github.com/aws/aws-lambda-go/lambda/messages"
)

func makeHandler(cFunction unsafe.Pointer) func(ctx, payload []byte) (result []byte, err error) {
  return func(ctx, payload []byte) (result []byte, err error) {
    C.callCfuncPointer((*[0]byte)(cFunction), C.CString(string(ctx[:])), C.CString(string(payload[:])));
    result = []byte("Here is a result....")
    err = nil
    return
  }
}

type Function struct {
	/* C function pointer */
	handler func(ctx, payload []byte) (result []byte, err error)
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
	_, cancel := context.WithDeadline(context.Background(), deadline)
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

  function := new(Function)
  function.handler = makeHandler(userHandler)

  rpc.Register(function)
  rpc.Accept(lis)
  log.Fatal("accept should not have returned")
}

func main() { }
