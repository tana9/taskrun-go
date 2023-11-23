package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"os/exec"
)

func msg(err error) int {
	scanner := bufio.NewScanner(os.Stdin)
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", string(exitErr.Stderr))
			fmt.Println(exitErr.ExitCode())
			return exitErr.ExitCode()
		}
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		fmt.Println("エラーが発生しました")
		scanner.Scan()
		return 1
	}
	fmt.Println("完了しました")
	scanner.Scan()
	return 0
}

func taskDir(args []string) string {
	if 2 <= len(args) {
		return args[1]
	}
	return "."
}

func run(args []string) error {
	dir := taskDir(args)
	task := TaskRunner{Dir: dir}
	tasks, err := task.ListAll()
	if err != nil {
		return err
	}
	prompt := promptui.Select{Label: "Select Task", Items: tasks}
	i, _, err := prompt.Run()
	if err != nil {
		return err
	}
	if err := task.Run(tasks[i]); err != nil {
		return err
	}
	return nil
}

func main() {
	msg(run(os.Args))
}
