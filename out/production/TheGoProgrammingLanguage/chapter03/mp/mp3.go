/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-20
 * Time: 下午2:08
 * To change this template use File | Settings | File Templates.
 */
package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	process int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing Mp3 music",source);

	p.process = 0;

	for p.process < 100 {
		time.Sleep(100 * time.Millisecond); // pretend being playing now
		fmt.Println(".");
		p.process += 10;
	}

	fmt.Println("\nFinished playing",source);
}

