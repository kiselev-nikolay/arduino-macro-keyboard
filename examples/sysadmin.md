# Sysadmin keymap

```go
switch message {
case "0":
    kb.HasCTRL(true)
    kb.HasALT(true)
    kb.AddKey(keybd_event.VK_DELETE)
    log.Println("CTRL+ALT+DEL pressed")
case "1":
    kb.AddKey(keybd_event.VK_S)
    kb.AddKey(keybd_event.VK_U)
    kb.AddKey(keybd_event.VK_D)
    kb.AddKey(keybd_event.VK_O)
    kb.AddKey(keybd_event.VK_SPACE)
    log.Println("`sudo ` inserted")
case "2":
    kb.AddKey(keybd_event.VK_R)
    kb.AddKey(keybd_event.VK_E)
    kb.AddKey(keybd_event.VK_B)
    kb.AddKey(keybd_event.VK_O)
    kb.AddKey(keybd_event.VK_O)
    kb.AddKey(keybd_event.VK_T)
    kb.AddKey(keybd_event.VK_ENTER)
    log.Println("`reboot` inserted and enter pressed")
case "3":
    kb.AddKey(keybd_event.VK_H)
    kb.AddKey(keybd_event.VK_I)
    kb.AddKey(keybd_event.VK_S)
    kb.AddKey(keybd_event.VK_T)
    kb.AddKey(keybd_event.VK_O)
    kb.AddKey(keybd_event.VK_R)
    kb.AddKey(keybd_event.VK_Y)
    kb.AddKey(keybd_event.VK_SPACE)
    kb.AddKey(keybd_event.VK_MINUS)
    kb.AddKey(keybd_event.VK_C)
    kb.AddKey(keybd_event.VK_ENTER)
    log.Println("`history -c` inserted and enter pressed")
default:
    continue
}
```