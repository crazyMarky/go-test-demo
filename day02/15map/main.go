package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//随机种子
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int,200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d",i)
		scoreMap[key] = rand.Intn(100)
	}

	//fmt.Println(scoreMap)
	var keys = make([]string,0,200)

	for key := range scoreMap {
		keys = append(keys,key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key,scoreMap[key])
	}

}
