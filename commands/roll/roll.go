package roll

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"math/rand"
	"strings"
	"time"
	"strconv"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func roll(command *bot.Cmd, matches []string) (msg string, err error) {
	var r []string
	rollnum, _ := strconv.Atoi(matches[1])
	for i := 0; i < rollnum; i++ {
		r = append(r, strconv.Itoa(random(1, 10)))
		time.Sleep(1 * time.Millisecond)
	}
	msg = fmt.Sprintf("%s rolls %sd10 : %s", command.Nick, matches[1], strings.Join(r, ", "))
	return
}
func init() {
	bot.RegisterCommand(
		"^r(?:oll)? ([0-9]{1,2})$",
		roll)
}
