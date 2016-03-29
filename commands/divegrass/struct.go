package divegrass

import (
	"time"
)

type Fixtures struct {
	TimeFrameStart string `json:"timeFrameStart"`
	TimeFrameEnd string `json:"timeFrameEnd"`
	Count int `json:"count"`
	Fixtures []struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Soccerseason struct {
				Href string `json:"href"`
			} `json:"soccerseason"`
			HomeTeam struct {
				Href string `json:"href"`
			} `json:"homeTeam"`
			AwayTeam struct {
				Href string `json:"href"`
			} `json:"awayTeam"`
		} `json:"_links"`
		Date time.Time `json:"date"`
		Status string `json:"status"`
		Matchday int `json:"matchday"`
		HomeTeamName string `json:"homeTeamName"`
		AwayTeamName string `json:"awayTeamName"`
		Result struct {
			GoalsHomeTeam int `json:"goalsHomeTeam"`
			GoalsAwayTeam int `json:"goalsAwayTeam"`
		} `json:"result"`
	} `json:"fixtures"`
}

