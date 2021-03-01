package behappy

import (
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/dao"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// buildVer represents 'behappy' build version
	buildVer string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "behappy",
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
		Long:  `查看 behappy 的版本号`,
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Info(buildVer)
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
	config.InitConfig()
	dao.InitDB()

	logrus.Info("启动定时任务")
	startOpenDota()

	logrus.Info("启动 Mirai HTTP 监听器")
	serveMirai()
}

// startOpenDota 定时任务相关逻辑
func startOpenDota() {
	if !dao.HasHeroData() {
		InitHeros()
	}

	c := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger), // or use cron.DefaultLogger
	))
	c.AddFunc("@every 5m", func() {
		SubscribeFunc()
	})

	c.AddFunc("0 0 14 * * ?", TellJoke)
	//c.AddFunc("@every 3s", TellJoke)

	c.Start()
}
