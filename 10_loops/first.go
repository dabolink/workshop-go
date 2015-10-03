 package main

import (
    "fmt"
    "time"
)

type Computer struct {
    Brand string
    Model string
    Price int
}

func (c *Computer) Describe() {
    fmt.Printf("%s %s $%d\n", c.Brand, c.Model, c.Price)
}

func (c *Computer) StartTimer(channel chan string, t time.Duration) {
    fmt.Println("Starting timer...")
    time.Sleep(t)
    channel <- "Time up!"
}

func (c *Computer) StartInterval(channel chan string, t time.Duration) {
    fmt.Println("Starting interval...")
    for i := 0*time.Second; i < t; i = i + time.Second{
        time.Sleep(1*time.Second)
        fmt.Println("Interval ", i, "ended")
    }
    close(channel)
    // TODO: Use a "for" loop to push a message to the channel every t seconds

    // TODO: Close the channel after the loop finishes
}

func main() {
    computer := Computer{
        Brand: "Apple",
        Model: "Macbook",
        Price: 1000,
    }
    
    channel := make(chan string)
    
    t := 10 * time.Second
    go computer.StartInterval(channel, t)

    computer.Describe()
    // TODO: Use a "for" loop to watch for messages on the channel
    for message := range channel{
        fmt.Println(message)
    }
    fmt.Println("Exited")
}