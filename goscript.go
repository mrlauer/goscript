package main

import(
    "os"
    "log"
    "io/ioutil"
    "path"
    "exec"
)

func execute(c string, args ... string) {
    cmd := exec.Command(c, args...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Fatal(err.String())
    }
}

func main() {
    // Get the arguments
    osargs := os.Args
    if len(osargs) < 2 {
        log.Fatal("No go file specified\n")
    }
    gofile := osargs[1]
    args := osargs[2:]

    // Make a temp directory
    tempdir, err := ioutil.TempDir(os.TempDir(), "goscript")
    if err != nil {
        log.Fatal(err.String())
    }
    defer os.RemoveAll(tempdir)

    compiledFile := path.Join(tempdir, "go.6")
    exe := path.Join(tempdir, "6.out")

    // Run 6g
    execute("6g", "-o", compiledFile, gofile)

    // Run 6l
    execute("6l", "-o", exe, compiledFile)

    // Run the executable
    execute(exe, args...)

}
