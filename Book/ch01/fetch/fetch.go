package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// func main() {
// 	for _, url := range os.Args[1:] {
// 		if strings.HasPrefix(url, "http://") {
// 			resp, err := http.Get(url)

// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 				os.Exit(1)
// 			}
// 			b, err := io.Copy(os.Stdout, resp.Body)
// 			resp.Body.Close()
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "ferch: %v\n", err)
// 				os.Exit(1)
// 			}
// 			fmt.Printf("%v\n%v", resp.Status, b)
// 		} else {
// 			resp, err := http.Get("http://" + url)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 				os.Exit(1)
// 			}
// 			b, err := io.Copy(os.Stdout, resp.Body)
// 			resp.Body.Close()
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "ferch: %v\n", err)
// 				os.Exit(1)
// 			}
// 			fmt.Printf("%v\n%v\n", resp.Status, b)
// 		}
// 	}
// }
