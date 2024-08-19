package cmd

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/humbornjo/wango/pkg/config"
	"github.com/humbornjo/wango/pkg/latea"
	"github.com/urfave/cli/v2"
)

var (
	seed    uint
	palette string
)

func Run() {
	app := &cli.App{
		Name:  "wango",
		Usage: "wang's tile artwork generator impl by go",
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:        "seed",
				Aliases:     []string{"s"},
				Usage:       "seed for wango",
				Value:       3407,
				DefaultText: "3407",
				Destination: &seed,
			},
			&cli.StringFlag{
				Name:        "palette",
				Aliases:     []string{"p"},
				Usage:       "palette for some shader",
				Value:       "#ff0000;#00ffff",
				DefaultText: "#ff0000;#00ffff",
				Destination: &palette,
			},
		},
		Action: func(*cli.Context) error {
			p := tea.NewProgram(
				latea.InitModel(),
				tea.WithAltScreen(),
				tea.WithMouseCellMotion(),
			)

			config.Rng = rand.New(rand.NewSource(int64(seed)))
			for _, hex := range strings.Split(palette, ";") {
				clr, err := latea.ParseColor(hex, color.RGBA{})
				if err != nil {
					return err
				}
				config.Palette = append(config.Palette, &clr)
			}

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
	fmt.Println("\ntill we meet again, ciallo～(∠・ω< )⌒★")
}
