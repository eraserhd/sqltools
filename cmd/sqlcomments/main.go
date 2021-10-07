package main

import (
	"github.com/2uinc/nexus-tools/pkg/sqlcomments"
	"os"
)

func main() {
	sqlcomments.Remove(os.Stdin, os.Stdout)
}
