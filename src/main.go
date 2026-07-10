package main

import (
	"fmt"
	"os"
)

func main() {
	parser, err := Argparser();
	if err != nil {
		return;
	}

	if parser.Upload != "" {
		// check if file or folder
		info, err := os.Stat(parser.Upload);
		if err != nil {
			fmt.Printf("[!] Error: %t\n", err)
			return;
		}

		if info.IsDir() {
			// its dir
		} else if info.Mode().IsRegular() {
			// Its a regular file
		} else {
			fmt.Printf("[!] Error: %s is not a file or folder\n", parser.Upload);
			return;
		}
	} else if parser.Del != "" {
		// check if file or folder
		info, err := os.Stat(parser.Del);
		if err != nil {
			fmt.Printf("[!] Error: %t\n", err)
			return;
		}

		if info.IsDir() {
			// its dir
		} else if info.Mode().IsRegular() {
			// Its a regular file
		} else {
			fmt.Printf("[!] Error: %s is not a file or folder\n", parser.Del);
			return;
		}
	} else if parser.Download != "" {
		// check if file or folder
		info, err := os.Stat(parser.Download);
		if err != nil {
			fmt.Printf("[!] Error: %t\n", err)
			return;
		}

		if info.IsDir() {
			// its dir
		} else if info.Mode().IsRegular() {
			// Its a regular file
		} else {
			fmt.Printf("[!] Error: %s is not a file or folder\n", parser.Download);
			return;
		}
	}
}

/*

func ManageFiles(file, bool: is dir or not) [string] {
	if dir then get all the files recursively
	if not the just that file
	return
}


*/
