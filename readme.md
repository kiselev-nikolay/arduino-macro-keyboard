# Arduino Macro Keyboard

![photo_2021-08-21_19-08-28](https://user-images.githubusercontent.com/55307887/130328007-18ca1915-a120-4e6b-b7eb-d9ecc9cfcbec.jpg)

### Recipe

0. Buy the kind of Arduino-things you like; buy about 12 tea spoons of Kailh Red mechanical switches
1. Cram as many buttons as the arduino allows (paralleling to the circuit like `5v->button->pin`). The Arduino Uno, for example, allows you to connect 12 keys in such a barbaric way
2. Download [`sketch_aug21c.ino`](https://github.com/kiselev-nikolay/arduino-macro-keyboard/blob/main/sketch_aug21c.ino) and upload it to Arduino via IDE
3. Figure out what COM port arduino is using, and run comand like `go run .\listener.go COM3` or `go run .\listener.go COM7`
4. Button will send to your OS and Go app will read it and perform <kbd>ctrl</kbd> + <kbd>shift</kbd> + <kbd>alt</kbd> + <kbd>F1</kbd> -like key combinations....
5. ... hey, actually you can set what you need by your self. Just edit this nooby switch at [listener.go:117](https://github.com/kiselev-nikolay/arduino-macro-keyboard/blob/main/listener.go#L117) *(or you can get one of the [examples/\*](https://github.com/kiselev-nikolay/arduino-macro-keyboard/tree/main/examples))*
