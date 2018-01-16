package main

import (
  "io"
  "net"
  "net/rpc"
)

import "C"

//export Accept
func Accept(lis net.Listener) {
  rpc.Accept(lis);
}

//export HandleHTTP
func HandleHTTP() {
  rpc.HandleHTTP();
}

//export Register
func Register(rcvr interface{}) error {
  return rpc.Register(rcvr);
}

//export RegisterName
func RegisterName(name string, rcvr interface{}) error {
  return rpc.RegisterName(name, rcvr);
}

//export ServeCodec
func ServeCodec(codec rpc.ServerCodec) {
  rpc.ServeCodec(codec);
}

//export ServeConn
func ServeConn(conn io.ReadWriteCloser) {
  rpc.ServeConn(conn);
}

//export ServeRequest
func ServeRequest(codec rpc.ServerCodec) error {
  return rpc.ServeRequest(codec);
}

//export Dial
func Dial(network, address string) (*rpc.Client, error) {
  return rpc.Dial(network, address);
}

//export DialHTTPPath
func DialHTTPPath(network, address, path string) (*rpc.Client, error) {
  return rpc.DialHTTPPath(network, address, path);
}

//export NewClient
func NewClient(conn io.ReadWriteCloser) *rpc.Client {
  return rpc.NewClient(conn);
}

//export NewClientWithCodec
func NewClientWithCodec(codec rpc.ClientCodec) *rpc.Client {
  return rpc.NewClientWithCodec(codec);
}

//export NewServer
func NewServer() *rpc.Server {
  return rpc.NewServer();
}

func main() { }
