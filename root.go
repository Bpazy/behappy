package really

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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
			log.Println(buildVer)
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
	db := initDB()

	client := loginDotaMax()

	startCron(client, db)

	serveMirai()
}

// initRestyClient 初始化 http client
func initRestyClient(c *configuration) (*resty.Client, *url.URL) {
	client := resty.New()

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	parseUrl, err := url.Parse("http://dotamax.com/")
	if err != nil {
		panic(err)
	}

	var cookies []*http.Cookie
	if c.DotaMax.Cookies != "" {
		err := json.Unmarshal([]byte(c.DotaMax.Cookies), &cookies)
		if err != nil {
			log.Fatalf("Cookie 格式错误: %+v", err)
		}
		jar.SetCookies(parseUrl, cookies)
	}
	client.SetCookieJar(jar)

	return client, parseUrl
}

// startCron 定时任务相关逻辑
func startCron(client *resty.Client, db *gorm.DB) {
	c := cron.New()
	c.AddFunc("@every 1m", loadRecentMatches(client, db))

	c.Start()
}
