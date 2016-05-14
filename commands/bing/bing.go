package bing

import (
	"encoding/json"
	"fmt"
	"github.com/Seventy-Two/Cara"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	searchURL = "https://api.datamarket.azure.com/Bing/Search/v1/Web?Query='%s'&Options='DisableLocationDetection'&Market='en-GB'&Adult='Off'&$format=json"
)

type SearchResults struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ID          string `json:"ID"`
			Title       string `json:"Title"`
			Description string `json:"Description"`
			DisplayURL  string `json:"DisplayUrl"`
			URL         string `json:"Url"`
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func bing(command *bot.Cmd, matches []string) (msg string, err error) {
	var resNo int
	var querystr string
	if len(matches) > 1 {
		resNo, _ = strconv.Atoi(matches[1])
		querystr = matches[2]
	} else {
		querystr = matches[1]
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf(searchURL, url.QueryEscape(querystr)), nil)
	request.SetBasicAuth("", bot.Config.TranslateClient)
	
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var results SearchResults
	json.Unmarshal(body, &results)

	if err != nil {
		return fmt.Sprintf("No results for %s", querystr), nil
	}

	if len(results.D.Results) == 0 {
		return fmt.Sprintf("No results for %s", querystr), nil
	}

	output := fmt.Sprintf("%s | %s",
		results.D.Results[resNo].Title,
		results.D.Results[resNo].URL)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^b(?:ing)?([0-9])? (.+)$",
		bing)

	bot.RegisterCommand(
		"^g(?:oogle)?([0-9])? (.+)$",
		bing)
}
