package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "6774983427:AAG4X93UVpxigwiUcsZu_xHzJ1KIgYHEvYM"

	// Create Bot with debug on
	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Stop reviving updates from updates channel
	defer bot.StopLongPolling()

	// Loop through all updates when they came
	for update := range updates {
		// Check if update contains message
		if update.Message != nil {
			// Get chat ID from message
			chatID := tu.ID(update.Message.Chat.ID)

			bot.SendMessage(&telego.SendMessageParams{
				ChatID: chatID,
				Text:   "hello",
			})

			f, err := os.Open("C:\\Users\\Valentyn\\Pictures\\Screenshots\\Знімок екрана 2024-05-12 222852.png")
			if err != nil {
				log.Fatal("Cannt read file")
			}
			bot.SendPhoto(&telego.SendPhotoParams{
				ChatID: chatID,
				Photo:  telego.InputFile{File: f},
			})

		}

	}

}

type Game interface {

}

func GetCharecters(Id int, Turn bool) {

}

type Charecters interface {
	// Anemi  []string
	// Herous []string
}


func (g *Game) GetGame() {}
