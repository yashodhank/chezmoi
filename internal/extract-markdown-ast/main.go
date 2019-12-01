package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
)

func run() error {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	m := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	m.Parser().Parse(text.NewReader(data)).Dump(data, 0)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
