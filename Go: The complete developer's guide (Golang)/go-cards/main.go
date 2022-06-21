package main

import (
	"fmt"
	"sync"
)

//  go run main.go deck.go card.go
func main() {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	cards := newDeck()
	cards.shuffle()

	game := make([]deck, 0, 4)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()

			fmt.Println(len(cards))

			hand, remainingCards := deal(cards, 4)
			hand.shuffle()

			cards = remainingCards
			game = append(game, hand)

			wg.Done()
		}()
	}
	wg.Wait()

	for _, hand := range game {
		hand.print()
		fmt.Println()
	}
}
