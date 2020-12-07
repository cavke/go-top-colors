package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/hashicorp/golang-lru"
)

func main() {
	// parse input arguments
	inputPtr := flag.String("input", "input.txt", "input file containing URLS")
	outputPtr := flag.String("output", "result.csv", "output CSV file that will hold results")
	// set maximum parallel goroutines to the number of CPUs
	maxParallelsPtr := flag.Int("max-parallels", runtime.NumCPU(), "limits the maximum number of parallel goroutines for image processing")
	// set cache size
	cacheSizePtr := flag.Int("cache-size", 50, "maximum number of elements in the LRU cache")
	flag.Parse()

	inputFilePath := *inputPtr
	outputFilePath := *outputPtr
	maxParallels := *maxParallelsPtr
	cacheSize := *cacheSizePtr

	// results are cached per url in the LRU cache
	cache, err := lru.New(cacheSize)
	if err != nil {
		log.Fatal(err)
	}

	executionStart := time.Now()
	// process data
	c :=
		processData(
			generateData(inputFilePath), cache, maxParallels,
		)

	// write results
	outputFile, err := os.Create(outputFilePath)
	defer outputFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(outputFile)

	dataCnt := 0
	for data := range c {
		dataCnt++
		err := csvWriter.Write(data.getRow())
		if err != nil {
			log.Fatal(err)
		}
	}

	csvWriter.Flush()

	log.Printf("ALL DONE. Images processed: %d, ExecutionTime: %s", dataCnt, time.Since(executionStart))
}
