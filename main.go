package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Shyp/bump_version/lib"
)

const VERSION = 1.0

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: bump_version <major|minor|patch> <filename>\n")
}

// runCommand execs the given command and exits if it fails.
func runCommand(binary string, args ...string) {
	out, err := exec.Command(binary, args...).CombinedOutput()
	if err != nil {
		log.Fatalf("Error when running command: %s.\nOutput was:\n%s", err.Error(), string(out))
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		usage()
		os.Exit(2)
	}
	versionTypeStr := args[0]
	filename := args[1]

	version, err := bump_version.BumpInFile(bump_version.VersionType(versionTypeStr), filename)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(os.Stderr, "Bumped version to %s\n", version)
	}
	runCommand("git", "add", filename)
	runCommand("git", "commit", "-m", version.String())
	runCommand("git", "tag", version.String(), "--annotate", "--message", version.String())
	fmt.Fprintf(os.Stderr, "Added new commit and tagged version %s.\n", version)
}
