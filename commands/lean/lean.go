package lean

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"math/rand"
	"time"
	"os"
	"bufio"
	"github.com/kardianos/osext"
)

const (
	leanPath = "\\leanlines"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func leandoer(command *bot.Cmd, matches []string) (msg string, err error) {
	dir, _ := osext.ExecutableFolder()
	quotePath := fmt.Sprintf("%s%s", dir, leanPath)
	file, err := os.Open(quotePath)
  	if err != nil {
    	return "", nil
  	}
	defer file.Close()
	var choices []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		choices = append(choices, scanner.Text())
	}
	chosen := random(0, len(choices))
	msg = fmt.Sprintf("%s", choices[chosen])
	return
}

func init() {
	bot.RegisterCommand(
		"^lean$",
		leandoer)
}
