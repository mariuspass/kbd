package kbd

// KeyEvent specifies various aspects of a keystroke
type KeyEvent uint32

// Modifiers is a bitmask representing a set of modifier keys.
type Modifiers byte

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms646270(v=vs.85).aspx
type wInput struct {
	typ uint32
	ki  wKeybdInput
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms646271(v=vs.85).aspx
type wKeybdInput struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// Code is a Virtual-Key Code
// https://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
type Code uint16
