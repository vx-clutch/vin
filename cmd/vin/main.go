package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if os.Args[1] == "--src" {
		customLangInput := os.Args[2]
		customName := os.Args[3]
		customLang(customName, customLangInput, false)
	} else if os.Args[1] == "--src-comp" {
		customLangInput := os.Args[2]
		customName := os.Args[3]
		customLang(customName, customLangInput, true)
	}
	lang := os.Args[1]
	name := os.Args[2]
	switch strings.ToLower(lang) {
	default:
		fmt.Printf("vin: We do not supported %v yet. Make a pr if you want\n", name)
		os.Exit(0)
	case "go":
		os.Mkdir(name, 0755)
		os.Chdir(name)
		run("go", []string{"mod", "init", name})
		os.MkdirAll(fmt.Sprintf("cmd/%v", name), 0755)
		os.Mkdir("pkg", 0755)
		os.MkdirAll("go-out/bin/", 0755)
		os.WriteFile(fmt.Sprintf("cmd/%v/main.go", name), []byte(get("go")), 0755)
		os.WriteFile("makefile", []byte(fmt.Sprintf("main=%v\ncomp=go build\n\nall: build\n\nbuild:\n\t@$(comp) -o ./%v-out/bin/ $(main)", fmt.Sprintf("cmd/%v/main.go", name), lang)), 0755)
	case "c":
		src("c", "c", true, name)
	case "cpp", "c++":
		src("c", "cpp", true, name)
	case "python":
		src("python", "py", false, name)
	case "js", "javascript":
		src("javascript", "js", false, name)
	case "rust":
		run("cargo", []string{"new", name})

	}
}

func run(command string, args []string) {
	cmd := exec.Command(command, args...)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}

func src(lang string, ext string, compiled bool, name string) {
	os.Mkdir(name, 0755)
	os.Chdir(name)
	os.Mkdir("src", 0755)
	if compiled {
		os.MkdirAll(fmt.Sprintf("%v-out/bin", lang), 0755)
	}
	os.WriteFile(fmt.Sprintf("src/main.%v", ext), []byte(get(lang)), 0755)
	os.WriteFile("makefile", []byte(fmt.Sprintf("main=%v\ncomp=\n\nall: build\n\nbuild:\n\t@$(comp) -o ./%v-out/bin/ $(main)", fmt.Sprintf("src/%v/main.%v", name, ext), lang)), 0755)
}

func get(lang string) string {
	switch lang {
	case "go":
		return "package main\n\nfunc main() {\n\treturn\n}"
	case "c", "cpp":
		return "int main() {\n\treturn 0\n}"
	case "python":
		return "print('Hello, World!')"
	case "javascript":
		return "console.log('Hello, World!')"
	}
	return ""
}

func customLang(name string, lang string, comp bool) {
	os.Mkdir(name, 0755)
	os.Chdir(name)
	os.Mkdir("src", 0755)
	if comp {
		os.MkdirAll(fmt.Sprintf("%v-out/bin/", lang), 0755)
	}
	os.WriteFile("makefile", []byte(fmt.Sprintf("main=\ncomp=\n\nall: build\n\nbuild:\n\t@$(comp) -o ./%v-out/bin/ $(main)", lang)), 0755)
	os.Exit(0)
}
