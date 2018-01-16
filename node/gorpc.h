#include <v8.h>
#include <node.h>

using v8::FunctionCallbackInfo;
using v8::Isolate;
using v8::Local;
using v8::Object;
using v8::String;
using v8::Value;

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef __complex float GoComplex64;
typedef __complex double GoComplex128;
typedef struct { char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#ifdef __cplusplus
extern "C" {
#endif

//extern void Accept(v8::Local<v8::Value> lis); // __asm__ ("gorpc.main.Accept");
//extern void Accept(v8::Local<v8::Value> lis) __asm__ ("net/rpc.(*Server).Accept");
extern void Accept(GoInterface lis);
extern struct Listener Listen(GoString network, GoString, address);

#ifdef __cplusplus
}
#endif
