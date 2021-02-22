package really

import (
	"errors"
	"fmt"
	"github.com/Bpazy/dotamax"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strings"
)

var r = regexp.MustCompile("\\d+")

func SubscribeFunc() {
	var sps []*SubscribePlayer
	if err := db.Find(&sps).Error; err != nil {
		log.Println("没有订阅的玩家")
	}

	for _, sp := range sps {
		loadRecentMatches(sp.PlayerId)
	}
}

func loadRecentMatches(playerId string) {
	playerDetailRes, err := client.R().Get(fmt.Sprintf("http://dotamax.com/player/detail/%s/", playerId))
	if err != nil {
		log.Printf("获取用户详情失败: %+v, 尝试重新登录\n", err)
		client = loginDotaMax(false)
		return
	}
	s := playerDetailRes.String()
	if strings.Contains(s, "社交账号登录") {
		log.Printf("获取用户详情失败: %+v, 尝试重新登录\n", err)
		client = loginDotaMax(false)
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(s))
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
		if err := db.Where(s).First(&MatchPlayer{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			// 新比赛
			log.Printf("探测到新的比赛：%s", mp)
			db.Create(mp)
		}
	}
}

func getMatchPlayers(dom *goquery.Document, playerDetails *goquery.Selection, playerId string) []*MatchPlayer {
	// Get(0): 常用英雄
	// Get(1): 最近比赛
	// Get(2): 最高记录
	if playerDetails.Length() == 0 {
		log.Println("查询玩家详情异常")
		return []*MatchPlayer{}
	}
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

func loginDotaMax(useCookie bool) *resty.Client {
	config := InitConfig()
	client, u := initRestyClient(config)

	// 如果有 Cookie 则跳过登录
	if useCookie && len(client.GetClient().Jar.Cookies(u)) != 0 {
		log.Printf("使用 Cookie 登录\n")
		return client
	}

	getLoginPageRes, err := client.R().Get("http://dotamax.com/login/")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"csrfmiddlewaretoken": getCsrfToken(getLoginPageRes.String()),
		"phoneNumCipherb64":   dotamax.Encrypt(""),
		"usernameCipherb64":   dotamax.Encrypt(config.DotaMax.Username),
		"passwordCipherb64":   dotamax.Encrypt(config.DotaMax.Password),
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
