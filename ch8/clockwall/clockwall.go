package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	timeTable := make(map[string]string)
	var mu sync.Mutex
	cities := make([]string, 0)
	for i := 1; i < len(os.Args); i++ {
		val := strings.Split(os.Args[i], "=")
		city, port := val[0], val[1]
		cities = append(cities, city)
		go func() {
			conn, err := net.Dial("tcp", port)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				mu.Lock()
				timeTable[city] = scanner.Text()
				mu.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 3)
	for {
		fmt.Fprintf(os.Stdout, "%-15s | %-15s\n", "City", "Time")
		for _, city := range cities {
			time := timeTable[city]
			fmt.Fprintf(os.Stdout, "%-15s : %-15s\n", city, time)
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
