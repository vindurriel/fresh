package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"flag"
	"strings"
	"fmt"
)

var buildArgs string

func init() {
	flag.StringVar(&buildArgs, "b", "", "Command line arguments for build")
}

func build() (string, bool) {

	args := []string{
		"build",
	}
	if buildArgs != "" {
		args = append(args[:1], strings.Split(buildArgs, ";")...)
	}
	args = append(args, "-i", "-o", buildPath(), root())

	buildLog(fmt.Sprintf("Building, args: %v", args))

	cmd := exec.Command("go", args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
