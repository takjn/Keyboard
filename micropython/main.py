from machine import Pin
from machine import UART
from key import Key
from rn42 import RN42


class Keyboard:
    cols = ['PIN10', 'PIN11', 'PIN9', 'PIN8', 'PIN7', 'PIN4', 'PIN17', 'PIN16', 'PIN15', 'PIN14', 'PIN13', 'PIN12']
    rows = ['PIN18', 'PIN19', 'PIN2', 'PIN3', 'PIN5', 'PIN6']
    key_map = [
        [
            [Key.NONE, Key.NONE, Key.APOSTROPHE, Key.LEFTBRACE, Key.MINUS, Key.GRAVE, Key.BACKSLASH, Key.EQUAL, Key.RIGHTBRACE, Key.SLASH, Key.NONE, Key.NONE],
            [Key.ESC, Key.K1, Key.K2, Key.K3, Key.K4, Key.K5, Key.K6, Key.K7, Key.K8, Key.K9, Key.K0, Key.DELETE],
            [Key.TAB, Key.Q, Key.W, Key.E, Key.R, Key.T, Key.Y, Key.U, Key.I, Key.O, Key.P, Key.BACKSPACE],
            [Key.LEFTCTRL, Key.A, Key.S, Key.D, Key.F, Key.G, Key.H, Key.J, Key.K, Key.L, Key.SEMICOLON, Key.ENTER],
            [Key.LEFTSHIFT, Key.Z, Key.X, Key.C, Key.V, Key.B, Key.N, Key.M, Key.COMMA, Key.DOT, Key.UP, Key.RIGHTSHIFT],
            [Key.LEFTCTRL, Key.SYSRQ, Key.LEFTMETA, Key.LEFTALT, Key.NONE, Key.SPACE, Key.SPACE, Key.NONE, Key.COMPOSE, Key.LEFT, Key.DOWN, Key.RIGHT],
        ],
        [
            [Key.NONE, Key.NONE, Key.APOSTROPHE, Key.LEFTBRACE, Key.MINUS, Key.GRAVE, Key.BACKSLASH, Key.EQUAL, Key.RIGHTBRACE, Key.SLASH, Key.NONE, Key.NONE],
            [Key.ESC, Key.F1, Key.F2, Key.F3, Key.F4, Key.F5, Key.F6, Key.F7, Key.F8, Key.F9, Key.F10, Key.DELETE],
            [Key.TAB, Key.Q, Key.W, Key.E, Key.R, Key.T, Key.Y, Key.U, Key.I, Key.F11, Key.F12, Key.BACKSPACE],
            [Key.LEFTCTRL, Key.A, Key.S, Key.D, Key.F, Key.G, Key.H, Key.J, Key.K, Key.L, Key.SEMICOLON, Key.ENTER],
            [Key.LEFTSHIFT, Key.Z, Key.X, Key.C, Key.V, Key.B, Key.N, Key.M, Key.COMMA, Key.DOT, Key.PAGEUP, Key.RIGHTSHIFT],
            [Key.LEFTCTRL, Key.SYSRQ, Key.LEFTMETA, Key.LEFTALT, Key.NONE, Key.SPACE, Key.SPACE, Key.NONE, Key.COMPOSE, Key.HOME, Key.PAGEDOWN, Key.END],
        ]
    ]

    # layer holds the current key map layer (0 or 1)
    layer = 0

    def __init__(self):
        # Set up pins
        for row in self.rows:
            r = Pin(row, Pin.OUT)
            r.on()

        for col in self.cols:
            _ = Pin(col, Pin.IN, Pin.PULL_UP)

        # workaround for MicroPython v1.9.4-1354-gebbaac271-dirty on 2019-11-24; GR-CITRUS with RX631
        for row in ['PIN5_A', 'PIN6_A']:
            r = Pin(row, Pin.OUT)
            r.on()

    def read_matrix(self):
        scan_codes = bytearray(b'\x00\x00\x00\x00\x00\x00')
        idx = 0
        modifier = 0
        layer_switch = False

        for i, row in enumerate(self.rows):
            r = Pin(row, Pin.OUT)
            r.off()

            # workaround for MicroPython v1.9.4-1354-gebbaac271-dirty on 2019-11-24; GR-CITRUS with RX631
            if (row == 'PIN5'):
                p = Pin('PIN5_A', Pin.OUT)
                p.off()
            if (row == 'PIN6'):
                p = Pin('PIN6_A', Pin.OUT)
                p.off()

            for j, col in enumerate(self.cols):
                c = Pin(col, Pin.IN, Pin.PULL_UP)

                # A pin is pulled up (set high) by default.
                # It should be false (low) when the key is pressed.
                if c.value() == 1:
                    continue

                # Get the scan code
                scan_code = self.key_map[self.layer][i][j]

                # Check layer key is pressed
                if scan_code == Key.NONE:
                    layer_switch = True

                # Check modifier key is pressed
                if scan_code >= 0xe0 and scan_code <= 0xe7:
                    shift = scan_code & 0b111
                    mask = 1 << shift
                    modifier = modifier | mask

                # Set the scan code
                if idx < len(scan_codes):
                    scan_codes[idx] = scan_code
                    idx += 1
                else:
                    for k, _ in enumerate(scan_codes):
                        scan_codes[k] = Key.ERROVF

                # print("row: {}, col: {}".format(i, j))
            r.on()

        # workaround for MicroPython v1.9.4-1354-gebbaac271-dirty on 2019-11-24; GR-CITRUS with RX631
        if (row == 'PIN5'):
            p = Pin('PIN5_A', Pin.OUT)
            p.on()
        if (row == 'PIN6'):
            p = Pin('PIN6_A', Pin.OUT)
            p.on()

        # Set layer
        if layer_switch:
            self.layer = 1
        else:
            self.layer = 0

        return scan_codes, modifier


def main():
    uart = UART(0, 115200)
    hid = RN42(uart)

    k = Keyboard()

    # last_scan_codes hold the last scan codes to detect key press and key release
    last_scan_codes = bytearray(b'\x00\x00\x00\x00\x00\x00')

    while 1:
        scan_codes, modifier = k.read_matrix()

        # Send scan codes if the scan codes is different from the last data.
        if last_scan_codes != scan_codes:
            hid.send_keyboard_report(scan_codes, modifier)
            print(scan_codes, modifier)
        last_scan_codes = scan_codes

        # pyb.delay(5)


main()
