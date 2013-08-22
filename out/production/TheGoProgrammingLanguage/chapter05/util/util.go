/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-22
 * Time: 下午1:51
 * To change this template use File | Settings | File Templates.
 */
package util

import (
	"os"
	"fmt"
	"io"
	"bytes"
	"net"
)

func readFully(conn net.Conn) ([]byte,error){
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512] byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	return result.Bytes(), nil
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
}
