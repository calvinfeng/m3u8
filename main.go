package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

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

type Task struct {
	URL      string `json:"url"`
	Dir      string `json:"dir"`
	Filename string `json:"filename"`
	Done     bool   `json:"done"`
}

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}

	jsonData, err := os.ReadFile(filepath.Join(workDir, "resources.json"))
	if err != nil {
		panic(err)
	}

	tasks := make([]*Task, 0)
	if err := json.Unmarshal(jsonData, &tasks); err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.Done {
			fmt.Printf("skip %s\n", task.Filename)
			continue
		}
		task.Dir = filepath.Join(workDir, "downloads")
		downloader, err := downloader.New(task.Dir, task.URL)
		if err != nil {
			panic(err)
		}
		if err := downloader.Start(chanSize, task.Filename); err != nil {
			panic(err)
		}
		fmt.Printf("%s is done!", task.Filename)
		task.Done = true
	}

	jsonData, err = json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(workDir, "resources.json"), jsonData, 0777); err != nil {
		panic(err)
	}
}
