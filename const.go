package kbd

// Key events
const (
	KeyDown     KeyEvent = 0
	keyExtended KeyEvent = 1
	KeyUp       KeyEvent = 2
	keyUnicode  KeyEvent = 4
	keyScancode KeyEvent = 8
)

// Modifier keys
const (
	ModNone    Modifiers = 0
	ModShift   Modifiers = 1 << 0
	ModControl Modifiers = 1 << 1
	ModAlt     Modifiers = 1 << 2
	ModMeta    Modifiers = 1 << 3 // Windows key
)

const (
	inputKeyboard = 1
)
