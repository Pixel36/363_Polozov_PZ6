package main
import (
	"fmt"
	"sync"
)
type Store struct {
	sync.RWMutex
	items map[string]int
}
func main() {
	store := &Store{items: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			store.Lock()
			store.items["Самса"] += 5
			fmt.Printf("Поступление %d: +5 товаров\n", id)
			store.Unlock()
		}(i)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			store.Lock()
			if store.items["Самса"] >= 5 {
				store.items["Самса"] -= 5
				fmt.Printf("Продажа %d: -5 товаров\n", id)
			} else {
				fmt.Printf("Продажа %d: недостаточно товаров\n", id)
			}
			store.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Остаток товаров: %d\n", store.items["Самса"])
}
