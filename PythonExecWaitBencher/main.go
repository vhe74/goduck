package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	log.Println("started")

	for i := 1; i <= 5; i++ {
		log.Println(i)
		x := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			DoRequest(x)
		}()
	}
	wg.Wait()
}

func DoRequest(i int) {
	url := "http://localhost:8000/"

	log.Printf("[%d] Requesting \n", i)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("[%d] Response received \n", i)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf("[%d] %s\n", i, sb)
}
