#include "gorpc.h"
#include <node.h>

namespace gorpc {

using v8::FunctionCallbackInfo;
using v8::Isolate;
using v8::Local;
using v8::Object;
using v8::String;
using v8::Value;

void rpcAccept(const FunctionCallbackInfo<Value>& args) {
  Accept(args[0]);
}

void init(Local<Object> exports) {
  NODE_SET_METHOD(exports, "accept", rpcAccept);
}

NODE_MODULE(NODE_GYP_MODULE_NAME, init)

}
