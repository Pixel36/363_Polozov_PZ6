package main

import (
    "fmt"
    "sync"
    "time"
)
type Cinema struct {
    seats     [38]bool
    mu        sync.Mutex
    booked    int
    available int
}
func (c *Cinema) BookSeat(seatNumber int, user string) bool {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    if seatNumber < 0 || seatNumber >= len(c.seats) {
        return false
    }
    if !c.seats[seatNumber] {
        c.seats[seatNumber] = true
        c.booked++
        c.available = len(c.seats) - c.booked
        fmt.Printf("Пользователь %s забранировал место %d\n", user, seatNumber)
        return true
    }
    fmt.Printf("Ошибка: Пользователь %s: место %d уже забронировано\n", user, seatNumber)
    return false
}
func (c *Cinema) GetAvSeats() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.available
}
func (c *Cinema) ShowStatus() {
    c.mu.Lock()
    defer c.mu.Unlock()
    fmt.Printf("Забронировано: %d, Свободно: %d\n", c.booked, c.available)
    for i, booked := range c.seats {
        status := "свободно"
        if booked {
            status = "забронировано"
        }
        fmt.Printf("Место %d: %s\n", i+1, status)
    }
}
func main() {
    cinema := Cinema{available: 38}
    var wg sync.WaitGroup
    users := []string{"Карина", "Матвей", "Хомяк", "Усатый", "Ивантуз"}
    for _, user := range users {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()
            for i := 0; i < 10; i++ {
                seat := i % 38
                cinema.BookSeat(seat, u)
                time.Sleep(time.Millisecond * 10)
            }
        }(user)
    }
    wg.Wait()
    cinema.ShowStatus()
}
