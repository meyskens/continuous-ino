Continuous Ino - CI server for Arduino  [![Build Status](https://travis-ci.org/meyskens/continuous-ino.svg?branch=master)](https://travis-ci.org/meyskens/continuous-ino)
======================================

Continuous Ino is a CI software designed to run unit tests in Arduino based hardware using frameworks like [arduinounit](https://github.com/mmurdoch/arduinounit).  
Continuous Ino will upload the code from a GitHub repository, run the tests and even emulates components. Continuous Ino itself runs on a Raspberry Pi, using a special designed PCB it's GPIO pins power an Arduino (nano) and is able to send and receive to the Arduino's interfaces. Both analog and digital signals are supported. 

## Background
Continuous Ino is developed for an international collaboration between TH Köln and Thomas More Geel. Since we have never seen a solution like this we've decided to open source our work under an Apache License.

## Project by
- Ben De Breuker
- Maartje Eyskens
- Kristie Lim