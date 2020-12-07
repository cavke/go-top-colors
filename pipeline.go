package main

import (
	"bufio"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru"
)

// function that reads URLs from input file and sends them to channel
func generateData(inputFilePath string) <-chan string {
	c := make(chan string)
	go func() {
		file, _ := os.Open(inputFilePath)
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			url, err := reader.ReadString('\n')
			if url != "" {
				url = strings.TrimSuffix(url, "\n")
				c <- url
			}
			if err == io.EOF {
				break
			}
		}

		close(c)
	}()
	return c
}

type ResultData struct {
	url      string
	topThree []string
}

func (d *ResultData) getRow() []string {
	row := make([]string, 4)
	row[0] = d.url

	for i, c := range d.topThree {
		if c != "" {
			row[i+1] = d.topThree[i]
		}
	}

	return row
}

// function that parallelize data processing
func processData(ic <-chan string, cache *lru.Cache, maxParallels int) <-chan ResultData {
	oc := make(chan ResultData)

	go func() {
		wg := &sync.WaitGroup{}

		// limit number of parallel executions
		maxChan := make(chan bool, maxParallels)

		for url := range ic {
			maxChan <- true

			wg.Add(1)
			go processImageFromUrl(url, oc, wg, maxChan, cache)
		}

		wg.Wait()
		close(oc)
	}()

	return oc
}

// function that retrieves image from URL and counts image colors
func processImageFromUrl(url string, output chan<- ResultData, wg *sync.WaitGroup, maxChan chan bool, cache *lru.Cache) {
	defer func(maxChan chan bool) { <-maxChan }(maxChan)
	defer wg.Done()

	executionStart := time.Now()

	// check if image URL is already parsed
	var resultObj, ok = cache.Get(url)
	if ok {
		// return result from cache
		result := resultObj.(*ResultData)
		output <- *result
	} else {
		// get image from url
		res, err := http.Get(url)
		if err != nil {
			log.Printf("Error fetching %s: %v", url, err)
		} else {
			image, _, err := image.Decode(res.Body)
			defer res.Body.Close()
			if err != nil {
				log.Printf("Error decoding %s: %v", url, err)
			} else {
				// count the number of colors
				colorMap := getColorMapFromImageWithPix(image)
				topThree := getTopThreeColors(colorMap)
				resultData := ResultData{url, topThree}

				// put result to cache
				cache.Add(url, &resultData)

				output <- resultData
			}
		}
	}

	cStr := ""
	if ok {
		// mark that the result was retrieved from cache
		cStr = "(C)"
	}
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	log.Printf("Image processed: %s ExecutionTime: %s, Alloc = %v MiB, url=%s", cStr, time.Since(executionStart), bToMb(memStats.Alloc), url)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
