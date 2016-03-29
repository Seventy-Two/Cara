package dota

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"strings"
)

const (
	dotaLeagueURL = "http://api.steampowered.com/IDOTA2Match_570/GetLiveLeagueGames/v1/?key=%s"
	dotaListingURL = "http://api.steampowered.com/IDOTA2Match_570/GetLeagueListing/v1/?key=%s"
	dotaMatchURL = "http://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v1/?key=%s"
)

func dotamatches(command *bot.Cmd, matches []string) (msg []string, err error) {
	data := &LeagueGames{}
	listing := &LeagueListing{}
	var radiantNet int 
	var direNet int
	var worth int
	err = web.GetJSON(fmt.Sprintf(dotaLeagueURL, bot.Config.API.Dota), data)
	if err != nil {
		msg = append(msg, fmt.Sprintf("Could not retrieve matches."))
		return msg, nil
	}
	err = web.GetJSON(fmt.Sprintf(dotaListingURL, bot.Config.API.Dota), listing)
	if err != nil {
		msg = append(msg, fmt.Sprintf("Could not retrieve league listings."))
		return msg, nil
	}
	var str []string
	var leaguename string
	for i := 0; i < len(data.Result.Games) ; i++ {
		worth = 0
		radiantNet = 0
		direNet = 0
		if data.Result.Games[i].LeagueTier == 2 && data.Result.Games[i].Spectators >= 500 {
			radiant := data.Result.Games[i].RadiantTeam.TeamName
			dire 	:= data.Result.Games[i].DireTeam.TeamName
			radiantScore := data.Result.Games[i].Scoreboard.Radiant.Score
			direScore := data.Result.Games[i].Scoreboard.Dire.Score
			game := data.Result.Games[i].RadiantSeriesWins + data.Result.Games[i].DireSeriesWins + 1
			viewers := data.Result.Games[i].Spectators
			for k :=0; k < len(listing.Result.Leagues); k++ {
				if data.Result.Games[i].LeagueID == listing.Result.Leagues[k].Leagueid {
					leaguename = listing.Result.Leagues[k].Name
					leaguename = strings.Replace(leaguename, "#DOTA_Item", "", -1)
					leaguename = strings.Replace(leaguename, "_", " ", -1)
					leaguename = strings.TrimSpace(leaguename)
				}
			}

			for j := 0; j < len(data.Result.Games[i].Scoreboard.Radiant.Players); j++ {
				radiantNet += data.Result.Games[i].Scoreboard.Radiant.Players[j].NetWorth
			}
			for j := 0; j < len(data.Result.Games[i].Scoreboard.Dire.Players); j++ {
				direNet += data.Result.Games[i].Scoreboard.Dire.Players[j].NetWorth
			}
			worth = radiantNet - direNet

			if worth > 0 {
				str = append(str, fmt.Sprintf("\x0307net+%d \x0303%s\x03 %d-%d \x0304%s\x03 [Game %d, %d viewers, League: %s]" , 
																												     worth, 
																													 radiant, 
																													 radiantScore, 
																													 direScore, 
																													 dire, 
																													 game, 
																													 viewers,
																													 leaguename))
			} else {
				worth = -worth
				str = append(str, fmt.Sprintf("\x0303%s\x03 %d-%d \x0304%s \x0307net+%d\x03 [Game %d, %d viewers, League: %s]", 
																												     radiant, 
																													 radiantScore, 
																													 direScore, 
																													 dire, 
																													 worth, 
																													 game, 
																													 viewers,
																													 leaguename))
			}
		}
	}
	if len(str) == 0 {
		msg = append(msg, fmt.Sprintf("No games found."))
		return msg, nil
	}
	return str, nil
}

func init() {
	bot.RegisterMultiCommand(
		"^(d2|dota)$",
		dotamatches)
}