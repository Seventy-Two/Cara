package bot

import (
	"regexp"
	"strings"
)

var (
	re = regexp.MustCompile("\\s+") // Matches one or more spaces
)

func parse(s string, channel string, nick string) *Cmd {
	c := &Cmd{Raw: s}
	hyp := []rune(s) // 

	// Remove those retarded fullwidth characters that one guy uses
	for i := 0; i < len(hyp); i++ {
		if hyp[i] >= 65281 && hyp[i] <= 65370{
	        hyp[i] -= 65248
	    } 
	    if hyp[i] == 12288 {
	    	hyp[i] = 32
	    }
    }
    s = string(hyp)

	s = strings.TrimSpace(s)

	if !strings.HasPrefix(s, Config.User) && !strings.HasPrefix(s, Config.Prefix) &&!strings.HasPrefix(s, Config.AltPrefix){
		return nil
	}

	c.Channel = strings.TrimSpace(channel)
	c.Nick = strings.TrimSpace(nick)
	// Trim the prefix and extra spaces
	if strings.HasPrefix(s, Config.User){
		c.Message = strings.TrimPrefix(s, Config.User + ":")
		c.Message = strings.TrimPrefix(c.Message, Config.User + ",")
	} else if !GetChannelKey(channel, "prefix") {
		c.Message = strings.TrimPrefix(s, Config.Prefix)
	} else {
		c.Message = strings.TrimPrefix(s, Config.AltPrefix)
	}
	c.Message = strings.TrimSpace(c.Message)

	// check if we have the command and not only the prefix
	if c.Message == "" {
		return nil
	}

	// get the command
	pieces := strings.SplitN(c.Message, " ", 2)
	c.Command = pieces[0]

	if len(pieces) > 1 {
		// get the arguments and remove extra spaces
		c.FullArg = removeExtraSpaces(pieces[1])
		c.Args = strings.Split(c.FullArg, " ")
	}

	return c
}

func removeExtraSpaces(args string) string {
	return re.ReplaceAllString(strings.TrimSpace(args), " ")
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
