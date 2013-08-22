/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-20
 * Time: 下午7:57
 * To change this template use File | Settings | File Templates.
 */
package ipc

import (
	"testing"
    "fmt"
)

type EchoServer struct {

}

func (server *EchoServer) Handle (method,params string) *Response {
	return &Response{"OK","ECHO:" + method + " ~ " + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{});
	client1 := NewIpcClient(server);
	client2 := NewIpcClient(server);

	resp1,_ := client1.Call("foo","From Client1");
	resp2,_ := client2.Call("foo","From Client2");

	fmt.Println("resp1",resp1.Body);
	fmt.Println("resp2",resp2.Body);

	client1.Close()
	client2.Close()

}

