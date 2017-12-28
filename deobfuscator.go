package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
    reader2 := bufio.NewReader(os.Stdin)
    fmt.Print("Obfuscated Text : ")
    text2, _ := reader2.ReadString('\n')
    newtext2 := text2[:len(text2)-2]

    deobfuscated := strings.NewReplacer("AbCfDdgSeZSAzDQQWfFQQW", "", "Vsfq1rtqvvqswfqwq1", "", "rfqvqwrfq1112qrqqq", "", "==", "")
    output := deobfuscated.Replace(newtext2)

    fmt.Println("Deobfuscated :", output)
}
