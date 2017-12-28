package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Obfuscated Text : ")
    text, _ := reader.ReadString('\n')
    newtext := text[:len(text)-2]

    deobfuscated := strings.NewReplacer("AbCfDdgSeZSAzDQQWfFQQW", "", "Vsfq1rtqvvqswfqwq1", "", "rfqvqwrfq1112qrqqq", "", "==", "")
    output := deobfuscated.Replace(newtext)

    fmt.Println("Deobfuscated :", output)
}
