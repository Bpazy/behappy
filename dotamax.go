package really

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile("\\d+")

func loadRecentMatches(client *resty.Client, db *gorm.DB) func() {
	return func() {
		playerId := "122155653"
		playerDetailRes, err := client.R().Get(fmt.Sprintf("http://dotamax.com/player/detail/%s/", playerId))
		if err != nil {
			log.Printf("获取用户详情失败: %+v, 尝试重新登录\n", err)
			client = loginDotaMax()
			return
		}
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(playerDetailRes.String()))
		if err != nil {
			log.Printf("解析用户详情 DOM 失败: %+v\n", err)
			return
		}

		matchPlayers := getMatchPlayers(dom, dom.Find(".table-player-detail"), playerId)

		for _, mp := range matchPlayers {
			s := map[string]interface{}{
				"match_id":  mp.MatchId,
				"player_id": playerId,
			}
			var matchPlayer MatchPlayer
			db.Where(s).First(&matchPlayer)
			if matchPlayer.ID == 0 {
				// 新比赛
				log.Printf("探测到新的比赛：%s", mp)
				db.Create(mp)
			}
		}
	}
}

func getMatchPlayers(dom *goquery.Document, playerDetails *goquery.Selection, playerId string) []*MatchPlayer {
	// Get(0): 常用英雄
	// Get(1): 最近比赛
	// Get(2): 最高记录
	s := dom.FindNodes(playerDetails.Get(1))

	var matchPlayers []*MatchPlayer
	s.Find("tr").Each(func(i int, cs *goquery.Selection) {
		// 每一场比赛
		var lines []string
		cs.Find("td").Each(func(i int, cs2 *goquery.Selection) {
			lines = append(lines, strings.TrimSpace(cs2.Text()))
		})
		hero := lines[0]
		matchId := r.FindString(lines[1])
		matchMode := strings.TrimSpace(strings.SplitAfter(lines[1], matchId)[1])
		result := lines[3]
		kda := lines[4]
		level := lines[5]

		matchPlayer := &MatchPlayer{
			MatchId:     matchId,
			PlayerId:    playerId,
			Hero:        hero,
			MatchMode:   matchMode,
			MatchResult: result,
			MatchKDA:    kda,
			MatchLevel:  level,
		}
		matchPlayers = append(matchPlayers, matchPlayer)
	})
	return matchPlayers
}

func loginDotaMax() *resty.Client {
	config := InitConfig()
	client, u := initRestyClient(config)

	// 如果有 Cookie 则跳过登录
	if len(client.GetClient().Jar.Cookies(u)) != 0 {
		log.Printf("使用 Cookie 登录\n")
		return client
	}

	getLoginPageRes, err := client.R().Get("http://dotamax.com/login/")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"csrfmiddlewaretoken": getCsrfToken(getLoginPageRes.String()),
		"phoneNumCipherb64":   encrypt(""),
		"usernameCipherb64":   encrypt(config.DotaMax.Username),
		"passwordCipherb64":   encrypt(config.DotaMax.Password),
		"account-type":        "2",
		"src":                 "None",
	}
	loginRes, err := client.R().
		SetFormData(data).
		Post("http://dotamax.com/accounts/login/")
	if err != nil {
		panic(err)
	}

	body := loginRes.String()
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	loginReply := dom.Find(".login-reply").Text()
	if loginReply != "" {
		log.Fatalf("登录 DotaMax 失败: %s\n", loginReply)
	}
	if strings.Contains(body, "随机征召") {
		log.Printf("登录 DotaMax 成功\n")
	}

	config.SetCookies(loginRes.Cookies())
	return client
}

func getCsrfToken(body string) string {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	csrfToken, exists := dom.Find("[name=csrfmiddlewaretoken]").Attr("value")
	if !exists {
		panic("未找到 csrfmiddlewaretoken")
	}
	return csrfToken
}

func encrypt(content string) string {
	rsaE := "10001"
	rsaN := "B81E72A33686A201B0AC009D679750990E3D168670DC6F9452C24E5A4C299AC002C6C89C3CB38784AEA95D66B7B3E9CA950EB9EEFB4EF61383EDDB67C35727F9CA87EE3238346C66D042B35812179501F472AD4F3BA19E701256FE0435AB856E5C5BEA24A2387153023CD4CD43CDA7260FCC1E2E49C14102C253F559F9A45D59DF5004A017B1239448A9A001D276CAD12535DEDE89FFBD57D75BBC9B575530DDD1B7FAD46064AD3C640CBD017F58981215B2EE17CBE175C36570C5235902818648577234E70E81133B088164F98E605D0D6E69A6095A32A72511E9AC901727B635CE2E8002A7B0EC8D012606903BCB825E60C7B6619FFCED4401E693F5EC68AB"

	n := new(big.Int)
	n, ok := n.SetString(rsaN, 16)
	if !ok {
		panic("public key should be hexadecimal")
	}

	hexRsaE, err := strconv.ParseInt(rsaE, 16, 64)
	if err != nil {
		panic(err)
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, &rsa.PublicKey{
		N: n,
		E: int(hexRsaE),
	}, []byte(content))
	if err != nil {
		panic(err)
	}
	return linebrk(base64.StdEncoding.EncodeToString(encryptedData), 64)
}

func linebrk(s string, n int) string {
	var ret = ""
	var i = 0
	for i+n < len(s) {
		ret += s[i:i+n] + "\n"
		i += n
	}
	return ret + s[i:]
}
