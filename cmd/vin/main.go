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
	default:
		fmt.Println("No lang found")
		os.Exit(0)
	case "go":
		run("go", []string{"mod", "init", name})
		os.MkdirAll(fmt.Sprintf("cmd/%v", name), 0755)
		os.Mkdir("pkg", 0755)
		os.MkdirAll("go-out/bin", 0755)
		os.WriteFile(fmt.Sprintf("cmd/%v/main.go", name), []byte(get("go")), 0755)
	case "c":
		src("c", "c")
	case "python":
		src("python", "py")
	}
}

func run(command string, args []string) {
	cmd := exec.Command(command, args...)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}

func src(lang string, ext string) {
	os.Mkdir("src", 0755)
	os.MkdirAll(fmt.Sprintf("%v-out/bin", lang), 0755)
	os.WriteFile(fmt.Sprintf("src/main.%v", ext), []byte(get(lang)), 0755)
}

func get(lang string) string {
	switch lang {
	case "go":
		return "package main\n\nfunc main() {\n\treturn\n}"
	case "c":
		return "int main() {\n\treturn 0\n}"
	case "python":
		return "print('Hello, World!')"
	}
	return ""
}
