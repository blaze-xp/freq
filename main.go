package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorBlue := "\033[34m"

	sc := bufio.NewScanner(os.Stdin)

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for domain := range jobs {
				resp, err := http.Get(domain)
				if err != nil {
					continue
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					continue
				}
				sb := string(body)

				// Initialize counts
				alertCount := strings.Count(sb, "alert(1)")
				confirmCount := strings.Count(sb, "confirm(1)")
				promptCount := strings.Count(sb, "prompt(1)")

				// Check for alerts
				if alertCount > 0 || confirmCount > 0 || promptCount > 0 {
					fmt.Printf(
						"%sXSS FOUND: %s\nAlert Count: %d, Confirm Count: %d, Prompt Count: %d%s\n",
						colorGreen,
						domain,
						alertCount,
						confirmCount,
						promptCount,
						colorReset,
					)
				} else {
					fmt.Printf(
						"%sNot Vulnerable: %s%s\n",
						colorRed,
						domain,
						colorReset,
					)
				}
			}
		}()
	}

	for sc.Scan() {
		domain := sc.Text()
		jobs <- domain
	}
	close(jobs)
	wg.Wait()
}
