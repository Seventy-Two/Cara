package omdb

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"net/url"
)

const (
	URL = "http://www.omdbapi.com/?t=%s&y=&plot=short&r=json"
)

func omdb(command *bot.Cmd, matches []string) (msg string, err error){
	data := &Omdb{}
	err = web.GetJSON(fmt.Sprintf(URL, url.QueryEscape(matches[1])), data)

	if err != nil {
		return fmt.Sprintf("There was a problem with your request."), nil
	}
	if data.Title == "" {
		return fmt.Sprintf("Not found."), nil
	}
	return fmt.Sprintf("%s (%s) | %s | %s | %s", data.Title, data.Year, data.Genre, data.ImdbRating, data.Plot), nil
}


func init() {

	bot.RegisterCommand(
		"^m(?:db|ovie)? (.+)$",
		omdb)
}