package dictionary

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"net/url"
)

const (
	nikURL = "http://api.wordnik.com/v4/word.json/%s/definitions?limit=3&includeRelated=true&sourceDictionaries=all&useCanonical=true&includeTags=false&api_key=%s"
	wotdURL = "http://api.wordnik.com:80/v4/words.json/wordOfTheDay?api_key=%s"
)


func dict(command *bot.Cmd, matches []string) (msg []string, err error){
	text := url.QueryEscape(matches[1])
	var data []Wordnik
	var result []string
	
	err = web.GetJSON(fmt.Sprintf(nikURL, text, bot.Config.API.Wordnik), &data)
	if err != nil {
		result = append(result, fmt.Sprintf("There was a problem with your request."))
		return result, nil
	}
	if len(data) == 0 {
		result = append(result, fmt.Sprintf("Word/phrase not found."))
		return result, nil
	}
	cap := len(data)	// never >3 because limit=3 in URL
	for i := 0; i < cap; i++{
		result = append(result, fmt.Sprintf("%s | %s | %s", data[i].Word, data[i].PartOfSpeech, data[i].Text))
	}
	return result, nil
}

func wotd(command *bot.Cmd, matches []string) (msg string, err error){
	data := &Wotd{}
	err = web.GetJSON(fmt.Sprintf(wotdURL, bot.Config.API.Wordnik), data)
	if err != nil {
		return fmt.Sprintf("There was a problem with your request."), nil
	}
	return fmt.Sprintf("Word of the day: %s - %s - %s", data.Word, data.Note, data.Definitions[0].Text), nil // I really hate doing [0] but we only want one definition. I hate comments that cause horizontal scroll also.
}


func init() {

	bot.RegisterMultiCommand(
		"^dict (.+)$",
		dict)
	bot.RegisterCommand(
		"^(wotd|word)$",	// Word of the day
		wotd)
}