# TinyGo for Arduino Uno

This repository contains TinyGo examples for the Arduino Uno microcontroller.

## Dependencies

`avr-gcc` and `avrdude` are needed in order to cross-compile and flash the program to the Arduino Uno.  Both of these dependencies are provided by the [Arduino IDE].

[Arduino IDE]: https://www.arduino.cc/en/software

## VS Code

It is recommended that VS Code users install the TinyGo extension so that the `go.toolsEnvVars` setting is properly configured.

## Usage

To flash the example program, simply execute the following:

`tinygo flash -target arduino -port <PORT> github.com/cheebz/uno/blink`

Where `<PORT>` is the port that your Arduino Uno is using. You can check the port that the Arduino Uno is using from within the Arduino IDE.

