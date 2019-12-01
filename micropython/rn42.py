class RN42:
    def __init__(self, uart):
        # Set up RN42
        self.uart = uart


    def write(self, key):
        self.uart.write(key)


    def send_keyboard_report(self, data, modifier):
        if len(data) != 6:
            return
        
        header = bytearray(b'\xFD\x09\x01\x00\x00')
        header[3] = modifier
        self.uart.write(header)
        self.uart.write(data)
