package utils

import (
  "fmt"
  "os"
)

func CreateAnswersFileIfNotExists (path string) {
  _, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
  }
}
