package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	disGo "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func startBot() {
	godotenv.Load()

	var pref = "~lesgo"
	token := os.Getenv("DISCORD_TOKEN")
	sess, err := disGo.New(token)
	if err != nil {
		log.Panicf("Sess error : %v", err)
	}

	sess.AddHandler(func(s *disGo.Session, m *disGo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")
		if args[0] != pref {
			return
		}

		if args[1] == "boom" {
			apiResponse, err := fetchAnime()
			if err != nil {
				fmt.Println(err)
				return
			}
			discordEm := disGo.MessageEmbed{
				Title:       apiResponse.Data.Title,
				Description: apiResponse.Data.Synopsis,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, &discordEm)
		}
	})

	sess.Identify.Intents = disGo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		return
	}
	defer sess.Close()

	fmt.Println("our lil bot is online yeee")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, os.Interrupt)
	<-sc
}
