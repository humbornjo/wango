package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/humbornjo/wango/pkg/latea"
	"github.com/urfave/cli/v2"
)

var ()

func Run() {
	app := &cli.App{
		Name:  "wango",
		Usage: "wang's tile artwork generator impl by go",
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:        "seed",
				Aliases:     []string{"s"},
				Usage:       "seed for wango",
				DefaultText: "3407",
			},
		},
		Action: func(*cli.Context) error {
			p := tea.NewProgram(
				latea.InitModel(),
				tea.WithAltScreen(),
				tea.WithMouseCellMotion(),
			)
			if _, err := p.Run(); err != nil {
				log.Fatalf("Comrade, we got an error: %v", err)
				os.Exit(1)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	if latea.Err != nil {
		fmt.Printf("failed: %v", latea.Err)
		return
	}
	if latea.Success {
		fmt.Printf("success: saved as %v", latea.Path)
		return
	}
	fmt.Println("till we meet again, ciallo～(∠・ω< )⌒★")
}
