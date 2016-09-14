package quote

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"math/rand"
	"time"
	"os"
	"bufio"
	"strings"
	"regexp"
	"github.com/kardianos/osext"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func quote(command *bot.Cmd, matches []string) (msg string, err error) {
	dir, _ := osext.ExecutableFolder()
	quotePath := fmt.Sprintf("%s\\logs\\%s.log", dir, strings.Replace(strings.Replace(command.Channel, "/", "–", -1), "#", "", 1))
	reg, err := regexp.Compile("\\[.*\\]")
	file, err := os.Open(quotePath)
  	if err != nil {
    	return "Error getting logs", nil
  	}
	defer file.Close()
	var choices []string
	user := "<" + matches[1] + ">" 
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			if strings.HasPrefix(strings.TrimSpace(reg.ReplaceAllString(scanner.Text(), "")), user) && !strings.HasPrefix(strings.TrimSpace(reg.ReplaceAllString(scanner.Text(), "")), user + " .") && (strings.Count(scanner.Text(), " ") > 2 || (strings.Count(scanner.Text(), "　") > 2)) {
				choices = append(choices, strings.TrimSpace(reg.ReplaceAllString(scanner.Text(), "")))
			}
	}
	if len(choices) > 0 {
		chosen := random(0, len(choices))
		msg = fmt.Sprintf("%s", choices[chosen])
		return msg, nil
	}
	return "No quotes found", nil
}

func init() {
	bot.RegisterCommand(
		"^quote (\\S+)$",
		quote)
}
