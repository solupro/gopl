package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"os"
	"net/http"
	"io"
)

func main()  {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fecth(url, ch)
	}

	for range os.Args[1:] {
		fmt.Printf(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fecth(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if nil != err {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if nil != err {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}


