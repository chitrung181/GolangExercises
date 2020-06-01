package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	pFlag := flag.Bool("f", false, "print to file flag")
	flag.Parse()
	for _, url := range flag.Args() {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("HTTP status: %v\n", resp.Status)
		if *pFlag {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			err = ioutil.WriteFile("output.txt", b, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Print: writing %s: %v\n", url, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("HTTP status: %s\n", resp.Status)
			_, err = io.Copy(os.Stdout, resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
		}

	}
}
