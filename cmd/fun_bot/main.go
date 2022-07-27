package main

import (
  "fmt"
  "time"
  "log"
  "strings"
  "os"
	. "fun_bot/pkg/Config"
  tele "gopkg.in/telebot.v3"
)

func main() {
	var cfg Config
  fmt.Println("Start")
	cfg.GetConfig("configs/config.yaml")
  pref := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
  b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

  b.Handle(tele.OnText, func(c tele.Context) error {
	// All the text messages that weren't
	// captured by existing handlers.

  	var (
  		user = c.Sender()
  		text = c.Text()
  	)
    fb, err := os.ReadFile("configs/pwords.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    str := string(fb) // convert content to a 'string'

    words := strings.Split(text, " ")
    for _, v := range words {
        if strings.Contains(str, v)  {
            text = strings.Replace(text, v, "test", -1)
            b.Edit(c.Message, "ok")
        }

    }
  	// Use full-fledged bot's functions
  	// only if you need a result:
    _, err = b.Send(user, text)
  	if err != nil {
  		return err
  	}
    //fmt.Print(msg)
  	// Instead, prefer a context short-hand:
  	return nil
  })

	b.Start()
	fmt.Print(cfg.Token)
}
