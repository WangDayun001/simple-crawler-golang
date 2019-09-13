package main

import (
  "strings"
  "fmt"
)

func parseToken(rawToken string) string {
  fmt.Println("Raw token", rawToken)
  tokenCharacters := make([]string, len(rawToken))

  charactersReplacements := map[string]string {
    "a": "z",
    "b": "y",
    "c": "x",
    "d": "w",
    "e": "v",
    "f": "u",
    "g": "t",
    "h": "s",
    "i": "r",
    "j": "q",
    "k": "p",
    "l": "o",
    "m": "n",
    "n": "m",
    "o": "l",
    "p": "k",
    "q": "j",
    "r": "i",
    "s": "h",
    "t": "g",
    "u": "f",
    "v": "e",
    "w": "d",
    "x": "c",
    "y": "b",
    "z": "a",
  }

  for pos, char := range rawToken {
    tokenCharacters[pos] = charactersReplacements[string(char)]

    if tokenCharacters[pos] == "" {
      tokenCharacters[pos] = string(char)
    }
  }

  formatedToken := strings.Join(tokenCharacters, "")

  fmt.Println("Formated token", formatedToken)

  return formatedToken
}

func main() {
  token := "89wx7530z3148857829z2uv1017u1047"
  formatedToken := parseToken(token)

  fmt.Println("Formated token", formatedToken)
}
