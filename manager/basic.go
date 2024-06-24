package manager

import (
	"github.com/spf13/cobra"
	"react-scaff.com/temps"
)

var (
	basicstruct = &cobra.Command{
		Use:   "basic",
		Short: "generate basic folder structure for react vite with daisy uiproject",
		Long:  `Generate basic folder structure for your react with daisy UI and tailwind configured.`,
		Run: func(cmd *cobra.Command, args []string) {
			basiccmd()
		},
	}
)

func basiccmd() {
	temps.CommonFrame()
	temps.ReactFrame()
}

func init() {
	goFrame.AddCommand(basicstruct)

}
