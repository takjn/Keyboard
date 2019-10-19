package key

const (
	// Modifier masks - used for the first byte in the HID report.
	// NOTE: The second byte in the report is reserved, 0x00

	MODLCTRL  = 0x01 // Modifier Left Control
	MODLSHIFT = 0x02 // Modifier Left Shift
	MODLALT   = 0x04 // Modifier Left Alt
	MODLMETA  = 0x08 // Modifier Left GUI
	MODRCTRL  = 0x10 // Modifier Right Control
	MODRSHIFT = 0x20 // Modifier Right Shift
	MODRALT   = 0x40 // Modifier Right Alt
	MODRMETA  = 0x80 // Modifier Right GUI

	// Scan codes - last N slots in the HID report (usually 6).
	// 0x00 if no key pressed.
	//
	// If more than N keys are pressed, the HID reports
	// KEY_ERR_OVF in all slots to indicate this condition.

	NONE   = 0x00 // No key pressed
	ERROVF = 0x01 // Keyboard Error Roll Over - used for all slots if too many keys are pressed ("Phantom key")
	A      = 0x04 // Keyboard a and A
	B      = 0x05 // Keyboard b and B
	C      = 0x06 // Keyboard c and C
	D      = 0x07 // Keyboard d and D
	E      = 0x08 // Keyboard e and E
	F      = 0x09 // Keyboard f and F
	G      = 0x0a // Keyboard g and G
	H      = 0x0b // Keyboard h and H
	I      = 0x0c // Keyboard i and I
	J      = 0x0d // Keyboard j and J
	K      = 0x0e // Keyboard k and K
	L      = 0x0f // Keyboard l and L
	M      = 0x10 // Keyboard m and M
	N      = 0x11 // Keyboard n and N
	O      = 0x12 // Keyboard o and O
	P      = 0x13 // Keyboard p and P
	Q      = 0x14 // Keyboard q and Q
	R      = 0x15 // Keyboard r and R
	S      = 0x16 // Keyboard s and S
	T      = 0x17 // Keyboard t and T
	U      = 0x18 // Keyboard u and U
	V      = 0x19 // Keyboard v and V
	W      = 0x1a // Keyboard w and W
	X      = 0x1b // Keyboard x and X
	Y      = 0x1c // Keyboard y and Y
	Z      = 0x1d // Keyboard z and Z

	K1 = 0x1e // Keyboard 1 and !
	K2 = 0x1f // Keyboard 2 and @
	K3 = 0x20 // Keyboard 3 and #
	K4 = 0x21 // Keyboard 4 and $
	K5 = 0x22 // Keyboard 5 and %
	K6 = 0x23 // Keyboard 6 and ^
	K7 = 0x24 // Keyboard 7 and &
	K8 = 0x25 // Keyboard 8 and *
	K9 = 0x26 // Keyboard 9 and (
	K0 = 0x27 // Keyboard 0 and )

	ENTER      = 0x28 // Keyboard Return (ENTER)
	ESC        = 0x29 // Keyboard ESCAPE
	BACKSPACE  = 0x2a // Keyboard DELETE (Backspace)
	TAB        = 0x2b // Keyboard Tab
	SPACE      = 0x2c // Keyboard Spacebar
	MINUS      = 0x2d // Keyboard - and _
	EQUAL      = 0x2e // Keyboard = and +
	LEFTBRACE  = 0x2f // Keyboard [ and {
	RIGHTBRACE = 0x30 // Keyboard ] and }
	BACKSLASH  = 0x31 // Keyboard \ and |
	HASHTILDE  = 0x32 // Keyboard Non-US # and ~
	SEMICOLON  = 0x33 // Keyboard ; and :
	APOSTROPHE = 0x34 // Keyboard ' and "
	GRAVE      = 0x35 // Keyboard ` and ~
	COMMA      = 0x36 // Keyboard , and <
	DOT        = 0x37 // Keyboard . and >
	SLASH      = 0x38 // Keyboard / and ?
	CAPSLOCK   = 0x39 // Keyboard Caps Lock

	F1  = 0x3a // Keyboard F1
	F2  = 0x3b // Keyboard F2
	F3  = 0x3c // Keyboard F3
	F4  = 0x3d // Keyboard F4
	F5  = 0x3e // Keyboard F5
	F6  = 0x3f // Keyboard F6
	F7  = 0x40 // Keyboard F7
	F8  = 0x41 // Keyboard F8
	F9  = 0x42 // Keyboard F9
	F10 = 0x43 // Keyboard F10
	F11 = 0x44 // Keyboard F11
	F12 = 0x45 // Keyboard F12

	SYSRQ      = 0x46 // Keyboard Print Screen
	SCROLLLOCK = 0x47 // Keyboard Scroll Lock
	PAUSE      = 0x48 // Keyboard Pause
	INSERT     = 0x49 // Keyboard Insert
	HOME       = 0x4a // Keyboard Home
	PAGEUP     = 0x4b // Keyboard Page Up
	DELETE     = 0x4c // Keyboard Delete Forward
	END        = 0x4d // Keyboard End
	PAGEDOWN   = 0x4e // Keyboard Page Down
	RIGHT      = 0x4f // Keyboard Right Arrow
	LEFT       = 0x50 // Keyboard Left Arrow
	DOWN       = 0x51 // Keyboard Down Arrow
	UP         = 0x52 // Keyboard Up Arrow

	COMPOSE = 0x65 // Keyboard Application

	LEFTCTRL   = 0xe0 // Keyboard Left Control
	LEFTSHIFT  = 0xe1 // Keyboard Left Shift
	LEFTALT    = 0xe2 // Keyboard Left Alt
	LEFTMETA   = 0xe3 // Keyboard Left GUI
	RIGHTCTRL  = 0xe4 // Keyboard Right Control
	RIGHTSHIFT = 0xe5 // Keyboard Right Shift
	RIGHTALT   = 0xe6 // Keyboard Right Alt
	RIGHTMETA  = 0xe7 // Keyboard Right GUI
)
