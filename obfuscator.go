package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text : ")
    text, _ := reader.ReadString('\n')
    newtext := text[:len(text)-2]

    var s string

    for k := 0; k < len(newtext); k++ {
      s += fmt.Sprintf("%s%s%s%s%s", "AbCfDdgSeZSAzDQQWfFQQW", string(newtext[k]), "Vsfq1rtqvvqswfqwq1", "rfqvqwrfq1112qrqqq", "==")
    }

    fmt.Println("Obfuscated :", s)
}
