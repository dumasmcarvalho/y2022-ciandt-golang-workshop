package main

import (
	"fmt"
	"log"
	"net/http"
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
		fmt.Printf("Arquivo não existe")
	}

	return err
}

func checkAPI(ctx *cli.Context) error {
	endpoint := ctx.String("endpoint")

	if endpoint != "" {
		resp, err := http.Get(endpoint)

		if err != nil {
			return err
		}

		fmt.Println("Código do status da requisição:", resp.StatusCode)
	}

	return nil
}

func aptUpdate(ctx *cli.Context) error {
	fmt.Println("Serão instalados alguns pacotes no seu ambiente...")
	fmt.Println("Caso necessário, informe sua senha de super usuário para que a instalação possa prosseguir.")
	fmt.Println()
	bashCmd := `sudo apt-get update`
	cmd := exec.Command("bash", "-c", bashCmd)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("%s", string(output))

	return nil
}
