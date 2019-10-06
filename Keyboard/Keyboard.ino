// Based on
// http://eucalyn.hatenadiary.jp/entry/original-keyboard-08
#include "Keyboard.h"

const int rowNum = 6;
const int colNum = 12;

const int rowPin[rowNum] = { 5, 1, 0, 2, 3, 4 };
const int colPin[colNum] = { 10, 16, 9, 8, 7, 6, 21, 20, 19, 18, 15, 14 };

// Key map definition
// https://www.arduino.cc/reference/en/language/functions/usb/keyboard/keyboardmodifiers/
// http://www.asciitable.com/
// NOTE: use 0x00 for layer switch

const byte keyMap1[rowNum][colNum] = {
  { 0x00,           0x00, 0x27, 0x5B, 0x2D, 0x60                         ,0x5C, 0x3D, 0x5D, 0x2F, 0x00, 0x00},
  { KEY_ESC,        0x31, 0x32, 0x33, 0x34, 0x35                         ,0x36, 0x37, 0x38, 0x39, 0x30, KEY_DELETE},
  { KEY_TAB,        0x71, 0x77, 0x65, 0x72, 0x74                         ,0x79, 0x75, 0x69, 0x6F, 0x70, KEY_BACKSPACE},
  { KEY_LEFT_CTRL,  0x61, 0x73, 0x64, 0x66, 0x67                         ,0x68, 0x6A, 0x6B, 0x6C, 0x3B, KEY_RETURN},
  { KEY_LEFT_SHIFT, 0x7A, 0x78, 0x63, 0x76, 0x62                         ,0x6E, 0x6D, 0x2C, 0x2E, KEY_UP_ARROW, KEY_RIGHT_SHIFT},
  { KEY_LEFT_CTRL, KEY_LEFT_CTRL, KEY_LEFT_GUI, KEY_LEFT_ALT, 0x00, 0x20 ,0x20, 0x00, KEY_RIGHT_ALT, KEY_LEFT_ARROW, KEY_DOWN_ARROW, KEY_RIGHT_ARROW}
};

const byte keyMap2[rowNum][colNum] = {
  { 0x00,           0x00, 0x22, 0x7B, 0x5F, 0x7E                                               ,0x7C, 0x2B, 0x7D, 0x2F, 0x00, 0x00},
  { KEY_ESC,        KEY_F1, KEY_F2, KEY_F3, KEY_F4, KEY_F5                                     ,KEY_F6, KEY_F7, KEY_F8, KEY_F9, KEY_F10, KEY_INSERT},
  { KEY_TAB,        KEY_HOME, KEY_UP_ARROW, KEY_END, KEY_PAGE_UP, 0x74                         ,KEY_HOME, KEY_UP_ARROW, KEY_END, KEY_PAGE_UP, 0x70, KEY_BACKSPACE},
  { KEY_LEFT_CTRL,  KEY_LEFT_ARROW, KEY_DOWN_ARROW, KEY_RIGHT_ARROW, KEY_PAGE_DOWN, 0x67       ,KEY_LEFT_ARROW, KEY_DOWN_ARROW, KEY_RIGHT_ARROW, KEY_PAGE_DOWN, 0x3B, KEY_RETURN},
  { KEY_LEFT_SHIFT, 0x7A, 0x78, 0x63, 0x76, 0x62                                               ,0x6E, 0x6D, 0x2C, 0x2E, KEY_PAGE_UP, KEY_RIGHT_SHIFT},
  { KEY_LEFT_CTRL, KEY_LEFT_CTRL, KEY_LEFT_GUI, KEY_LEFT_ALT, 0x00, 0x20                       ,0x20, 0x00, KEY_RIGHT_ALT, KEY_HOME, KEY_PAGE_DOWN, KEY_END}
};

byte *keyMap;

bool currentState[rowNum][colNum];
bool beforeState[rowNum][colNum];

int i,j;

void setup() {
  keyMap = *keyMap1;
  
  for( i = 0; i < rowNum; i++){
    pinMode(rowPin[i],OUTPUT);
  }

  for( i = 0; i < colNum; i++){
    pinMode(colPin[i],INPUT_PULLUP);
  }

  for( i = 0; i < rowNum; i++){
    for( j = 0; j < colNum; j++){
      currentState[i][j] = HIGH;
      beforeState[i][j] = HIGH;
    }
    digitalWrite(rowPin[i],HIGH);
  }

  Serial.begin(9600);
  Keyboard.begin();
}

void loop() {
  delay(5);
  
  for( i = 0; i < rowNum; i++){
    digitalWrite( rowPin[i], LOW );

    for( j = 0; j < colNum; j++){
      currentState[i][j] = digitalRead(colPin[j]);

      if ( currentState[i][j] != beforeState[i][j] ){

        Serial.print("key(");
        Serial.print(i);
        Serial.print(",");
        Serial.print(j);
        Serial.print(")");

        if ( currentState[i][j] == LOW){
          Serial.println(" Push!");
          Keyboard.press( keyMap[i*colNum+j] );

          // layer switch
          if (keyMap[i*colNum+j] == 0x00) {
            keyMap = *keyMap2;
          }
        } else {
          Serial.println(" Release!");
          Keyboard.release( keyMap[i*colNum+j] );

          //layer switch
          if (keyMap[i*colNum+j] == 0x00) {
            keyMap = *keyMap1;
          }
        }
      beforeState[i][j] = currentState[i][j];
      }
    }
    digitalWrite( rowPin[i], HIGH );
  }
}
