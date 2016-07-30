package lastfm

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"strings"
	"time"
)

func whosPlaying(command *bot.Cmd, matches []string) (msg []string, err error) {
	var users []string
	users = bot.GetNames(strings.ToLower(command.Channel))

	for _, user := range users {
		if bot.GetUserKey(user, "lastfm") != "" {
			username := checkLastfm(user, user)

			data := &NowPlaying{}
			err = web.GetJSON(fmt.Sprintf(NowPlayingURL, username, bot.Config.API.Lastfm), data)
			if err != nil || data.Error > 0 {
				continue
			}
			if data.Recenttracks.Attr.Total == "0" {
				continue
			}

			if data.Recenttracks.Track[0].Attr.Nowplaying != "true" {
				continue
			}

			var fmttags string
			tags := &ArtistTags{}
			if len(data.Recenttracks.Track[0].Artist.Mbid) > 10 {
				err = web.GetJSON(fmt.Sprintf(MBIDTagsURL, data.Recenttracks.Track[0].Artist.Mbid, bot.Config.API.Lastfm), tags)
			} else {
				err = web.GetJSON(fmt.Sprintf(ArtistTagsURL, data.Recenttracks.Track[0].Artist.Text, bot.Config.API.Lastfm), tags)
			}

			if err != nil {
				continue
			}

			maxlen := len(tags.Toptags.Tag)
			if maxlen > 4 {
				maxlen = 4
			}
			for i := range tags.Toptags.Tag[:maxlen] {
				fmttags += fmt.Sprintf("%s, ", tags.Toptags.Tag[i].Name)
			}
			fmttags = strings.TrimSuffix(fmttags, ", ")

			nick := bot.GetUserKey(user, "nickname")
			if nick != "" {
				user = nick
			}
			msg = append(msg, fmt.Sprintf("%s (%s) | “%s” by %s | %s",
				user,
				username,
				data.Recenttracks.Track[0].Name,
				data.Recenttracks.Track[0].Artist.Text,
				fmttags))

			time.Sleep(10 * time.Millisecond)
		}
	}
	return msg, nil
}

func init() {
	bot.RegisterMultiCommand(
		"^wp$",
		whosPlaying)
}
