/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-22
 * Time: 上午11:08
 * To change this template use File | Settings | File Templates.
 */
package main
import (
	"net"
	"os"
	"io"
	"bytes"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage:%s host:port\n",os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	conn, err := net.Dial("tcp",service)
	checkError(err)

	x,err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	x += 1;
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

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
