package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

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
