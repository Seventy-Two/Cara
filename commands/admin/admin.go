package admin

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/steveyen/gkvlite"
	"strings"
)

const (
	helpURL = "https://github.com/Seventy-Two/Cara#functions"
)

func help(command *bot.Cmd, matches []string) (msg string, err error) {
	return fmt.Sprintf("%s: %s", command.Nick, helpURL), nil
}

func setIgnore(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return
	}
	bot.SetUserKey(strings.TrimSpace(matches[1]), "ignore", "true")
	return "I never liked him anyway", nil
}

func setUnignore(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return
	}
	bot.DeleteUserKey(strings.TrimSpace(matches[1]), "ignore")
	return "Sorry about that", nil
}

func listChannels(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return
	}
	output := "I'm in:"

	bot.Channels.VisitItemsAscend([]byte(""), true, func(i *gkvlite.Item) bool {
		if bot.GetChannelKey(string(i.Key), "auto_join") == true {
			output = fmt.Sprintf("%s %s", output, string(i.Key))
		}
		return true
	})

	return output, nil
}

func setPrefix(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return "", nil
	}

	onOff := matches[1]
	channelToToggle := matches[2]

	if onOff == "on" {
		bot.SetChannelKey(channelToToggle, "prefix", true)
		return fmt.Sprintf("Alternate prefix enabled in %s", channelToToggle), nil
	} else if onOff == "off" {
		bot.SetChannelKey(channelToToggle, "prefix", false)
		return fmt.Sprintf("Alternate prefix disabled in %s", channelToToggle), nil
	}
	return "", nil
}

func init() {
	bot.RegisterCommand("^help",
		help)

	bot.RegisterCommand(
		"^set ignore (\\S+)$",
		setIgnore)

	bot.RegisterCommand(
		"^set unignore (\\S+)$",
		setUnignore)

	bot.RegisterCommand(
		"^list channels$",
		listChannels)

	bot.RegisterCommand(
		"^set altprefix (\\S+) (\\S+)$",
		setPrefix)
}
