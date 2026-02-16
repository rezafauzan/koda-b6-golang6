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
	{Nama: "Aetherion", Waktu: 4},
	{Nama: "Zythera", Waktu: 6},
	{Nama: "Vaelthryn", Waktu: 1},
	{Nama: "Eryndor", Waktu: 9},
}

func PrintAntrian(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Memulai antrian")
	for x := range antrian {
		time.Sleep(time.Duration(antrian[x].Waktu) * time.Second)
		fmt.Printf("Pesanan %s selesai \n" ,antrian[x].Nama)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go PrintAntrian(&wg)
	wg.Wait()
}