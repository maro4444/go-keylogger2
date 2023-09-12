package main

import (
	"fmt"
	"github.com/BKasin/go-keylogger"
	"strconv"
	"strings"
	"time"
)

const (
	delayKeyfetchMS = 0
)

func main() {
	var tablica = []rune{}
	var kluczyk = [10]int{2, 0, 5, 6, 5, 7, 1, 8, 5, 0}
	var kluczykString3 = "2056571850"

	fmt.Println(len(kluczyk))
	var kluczykString = []string{}

	for _, poka := range kluczyk {
		kluczykString = append(kluczykString, strconv.Itoa(poka))
	}
	fmt.Println(kluczykString)

	//2056571850
	//var kluczNFC01 = []keylogger.Key{{false, 2, 2}, {false, 0, 0}, {false, 5, 5}, {false, 6, 6}, {false, 5, 5}, {false, 7, 7}, {false, 1, 1}, {false, 8, 8}, {false, 5, 5}, {false, 0, 0}}

	kl := keylogger.NewKeylogger()
	emptyCount := 0

	for {
		key := kl.GetKey()

		if !key.Empty {
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)
			tablica = append(tablica, key.Rune)
			if key.Keycode == 13 {
				//i := strconv.Itoa(kluczyk[0])
				var tablicaString = []string{}
				for _, poka := range tablica {
					//tablicaString += (strconv.Itoa(poka))
					//tablicaString = append(tablicaString, strconv.Itoa(poka))
					tablicaString = append(tablicaString, string(poka))
				}

				pierwsze := [10]string{}
				drugie := [10]string{}
				copy(pierwsze[:], kluczykString)
				copy(drugie[:], tablicaString)
				var pierwsze3 string = strings.Join(kluczykString, "")

				if pierwsze == drugie {
					fmt.Println("TAAAAAAAAAAAAAAAAAAAAAAAK")
				}
				if pierwsze3 == kluczykString3 {
					fmt.Println("TAAAAAAAAAAAAAAAAAAAAAAAK")
				}

				println("kluczykString: ", len(kluczykString), " tablicaString: ", len(tablicaString))
				println("pierwsze: ", len(pierwsze), " drugie: ", len(drugie))
				fmt.Println("Pierwsze: ", pierwsze)
				fmt.Println("Drugie:   ", drugie)
			}
		}

		emptyCount++

		fmt.Printf("Empty count: %d\r", emptyCount)

		//time.Sleep(delayKeyfetchMS * time.Millisecond)
		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}

//205657185020565718502056571850205657185020565718502056571850205657185020565718502056571850205657185020565718502056571850
//20565718502056571850205657185020565718502056571850205657185020565718502056571850

//
//
//
