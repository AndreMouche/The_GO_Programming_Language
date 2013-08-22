/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-22
 * Time: 下午6:16
 * To change this template use File | Settings | File Templates.
 */
package main
import (
	"fmt"
	"net/http"
	"log"
)

const SERVER_PORT = 8089
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "Hello"

func rootHandler(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Content-Type","test/html")
	w.Header().Set("Content-Length",fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}
func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/",SERVER_DOMAIN,SERVER_PORT),rootHandler)
	err:=http.ListenAndServeTLS(fmt.Sprintf(":%d",SERVER_PORT),"rui.crt","rui.key",nil)
	if err != nil {
		log.Fatal("Listen and Serve:",err.Error())

	}
}

