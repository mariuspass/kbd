// +build windows

package kbd

import (
	"errors"
	"syscall"
)

var (
	procKeybdEvent = syscall.NewLazyDLL("user32.dll").NewProc("keybd_event")

	// KeysMap contains a map of all key code by name
	KeysMap = map[string]Code{}
)

func init() {
	for code, name := range _Code_map {
		KeysMap[name] = code
	}
}

// SendInput function send inputs events to user32.dll
func SendInput(inputs []Input) (err error) {
	for _, i := range inputs {
		_, _, ret := procKeybdEvent.Call(
			0,
			uintptr(i.Code),
			uintptr(i.Evnt),
			0,
		)

		if ret.Error() != "The operation completed successfully." {
			err = ret
			break
		}
	}

	return err
}

// GetInputByCode return a Input by a Code
func GetInputByCode(code Code, evt KeyEvent) Input {
	if code > 255 {
		evt = evt | keyExtended
	}

	return Input{
		Code: code,
		Evnt: evt | keyScancode,
	}
}

// GetInputByRune return a Input by a rune
func GetInputByRune(r int32, evt KeyEvent) Input {
	return Input{
		Code: Code(r),
		Evnt: evt | keyUnicode,
	}
}

// TapKeys press and release a key or more with or without modifiers.
// Modifers are: ModCtrl, ModShift, ModAlt and ModMeta
func TapKeys(codes []Code, mod Modifiers) error {
	pressInputs := []Input{}
	releaseInputs := []Input{}

	if mod > 0 {
		if mod&ModShift == ModShift {
			pressInputs = append(pressInputs, GetInputByCode(LeftShift, KeyDown))
			releaseInputs = append(releaseInputs, GetInputByCode(LeftShift, KeyUp))
		}

		if mod&ModControl == ModControl {
			pressInputs = append(pressInputs, GetInputByCode(LeftControl, KeyDown))
			releaseInputs = append(releaseInputs, GetInputByCode(LeftControl, KeyUp))
		}

		if mod&ModAlt == ModAlt {
			pressInputs = append(pressInputs, GetInputByCode(LeftAlt, KeyDown))
			releaseInputs = append(releaseInputs, GetInputByCode(LeftAlt, KeyUp))
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

	inputs := []Input{}

	if mod > 0 {
		if mod&ModShift == ModShift {
			inputs = append(inputs, GetInputByCode(LeftShift, evt))
		}

		if mod&ModControl == ModControl {
			inputs = append(inputs, GetInputByCode(LeftControl, evt))
		}

		if mod&ModAlt == ModAlt {
			inputs = append(inputs, GetInputByCode(LeftAlt, evt))
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

	inputs := []Input{}

	for _, s := range str {
		inputs = append(inputs, GetInputByRune(s, KeyDown))
		inputs = append(inputs, GetInputByRune(s, KeyUp))
	}

	return SendInput(inputs)
}
