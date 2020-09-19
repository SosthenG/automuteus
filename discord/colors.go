package discord

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

// Color : Int constant mapping
const (
	Red    = 0
	Blue   = 1
	Green  = 2
	Pink   = 3
	Orange = 4
	Yellow = 5
	Black  = 6
	White  = 7
	Purple = 8
	Brown  = 9
	Cyan   = 10
	Lime   = 11
)

// ColorStrings for lowercase, possibly for translation if needed
var ColorStrings = map[string]int{
	"red":    Red,
	"blue":   Blue,
	"green":  Green,
	"pink":   Pink,
	"orange": Orange,
	"yellow": Yellow,
	"black":  Black,
	"white":  White,
	"purple": Purple,
	"brown":  Brown,
	"cyan":   Cyan,
	"lime":   Lime,
}

// GetColorStringForInt does what it sounds like
func GetColorStringForInt(colorint int) string {
	for str, idx := range ColorStrings {
		if idx == colorint {
			return str
		}
	}
	return ""
}

func IsColorString(test string) bool {
	_, ok := ColorStrings[test]
	return ok
}

// Emoji struct for discord
type Emoji struct {
	Name string
	ID   string
}

func (e *Emoji) FormatForReaction() string {
	return "<:" + e.Name + ":" + e.ID
}

func (e *Emoji) GetDiscordCDNUrl() string {
	return "https://cdn.discordapp.com/emojis/" + e.ID + ".png"
}

func (e *Emoji) DownloadAndBase64Encode() string {
	url := e.GetDiscordCDNUrl()
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	encodedStr := base64.StdEncoding.EncodeToString(bytes)
	return encodedStr
}

// AlivenessColoredEmojis keys are IsAlive, Color
var AlivenessColoredEmojis = map[bool]map[int]Emoji{
	true: map[int]Emoji{
		Red: {
			Name: "red",
			ID:   "756202732301320325",
		},
		Blue: {
			Name: "blue",
			ID:   "756201148154642642",
		},
		Green: {
			Name: "green",
			ID:   "756202732099993753",
		},
		Pink: {
			Name: "pink",
			ID:   "756200620049956864",
		},
		Orange: {
			Name: "orange",
			ID:   "756202732523618435",
		},
		Yellow: {
			Name: "yellow",
			ID:   "756202732678938624",
		},
		Black: {
			Name: "black",
			ID:   "756202732758761522",
		},
		White: {
			Name: "white",
			ID:   "756202732343394386",
		},
		Purple: {
			Name: "purple",
			ID:   "756202732624543770",
		},
		Brown: {
			Name: "brown",
			ID:   "756202732594921482",
		},
		Cyan: {
			Name: "cyan",
			ID:   "756202732511297556",
		},
		Lime: {
			Name: "lime",
			ID:   "756202732360040569",
		},
	},
	false: map[int]Emoji{
		Red: {
			Name: "reddead",
			ID:   "756404218163888200",
		},
		Blue: {
			Name: "bluedead",
			ID:   "756552864309969057",
		},
		Green: {
			Name: "greendead",
			ID:   "756552867275604008",
		},
		Pink: {
			Name: "pinkdead",
			ID:   "756552867413753906",
		},
		Orange: {
			Name: "orangedead",
			ID:   "756404218436517888",
		},
		Yellow: {
			Name: "yellowdead",
			ID:   "756404218339786762",
		},
		Black: {
			Name: "blackdead",
			ID:   "756552864171557035",
		},
		White: {
			Name: "whitedead",
			ID:   "756552867200106596",
		},
		Purple: {
			Name: "purpledead",
			ID:   "756552866491138159",
		},
		Brown: {
			Name: "browndead",
			ID:   "756552864620347422",
		},
		Cyan: {
			Name: "cyandead",
			ID:   "756204054698262559",
		},
		Lime: {
			Name: "limedead",
			ID:   "756552864847102042",
		},
	},
}
