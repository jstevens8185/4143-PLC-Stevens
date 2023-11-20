package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

/////////////////////////////////////////////////////
// readURLsFromFile
//
// Reads URLs from a text file and returns them in a slice.
//
// Parameters:
//   - string: filePath - The path to the text file
//					      containing URLs.
//
// Returns:
//   - []string: A slice of URLs read from the file.
//   - error: An error, if any, encountered during the process.
/////////////////////////////////////////////////////

func readURLsFromFile(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line (URL) to the slice
		urls = append(urls, scanner.Text())
	}

	// Check for any reading errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

/////////////////////////////////////////////////////
// downloadImage
//
// Downloads an image from the given URL and
// saves it to disk with a unique name.
//
// Parameters:
//   - string: url - The URL of the image to download.
//   - *sync.WaitGroup: wg - A WaitGroup to synchronize
//							 goroutines.
//
// Returns: None
/////////////////////////////////////////////////////

func downloadImage(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Download the image
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	// Create a unique filename for the image
	fileName := fmt.Sprintf("image_%d.jpg", time.Now().UnixNano())
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file for %s: %s\n", url, err)
		return
	}
	defer file.Close()

	// Save the image to disk
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Error saving image for %s: %s\n", url, err)
		return
	}

	fmt.Printf("Downloaded %s and saved as %s\n", url, fileName)
}

/////////////////////////////////////////////////////
// downloadSequential
//
// Downloads and saves images sequentially.
//
// Parameters:
//   - slice: urls - A slice of image URLs to download.
//
// Returns: None
/////////////////////////////////////////////////////

func downloadSequential(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		downloadImage(url, &wg)
	}

	wg.Wait()
}

/////////////////////////////////////////////////////
// downloadConcurrent
//
// Downloads and saves images concurrently using goroutines.
//
// Parameters:
//   - slice: urls - A slice of image URLs to download.
//
// Returns: None
/////////////////////////////////////////////////////

func downloadConcurrent(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go downloadImage(url, &wg)
	}

	wg.Wait()
}

func main() {
	filePath := "urls.txt" // Replace file path as needed
	// List of image URLs to download
	urls, err := readURLsFromFile(filePath)
	if err != nil {
		panic(err)
	}

	// Record the start time for sequential download
	startTime := time.Now()

	fmt.Println("Downloading sequentially:")
	downloadSequential(urls)

	// Calculate and print the time taken for sequential download
	sequentialTime := time.Since(startTime)
	fmt.Printf("\nTime taken for sequential download: %s\n", sequentialTime)

	// Record the start time for concurrent download
	startTime = time.Now()

	fmt.Println("\nDownloading concurrently:")
	downloadConcurrent(urls)

	// Calculate and print the time taken for concurrent download
	concurrentTime := time.Since(startTime)
	fmt.Printf("\nTime taken for concurrent download: %s\n", concurrentTime)
}
