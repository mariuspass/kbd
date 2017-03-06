package kbd

import (
	// #include <wtypes.h>
	"C"
	"errors"
	"syscall"
	"unsafe"
)

var (
	procSendInput = syscall.NewLazyDLL("user32.dll").NewProc("SendInput")

	// KeysMap contains a map of all key code by name
	KeysMap = map[string]Code{}
)

func init() {
	for code, name := range _Code_map {
		KeysMap[name] = code
	}
}

// SendInput function send inputs events to user32.dll
func SendInput(inputs []C.INPUT) error {
	ret, _, _ := procSendInput.Call(
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		uintptr(unsafe.Sizeof(C.INPUT{})),
	)

	if ret == 0 {
		return errors.New("The input was already blocked by another thread")
	}

	return nil
}

// GetInputByCode return a C.INPUT by a Code
func GetInputByCode(code Code, evt KeyEvent) C.INPUT {
	input := C.INPUT{_type: C.DWORD(inputKeyboard)}
	(*wInput)(unsafe.Pointer(&input)).ki = wKeybdInput{
		WVk:     uint16(code),
		DwFlags: uint32(evt),
	}

	return input
}

// GetInputByRune return a C.INPUT by a rune
func GetInputByRune(r int32, evt KeyEvent) C.INPUT {
	input := C.INPUT{_type: C.DWORD(inputKeyboard)}
	(*wInput)(unsafe.Pointer(&input)).ki = wKeybdInput{
		WScan:   uint16(r),
		DwFlags: uint32(evt) | uint32(keyUnicode),
	}

	return input
}

// TapKeys press and release a key or more with or without modifiers.
// Modifers are: ModCtrl, ModShift, ModAlt and ModMeta
func TapKeys(codes []Code, mod Modifiers) error {
	pressInputs := []C.INPUT{}
	releaseInputs := []C.INPUT{}

	if mod > 0 {
		if mod&ModShift == ModShift {
			pressInputs = append(pressInputs, GetInputByCode(Shift, KeyDown))
			releaseInputs = append(releaseInputs, GetInputByCode(Shift, KeyUp))
		}

		if mod&ModControl == ModControl {
			pressInputs = append(pressInputs, GetInputByCode(Control, KeyDown))
			releaseInputs = append(releaseInputs, GetInputByCode(Control, KeyUp))
		}

		if mod&ModAlt == ModAlt {
			pressInputs = append(pressInputs, GetInputByCode(Alt, KeyDown))
			releaseInputs = append(releaseInputs, GetInputByCode(Alt, KeyUp))
		}

		if mod&ModMeta == ModMeta {
			pressInputs = append(pressInputs, GetInputByCode(LeftWin, KeyDown))
			releaseInputs = append(releaseInputs, GetInputByCode(LeftWin, KeyUp))
		}
	}

	for _, c := range codes {
		pressInputs = append(pressInputs, GetInputByCode(c, KeyDown))
		releaseInputs = append(releaseInputs, GetInputByCode(c, KeyUp))
	}

	return SendInput(append(pressInputs, releaseInputs...))
}

// ToggleKey press/release a key with or without modifiers.
// Modifers are: ModCtrl, ModShift, ModAlt and ModMeta
func ToggleKey(code Code, evt KeyEvent, mod Modifiers) error {
	if evt != KeyDown && evt != KeyUp {
		return errors.New("KeyEvent must be KeyDown or KeyUp")
	}

	inputs := []C.INPUT{}

	if mod > 0 {
		if mod&ModShift == ModShift {
			inputs = append(inputs, GetInputByCode(Shift, evt))
		}

		if mod&ModControl == ModControl {
			inputs = append(inputs, GetInputByCode(Control, evt))
		}

		if mod&ModAlt == ModAlt {
			inputs = append(inputs, GetInputByCode(Alt, evt))
		}

		if mod&ModMeta == ModMeta {
			inputs = append(inputs, GetInputByCode(LeftWin, evt))
		}
	}

	inputs = append(inputs, GetInputByCode(code, evt))

	return SendInput(inputs)
}

// TypeString send a string to foreground window.
func TypeString(str string) error {
	if len(str) == 0 {
		return errors.New("The string cannot be empty")
	}

	inputs := []C.INPUT{}

	for _, s := range str {
		inputs = append(inputs, GetInputByRune(s, KeyDown))
		inputs = append(inputs, GetInputByRune(s, KeyUp))
	}

	return SendInput(inputs)
}
