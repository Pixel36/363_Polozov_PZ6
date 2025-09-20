package main

import (
 "fmt"
 "sync"
)

type Queue struct {
 items []string
 mu    sync.Mutex
}

func (q *Queue) Enqueue(item string) {
 q.mu.Lock()
 defer q.mu.Unlock()
 q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (string, bool) {
 q.mu.Lock()
 defer q.mu.Unlock()
 if len(q.items) == 0 {
  return "", false
 }
 item := q.items[0]
 q.items = q.items[1:]
 return item, true
}

func main() {
 var wg sync.WaitGroup
 q := &Queue{}

 for i := 0; i < 5; i++ {
  wg.Add(1)
  go func(n int) {
   defer wg.Done()
   q.Enqueue(fmt.Sprintf("задача %d", n))
  }(i)
 }

 wg.Wait()

 for {
  item, ok := q.Dequeue()
  if !ok {
   break
  }
  fmt.Println("Очередь:", item)
 }
}
