package bing

import (
	"encoding/json"
	"fmt"
	"github.com/Seventy-Two/Cara"
	"github.com/dustin/go-humanize"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	imageURL = "https://api.datamarket.azure.com/Bing/Search/v1/Image?Query='%s'&$format=json"
)

type ImageResults struct {
	D struct {
		Results []struct {
			MediaURL    string `json:"MediaUrl"`
			FileSize    string `json:"FileSize"`
			ContentType string `json:"ContentType"`
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func image(command *bot.Cmd, matches []string) (msg string, err error) {
	var resNo int
	var querystr string
	if len(matches) > 1 {
		resNo, _ = strconv.Atoi(matches[1])
		querystr = matches[2]
	} else {
		querystr = matches[1]
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf(imageURL, url.QueryEscape(querystr)), nil)
	request.SetBasicAuth("", bot.Config.TranslateClient)

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var results ImageResults
	json.Unmarshal(body, &results)

	if err != nil {
		return fmt.Sprintf("No results for %s", querystr), nil
	}

	if len(results.D.Results) == 0 {
		return fmt.Sprintf("No results for %s", querystr), nil
	}

	size, _ := strconv.ParseUint(results.D.Results[0].FileSize, 10, 64)
	humanize.Bytes(size)

	output := fmt.Sprintf("%s → %s %s | %s", querystr,
		results.D.Results[resNo].ContentType,
		humanize.Bytes(size),
		results.D.Results[resNo].MediaURL)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^img([0-9])? (.+)$",
		image)
}
