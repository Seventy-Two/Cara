package divegrass

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

const (
	URL = "http://api.football-data.org/v1/fixtures?timeFrame=%s&league=%s"
)

var (
	frame  = "p1"
) 

func modifier(command *bot.Cmd, matches []string) (msg []string, err error) {
	frame = matches[1]
	return divegrass(command, matches)
}

func divegrass(command *bot.Cmd, matches []string) (msg []string, err error) {
	data := &Fixtures{}
	client := &http.Client{}
	var str []string
	leagues := []string{"PL", "CL"}

	for i := 0; i < len(leagues); i++ {	// Loop through the leagues we want
		url := fmt.Sprintf(URL, url.QueryEscape(frame), url.QueryEscape(leagues[i]))
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("X-Auth-Token", bot.Config.API.FootballData)
		req.Header.Add("X-Response-Control", `minified`)
		resp, err := client.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		json.Unmarshal(body, data)
		if err != nil {
			str = append(str, fmt.Sprintf("There was a problem with your request"))
			return str, nil
		}
		if data.Count <= 0 && i == len(leagues) - 1 && len(str) == 0{
			str = append(str, fmt.Sprintf("There are no games available."))
			return str, nil
		}

		for j := 0; j < data.Count; j++ {
			date := data.Fixtures[j].Date
//			if date.Date() != time.Now().Date() {
				_,month,day := date.Date()
				fmtdate := fmt.Sprintf("%d/%d", day ,month)
//			} else {
//				hour,min,_ := date.Clock()
//				fmtdate := fmt.Sprintf("%d:%d", hour ,min)
//			}
			hName := data.Fixtures[j].HomeTeamName
			aName := data.Fixtures[j].AwayTeamName
			hScore := data.Fixtures[j].Result.GoalsHomeTeam
			aScore := data.Fixtures[j].Result.GoalsAwayTeam	
			str = append(str, strings.Replace(strings.Replace(fmt.Sprintf("%s %s %d - %d %s", fmtdate, hName, hScore, aScore, aName), " AFC", "", 2), " FC", "", 2))
		}
	}
	frame = "p1" // Yeah awful, I know.
	return str, nil
}

func init() {
	bot.RegisterMultiCommand(
		"^f(?:ooty)?$",
		divegrass)

	bot.RegisterMultiCommand(
		"^f(?:ooty)?\\s([p|n][0-9])$",
		modifier)
}