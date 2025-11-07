package concurrency

import (
	"fmt"
	"library_management/services"
	"sync"
)

func SimulateConcurrentReservations(lib *services.Library) {
	var wg sync.WaitGroup

	members := []int{101, 102, 103}
	bookID := 1

	for _, mID := range members {
		wg.Add(1)
		go func(memberID int) {
			defer wg.Done()
			err := lib.ReserveBook(bookID, memberID)
			if err != nil {
				fmt.Printf("Member %d: %v\n", memberID, err)
			} else {
				fmt.Printf("Member %d successfully reserved book %d\n", memberID, bookID)
			}
		}(mID)
	}

	wg.Wait()
}
