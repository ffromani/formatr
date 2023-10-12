package main

import (
	"fmt"

	"github.com/ffromani/formatr"
)

func main() {
	fmtr := formatr.NewFormatterKlog(formatr.Options{})
	kvList := []any{"foo", 1, "bar", 0.5, "fizz", "buzz"}
	_, args := fmtr.FormatInfo(0, "test message", kvList)
	fmt.Println(args)
}
