@ECHO OFF

set PWD=%~dp0
set BOARD="arduino:avr:nano:cpu=atmega328"

arduino_debug.exe --verify --pref sketchbook.path=%PWD% --board %BOARD% %1
arduino_debug.exe --upload --pref sketchbook.path=%PWD% --board %BOARD% %1