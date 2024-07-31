package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	lang := os.Args[1]
	name := os.Args[2]
	switch lang {
	case "go":
		run("go", []string{"mod", "init", name})
		os.MkdirAll(fmt.Sprintf("cmd/%v", name), 0666)
		os.Mkdir("pkg", 0666)
		os.MkdirAll("go-out/bin", 0666)
		os.WriteFile(fmt.Sprintf("cmd/%v/main.go", name), []byte("package main\nfunc main() {\n\treturn\n}"), 0666)
	}
}

func run(command string, args []string) {
	cmd := exec.Command(command, args...)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}
