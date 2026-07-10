package main

import (
	"github.com/akamensky/argparse"
	"fmt"
	"os"
)

type Parser struct {
	Upload    string
	Download  string
	Del       string
}

func Argparser() Parser {
	parser := argparse.NewParser(os.Args[0], "Github Api wrapper");
	docs := parser.String("u", "upload", &argparse.Options{
		Required: false,
		Help: "upload a file or folder",
	});

	del := parser.String("x", "delete", &argparse.Options{
		Required: false,
		Help: "Delete a file or folder",
	});

	download := parser.String("d", "download", &argparse.Options{
		Required: false,
		Help: "Download a file or folder",
	});

	err := parser.Parse(os.Args);
	if err != nil {
		fmt.Print(parser.Usage(err));
	}

	if *docs == "" && *del == "" && *download == "" {
		fmt.Print(parser.Usage(nil));
	}

	fmt.Println(*docs, *del, *download);

	return Parser{
		Upload: *docs,
		Download: *download,
		Del: *del,
	};
}

