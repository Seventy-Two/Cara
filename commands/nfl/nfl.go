package nfl

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"gopkg.in/xmlpath.v2"
	"net/http"
	"strings"
	"time"
)

const (
	URL = "http://www.nfl.com/liveupdate/scorestrip/ss.xml"
	PostSeasonURL = "http://static.nfl.com/liveupdate/scorestrip/postseason/ss.xml" //TODO
)

func nfl(command *bot.Cmd, matches []string) (msg []string, err error) {
	doc, _ := http.Get(fmt.Sprintf(URL))
	defer doc.Body.Close()
	root, err := xmlpath.Parse(doc.Body)
	var debug bool

	if (strings.EqualFold(matches[1], "debug")) {
		debug = true
	} else {
		debug = false
	}

	if err != nil {
		msg = append(msg, fmt.Sprintf("Could not retrieve matches."))
		return msg, nil
	}

	todaysdate := getToday()
	date := xmlpath.MustCompile("/ss/gms/g/@d")

	dateIter := date.Iter(root)
	var i int
	var timeStr string
	var awayScoreStr string
	var homeScoreStr string
	i = 1
	for dateIter.Next() {
		timeStr = ""
		awayScoreStr = ""
		homeScoreStr = ""
		
		if (strings.EqualFold(todaysdate, dateIter.Node().String())|| debug) {
			home := xmlpath.MustCompile(fmt.Sprintf("/ss/gms/g[%d]/@hnn", i))
			homeScore := xmlpath.MustCompile(fmt.Sprintf("/ss/gms/g[%d]/@hs", i))
			away := xmlpath.MustCompile(fmt.Sprintf("/ss/gms/g[%d]/@vnn", i))
			awayScore := xmlpath.MustCompile(fmt.Sprintf("/ss/gms/g[%d]/@vs", i))
			quarter := xmlpath.MustCompile(fmt.Sprintf("/ss/gms/g[%d]/@q", i))
		 	time := xmlpath.MustCompile(fmt.Sprintf("/ss/gms/g[%d]/@t", i))

		 	homeStr, _ := home.String(root)
		 	awayStr, _ := away.String(root)
		 	quarterStr, _ := quarter.String(root)
		 	if (strings.EqualFold(quarterStr, "P")) {
			 	timeStr, _ = time.String(root)
			 	timeStr = timeStr + " ET"
			 } else {
 			 	homeScoreStr, _ = homeScore.String(root)
	 		 	awayScoreStr, _ = awayScore.String(root)
			 }

		 	homeStr = getTeamColour(homeStr)
		 	awayStr = getTeamColour(awayStr)
		 	quarterStr = fixQuarter(quarterStr)

		 	msg = append(msg, fmt.Sprintf(homeStr + " " + homeScoreStr + " - " + awayScoreStr + " " + awayStr + " [" + quarterStr + timeStr + "]"))
		}
		i++
	}

	return msg, nil
}

func getToday() (date string) {
	now := time.Now()
	nowUTC := now.UTC()
	loc, _ := time.LoadLocation("America/New_York")
	jst := nowUTC.In(loc)
	return jst.Format("Mon")
}

func getTeamColour(team string) (colouredTeam string) {
	switch(team) {
		case "cardinals":
			return "\x0304Cardinals\x03"
		case "falcons":
			return "\x0304Falcons\x03"
		case "panthers":
			return "\x0311Panthers\x03"
		case "bears":
			return "\x0305Bears\x03"
		case "cowboys":
			return "\x0312Cowboys\x03"
		case "lions":
			return "\x0311Lions\x03"
		case "packers":
			return "\x0308Packers\x03"
		case "vikings":
			return "\x0306Vikings\x03"
		case "saints":
			return "\x0305Saints\x03"
		case "giants":
			return "\x0312Giants\x03"
		case "eagles":
			return "\x0303Eagles\x03"
		case "rams":
			return "\x0312Rams\x03"
		case "49ers":
			return "\x030449ers\x03"
		case "seahawks":
			return "\x0312Seahawks\x03"
		case "buccaneers":
			return "\x0304Buccaneers\x03"
		case "redskins":
			return "\x0304Redskins\x03"
		case "ravens":
			return "\x0306Ravens\x03"
		case "bills":
			return "\x0302Bills\x03"
		case "bengals":
			return "\x0307Bengals\x03"
		case "browns":
			return "\x0305Browns\x03"
		case "broncos":
			return "\x0307Broncos\x03"
		case "texans":
			return "\x0304Texans\x03"
		case "colts":
			return "\x0312Colts\x03"
		case "jaguars":
			return "\x0311Jaguars\x03"
		case "chiefs":
			return "\x0304Chiefs\x03"
		case "dolphins":
			return "\x0311Dolphins\x03"
		case "patriots":
			return "\x0312Patriots\x03"
		case "jets":
			return "\x0303Jets\x03"
		case "raiders":
			return "\x0301Raiders\x03"
		case "steelers":
			return "\x0308Steelers\x03"
		case "chargers":
			return "\x0312Chargers\x03"
		case "titans":
			return "\x0311Titans\x03"
		default:
			return team
	}
}

func fixQuarter(quarter string) (prettyQuarter string) {
	switch (quarter){
		case "P":
			return ""
		case "1":
			return "Q1"
		case "2":
			return "Q2"
		case "3":
			return "Q3"
		case "4":
			return "Q4"
		case "5":
			return "OT"
		case "F":
			return "Final"
		case "FO":
			return "Final (OT)"
		default:
			return quarter
	}
}

func init() {
	bot.RegisterMultiCommand(
		"^nfl\\s?(debug)?$",
		nfl)
}