# Cara

An IRC Bot written in Go - All credit to 0x263b, Cara is just a clowny Porygon2 with poorly implemented changes for a bunch of idiots

### Install
```
# Dependencies
go get -u github.com/thoj/go-ircevent github.com/steveyen/gkvlite github.com/PuerkitoBio/goquery github.com/dustin/go-humanize github.com/kennygrant/sanitize gopkg.in/xmlpath.v2 github.com/kurrik/oauth1a github.com/kurrik/twittergo

# Cara
go get github.com/Seventy-Two/Cara

#### APIs that require keys

* [Bing](https://datamarket.azure.com/dataset/bing/search): [commands/bing](commands/bing/images.go)
* [Giphy](https://github.com/Giphy/GiphyAPI): [commands/giphy](commands/giphy/giphy.go)
* [Google Geocode](https://developers.google.com/maps/documentation/geocoding/intro): [commands/weather](commands/weather/weather.go)
* [Last.fm](http://www.last.fm/api): [commands/lastfm](commands/lastfm/)
* [Forecast.io](https://developer.forecast.io/docs/v2): [commands/weather](commands/weather/weather.go)
* [Microsoft Translator](https://msdn.microsoft.com/en-us/library/hh454949.aspx): [commands/translate](commands/translate/translate.go)
* [Twitter](https://dev.twitter.com/rest/public): [commands/twitter](commands/twitter/twitter.go)
* [Wolfram Alpha](http://products.wolframalpha.com/api/): [commands/wolfram](commands/wolfram/wolfram.go)
* [Youtube](https://developers.google.com/youtube/v3/): [commands/youtube](commands/youtube/youtube.go)
* [Dota2](https://steamcommunity.com/dev/apikey): [commands/dota](commands/dota/dota.go)
* [Omdb](http://omdbapi.com/): [commands/omdb](commands/omdb/omdb.go)
* [Football/Soccer](http://api.football-data.org/): [commands/divegrass](commands/divegrass/divegrass.go)
* [Wordnik](http://api.wordnik.com): [commands/dictionary](commands/dictionary/dictionary.go)

***

### Functions

* [8ball](#8ball)
* [Bing](#bing)
* [Carapic](#carapic)
* [Dictionary](#dictionary)
* [DotA2](#DotA2)
* [Football/Soccer](#divegrass)
* [Lastfm](#lastfm)
* [Omdb](#omdb)
* [Random](#random)
* [Roll](#roll)
* [Stocks](#stocks)
* [Translate](#translate)
* [TVMaze](#tvmaze)
* [Twitter](#twitter)
* [Urban Dictionary](#urban-dictionary)
* [URL Parser](#url-parser)
* [User Profiles](#user-profiles)
* [Weather](#weather)
* [WolframAlpha](#wolframalpha)
* [Youtube](#youtube)

* [Admin functions](#admin-functions)


***

### 8Ball
Gives and 8ball style answer to a *question*

**Cara,** *question?*

	Cara, Am I going to score with this one girl I just finished talking to?
	My sources say no


### Bing
Gets the first result from [Bing](https://www.bing.com/) for *search query*

**.g/.google** *search query*

	.google Richard Stallman
	Google | Richard Stallman's Personal Page | http://stallman.org/

**.b/.bing** *search query*

	.bing Richard Stallman
	Bing | Richard Stallman's Personal Page | http://stallman.org/

Gets the first result from [Bing image search](https://www.bing.com/images) for *search query*

**.img** *search query*

	.img Richard Stallman
	Bing | Richard Stallman → image/jpeg 257 kB | http://www.straferight.com/photopost/data/500/richard-stallman.jpg

### Carapic
Returns a randomly selected image of Cara Delevingne from Big Dave's personal collection.

**.carapic**

### Dictionary 
Returns the word of the day from Wordnik

**.word/.wotd**

Returns the Wordnik dictionary results (up to 3) for the given query

**.dict** *search query*

## Divegrass
Returns the upcoming games for the given number of days

**.f n** *1-9*

Returns the scores of the games from the past number of days

**.f p** *1-9*

Simulates behaviour of n1

**.f**

## DotA2 
Returns information on the current games being played. For tier 3 (Premium) games, games with more than 200 viewers are returned. For tier 2 (Professional) games, games with more than 1000 viewers are returned.

**.d2/dota**

Returns heroes picked, along with the above

**.d2h** 

Returns scores along with the game data

**.d2s**

Returns all information 

**.d2hs**

### Last.fm
Associates your current irc nick with *user*.
Other lastfm functions will default to this nick if no user is provided.

**.set lastfm** *user*
	
	<joebloggs> .set lastfm JosefBloggs
	<Cara> joebloggs: last.fm user updated to: JosefBloggs
 

Weekly stats for *user*

**.charts** *user*

	.charts Cbbleh
	Last.fm | Top 5 Weekly artists for Cbbleh | Slayer (26), Iced Earth (25), Jean-Féry Rebel (23), Morbid Saint (15), Judas Priest (14)


Returns the currently playing/last scrobbled track for *user* and top artist tags

**.np** *user*
	
	.np Cbbleh
	Last.fm | cbbleh is playing: "Super X-9" by Daikaiju from Daikaiju | Surf, surf rock, instrumental, instrumental surf rock

Returns the currently playing track for the users in the channel that have set a lastfm account

**.wp**

	.wp 

Sets a nickname for the wp command so that you are not highlighted during wp

**.set nick** *user*

### Omdb
Returns tags, imdb + rt ratings, and short descriptions of the given query

**.m/.movie** *search query* 

### Random
Randomly picks an option from an array separated by |

**.rand** `one | two | three`

	.r do work | don't do work
	don't do work

### Roll
Rolls the given number of d10s

**.r/.roll** *0-99*

Rolls the given number of dice, of the given number of sides

**.r/.roll** *0-99*d*0-99*

## Stocks
Returns the current ask price, and the current change in % and USD from the NYSE of the given query. Query format must be a NYSE Symbol.

**.s/.stocks** *Query*

### TVMaze
Info for *tv show* with episode airtime if available **-tv** *tv show*

	-tv Better call saul
	TVmaze | Better Call Saul | Airtime: Monday 22:00 on AMC | Status: Running | Next Ep: S2E6 at 22:00 2016-03-21
	
	-tv Mr Robot
	TVmaze | Mr. Robot | Airtime: Wednesday 22:00 on USA Network | Status: Running


### Twitter
Latest tweet for *user* **.tw/.twitter** *user*

	.twitter Guardian
	Twitter | The Guardian (@guardian) | Aston Villa target Rémi Garde after sacking Tim Sherwood https://t.co/cqcgUpiEOJ via @guardian_sport | 31 seconds ago


### Urban Dictionary
Gets the first definition of *query* at [UrbanDictionary](http://www.urbandictionary.com/)

**.u/.ur/.urban** *query*

	.urban 4chan
	Urban Dictionary | 4chan | http://mnn.im/upucr | you have just entered the very heart, soul, and life force of the internet. this is a place beyond sanity, wild and untamed. there is nothing new here. "new" content on 4chan is not found; it is created from old material. every interesting, offensive, shoc…


Gets the *n*th definition for *query* (only works for definitions 1-7)

**.u/.ur/.urban** *n* *query*

	.urban 3 4chan
	UrbanDictionary | 4chan | 4chan.org is the absolute hell hole of the internet, but still amusing. Entering this website requires you leave your humanity behind before entering. WARNING: You will see things on /b/ that you wish you had never seen in your life.


### URL Parser
Returns the title of a page and the host for html URLs.
Returns the type, size, and (sometimes) filename of a file URL.

	https://news.ycombinator.com/
	Title | Hacker News | news.ycombinator.com

	https://41.media.tumblr.com/bca28cbcbba3718cd67fd20062df19b9/tumblr_nl8gekhnLU1tdhimpo1_1280.png
	File | image/png 272kB | 41.media.tumblr.com


### User Profiles
Returns the set variables for a *user*

	.whois qb
	qb | Twitter: @abnormcore | URL: https://dribbble.com/qb
	
Variables are set using **.set url** *url* or **.set twitter** *handle*

	.set twitter someone
	twitter updated to: someone

	.set url http://www.something.com/
	url updated to: http://www.something.com/


### Weather
[Yahoo Weather](http://weather.yahoo.com/) for *location*
**.w/.we/.weather** *location*

	.weather Washington, DC
	Weather | Washington | Cloudy 15°C. Wind chill: 15°C. Humidity: 72%

[Yahoo Weather Forecast](http://weather.yahoo.com/) for *location*
**.f/.fo/.forecast** *location*

	.forecast Washington, DC
	Forecast | Washington | Sun: Clouds Early/Clearing Late 16°C/10°C | Mon: Mostly Sunny 19°C/8°C | Tue: Mostly Sunny 23°C/11°C | Wed: Partly Cloudy 24°C/11°C
	
Associates your current irc nick with *location*.
Other weather functions will default to this location if none is provided.

**.set location** *location* 

	<joebloggs> .set location Washington, DC
	<Cara> joebloggs: location updated to: Washington, DC


### WolframAlpha
Finds the answer of *question* using [WolfarmAlpha](http://www.wolframalpha.com/)

**.wa** *question*

	.wa time in Bosnia
	Wolfram | current time in Bosnia and Herzegovina >>> 12:55:38 pm CEST | Tuesday, October 6, 2015


### Youtube
Gets the first result from [Youtube](https://www.youtube.com) for *search query* 

**.yt/.youtube** *search query*

	.yt Richard Stallman interject
	YouTube | I'd just like to interject... | 3m1s | https://youtu.be/QlD9UBTcSW4

***

### Admin functions
These functions are limited to bot admins and can only be used in a private message.
	
Ignore a user

**.set ignore** *nick*

	.set ignore Cbbleh
	<Cara> I never liked him anyway
	
Unignore a user

**.set unignore** *nick*

	.set unignore Cbbleh
	<Cara> Sorry about that
	
Toggles the URL parser for the channel

**.set urls on/off** *channel*

	.set urls on #lobby
	<Cara> Now reacting to URLs in #lobby
	
Toggles the file URL parser for the channel

**.set files on/off** *channel*

	.set files on #lobby
	<Cara> No longer displaying file info in #lobby	
Joins a channel and adds it to auto join

**.join** *channel*

	.join #foobar
	* Cara has joined #foobar
	
Parts a channel and removes it from auto join

**.part** *channel*

	.part #foobar
	* Cara has left the channel

