/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-20
 * Time: 下午1:32
 * To change this template use File | Settings | File Templates.
 */
package mp
import "fmt"

type Player interface {
	Play(source string)
}

func Play(source,memtype string) {
	var p Player

	switch  memtype {
	   case "MP3":
	       p = &MP3Player{};
	  /* case "WAV":
	       p = &WAVPlayer{};
	  */
	   default:
	       fmt.Println("Unsupported music type",memtype);
		   return
	}

	p.Play(source);

}

