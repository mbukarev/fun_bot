package main

import (
	"fmt"
	. "fun_bot/pkg/Config"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"strings"
	"time"
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
		text := c.Text()
		//fmt.Println(c.Message().MessageSig())
		fb, err := os.ReadFile("configs/pwords.txt") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		str := string(fb) // convert content to a 'string'
		words := strings.Split(text, " ")
		bad := false
		for _, v := range words {
			if strings.Contains(str, v) {
				text = strings.Replace(text, v, "*Пи-и-и*", -1)
				bad = true
			}

		}
		if bad {
			erred := b.Delete(c.Message())
			if erred != nil {
				fmt.Print(erred)
			}
			err := c.Send("Прилетело НЛО и удалило пахабное сообщение:\n" + text)
			if err != nil {
				return err
			}
		}
		return nil
	})

	b.Start()
}
