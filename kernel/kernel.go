package kernel

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-skeleton/cmd/handlers/cli"
	"go-skeleton/cmd/handlers/http"
	"go-skeleton/pkg"
	"time"
)

type kernel struct {
	rootCmd *cobra.Command
}

func NewKernel() *kernel {
	k := &kernel{}
	k.rootCmd = &cobra.Command{
		Use:   "go-skeleton",
		Short: "",
		Long:  ``,
		Run:   k.RootCmd,
	}
	return k
}

func (k *kernel) Start() error {
	return k.rootCmd.Execute()
}

func (k *kernel) Boot() {
	for index, dep := range pkg.KernelDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[kernel.Kernel] Booting %s", index))
	}

	pkg.Logger.Info("[kernel.Kernel] Booting application!")
	timezoneLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		pkg.Logger.Critical(err, "[kernel.Kernel] Unable to identify timezone")
	}
	time.Local = timezoneLocation

	k.rootCmd.AddCommand(
		&cobra.Command{
			Use:               "http",
			Short:             "Start a http server (API)",
			Long:              ``,
			Run:               k.BootServer,
			PersistentPostRun: k.startServer,
		},
	)

	cliCmd := &cobra.Command{
		Use:   "cli",
		Short: "",
		Long:  ``,
	}
	k.startCli(cliCmd)
	k.rootCmd.AddCommand(cliCmd)
}

func (k *kernel) RootCmd(cmd *cobra.Command, args []string) {
	pkg.Logger.Info(fmt.Sprintf("Go Skeleton Version %v", pkg.Config.Version))
	pkg.Logger.Info("Use --help to check witch commands are available")
}

func (k *kernel) startServer(cmd *cobra.Command, args []string) {
	http.NewServer(pkg.Config.Environment).Start(":1323")
}

func (k *kernel) startCli(cmd *cobra.Command) {
	cli.NewCli(pkg.Config.Environment).RegisterCommands(cmd)
}

func (k *kernel) BootServer(cmd *cobra.Command, args []string) {
	for index, dep := range pkg.ServerDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[kernel.Kernel] Booting %s", index))
	}
	pkg.Logger.Info("[kernel.Kernel] Done!")
}
