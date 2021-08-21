/*
Copyright (c) 2021 Kiselev Nikolai

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
	"go.bug.st/serial"
)

func getKeyer() *keybd_event.KeyBonding {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	// > For linux, it is very important to wait 2 seconds
	// (c) https://github.com/micmonay/keybd_event/blob/9db684e1b3b474fdce1c7dfd23b35e91d6b30f01/README.md#L28
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	return &kb
}

func getArduinoPort(portName string) serial.Port {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, availablePortName := range ports {
		if availablePortName == portName {
			port, err := serial.Open(portName, &serial.Mode{
				BaudRate: 115200,
				Parity:   serial.EvenParity,
				DataBits: 7,
				StopBits: serial.OneStopBit,
			})
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Port open")
			return port
		}
	}
	log.Fatal("Not found")
	return nil
}

func main() {
	quit := make(chan struct{})

	kb := getKeyer()
	arduinoPort := getArduinoPort(os.Args[1])

	messages := make(chan string)

	go func() {
		msg := bytes.NewBufferString("")
		buff := make([]byte, 1024)
		for {
			n, err := arduinoPort.Read(buff)
			if err != nil {
				panic(err)
			}
			if n == 0 {
				log.Println("\nEOF", kb)
				break
			}
			msg.Write(buff[:n])
			s := msg.String()
			if strings.Contains(s, "\n") {
				msgLines := strings.Split(s, "\n")
				for _, v := range msgLines[:len(msgLines)-1] {
					messages <- strings.TrimSpace(v)
				}
				msg.Reset()
				msg.WriteString(msgLines[len(msgLines)-1])
			}
		}
	}()

	go func() {
		for {
			message := <-messages
			kb.Clear()
			kb.HasCTRL(true)
			kb.HasSHIFT(true)
			kb.HasALT(true)
			// TODO: Populate this switch with smarter keys.
			switch message {
			case "0":
				kb.AddKey(keybd_event.VK_F1)
				log.Println("CTRL+SHIFT+ALT+F1 pressed")
			case "1":
				kb.AddKey(keybd_event.VK_F2)
				log.Println("CTRL+SHIFT+ALT+F2 pressed")
			case "2":
				kb.AddKey(keybd_event.VK_F3)
				log.Println("CTRL+SHIFT+ALT+F3 pressed")
			case "3":
				kb.AddKey(keybd_event.VK_F4)
				log.Println("CTRL+SHIFT+ALT+F4 pressed")
			case "4":
				kb.AddKey(keybd_event.VK_F5)
				log.Println("CTRL+SHIFT+ALT+F5 pressed")
			case "5":
				kb.AddKey(keybd_event.VK_F6)
				log.Println("CTRL+SHIFT+ALT+F6 pressed")
			case "6":
				kb.AddKey(keybd_event.VK_F7)
				log.Println("CTRL+SHIFT+ALT+F7 pressed")
			case "7":
				kb.AddKey(keybd_event.VK_F8)
				log.Println("CTRL+SHIFT+ALT+F8 pressed")
			case "8":
				kb.AddKey(keybd_event.VK_F9)
				log.Println("CTRL+SHIFT+ALT+F9 pressed")
			case "9":
				kb.AddKey(keybd_event.VK_F10)
				log.Println("CTRL+SHIFT+ALT+F10 pressed")
			case "10":
				kb.AddKey(keybd_event.VK_F11)
				log.Println("CTRL+SHIFT+ALT+F11 pressed")
			case "11":
				kb.AddKey(keybd_event.VK_F12)
				log.Println("CTRL+SHIFT+ALT+F12 pressed")
			default:
				continue
			}
			err := kb.Launching()
			if err != nil {
				panic(err)
			}
		}
	}()

	<-quit
}
