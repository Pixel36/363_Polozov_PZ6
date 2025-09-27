package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var (
    golosa   map[string]int
    mu sync.Mutex
)

func GeneratorGolosov() {
    candidati := []string{"1", "2", "3"}
    for {
        cand := candidati[rand.Intn(len(candidati))]
        mu.Lock()
        golosa[cand]++
        mu.Unlock()
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
    }
}

func main() {
    golosa = make(map[string]int)
    for i := 0; i < 3; i++ {
        go GeneratorGolosov()
    }
    time.Sleep(2 * time.Second)
    mu.Lock()
    fmt.Println("Результаты:", golosa)
    mu.Unlock()
}
