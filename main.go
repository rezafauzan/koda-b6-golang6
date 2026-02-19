package main

import (
	"fmt"
	"sync"
	"time"
)

type Antrian struct {
	Nama  string
	Waktu int
}

var antrian = []Antrian{
	{Nama: "Aetherion", Waktu: 1},
	{Nama: "Zythera", Waktu: 2},
	{Nama: "Vaelthryn", Waktu: 4},
	{Nama: "Eryndor", Waktu: 4},
}

func PrintAntrian(data chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Memulai antrian")
	for x := range antrian {
		time.Sleep(time.Duration(antrian[x].Waktu) * time.Second)
		fmt.Printf("Pesanan %s selesai (%d) \n" ,antrian[x].Nama, antrian[x].Waktu)
	}
}

func main() {
	data := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go PrintAntrian(data, &wg)
	wg.Wait()
	fmt.Println("Antrian Selesai")
}