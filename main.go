package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp) 
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4434836363509-4450474566833-bIHWERkxEKt8xeaVygsKJRbv ")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04CVLDUYN7-4437803409682-7d287d1e121fcf12a2415b895530a9e063ab80d9c2030e09dfed6fd5e58358f8")
  

     bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	 go printCommandEvents(bot.CommandEvents())

	 bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
	
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err :=strconv.Atoi(year)
			if err!= nil {
				println("error")
			}
			age:= 2021-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	 })


	 ctx, cancel := context.WithCancel(context.Background())
	 defer cancel()

	 err := bot.Listen(ctx)
	 if err != nil {
	     log.Fatal(err)
	 }
}