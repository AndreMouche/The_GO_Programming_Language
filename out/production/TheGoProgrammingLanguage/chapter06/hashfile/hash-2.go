/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-22
 * Time: 下午6:03
 * To change this template use File | Settings | File Templates.
 */
package main
import (
	"io"
	"fmt"
	"os"
	"crypto/md5"
	"crypto/sha1"
)

func main() {
	TestFile := "/home/fun/db.sql"
	infile,inerr := os.Open(TestFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h,infile)
		fmt.Println("%x %s\n",md5h.Sum([]byte("")),TestFile)

		sha1h := sha1.New()
		io.Copy(sha1h,infile)
		fmt.Println("%x %s\n",sha1h.Sum([]byte("")),TestFile)
	} else {
		fmt.Println(inerr);
		os.Exit(1)
	}
}

