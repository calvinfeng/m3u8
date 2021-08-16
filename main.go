package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/calvinfeng/m3u8/downloader"
)

var (
	url      string
	output   string
	chanSize int
	filename string
)

func init() {
	flag.StringVar(&url, "u", "", "M3U8 URL, required")
	flag.IntVar(&chanSize, "c", 25, "Maximum number of occurrences")
	flag.StringVar(&output, "o", "", "Output folder, required")
	flag.StringVar(&filename, "f", "", "Output filename, required")
}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	if url == "" {
		panicParameter("u")
	}
	if output == "" {
		panicParameter("o")
	}
	if filename == "" {
		panicParameter("f")
	}
	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	downloader, err := downloader.New(output, url)
	if err != nil {
		panic(err)
	}
	if err := downloader.Start(chanSize, filename); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func panicParameter(name string) {
	panic("parameter '" + name + "' is required")
}
