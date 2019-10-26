// github.com/tinygo-org/tinygo version 0.10.0-dev linux/amd64 (using go version go1.13.3)
//
// tinygo flash -target=arduino -port /dev/ttyUSB0 ./main.go
// screen /dev/ttyUSB0 115200
// cat /dev/ttyUSB0 | od -tx1z -Ax
package main

import (
	"time"

	"machine"

	"github.com/takjn/Keyboard/tinygo/rn42hid"
	"github.com/takjn/Keyboard/tinygo/rn42hid/key"
)

var (
	hid rn42hid.Device

	// pin assignment
	row = [6]machine.Pin{8, 9, 10, 11, 12, 13}
	col = [12]machine.Pin{19, 18, 17, 16, 15, 14, 7, 6, 5, 4, 3, 2}

	// key mapping
	keyMap = [2][6][12]byte{
		{
			{key.NONE, key.NONE, key.APOSTROPHE, key.LEFTBRACE, key.MINUS, key.GRAVE, key.BACKSLASH, key.EQUAL, key.RIGHTBRACE, key.SLASH, key.NONE, key.NONE},
			{key.ESC, key.K1, key.K2, key.K3, key.K4, key.K5, key.K6, key.K7, key.K8, key.K9, key.K0, key.DELETE},
			{key.TAB, key.Q, key.W, key.E, key.R, key.T, key.Y, key.U, key.I, key.O, key.P, key.BACKSPACE},
			{key.LEFTCTRL, key.A, key.S, key.D, key.F, key.G, key.H, key.J, key.K, key.L, key.SEMICOLON, key.ENTER},
			{key.LEFTSHIFT, key.Z, key.X, key.C, key.V, key.B, key.N, key.M, key.COMMA, key.DOT, key.UP, key.RIGHTSHIFT},
			{key.LEFTCTRL, key.SYSRQ, key.LEFTMETA, key.LEFTALT, key.NONE, key.SPACE, key.SPACE, key.NONE, key.COMPOSE, key.LEFT, key.DOWN, key.RIGHT},
		},
		{
			{key.NONE, key.NONE, key.APOSTROPHE, key.LEFTBRACE, key.MINUS, key.GRAVE, key.BACKSLASH, key.EQUAL, key.RIGHTBRACE, key.SLASH, key.NONE, key.NONE},
			{key.ESC, key.F1, key.F2, key.F3, key.F4, key.F5, key.F6, key.F7, key.F8, key.F9, key.F10, key.DELETE},
			{key.TAB, key.Q, key.W, key.E, key.R, key.T, key.Y, key.U, key.I, key.F11, key.F12, key.BACKSPACE},
			{key.LEFTCTRL, key.A, key.S, key.D, key.F, key.G, key.H, key.J, key.K, key.L, key.SEMICOLON, key.ENTER},
			{key.LEFTSHIFT, key.Z, key.X, key.C, key.V, key.B, key.N, key.M, key.COMMA, key.DOT, key.PAGEUP, key.RIGHTSHIFT},
			{key.LEFTCTRL, key.SYSRQ, key.LEFTMETA, key.LEFTALT, key.NONE, key.SPACE, key.SPACE, key.NONE, key.COMPOSE, key.HOME, key.PAGEDOWN, key.END},
		},
	}

	// layer holds the current key map layer (0 or 1)
	layer = 0
)

func init() {
	// Set up RN42
	config := machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	}
	machine.UART0.Configure(config)
	hid = rn42hid.New(machine.UART0)

	// Set up pins
	for _, r := range row {
		r.Configure(machine.PinConfig{Mode: machine.PinOutput})
		r.High()
	}
	for _, c := range col {
		c.Configure(machine.PinConfig{Mode: machine.PinInput})
		c.High() // HACK: InputPullUp
	}
}

// readKeyMatrix gets 6 key scan codes
func readKeyMatrix() ([6]byte, byte) {
	scanCodes := [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	var idx int
	var modifier byte
	var layerSwitch = false

	// Read key matrix
	for i, r := range row {
		r.Low()
		for j, c := range col {
			// A pin is pulled up (set high) by default.
			// It should be false (low) when the key is pressed.
			if c.Get() == true {
				continue
			}

			// Get the scan code
			scanCode := keyMap[layer][i][j]

			// Check layer key is pressed
			if scanCode == key.NONE {
				layerSwitch = true
			}

			// Check modifier key is pressed
			if scanCode >= 0xe0 && scanCode <= 0xe7 {
				shift := scanCode & 0b111
				mask := byte(1 << shift)
				modifier = modifier | mask
			}

			// Set the scan code
			if idx < len(scanCodes) {
				scanCodes[idx] = scanCode
				idx++
			} else {
				for i := 0; i < len(scanCodes); i++ {
					scanCodes[idx] = key.ERROVF
				}
			}
		}
		r.High()
	}

	// Set layer
	if layerSwitch {
		layer = 1
	} else {
		layer = 0
	}

	return scanCodes, modifier
}

func main() {
	// lastScanCodes hold the last scan codes to detect key press and key release
	lastScanCodes := [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	for {
		// Get scan codes - RN42 has 6 slots in the HID report
		scanCodes, modifier := readKeyMatrix()

		// Send scan codes if the scan codes is different from the last data.
		if lastScanCodes != scanCodes {
			if err := hid.SendKeyboardReport(scanCodes[:], modifier); err != nil {
				println(err.Error())
			}
		}
		lastScanCodes = scanCodes

		// wait to avoid key chattering
		time.Sleep(time.Millisecond * 5)
	}
}
