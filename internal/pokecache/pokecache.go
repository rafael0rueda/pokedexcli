package pokecache

import ( 
	"time"
	"fmt"
)

type cacheEntry struct {
	createAt time.Time
	val []byte
}

func ShowTime() {
	fmt.Println(time.Now())
	return
}
