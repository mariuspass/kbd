// +build windows

//go:generate stringer -type=Code

package kbd

// Key scancode
// http://www.win.tue.nl/~aeb/linux/kbd/scancodes-10.html#correspondence
const (
	Tilde           Code = 0x29   // ` ~
	Key1            Code = 0x02   // 1 !
	Key2            Code = 0x03   // 2 @
	Key3            Code = 0x04   // 3 #
	Key4            Code = 0x05   // 4 $
	Key5            Code = 0x06   // 5 %
	Key6            Code = 0x07   // 6 ^
	Key7            Code = 0x08   // 7 &
	Key8            Code = 0x09   // 8 *
	Key9            Code = 0x0a   // 9 (
	Key0            Code = 0x0b   // 0 )
	Minus           Code = 0x0c   // - _
	Plus            Code = 0x0d   // = +
	Backspace       Code = 0x0e   // Backspace
	Tab             Code = 0x0f   // Tab
	KeyQ            Code = 0x10   // Q
	KeyW            Code = 0x11   // W
	KeyE            Code = 0x12   // E
	KeyR            Code = 0x13   // R
	KeyT            Code = 0x14   // T
	KeyY            Code = 0x15   // Y
	KeyU            Code = 0x16   // U
	KeyI            Code = 0x17   // I
	KeyO            Code = 0x18   // O
	KeyP            Code = 0x19   // P
	OpenBracket     Code = 0x1a   // [ {
	CloseBracket    Code = 0x1b   // ] }
	Backslash       Code = 0x2b   // \ |
	CapsLock        Code = 0x3a   // CapsLock
	KeyA            Code = 0x1e   // A
	KeyS            Code = 0x1f   // S
	KeyD            Code = 0x20   // D
	KeyF            Code = 0x21   // F
	KeyG            Code = 0x22   // G
	KeyH            Code = 0x23   // H
	KeyJ            Code = 0x24   // J
	KeyK            Code = 0x25   // K
	KeyL            Code = 0x26   // L
	Colon           Code = 0x27   // ; :
	Quote           Code = 0x28   // ' "
	Return          Code = 0x1c   // Enter
	LeftShift       Code = 0x2a   // LShift
	KeyZ            Code = 0x2c   // Z
	KeyX            Code = 0x2d   // X
	KeyC            Code = 0x2e   // C
	KeyV            Code = 0x2f   // V
	KeyB            Code = 0x30   // B
	KeyN            Code = 0x31   // N
	KeyM            Code = 0x32   // M
	Comma           Code = 0x33   // , <
	Period          Code = 0x34   // . >
	Slash           Code = 0x35   // / ?
	RightShift      Code = 0x36   // RShift
	LeftControl     Code = 0x1d   // LCtrl
	LeftAlt         Code = 0x38   // LAlt
	Space           Code = 0x39   // space
	RightAlt        Code = 0xe038 // RAlt
	RightControl    Code = 0xe01d // RCtrl
	Insert          Code = 0xe052 // Insert
	Delete          Code = 0xe053 // Delete
	Home            Code = 0xe047 // Home
	End             Code = 0xe04f // End
	PageUp          Code = 0xe049 // PgUp
	PageDown        Code = 0xe051 // PgDn
	Left            Code = 0xe04b // Left
	Up              Code = 0xe048 // Up
	Down            Code = 0xe050 // Down
	Right           Code = 0xe04d // Right
	NumLock         Code = 0x45   // NumLock
	Numpad7         Code = 0x47   // KP-7 / Home
	Numpad4         Code = 0x4b   // KP-4 / Left
	Numpad1         Code = 0x4f   // KP-1 / End
	NumpadDivide    Code = 0xe035 // KP-/
	Numpad8         Code = 0x48   // KP-8 / Up
	Numpad5         Code = 0x4c   // KP-5
	Numpad2         Code = 0x50   // KP-2 / Down
	Numpad0         Code = 0x52   // KP-0 / Ins
	NumpadMultiply  Code = 0x37   // KP-*
	Numpad9         Code = 0x49   // KP-9 / PgUp
	Numpad6         Code = 0x4d   // KP-6 / Right
	Numpad3         Code = 0x51   // KP-3 / PgDn
	NumpadDecimal   Code = 0x53   // KP-. / Del
	NumpadSubstract Code = 0x4a   // KP--
	NumpadAdd       Code = 0x4e   // KP-+
	NumpadReturn    Code = 0xe01c // KP-Enter
	Escape          Code = 0x01   // Esc
	F1              Code = 0x3b   // F1
	F2              Code = 0x3c   // F2
	F3              Code = 0x3d   // F3
	F4              Code = 0x3e   // F4
	F5              Code = 0x3f   // F5
	F6              Code = 0x40   // F6
	F7              Code = 0x41   // F7
	F8              Code = 0x42   // F8
	F9              Code = 0x43   // F9
	F10             Code = 0x44   // F10
	F11             Code = 0x57   // F11
	F12             Code = 0x58   // F12
	F13             Code = 0x64   // F13
	F14             Code = 0x65   // F14
	F15             Code = 0x66   // F15
	F16             Code = 0x67   // F16
	F17             Code = 0x68   // F17
	F18             Code = 0x69   // F18
	F19             Code = 0x6A   // F19
	F20             Code = 0x6B   // F20
	F21             Code = 0x6C   // F21
	F22             Code = 0x6D   // F22
	F23             Code = 0x6E   // F23
	F24             Code = 0x76   // F24
	PrintScreen     Code = 0xe037 // PrtScr
	AltSysRq        Code = 0x54   // Alt+SysRq
	Scroll          Code = 0x46   // ScrollLock
	CtrlBreak       Code = 0xe046 // Ctrl+Break
	LeftWin         Code = 0xe05b // LWin (USB: LGUI)
	RightWin        Code = 0xe05c // RWin (USB: RGUI)
	Apps            Code = 0xe05d // Menu
	Sleep           Code = 0xe05f // Sleep
	Power           Code = 0xe05e // Power
	Wake            Code = 0xe063 // Wake

	// Pause           Code = 0xe11d45 // Pause
)
