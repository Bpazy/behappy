package really

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"log"
	"math/big"
	"net/http/cookiejar"
	"strings"
)

var (
	// buildVer represents 'really' build version
	buildVer string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "really",
		Short: "TODO",
		Long: `TODO
`,
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "版本号",
		Long:  `查看 really 的版本号`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(buildVer)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func Run() {
	go func() {
		client := resty.New()
		jar, err := cookiejar.New(nil)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		client.SetCookieJar(jar)

		getLoginPageRes, err := client.R().Get("http://dotamax.com/login/")
		if err != nil {
			panic(err)
		}

		body := getLoginPageRes.String()
		fmt.Println(body)
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			log.Fatal(err)
		}
		csrfToken, exists := dom.Find("[name=csrfmiddlewaretoken]").Attr("value")
		if !exists {
			panic("未找到 csrfmiddlewaretoken")
		}

		data := map[string]string{
			"csrfmiddlewaretoken": csrfToken,
			"phoneNumCipherb64":   encrypt(""),
			"usernameCipherb64":   encrypt(""),
			"passwordCipherb64":   encrypt(""),
			"account-type":        "2",
			"src":                 "None",
		}
		loginRes, err := client.R().
			SetFormData(data).
			Post("http://dotamax.com/accounts/login/")
		if err != nil {
			panic(err)
		}

		fmt.Println(data)
		fmt.Println(loginRes.Header())
	}()

	serve()
}

func encrypt(content string) string {
	rsaE := 10001
	rsaN := "B81E72A33686A201B0AC009D679750990E3D168670DC6F9452C24E5A4C299AC002C6C89C3CB38784AEA95D66B7B3E9CA950EB9EEFB4EF61383EDDB67C35727F9CA87EE3238346C66D042B35812179501F472AD4F3BA19E701256FE0435AB856E5C5BEA24A2387153023CD4CD43CDA7260FCC1E2E49C14102C253F559F9A45D59DF5004A017B1239448A9A001D276CAD12535DEDE89FFBD57D75BBC9B575530DDD1B7FAD46064AD3C640CBD017F58981215B2EE17CBE175C36570C5235902818648577234E70E81133B088164F98E605D0D6E69A6095A32A72511E9AC901727B635CE2E8002A7B0EC8D012606903BCB825E60C7B6619FFCED4401E693F5EC68AB"

	n := new(big.Int)
	n, ok := n.SetString(rsaN, 16)
	if !ok {
		panic("public key should be hexadecimal")
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, &rsa.PublicKey{
		N: n,
		E: rsaE,
	}, []byte(content))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(encryptedData)
}

// serve 开启事件上报监听器
func serve() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/post", func(c *gin.Context) {
		all, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(all))
		c.JSON(200, nil)
	})
	log.Fatal(r.Run("0.0.0.0:10000"))
}

type FriendMessage struct {
	SessionKey   string         `json:"sessionKey"`
	Target       int            `json:"target"`
	MessageChain []PlainMessage `json:"messageChain"`
}

type PlainMessage struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

func NewPlainMessage(text string) PlainMessage {
	return PlainMessage{
		Type: "Plain",
		Text: text,
	}
}

func NewFriendMessage(target int, text string) *FriendMessage {
	return &FriendMessage{
		SessionKey: "",
		Target:     target,
		MessageChain: []PlainMessage{
			NewPlainMessage(text),
		},
	}
}
