package main

import (
  "strings"
  "fmt"

  "github.com/gocolly/colly"
)

func parseToken(rawToken string) string {
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

  return formatedToken
}

func getAnswer () {
  // initial props
  c := colly.NewCollector()
  token := ""

  // methods
  c.OnRequest(func(r *colly.Request) {
    r.Headers.Set("Referer", "http://applicant-test.us-east-1.elasticbeanstalk.com/")
  })

  c.OnHTML("input[name=token]", func(e *colly.HTMLElement) {
    token = parseToken(e.Attr("value"))
  })

  c.OnHTML("#answer", func(e *colly.HTMLElement) {
    answer := e.Text

    fmt.Println("Answer", answer)
  })

  // init
  c.Visit("http://applicant-test.us-east-1.elasticbeanstalk.com/")

  c.Post("http://applicant-test.us-east-1.elasticbeanstalk.com/", map[string]string{"token": token})
}

func main() {
  getAnswer()
}
