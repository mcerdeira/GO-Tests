package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	Mstart := time.Now()
	c := make(chan string)
	n := 0
	do(func(url string) {
		n++
		go GetUrl(url, c)
	})

	for i := 0; i < n; i++ {
		select {
		case result := <-c:
			fmt.Println(result)
		case <-time.After(time.Second * 20):
			fmt.Println("Time out")
			break
		}
	}
	fmt.Println("Total Time: ", time.Since(Mstart).Seconds())
}

func do(f func(string)) {
	var urls [255]string = GetFile("urls.txt")
	for _, u := range urls {
		if u != "" {
			f(u)
		}
	}
}

func GetFile(path string) [255]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var u [255]string
	var i int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		u[i] = scanner.Text()
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return u
}

func GetUrl(url string, c chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf(url + ": " + err.Error())
		return
	}
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s [%.2f]", url, dt)
}
