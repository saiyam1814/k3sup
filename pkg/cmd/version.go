package cmd

import (
        "fmt"

        "github.com/spf13/cobra"
)

var (
        Version   string
        GitCommit string
)

func MakeVersion() *cobra.Command {
        var command = &cobra.Command{
                Use:          "version",
                Short:        "Print the version",
                Example:      `  k3sup version`,
                SilenceUsage: false,
        }
        command.Run = func(cmd *cobra.Command, args []string) {
                if len(Version) == 0 {
                        fmt.Println("Version: dev")
                } else {
                        fmt.Println("Version:", Version)
                }
                fmt.Println("Git Commit:", GitCommit)
        }
        return command
}

func printFiglet() {
        figletColoured := aec.BlueF.Apply(figletStr)
        if runtime.GOOS == "windows" {
                figletColoured = aec.GreenF.Apply(figletStr)
        }
        fmt.Printf(figletColoured)
}

const figletStr = `
 _    _____                 
| | _|___ / ___ _   _ _ __  
| |/ / |_ \/ __| | | | '_ \ 
|   < ___) \__ \ |_| | |_) |
|_|\_\____/|___/\__,_| .__/ 
                     |_|
`
