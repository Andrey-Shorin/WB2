package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("shell> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Usage: kill <PID>")
				continue
			}
			cmd := exec.Command("taskkill", "/PID", args[1], "/F")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "ps":
			cmd := exec.Command("tasklist")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		case "exit":
			os.Exit(0)
		default:
			doCommands(strings.Split(input, "|"))
		}
	}
}

func doCommands(conveyor []string) {

	cmds := make([]*exec.Cmd, 0, len(conveyor))

	for _, v := range conveyor {
		v = strings.TrimSpace(v)
		command := strings.Split(v, " ")

		cmds = append(cmds, exec.Command(command[0], command[1:]...))
	}

	if len(cmds) == 1 {
		cmds[0].Stdin = os.Stdin
		cmds[0].Stdout = os.Stdout
		cmds[0].Run()
		return
	}

	output, stdErr, err := Pipeline(cmds)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, output)
	io.Copy(os.Stderr, stdErr)

}

func Pipeline(cmds []*exec.Cmd) (*bytes.Buffer, *bytes.Buffer, error) {

	var output bytes.Buffer
	var stderr bytes.Buffer

	cmds[0].Stdin = os.Stdin

	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		var err error
		cmds[i+1].Stdin, err = cmd.StdoutPipe()
		if err != nil {
			return nil, nil, err
		}

		cmd.Stderr = &stderr
	}

	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			return nil, nil, err
		}
	}

	for _, cmd := range cmds {
		cmd.Wait()
	}

	return &output, &stderr, nil
}
