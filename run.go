package main

import (
	"context"
	"log"
	"os"

	"github.com/MagicalTux/gophp/core"
	_ "github.com/MagicalTux/gophp/ext/ctype"
	_ "github.com/MagicalTux/gophp/ext/gmp"
	_ "github.com/MagicalTux/gophp/ext/hash"
	_ "github.com/MagicalTux/gophp/ext/pcre"
	_ "github.com/MagicalTux/gophp/ext/standard"
)

func main() {
	// by default, run script test.php
	p := core.NewProcess("cli")
	p.CommandLine(os.Args)
	ctx := core.NewGlobal(context.Background(), p)
	if err := ctx.RunFile("test.php"); err != nil {
		log.Printf("failed to run test file: %s", err)
		os.Exit(1)
	}
}
