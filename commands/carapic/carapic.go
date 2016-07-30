package carapic

import (
	"fmt"
	"github.com/Seventy-Two/Cara"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func carapic(command *bot.Cmd, matches []string) (msg string, err error) {
	choices := []string{
		"b5T0uce",
		"bdsVc8H",
		"wawK4bi",
		"VjTVwgP",
		"UsZlip0",
		"GhHinMp",
		"CBaSMf9",
		"8dRIVd5",
		"xlZ80JW",
		"lTSGw3F",
		"kSbQQ8R",
		"sfx9FIP",
		"qSl4Ge3",
		"rdVmW1e",
		"6zKd5up",
		"gHTg6Fk",
		"xD3uqhz",
		"wA08hec",
		"xH0Cagd",
		"QKqtO8M",
		"P3PxFvi",
		"1gO8IeU",
		"WLsfQdC",
		"CGthgNJ",
		"earBdbj",
		"vBXyGD6",
		"C90AdOF",
		"HfdYHPO",
		"yuE19gq",
		"P7WTkGD",
		"FKx6Cl6",
		"KfEPdoI",
		"7tqkw1w",
		"AleOKpj",
		"R1nftaJ",
		"Y0g4ELR",
		"5V2dFtM",
		"jBxFuXv",
		"Nr9ocv9",
		"sl71IkK",
		"7c58mud",
		"tiHWFlY",
		"GwyUaGd",
		"n11sxUf",
		"yfetUM0",
		"4ujnFee",
		"A4JTpYU",
		"WPHTqwm",
		"qwK7IaZ",
		"4Ji2fcf",
		"HFmpVc9",
		"WMP0MYo",
		"lACHqyM",
		"iG3SlNR",
		"nZSUBPP",
		"zP5b7Bh",
		"ZJn1b51",
		"L5MazWh",
		"opk1Aas",
		"Qs8ENyS",
		"4mDH1uR",
		"C6hDXIv",
		"TP5zbUo",
		"uAhmg1h",
		"pWjA2KR",
		"rsYLudm",
		"wz3HIsu",
		"RKvwVYP",
		"3bn1IRT",
		"CcxNT7D",
		"ljemZ7x",
		"al6Tn8M",
		"EmWP98o",
		"ZTmCWXH",
		"wiMScVA",
		"VRsGkbY",
		"LhbPfFe",
		"Qmu6aor",
		"L950O9q",
		"xGLyc4f",
	}
	chosen := random(0, len(choices))
	msg = fmt.Sprintf("%s: http://i.imgur.com/%s.jpg", command.Nick, choices[chosen])
	return
}

func init() {
	bot.RegisterCommand(
		"^carapic$",
		carapic)
}
