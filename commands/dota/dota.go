package dota

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"strings"
	"time"
)

const (
	dotaLeagueURL = "http://api.steampowered.com/IDOTA2Match_570/GetLiveLeagueGames/v1/?key=%s"
	dotaListingURL = "http://api.steampowered.com/IDOTA2Match_570/GetLeagueListing/v1/?key=%s"
	dotaMatchURL = "http://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v1/?key=%s"
	dotaHeroesURL = "http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v1/?language=en_gb&key=%s"
	Timer = "15:04:05"
)

func dotamatches(command *bot.Cmd, matches []string) (msg []string, err error) {
	data := &LeagueGames{}
	listing := &LeagueListing{}
	getHeroes := &GetHeroes{}
	var radiantNet int 
	var direNet int
	var worth int
	heroes := false
	showScore := false
	if strings.Contains(matches[0], "h") {
		heroes = true
	}
	if strings.Contains(matches[0], "s") {
		showScore = true
	}
	err = web.GetJSON(fmt.Sprintf(dotaListingURL, bot.Config.API.Dota), listing)
	if err != nil {
		msg = append(msg, fmt.Sprintf("Could not retrieve league listings."))
		return msg, nil
	}
	err = web.GetJSON(fmt.Sprintf(dotaLeagueURL, bot.Config.API.Dota), data)
	if err != nil {
		msg = append(msg, fmt.Sprintf("Could not retrieve matches."))
		return msg, nil
	}
	web.GetJSON(fmt.Sprintf(dotaHeroesURL, bot.Config.API.Dota), getHeroes)
	if err != nil {
		msg = append(msg, fmt.Sprintf("Could not retrieve heroes."))
		return msg, nil
	}
	var str []string
	for i := 0; i < len(data.Result.Games) ; i++ {
		worth = 0
		radiantNet = 0
		direNet = 0
		if (data.Result.Games[i].LeagueTier == 2 && data.Result.Games[i].Spectators >= 1000) || (data.Result.Games[i].LeagueTier == 3 && data.Result.Games[i].Spectators >= 200) {
			var leaguename string
			herostr := ""
			radTower := ""
			direTower := ""
			radHeroes := []string{""}
			direHeroes := []string{""}
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

			duration := int(data.Result.Games[i].Scoreboard.Duration)
			t := fmt.Sprintf((time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC).Add(time.Duration(duration) * time.Second)).Format(Timer))


			for j := 0; j < len(data.Result.Games[i].Scoreboard.Radiant.Players); j++ {
				radiantNet += data.Result.Games[i].Scoreboard.Radiant.Players[j].NetWorth
			}
			if heroes {
				for k := 0; k < len(data.Result.Games[i].Scoreboard.Radiant.Picks); k++ {
					radHeroes = append(radHeroes, getHerofromID(data.Result.Games[i].Scoreboard.Radiant.Picks[k].HeroID, getHeroes))
				}
			}
			for j := 0; j < len(data.Result.Games[i].Scoreboard.Dire.Players); j++ {
				direNet += data.Result.Games[i].Scoreboard.Dire.Players[j].NetWorth
			}
			if heroes {
				for k := 0; k < len(data.Result.Games[i].Scoreboard.Dire.Picks); k++ {
					direHeroes = append(direHeroes, getHerofromID(data.Result.Games[i].Scoreboard.Dire.Picks[k].HeroID, getHeroes))
				}
			}
			worth = radiantNet - direNet
			if showScore {
				radTower, direTower = towerToString(data.Result.Games[i].Scoreboard.Radiant.TowerState, data.Result.Games[i].Scoreboard.Dire.TowerState)
			}
			if heroes {
	 			herostr = fmt.Sprintf("\x0303%s\x03\x0304%s\x03|", strings.Join(radHeroes[:], "\x03|\x0303"), strings.Join(direHeroes[:], "\x03|\x0304"))
	 		}
	 		if showScore && worth != 0 {
				if worth > 0 {
					str = append(str, fmt.Sprintf("\x0307net+%d \x0303%s\x03 %s %d-%d %s \x0304%s\x03 [%s Game %d, %d viewers, League: %s %s]" , 
																													     worth, 
																														 radiant,
																														 radTower,
																														 radiantScore, 
																														 direScore,
																														 direTower, 
																														 dire,
																														 t,
																														 game, 
																														 viewers,
																														 leaguename,
																														 herostr))
				} else {
					worth = -worth
					str = append(str, fmt.Sprintf("\x0303%s\x03 %s %d-%d %s \x0304%s \x0307net+%d\x03 [%s Game %d, %d viewers, League: %s %s]", 
																													     radiant, 
																													     radTower,
																														 radiantScore, 
																														 direScore,
																														 direTower, 
																														 dire, 
																														 worth,
																														 t, 
																														 game, 
																														 viewers,
																														 leaguename,
																														 herostr))
				}
			} else {
				str = append(str, fmt.Sprintf("\x0303%s\x03 - \x0304%s\x03 [Game %d, %d viewers, League: %s %s]", 
																												     radiant, 
																													 dire, 
																													 game, 
																													 viewers,
																													 leaguename,
																													 herostr))	
			}
		}
	}
	if len(str) == 0 {
		msg = append(msg, fmt.Sprintf("No games found."))
		return msg, nil
	}
	return str, nil
}

func towerToString(rad int, dire int) (radTower string, direTower string) {
	towerUp := "♜"
	towerDown := "♖"
	ancient := "♚"
	radstr := ancient + fmt.Sprintf("%011b", int64(rad))
	radstr = organise(radstr)
	radstr = strings.Replace(radstr, "0", towerDown, -1)
	radstr = strings.Replace(radstr, "1", towerUp, -1)

	direstr := ancient + fmt.Sprintf("%011b", int64(dire))
	direstr = organise(direstr)
	var tempstr string
	for _,v := range direstr {
  	  tempstr = string(v) + tempstr	// because allow runes and unicode 
  	}
  	direstr = tempstr
	direstr = strings.Replace(direstr, "0", towerDown, -1)
	direstr = strings.Replace(direstr, "1", towerUp, -1)

	return radstr, direstr
}

func organise(in string) (out string) {
	// volvo returns towers grouped by top/mid/bot, when we want towers grouped by tier
	tempin := []rune(in)
	var tempout []rune
	tempout = append(tempout, tempin[0])
	tempout = append(tempout, ' ')
	tempout = append(tempout, tempin[1])
	tempout = append(tempout, tempin[2])
	tempout = append(tempout, ' ')
	tempout = append(tempout, tempin[3])
	tempout = append(tempout, tempin[6])
	tempout = append(tempout, tempin[9])
	tempout = append(tempout, ' ')
	tempout = append(tempout, tempin[4])
	tempout = append(tempout, tempin[7])
	tempout = append(tempout, tempin[10])
	tempout = append(tempout, ' ')
	tempout = append(tempout, tempin[5])
	tempout = append(tempout, tempin[8])
	tempout = append(tempout, tempin[11])
	out = string(tempout)
	return out
}

func getHerofromID(id int, heroes *GetHeroes) (out string) {
	if id == 0 {
		return "PICK"
	}
	out = getShortHero(id)
	if out != "" {
		return out
	}
	for m := 0; m < len(heroes.Result.Heroes); m++ {
		if heroes.Result.Heroes[m].ID == id {
			out = heroes.Result.Heroes[m].LocalizedName
			return out
		}
	}
	return ""
}

func init() {
	bot.RegisterMultiCommand(
		"^(d2|dota)\\s?(h(?:ero)?)?\\s?(s(?:core)?)?$",
		dotamatches)
}