// +build windows

package kbd

// KeyEvent specifies various aspects of a keystroke
type KeyEvent uint32

// Modifiers is a bitmask representing a set of modifier keys.
type Modifiers byte

// Input struct
type Input struct {
	Code Code
	Evnt KeyEvent
}

// Code is a keyboard scancode
// https://en.wikipedia.org/wiki/Scancode
type Code uint16
