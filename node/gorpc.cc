#include "gorpc.h"
#include <node.h>

namespace gorpc {

using v8::Function;
using v8::FunctionCallbackInfo;
using v8::Isolate;
using v8::Local;
using v8::Object;
using v8::String;
using v8::Value;

void start(const FunctionCallbackInfo<Value>& args) {
  Isolate* isolate = args.GetIsolate();
  Local<Function> handler = Local<Function>::Cast(args[0]);
  Local<Function> callback = Local<Function>::Cast(args[2]);

  /* Provide a C function that is called from Golang,
     this invokes our javascript function. */
  auto wrappedHandler = [=](const char* context, const char* event) {
    Local<Value> jsContext = String::NewFromUtf8(isolate, context);
    Local<Value> jsEvent = String::NewFromUtf8(isolate, event);
    Local<Value> argv[] = {jsContext, jsEvent};

    handler->Call(Null(isolate), 2, argv);
    callback->Call(Null(isolate), 0, {});
    return;
  };

  /* We need to pass our C function to Go. Go will
     register a go-based handler holding the C function
     pointer. When the go-handler is called, it will
     invoke its own internal C function that calls the
     pointed function (Go cannot call function pointers) */
  void *ptr = &wrappedHandler;
  Start(ptr);
}

void init(Local<Object> exports) {
  NODE_SET_METHOD(exports, "start", start);
}

NODE_MODULE(NODE_GYP_MODULE_NAME, init)

}
