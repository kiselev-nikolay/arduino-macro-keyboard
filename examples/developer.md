# Junior developer keymap

```go
switch message {
case "0":
    kb.AddKey(keybd_event.VK_BROWSER_HOME)
    log.Println("`https://stackoverflow.com/` open")
case "1":
    kb.HasCTRL(true)
    kb.AddKey(keybd_event.VK_C)
    log.Println("production code just appear")
case "2":
    kb.HasCTRL(true)
    kb.AddKey(keybd_event.VK_V)
    log.Println("such a job done")
case "3":
    kb.AddKey(keybd_event.VK_G)
    kb.AddKey(keybd_event.VK_I)
    kb.AddKey(keybd_event.VK_T)
    kb.AddKey(keybd_event.VK_SPACE)
    kb.AddKey(keybd_event.VK_P)
    kb.AddKey(keybd_event.VK_U)
    kb.AddKey(keybd_event.VK_S)
    kb.AddKey(keybd_event.VK_H)
    kb.AddKey(keybd_event.VK_SPACE)
    kb.AddKey(keybd_event.VK_MINUS)
    kb.AddKey(keybd_event.VK_F)
    kb.AddKey(keybd_event.VK_SPACE)
    kb.AddKey(keybd_event.VK_M)
    kb.AddKey(keybd_event.VK_A)
    kb.AddKey(keybd_event.VK_S)
    kb.AddKey(keybd_event.VK_T)
    kb.AddKey(keybd_event.VK_E)
    kb.AddKey(keybd_event.VK_R)
    kb.AddKey(keybd_event.VK_ENTER)
    log.Println("`git push -f master` inserted and enter pressed")
default:
    continue
}
```