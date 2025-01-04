package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	disGo "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func startBot() {
	godotenv.Load()

	// Add your own pref and discord_token
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

		if args[1] == "anime" {
			apiResponse, err := fetchAnime()
			if err != nil {
				fmt.Println(err)
				return
			}

			discordEm := disGo.MessageEmbed{
				Fields: []*disGo.MessageEmbedField{
					{
						Name:   "id",
						Value:  strconv.Itoa(apiResponse.Data.MalID),
						Inline: true,
					},
					{
						Name:   "Score",
						Value:  strconv.Itoa(int(apiResponse.Data.Score)),
						Inline: true,
					},
					{
						Name:   "Title",
						Value:  apiResponse.Data.Title,
						Inline: false,
					},
					{
						Name:   "Description",
						Value:  apiResponse.Data.Synopsis,
						Inline: false,
					},
				},
				Image: &disGo.MessageEmbedImage{
					URL: apiResponse.Data.Image.Jpg.ImageURL,
				},
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
