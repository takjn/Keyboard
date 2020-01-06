class USB:
    def __init__(self, usb_hid):
        # Set up USBHID
        self.usb_hid = usb_hid

    def send_keyboard_report(self, data, modifier):
        if len(data) != 6:
            return

        d = bytearray(b'\x00\x00')
        d[0] = modifier
        d.extend(data)
        
        self.usb_hid.send(d)
