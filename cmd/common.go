package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func CopyFile(src string, dst string) error {

	source, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	fmt.Println("File copied successfully.")

	return nil
}

func RunCommand(cmd string, argsAndOptions ...interface{}) {
	printOutput := true
	var args []string

	for _, arg := range argsAndOptions {
		switch v := arg.(type) {
		case bool:
			printOutput = v
		case string:
			args = append(args, v)
		default:
			log.Fatal("Invalid argument type")
		}
	}

	command := exec.Command(cmd, args...)
	stderr, err := command.StderrPipe()
	if err != nil {
		log.Fatal(err)
		return
	}

	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := command.Start(); err != nil {
		log.Fatal(err)
		return
	}

	if printOutput {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		scanner = bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

	if err := command.Wait(); err != nil {
		log.Fatal(err)
	}
}

func RemoveIndex[K any](s []K, index int) []K {
	return append(s[:index], s[index+1:]...)
}

type Template struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	GitUrl      string `json:"git_url"`
	CreatedAt   string `json:"created_at"`
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error making request: \n\t", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading body: \n\t", err)
	}
	return data, nil
}
