package main

import (
 "fmt"
 "sync"
)

func main() {
 var shotchik int
 var mutex sync.Mutex
 var ojidalka sync.WaitGroup

 for i := 0; i < 10; i++ {
  ojidalka.Add(1)
  go func() {
   defer ojidalka.Done()
   for j := 0; j < 10000; j++ {
    mutex.Lock()
    shotchik++
    mutex.Unlock()
   }
  }()
 }

 ojidalka.Wait()
 fmt.Println("Количество просмотров:", shotchik)
}
