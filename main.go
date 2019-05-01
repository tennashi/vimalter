package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/shirou/gopsutil/process"
)

var (
	tabFlg = flag.Bool("tab", false, "open file in new tab")
)

func main() {
	flag.Parse()
	args := flag.Args()

	var vim *exec.Cmd
	vimRuntime := os.Getenv("VIMRUNTIME")
	if vimRuntime == "" {
		vim = exec.Command("vim", args...)
	}

	vimEnv := filepath.Dir(vimRuntime)
	switch filepath.Base(vimEnv) {
	case "vim":
		srvName := os.Getenv("VIM_SERVERNAME")
		if srvName == "" {
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, "open new file in vim8(-clientserver) is not supported")
				os.Exit(1)
			}
			vim = exec.Command("echo", "-e", "\x1b]51;[\"drop\",\""+strings.Join(args, " ")+"\"]\x07")
			break
		}
		path, err := evalVimPath()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if len(args) == 0 {
			if *tabFlg {
				args = []string{"--servername", srvName, "--remote-send", `<C-\><C-N>:tabnew<CR>`}
			} else {
				args = []string{"--servername", srvName, "--remote-send", `<C-\><C-N>:new<CR>`}
			}
		} else {
			if *tabFlg {
				args = append([]string{"--servername", srvName, "--remote-tab"}, args...)
			} else {
				args = append([]string{"--servername", srvName, "--remote"}, args...)
			}

		}
		vim = exec.Command(path, args...)

	case "nvim":
		if *tabFlg {
			args = append([]string{"--remote-tab"}, args...)
		}
		vim = exec.Command("nvr", args...)
	}

	vim.Stdin = os.Stdin
	vim.Stdout = os.Stdout
	vim.Stderr = os.Stderr

	vim.Run()
}

func evalVimPath() (string, error) {
	for ppid := int32(os.Getppid()); ppid != 1; {
		ps, err := process.NewProcess(ppid)
		if err != nil {
			return "", err
		}
		ppid, err = ps.Ppid()
		if err != nil {
			return "", err
		}

		execPath, err := ps.Exe()
		if err != nil {
			return "", err
		}

		path, err := filepath.EvalSymlinks(execPath)
		if err != nil {
			return "", err
		}

		switch filepath.Base(path) {
		case "vim", "vim.basic":
			return path, nil

		default:
			continue
		}
	}

	return "", errors.New("not found")
}
