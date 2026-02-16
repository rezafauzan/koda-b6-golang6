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

func main() {	}