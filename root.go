package behappy

import (
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/qq"
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
		Short: "Don't worry,be happy",
		Long: `风力掀天浪打头，只须一笑不须愁
`,
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "运行 behappy 主程序",
		Long:  `运行 behappy 主程序`,
		Run: func(cmd *cobra.Command, args []string) {
			if config.DebugMode {
				logrus.SetLevel(logrus.DebugLevel)
			}
			Run()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "查看 behappy 版本号",
		Long:  `查看 behappy 版本号`,
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Info(buildVer)
		},
	}
)

func Execute() error {
	rootCmd.PersistentFlags().BoolVar(&config.DebugMode, "debug", false, "Debug Mode")
	runCmd.Flags().StringVar(&config.Addr, "addr", "0.0.0.0:10000", "监听 Mirai 事件上报的端口号")

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)

	return rootCmd.Execute()
}

func Run() {
	config.InitConfig()
	dao.InitDB()

	startOpenDota()

	//qq.UploadGroupImage()

	qq.ServeMirai()
}

// startOpenDota 定时任务相关逻辑
func startOpenDota() {
	logrus.Info("启动定时任务")

	if !dao.HasHeroData() {
		InitHeros()
	}

	c := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger), // or use cron.DefaultLogger
	))
	c.AddFunc("@every 5m", func() {
		SubscribeFunc()
	})

	//_, err := c.AddFunc("30 14 * * *", TellJoke)
	//if err != nil {
	//	panic(err)
	//}

	c.Start()
}
