// github.com/takjn/tinygo version 0.10.0-dev linux/amd64 (using go version go1.13.3)
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
	device rn42hid.Device

	row = []machine.Pin{14, 15, 16, 17, 18, 19}
	col = []machine.Pin{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

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
)

func init() {
	// Set up RN42
	config := machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	}
	machine.UART0.Configure(config)
	device = rn42hid.New(machine.UART0)

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

func main() {
	// lastScanCodes hold the last scan codes to detect key press and key release
	lastScanCodes := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	// layer holds the current key map layer (0 or 1)
	layer := 0

	for {
		// Scan codes - RN42 has 6 slots in the HID report
		scanCodes := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		var idx int
		var modifier byte
		var layerSwitch = false

		// Read key matrix
		for i, r := range row {
			r.Low()
			for j, c := range col {
				scanCode := keyMap[layer][i][j]
				if c.Get() == false {
					// Check layer key
					if scanCode == key.NONE {
						layerSwitch = true
					}

					// Check modifier key
					if scanCode >= 0xe0 && scanCode <= 0xe7 {
						shift := 0b111
						mask := byte(1 << shift)
						modifier = modifier | mask
					}

					// Key press
					if idx < len(scanCodes) {
						scanCodes[idx] = scanCode
						idx++
					} else {
						for i := 0; i < len(scanCodes); i++ {
							scanCodes[idx] = key.ERROVF
						}
					}
				}
			}
			r.High()
		}

		// Send scan codes if the scan codes is different from the last data.
		// TODO: compare array
		changed := false
		for i := 0; i < len(lastScanCodes); i++ {
			if lastScanCodes[i] != scanCodes[i] {
				changed = true
			}
			lastScanCodes[i] = scanCodes[i]
		}
		if changed {
			err := device.SendKeyboardReport(scanCodes, modifier)
			if err != nil {
				println(err.Error())
			}
		}

		// Set layer
		if layerSwitch {
			layer = 1
		} else {
			layer = 0
		}

		time.Sleep(time.Millisecond * 5)
	}
}
