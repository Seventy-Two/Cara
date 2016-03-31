package stocks

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"strconv"
)

const (
	URL = "http://query.yahooapis.com/v1/public/yql?q=select+*+from+yahoo.finance.quotes+where+symbol+in+%s&env=http://datatables.org/alltables.env&format=json"
)

func stocks(command *bot.Cmd, matches []string) (msg string, err error){
	data := &Stocks{}
	formattedInput := fmt.Sprintf("(\"%s\")", matches[1])
	err = web.GetJSON(fmt.Sprintf(URL, formattedInput), data)
	if err != nil {
		return fmt.Sprintf("There was a problem with your request. %s", err), nil
	}
	if data.Query.Results.Quote.Ask == "" {
		return fmt.Sprintf("No results found."), nil
	}
	var change string
	var perc string
	changeval, _ := strconv.ParseFloat(data.Query.Results.Quote.Change,32)
	if changeval < 0 {
		change = fmt.Sprintf("\x0304%s %s\x03", data.Query.Results.Quote.Change, data.Query.Results.Quote.Currency)
		perc   = fmt.Sprintf("\x0304%s\x03", data.Query.Results.Quote.PercentChange)
	} else {
		change = fmt.Sprintf("\x0303%s %s\x03", data.Query.Results.Quote.Change, data.Query.Results.Quote.Currency)
		perc   = fmt.Sprintf("\x0303%s\x03", data.Query.Results.Quote.PercentChange)
	}

	return fmt.Sprintf("%s | %s %s | %s (%s)", data.Query.Results.Quote.Name,
													data.Query.Results.Quote.Ask,
													data.Query.Results.Quote.Currency,
													change,
													perc,
													), nil 
}


func init() {

	bot.RegisterCommand(
		"^s(?:tocks)? (.+)$",
		stocks)
}
