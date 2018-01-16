//include "gorpc.h"
#include "../go-build/_cgo_export.h"
#include <node.h>

namespace gorpc {

using v8::Function;
using v8::FunctionCallbackInfo;
using v8::Isolate;
using v8::Local;
using v8::Object;
using v8::String;
using v8::Value;

/* Go should pass us back the wrappedHandler to trigger the invocation in C */
/*void invoke(void *handler, GoInterface &context, GoInterface &event) {
  // Convert the go types to Javascript types.
  handler(context, event);
}*/

void start(const FunctionCallbackInfo<Value>& args) {
  Isolate* isolate = args.GetIsolate();
  Local<Function> handler = Local<Function>::Cast(args[0]);

  /* Create an anonymous function in C++,
     we can pass this into Go, which will need to pass it
     back to C.invoke in order to invoke the code. */
  auto wrappedHandler = [](GoInterface context, GoInterface event) {
    const unsigned argc = 2;
    Isolate* isolate = args.GetIsolate();
    Local<Object> context = Local<Object>::Cast(args[0]);
    Local<Object> event = Local<Object>::Cast(args[1]);
    handler->Call(Null(isolate), argc, argv);
  }

  Start(wrappedHandler);
}

void init(Local<Object> exports) {
  NODE_SET_METHOD(exports, "start", start);
}

NODE_MODULE(NODE_GYP_MODULE_NAME, init)

}
