package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/shirou/gopsutil/process"
	"github.com/tennashi/tapiexec"
)

var (
	tabFlg = flag.Bool("tab", false, "open file in new tab")
)

func main() {
	flag.Parse()
	args := flag.Args()

	vimRuntime := os.Getenv("VIMRUNTIME")
	if vimRuntime == "" {
		err := runVim(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	vimEnv := filepath.Dir(vimRuntime)
	switch filepath.Base(vimEnv) {
	case "vim":
		srvName := os.Getenv("VIM_SERVERNAME")

		if srvName == "" {
			err := runVimTapi(args)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			err := runVimCS(srvName, args)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

		os.Exit(0)

	case "nvim":
		err := runNvimNVR(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}

func runVim(args []string) error {
	vim := exec.Command("vim", args...)

	vim.Stdin = os.Stdin
	vim.Stdout = os.Stdout
	vim.Stderr = os.Stderr

	return vim.Run()
}

func runVimTapi(args []string) error {
	if *tabFlg {
		args = append([]string{"tabnew"}, args...)
	} else {
		args = append([]string{"split"}, args...)
	}
	tapi := tapiexec.CallAPI("Tapi_open_wait", args)

	err := tapi.Run()
	if err != nil {
		return err
	}
	return tapiexec.WaitMsg("done")
}

func runVimCS(srvName string, args []string) error {
	path, err := evalVimPath()
	if err != nil {
		return err
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
	vim := exec.Command(path, args...)
	vim.Stdin = os.Stdin
	vim.Stdout = os.Stdout
	vim.Stderr = os.Stderr

	return vim.Run()
}

func runNvimNVR(args []string) error {
	if *tabFlg {
		args = append([]string{"--remote-tab-wait"}, args...)
	} else {
		args = append([]string{"--remote-wait"}, args...)
	}
	vim := exec.Command("nvr", args...)

	vim.Stdin = os.Stdin
	vim.Stdout = os.Stdout
	vim.Stderr = os.Stderr

	return vim.Run()
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
