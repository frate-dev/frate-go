package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func RunCommand(cmd string, args ...string) {
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


	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := command.Wait(); err != nil {
		log.Fatal(err)
	}
}

func RemoveIndex[K any](s []K, index int) []K {
    return append(s[:index], s[index+1:]...)
}
