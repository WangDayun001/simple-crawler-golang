package main

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "os"
  "time"
  "github.com/gocolly/colly"

  Utils "./utils"
)

/*
 * Structs
 */
type Answers struct {
  Answers []Answer `json:"answers"`
}

type Answer struct {
  Value   string `json:"value"`
  Time   string `json:"time"`
}

/*
 * Global Props
 */
var answerFilePath = "answers.json"
var websiteUrl = "http://applicant-test.us-east-1.elasticbeanstalk.com/"

/*
 * Functions
 */
func getStoredAnswers () Answers {
  Utils.CreateAnswersFileIfNotExists(answerFilePath)

  jsonFile, err := os.Open(answerFilePath)
  if err != nil {
    fmt.Println(err)
  }
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  var answers Answers
  json.Unmarshal(byteValue, &answers)

  return answers
}

func getCurrentAnswer () Answer {
  // initial props
  c := colly.NewCollector()
  var token string
  var answerValue string

  // methods
  c.OnRequest(func(r *colly.Request) {
    r.Headers.Set("Referer", websiteUrl)
  })

  c.OnHTML("input[name=token]", func(e *colly.HTMLElement) {
    token = Utils.ParseToken(e.Attr("value"))
  })

  c.OnHTML("#answer", func(e *colly.HTMLElement) {
    currentAnswer := e.Text
    answerValue = currentAnswer
  })

  // init
  c.Visit(websiteUrl)

  c.Post(websiteUrl, map[string]string{"token": token})

  answer := Answer{
    Time: time.Now().String(),
    Value: answerValue,
  }

  fmt.Println("Answer Value: ", answer.Value)
  fmt.Println("Answer Time: ", answer.Time)

  return answer
}

func saveAnswers (answers Answers) {
  file, _ := json.MarshalIndent(answers, "", " ")
	_ = ioutil.WriteFile(answerFilePath, file, 0644)
}

func main() {
  storedAnswers := getStoredAnswers()
  currentAnswer := getCurrentAnswer()
  storedAnswers.Answers = append(storedAnswers.Answers, currentAnswer)
  saveAnswers(storedAnswers)
}
