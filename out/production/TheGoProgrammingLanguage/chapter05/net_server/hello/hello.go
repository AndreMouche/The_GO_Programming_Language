/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-22
 * Time: 下午2:14
 * To change this template use File | Settings | File Templates.
 */
package main
import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"Hello,World")
}

func main() {
	http.HandleFunc("/hello",helloHandler)
	err := http.ListenAndServe(":8089",nil)
	if err != nil {
		log.Fatal("Listen and Serve:",err.Error())
	}
}

