package dictionary 

type Wordnik struct {
		// TextProns []interface{} `json:"textProns"`	// I could leave these in but they appear to return nothing all the time
		// SourceDictionary string `json:"sourceDictionary"`
		// ExampleUses []interface{} `json:"exampleUses"`
		// RelatedWords []interface{} `json:"relatedWords"`
		// Labels []interface{} `json:"labels"`
		// Citations []interface{} `json:"citations"`
		Word string `json:"word"`
		// Sequence string `json:"sequence"`	// this is out because why would you have a count/sequence indicator as a string smdh
		PartOfSpeech string `json:"partOfSpeech"`
		// AttributionText string `json:"attributionText"`	// dont care
		Text string `json:"text"`
		// Score float64 `json:"score"` // dont care
}

type Wotd struct {
	ID int `json:"id"`
	Word string `json:"word"`
	PublishDate string `json:"publishDate"`
	ContentProvider struct {
		Name string `json:"name"`
		ID int `json:"id"`
	} `json:"contentProvider"`
	Note string `json:"note"`
	Examples []struct {
		URL string `json:"url"`
		Text string `json:"text"`
		ID int `json:"id"`
		Title string `json:"title"`
	} `json:"examples"`
	Definitions []struct {
		Text string `json:"text"`
		PartOfSpeech string `json:"partOfSpeech"`
		Source string `json:"source"`
	} `json:"definitions"`
}