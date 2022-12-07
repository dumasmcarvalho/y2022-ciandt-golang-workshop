# Golang CLI

## Instalando GO

Abrir a pagina oficial do Go, selecionar seu SO e seguir os passos recomendados:

https://go.dev/doc/install

Caso tenha dificuldades, algumas alternativas a instalacao original:

### Windows

https://medium.com/@rafaelmoraisdev/como-instalar-go-no-windows-10-7787faac3a7f

### Linux

https://www.edivaldobrito.com.br/linguagem-go-no-linux/

### Mac

https://www.digitalocean.com/community/tutorials/como-instalar-o-go-e-configurar-um-ambiente-de-programacao-local-no-macos-pt

## Setup Inicial

Estrutura de arquivos

```
root
|------ main.go
```

Codigo inicial

```
package main

func main() {
    println("Hi")
}
```

## Executando codigo GO

```
go run .\main.go
```

## Gerando arquivo de modulos

```
go mod init github.com/IAPOLINARIO/demos/golang_noobs/mycli
go mod tidy
```

### Exemplo de arquivo go.mod

```
module github.com/IAPOLINARIO/demos/golang_noobs/noobcli

go 1.18

require github.com/urfave/cli/v2 v2.23.2

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

```

## Baixando Dependencias

```
go get github.com/urfave/cli/v2
```

### Importando Pacotes

Dentro do arquivo main.go:

```
import (
	"github.com/urfave/cli/v2"
)
```

## Criando a primeira CLI

Main.go

```
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name:  "ola",
        Usage: "Exibe a mensagem da sua primeira CLI",
        Action: func(*cli.Context) error {
            fmt.Println("Ola, minha primeira CLI!")
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

```

### Executando

```
go run .\main.go
```

## Flags

Main.go:

```
func main() {
    app := &cli.App{
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:  "name",
                Value: "default",
                Usage: "Adiciona flag com o nome do usuario",
            },
        },
        Action: func(ctx *cli.Context) error {
            username := ctx.String("name")
            if username != "" {
                fmt.Println("Ola", username)
            } else {
                fmt.Println("Flag --name nao foi informada", username)
            }
            return nil

        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

```

## Comandos

main.go

```
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Valida diferentes recursos",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Recurso validado ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Cria um novo recurso",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Novo recurso criado: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}
```

### Sub-comandos

```
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
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Completa uma tarefa da lista",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Tarefa completada: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}
```

## Actions

```
app := &cli.App{
	Commands: []*cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "Valida diferentes recursos",
			Action: runAction,
		},
	},
}


func runAction() error {
	cmd := exec.Command("node", "-v")

	output, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Printf("Node version: %v", string(output))

	return nil
}
```

## Executar um comando no SO

```
cmd := exec.Command("ls", "-la")

	output, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Printf("Output: %v", string(output))

	return nil
```

## Checar existencia de um arquivo

```
if _, err := os.Stat("./my_project/sample.txt"); err == nil {
		fmt.Printf("Arquivo existe")
	} else {
		fmt.Printf("Arquivo nao existe")
	}

	return err
```

## Projeto Atualizado

Estrutura de arquivos

```
root
|------ checkCommand.go
|------ go.mod
|------ go.sum
|------ main.go
```

main.go

```
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
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Cria um novo recurso",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Nova tarefa criada: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

```

checkCommand.go

```
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func checkEnv(ctx *cli.Context) error {
	err := checkNode()
	err = checkNPM()
	err = checkVSCode()

	return err
}

func checkNode() error {
	cmd := exec.Command("node", "-v")

	output, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Printf("Node version: %v", string(output))

	return nil
}

func checkNPM() error {
	cmd := exec.Command("npm", "-v")

	output, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Printf("NPM version: %v", string(output))

	return nil
}

func checkVSCode() error {

	return checkDependency("code", "-v")
}

func checkDependency(mainCommand string, args ...string) error {
	cmd := exec.Command(mainCommand, args...)

	output, err := cmd.Output()

	if err != nil {
		return err
	}

	color.Yellow(fmt.Sprintf("%v version: %v", mainCommand, strings.Split(string(output), "\n")[0]))

	return nil
}

func checkProject(ctx *cli.Context) (err error) {

	if _, err := os.Stat("./my_project/sample.txt"); err == nil {
		fmt.Printf("Arquivo existe")
	} else {
		fmt.Printf("Arquivo nao existe")
	}

	return err
}

```

## Check API Function

```
func checkAPI(ctx *cli.Context) error {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		return err
	}

	fmt.Println("The status code we got is:", resp.StatusCode)

	return nil
}

```

## Configuration Management

Package ini provides INI file read and write functionality in Go.

```
go get gopkg.in/ini.v1
```

config.ini:

```
env = dev

[local-dependencies]
npm-packages = angular,vue,yarn
vscode-plugins = golang.go,hashicorp.terraform,zsh.go-snippets

[project-configuration]
required-files = main.txt,config.txt
default-project-name = my-project-go

[api-health-check]
critical-endpoints = www.google.com,www.g1.com.br,www.uol.comm.br
customer-endpoint = www.twitter.com
sales-endpoint = www.tesla.com
checkout-endpoint = www.amazon.com
dev-endpoint = www.github.com
prod-endpoint = www.ciandt.com

[github-repos]
devops-team = https://github.com/IAPOLINARIO/100-days-of-code
cloud-team = https://github.com/IAPOLINARIO/100-days-of-code,https://github.com/webinstall/webi-installers,https://github.com/adrg/xdg
android-team = https://github.com/segmentio/kafka-go,https://github.com/godot-rust/gdextension
```

main.go:

```
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Falha ao ler arquivo de configuracao: %v", err)
		os.Exit(1)
	}

	fmt.Println("Working on env:", cfg.Section("").Key("env").String())
	fmt.Println("Checando dependencias:", cfg.Section("local-dependencies").Key("npm-packages").String())

	app := &cli.App{
	.....
	.....
	}
```

## Util - Check Local Dependencies

CheckCommand.go

```
func checkLocalDependencies(ctx *cli.Context, npmPackages []string, vsCodePlugins []string) error {

	for _, item := range npmPackages {
		fmt.Printf("Checking item %s \n", item)
	}

	for _, item := range vsCodePlugins {
		fmt.Printf("Checking item %s \n", item)
	}

	return nil
}

```
