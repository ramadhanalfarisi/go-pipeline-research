package pipeline

import (
	"log"
	"os"
	"sync"

	"github.com/ramadhanalfarisi/go-concurrency-pipeline/model"
)

func WriteText(payloads chan model.Payload) chan string {
	filepaths := make(chan string, 4)
	go func(payloads chan model.Payload) {
		for payload := range payloads {
			d1 := []byte(payload.Content)
			err := os.WriteFile(payload.Filepath, d1, 0644)
			if err != nil {
				log.Fatal(err)
			}
			filepaths <- payload.Filepath
		}
		close(filepaths)
	}(payloads)
	return filepaths
}

func CollectPath(payloads ...chan string) []string {
	var paths []string
	var wg sync.WaitGroup
	var mut sync.Mutex
	for _, payload := range payloads {
		wg.Add(1)
		go func(payload chan string) {
			for item := range payload {
				mut.Lock()
				paths = append(paths, item)
				mut.Unlock()
			}
			wg.Done()
		}(payload)
	}
	wg.Wait()
	return paths
}
