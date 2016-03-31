package roll

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"math/rand"
	"strings"
	"time"
	"strconv"
)

func random(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1 // Since rand.Intn(1) = 0
}

func roll(command *bot.Cmd, matches []string) (msg string, err error) {
	var r []string
	var dice []string = strings.Split((strings.Split(matches[0], " ")[1]), "d")
	var sides int

	if len(dice) > 1 {
		if dice[1] == "0" {
			msg =fmt.Sprintf("You stupid cunt")
			return
		}
	}

	rollnum, _ := strconv.Atoi(dice[0])
	if (len(dice) > 1) {
		sides, _ = strconv.Atoi(dice[1])
	} else {
		sides = 10
	}

	for i := 0; i < rollnum; i++ {
		r = append(r, strconv.Itoa(random(sides)))
		time.Sleep(1 * time.Millisecond)
	}
	msg = fmt.Sprintf("%s rolls %dd%d : %s", command.Nick, rollnum, sides, strings.Join(r, ", "))
	return
}
func init() {
	bot.RegisterCommand(
		"^r(?:oll)? ([0-9]{1,2})((d)([0-9]{1,2}))?$",
		roll)
}
