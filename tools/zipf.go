package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	zipf := rand.NewZipf(rd, 1.1, 1.0, 10)
	cnt := [11]int{0}

	for i := 0; i < 10000; i++ {
		rnum := zipf.Uint64()
		cnt[rnum] += 1

		url := fmt.Sprintf("http://localhost:1323/sleep/%d/50", rnum*10)
		_, err := http.Get(url)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Print(cnt)
}
