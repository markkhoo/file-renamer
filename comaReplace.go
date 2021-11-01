package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	cur_dir, err := filepath.Abs(".")
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var file_slice []string
	var dir_slice []string

	for _, path := range files {
		fullPath := cur_dir + "\\" + path.Name()
		fi, err := os.Stat(fullPath)
		if err != nil {
			log.Fatal(err)
		}

		switch mode := fi.Mode(); {
		case mode.IsRegular():
			// do file stuff
			file_slice = append(file_slice, fullPath)
		case mode.IsDir():
			// do directory stuff
			dir_slice = append(dir_slice, fullPath)
		}
	}

	for len(dir_slice) != 0 {
		cur_dir = dir_slice[0]
		dir_slice = RemoveIndex(dir_slice, 0)
		files, err := os.ReadDir(cur_dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, path := range files {
			fullPath := cur_dir + "\\" + path.Name()
			fi, err := os.Stat(fullPath)
			if err != nil {
				log.Fatal(err)
			}

			switch mode := fi.Mode(); {
			case mode.IsRegular():
				// do file stuff
				file_slice = append(file_slice, fullPath)
			case mode.IsDir():
				// do directory stuff
				dir_slice = append(dir_slice, fullPath)
			}
		}
	}

	for _, file := range file_slice {
		e := os.Rename(file, strings.ReplaceAll(file, ",", "_"))
		if e != nil {
			log.Fatal(e)
		}
	}

	fmt.Println("File Renaming Complete")

}
