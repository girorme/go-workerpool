package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
	"workerpool/httpclient"

	"github.com/schollz/progressbar/v3"
)

func worker(
	wg *sync.WaitGroup,
	urlChannel <-chan string,
	resultChannel chan string,
	bar *progressbar.ProgressBar,
) {
	defer wg.Done()
	client := new(httpclient.HttpClient)
	for url := range urlChannel {
		resultChannel <- client.Get(url)
		bar.Add(1)
	}
}

func main() {
	start := time.Now()

	urls := []string{
		"https://google.com.br",
		"https://twitter.com",
		"https://mj-go.in/golang/async-http-requests-in-go",
		"https://wordpress.org",
		"https://cloudflare.com",
		"https://docs.google.com",
		"https://mozilla.org",
		"https://en.wikipedia.org",
		"https://maps.google.com",
		"https://accounts.google.com",
		"https://googleusercontent.com",
		"https://drive.google.com",
		"https://sites.google.com",
		"https://adobe.com",
		"https://plus.google.com1",
		"https://europa.eu",
		"https://bbc.co.uk",
		"https://vk.com",
		"https://es.wikipedia.org",
	}

	threads := flag.Int("t", 10, "Number of threads")
	flag.Parse()

	wg := new(sync.WaitGroup)
	wg.Add(*threads)
	urlChannel := make(chan string)
	resultChannel := make(chan string, len(urls))

	bar := progressbar.Default(int64(len(urls)))

	// Start the workers
	for i := 0; i < *threads; i++ {
		go worker(wg, urlChannel, resultChannel, bar)
	}

	// Send jobs to worker
	for _, url := range urls {
		urlChannel <- url
	}

	close(urlChannel)
	wg.Wait()
	close(resultChannel)

	// receive results from workes
	for result := range resultChannel {
		fmt.Print(result)
	}

	fmt.Printf("Took %s\n", time.Since(start))
}
