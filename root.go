package really

import (
	"github.com/go-resty/resty/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
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

var (
	client *resty.Client
	db     *gorm.DB
)

func Run() {
	db = initDB()
	client = initRestyClient()

	startOpenDota()
	serveMirai()
}

// initRestyClient 初始化 http client
func initRestyClient() *resty.Client {
	return resty.New()
}

// startOpenDota 定时任务相关逻辑
func startOpenDota() {
	if err := db.First(&Hero{}).Error; err != nil {
		InitHeros()
	}

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		SubscribeFunc()
	})

	c.Start()
}
