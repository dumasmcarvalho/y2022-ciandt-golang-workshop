package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Valida diferentes recursos",
				Subcommands: []*cli.Command{
					{
						Name:   "env",
						Usage:  "Valida os requisitos do ambiente",
						Action: checkEnv,
					},
					{
						Name:   "project",
						Usage:  "Valida a estrutura de um projeto",
						Action: checkProject,
					},
				},
			},
			{
				Name:    "api",
				Aliases: []string{"a"},
				Usage:   "Valida uma api",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "endpoint",
						Aliases: []string{"ep"},
						Value:   "",
						Usage:   "Informa o endpoint a ser validado",
					},
				},
				Action: checkAPI,
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Cria um novo recurso",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Nova tarefa criada: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "apt",
				Aliases: []string{"a"},
				Usage:   "Realiza o apt-get update",
				Action:  aptUpdate,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
