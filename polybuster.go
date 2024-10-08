package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const urlPlaylist string = "https://raw.githubusercontent.com/danielmiessler/SecLists/refs/heads/master/Discovery/Web-Content/directory-list-2.3-medium.txt"

func main() {
	var url string
	var wordlist string
	var help bool
	flag.BoolVar(&help, "h", "false", "ask for help")
	flag.StringVar(&url, "u", "", "url link, end with /")
	flag.StringVar(&wordlist, "w", "", "path of wordlist")
	flag.Parse()
if help {
	printhelp()
}

  _, err := os.Stat("seclist_wordlist.txt")
	if err == nil {
		fmt.Println("using seclist_wordlist.txt")
		wordlist = "seclist_wordlist.txt"
	}
	if url == "" {
  fmt.Println("please enter an url here:")
    fmt.Scan(&url)
	  if url == "" {
      fmt.Println("enter a url pls")
      return
      }
  }
	if wordlist == "" {
		var download string
		fmt.Printf("no wordlist detected, do you wanna download %s? (y/N)", urlPlaylist)
		fmt.Scan(&download)
		if download == "y" || download == "Y" {
			fmt.Println("downloading...")
			resp, err := http.Get(urlPlaylist)
			if err != nil {
				fmt.Println("error making the request to github:", err)
				return
			}
			defer resp.Body.Close()
      
			respBody, err := io.ReadAll(resp.Body)
      if err != nil {
				fmt.Println("error converting body to byte", err)
				return
			}
			os.Create("seclist_wordlist.txt")
			os.WriteFile("seclist_wordlist.txt", respBody, 0644)
		}
		wordlist = "seclist_wordlist.txt"

	}

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}

	fmt.Println("using directory listing on", url)
	fmt.Println("Attempting to open wordlist file:", wordlist)
	wordlistFile, err := os.OpenFile(wordlist, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal("error opening wordlist file:", err)
	}
	defer wordlistFile.Close()

	scanner := bufio.NewScanner(wordlistFile)
	fmt.Println("starting scans")
	for scanner.Scan() {
		line := scanner.Text()
			if strings.Contains(line, "#") {
      continue
    }
		request, err := http.Get(fmt.Sprintf("%s%s", url, line))
		if err != nil {
			log.Fatal("error making the request:", err)
		}

		defer request.Body.Close()
		if request.StatusCode == http.StatusOK {
			fmt.Println("directory found:", url, line)
		}

  }
  fmt.Println("no more words in the playlist :3")

}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -h        Show help information")
	fmt.Println("  -u <url>  Specify the URL to scan")
	fmt.Println("  -w <path> Specify the path to the wordlist file")
	fmt.Println()
	fmt.Println("If no wordlist is provided, the program will attempt to download a default wordlist. (directory-list-2.3-medium.txt from github)")
	fmt.Println("Example:")
	fmt.Println("  go run main.go -u https://example.com/ -w /path/to/wordlist.txt")
}

