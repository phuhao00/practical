package closure

import (
	"fmt"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	ForRangeAsyncFix2()
}

func ForRangeAsyncFix2() {
	type Person struct {
		Name string
	}
	group := []*Person{{"li"}, {"zhao"}}

	var wg sync.WaitGroup

	for _, p := range group {
		wg.Add(1)

		go func(person *Person) {
			defer wg.Done()
			fmt.Println("name =", person.Name)
		}(p)
	}
	wg.Wait()
}
