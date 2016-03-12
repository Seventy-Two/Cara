package lastfm

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"strings"
    "regexp"
)

const (
	GetInfoURL = "http://ws.audioscrobbler.com/2.0/?method=artist.getinfo&artist=%s&api_key=%s&format=json"
	bioLen = 350
	outLen = 450
)

func artist(command *bot.Cmd, matches []string) (msg string, err error) {
	artist := matches[1]

	info := &Info{}
	err = web.GetJSON(fmt.Sprintf(GetInfoURL, artist, bot.Config.API.Lastfm), info)
	if err != nil {
		return fmt.Sprintf("Could not get artist info for %s", artist), nil
	}
	var bio string = info.ArtInfo.Bio.Summary
	//var name string = info.ArtInfo.Name
	 var fmttags string
	 maxlen := len(info.ArtInfo.Tags.Tag)
	 if len(info.ArtInfo.Tags.Tag) > 4 {
	 	maxlen = 4
	 } 
	 for i := range info.ArtInfo.Tags.Tag[:maxlen] {
	 	fmttags += fmt.Sprintf("%s, ", info.ArtInfo.Tags.Tag[i].Name)
	 }
	 fmttags = strings.TrimSuffix(fmttags, ", ")
	bio = newLineTrim(bio, bioLen)
	bio = garbageClean(bio)
	msg = fmt.Sprintf("Last.fm | %s  | %s", bio, fmttags)
	msg = newLineTrim(msg, outLen)
	return 
}

func newLineTrim(in string, trimLen int) (out string) {
	reg, _ := regexp.Compile("\n")
	out = reg.ReplaceAllString(in, " ")
	if len(out) > trimLen{
		out = out[:trimLen]
		out += "..."
	}
    return out
}

func garbageClean(in string) (out string) {
	reg, _ := regexp.Compile("<a href=\"|\">.*</a>") // Yeah thanks RE2 fucking idiots
	out = reg.ReplaceAllString(in, "")
	return out
}
func init() {
	bot.RegisterCommand(
		"^artist (.+)$",
		artist)
}
