package main

import (
	"github.com/eraserhd/sqltools/pkg/sqlcomments"
	"os"
)

func main() {
	sqlcomments.Remove(os.Stdin, os.Stdout)
}
