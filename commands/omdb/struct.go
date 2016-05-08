package omdb

type Omdb struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	Rated string `json:"Rated"`
	Released string `json:"Released"`
	Runtime string `json:"Runtime"`
	Genre string `json:"Genre"`
	Director string `json:"Director"`
	Writer string `json:"Writer"`
	Actors string `json:"Actors"`
	Plot string `json:"Plot"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Awards string `json:"Awards"`
	Poster string `json:"Poster"`
	Metascore string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes string `json:"imdbVotes"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	TomatoMeter string `json:"tomatoMeter"`
	TomatoImage string `json:"tomatoImage"`
	TomatoRating string `json:"tomatoRating"`
	TomatoReviews string `json:"tomatoReviews"`
	TomatoFresh string `json:"tomatoFresh"`
	TomatoRotten string `json:"tomatoRotten"`
	TomatoConsensus string `json:"tomatoConsensus"`
	TomatoUserMeter string `json:"tomatoUserMeter"`
	TomatoUserRating string `json:"tomatoUserRating"`
	TomatoUserReviews string `json:"tomatoUserReviews"`
	TomatoURL string `json:"tomatoURL"`
	DVD string `json:"DVD"`
	BoxOffice string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website string `json:"Website"`
	Response string `json:"Response"`
}