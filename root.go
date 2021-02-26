package really

import (
	"github.com/Bpazy/really/dao"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"runtime/debug"
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

	c := cron.New()
	c.AddFunc("@every 5m", func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("run time panic: %s", string(debug.Stack()))
			}
		}()

		SubscribeFunc()
	})

	c.Start()
}
