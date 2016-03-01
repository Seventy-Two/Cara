package google

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/Seventy-Two/Cara/web"
	"github.com/kennygrant/sanitize"
	"net/url"
)

const (
	googleURL = "https://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=%s"
)

type SearchResults struct {
	Responsedata struct {
		Results []struct {
			Gsearchresultclass string `json:"GsearchResultClass"`
			Unescapedurl       string `json:"unescapedUrl"`
			URL                string `json:"url"`
			Visibleurl         string `json:"visibleUrl"`
			Cacheurl           string `json:"cacheUrl"`
			Title              string `json:"title"`
			Titlenoformatting  string `json:"titleNoFormatting"`
			Content            string `json:"content"`
		} `json:"results"`
		Cursor struct {
			Resultcount string `json:"resultCount"`
			Pages       []struct {
				Start string `json:"start"`
				Label int    `json:"label"`
			} `json:"pages"`
			Estimatedresultcount string `json:"estimatedResultCount"`
			Currentpageindex     int    `json:"currentPageIndex"`
			Moreresultsurl       string `json:"moreResultsUrl"`
			Searchresulttime     string `json:"searchResultTime"`
		} `json:"cursor"`
	} `json:"responseData"`
	Responsedetails interface{} `json:"responseDetails"`
	Responsestatus  int         `json:"responseStatus"`
}

func google(command *bot.Cmd, matches []string) (msg string, err error) {
	results := &SearchResults{}
	err = web.GetJSON(fmt.Sprintf(googleURL, url.QueryEscape(matches[1])), results)
	if err != nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	if len(results.Responsedata.Results) == 0 {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	title := sanitize.HTML(results.Responsedata.Results[0].Title)
	link, _ := url.QueryUnescape(results.Responsedata.Results[0].URL)

	output := fmt.Sprintf("Google | %s | %s", title, link)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^g(?:oogle)? (.+)$",
		google)
}
