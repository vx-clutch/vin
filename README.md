# vin

An opinionated init tool for many languages

# Native Languages
    - Go
    - Python
    - C/C++
    - JavaScript
    - Rust

# Usage

    vin {lang} {name}
        - vin go hello
    This creates a directory and then adds the files inside needed for the language. The command vin go hello would make

    /hello
        /cmd/hello/main.go
        /pkg/
        /go-out/bin/
        makefile
    
    For not native supported languages you can use the flags --src {lang} {name}, or --src-comp {lang} {name}. This will create the directorys but not create any files, or add boilerplate code.
