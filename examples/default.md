# Default keymap

```go
switch message {
case "0":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F1)
    log.Println("CTRL+SHIFT+ALT+F1 pressed")
case "1":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F2)
    log.Println("CTRL+SHIFT+ALT+F2 pressed")
case "2":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F3)
    log.Println("CTRL+SHIFT+ALT+F3 pressed")
case "3":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F4)
    log.Println("CTRL+SHIFT+ALT+F4 pressed")
case "4":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F5)
    log.Println("CTRL+SHIFT+ALT+F5 pressed")
case "5":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F6)
    log.Println("CTRL+SHIFT+ALT+F6 pressed")
case "6":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F7)
    log.Println("CTRL+SHIFT+ALT+F7 pressed")
case "7":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F8)
    log.Println("CTRL+SHIFT+ALT+F8 pressed")
case "8":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F9)
    log.Println("CTRL+SHIFT+ALT+F9 pressed")
case "9":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F10)
    log.Println("CTRL+SHIFT+ALT+F10 pressed")
case "10":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F11)
    log.Println("CTRL+SHIFT+ALT+F11 pressed")
case "11":
    kb.HasCTRL(true)
    kb.HasSHIFT(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_F12)
    log.Println("CTRL+SHIFT+ALT+F12 pressed")
default:
    continue
}
```