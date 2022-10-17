package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type CMD struct {
	ID int
}

func main() {
	c := cron.New(
		cron.WithParser(cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)),
	)
	for i := 0; i < 3; i++ {
		cmd := CMD{
			ID: i,
		}
		v := fmt.Sprintf("%v", cmd)
		fmt.Println("MMMM")
		fmt.Println(v)
		fmt.Println("WWWW")
		c.AddFunc("* * * * * *", job(cmd))
	}
	c.Start()

	for _, e := range c.Entries() {
		fmt.Println(e.ID)
	}
	select {}
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(spec.Next(time.Now()).Format("2006-01-02 15:04:05"))
}

func job(cmd CMD) func() {
	return func() {
		fmt.Printf("Hi %d\n", cmd.ID)
	}
}
