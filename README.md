# kbd

### a simple library to simulate key press in Windows

**Example:**

```golang
package main

import (
	"time"

	"github.com/mariuspass/kbd"
)

func main() {

	time.Sleep(time.Second * 2)

	// will type 'Lorem ipsum dolor sit amet.' to the foreground window
	kbd.TypeString(`Lorem ipsum dolor sit amet.`)

	time.Sleep(time.Second * 2)

	// will press Windows Key + E
	kbd.ToggleKey(kbd.KeyE, kbd.KeyDown, kbd.ModMeta)

	// will release Windows Key + E
	kbd.ToggleKey(kbd.KeyE, kbd.KeyUp, kbd.ModMeta)
	// Windows Explorer should be open

	time.Sleep(time.Second * 2)

	// will tap(press and release) Ctrl+Shift+Escape
	kbd.TapKeys([]kbd.Code{kbd.Escape}, kbd.ModShift|kbd.ModControl)
	// Task Manager should be open

}```
