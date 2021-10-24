#!/bin/sh
tinygo flash -target arduino -port COM3 github.com/cheebz/uno/${2-blink}