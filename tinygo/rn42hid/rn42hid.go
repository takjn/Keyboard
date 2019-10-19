// Package rn42hid provides a driver for the RN42HID Bluetooth module
//
// Datasheet:
// http://ww1.microchip.com/downloads/en/DeviceDoc/50002328A.pdf
// http://ww1.microchip.com/downloads/en/DeviceDoc/bluetooth_cr_UG-v1.0r.pdf
package rn42hid

import (
	"machine"
)

// Device wraps UART connection to RN42HID device.
type Device struct {
	bus machine.UART
}

// New creates a new RN42HID connection. The UART must already be configured.
//
// This function only creates the Device object, it does not touch the device.
func New(u machine.UART) Device {
	return Device{
		bus: u,
	}
}

// Write raw bytes to the UART.
func (d *Device) Write(b []byte) (n int, err error) {
	return d.bus.Write(b)
}

// Read raw bytes from the UART.
func (d *Device) Read(b []byte) (n int, err error) {
	return d.bus.Read(b)
}

// Send sends HID reports for printable ASCII characters.
// See 5.3.1 Translation Mode in datasheet.
func (d Device) Send(cmd string) error {
	_, err := d.Write([]byte(cmd))
	return err
}

// SendKeyboardReport sends a raw keyboard report for scan codes and modifier keys.
// The data byte array should be specified in scan codes or encoded values
// and the length should be 6.
// The modifier byte is a bit mask value that represents modifier key status.
// See 5.3.3 Raw Report Mode in datasheet.
func (d Device) SendKeyboardReport(data []byte, modifier byte) error {
	if len(data) != 6 {
		return nil
		// TODO return error
		// return fmt.Errorf("length of data should be 6")
	}

	header := []byte{0xFD, 0x09, 0x01, modifier, 0x00}
	_, err := d.Write(header)
	if err != nil {
		return err
	}
	_, err = d.Write(data)
	return nil
}
