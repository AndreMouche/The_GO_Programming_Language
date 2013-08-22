/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-22
 * Time: 下午1:47
 * To change this template use File | Settings | File Templates.
 */
package main
import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
	"chapter05/util"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage: %s host:port\n",os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
	util.CheckError(err)

	conn, err := net.DialTCP("tcp",nil,tcpAddr)
	util.CheckError(err)

	_,err = conn.Write([]byte("HEAD / HTTP/1.0 \r\n\r\n"))
    util.CheckError(err)

	result, err := ioutil.ReadAll(conn)
	util.CheckError(err)

	fmt.Println(string(result))
	os.Exit(0)

}
