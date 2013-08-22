/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-20
 * Time: 下午7:39
 * To change this template use File | Settings | File Templates.
 */
package ipc

import (
	"encoding/json"
	"fmt"

)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect();
	return &IpcClient{c}
}

func (client *IpcClient) Call (method,params string) (resp *Response,err error) {
	req := &Request{method,params}
	var b []byte
	    b,err = json.Marshal(req);
	fmt.Println("b:" + string(b));
	if err != nil {
		return
	}

	client.conn <- string(b)   //写入b值
	str := <-client.conn //等待返回值
	var resp1 Response
	err = json.Unmarshal([]byte(str),&resp1)
	resp = &resp1
	return
}


func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
