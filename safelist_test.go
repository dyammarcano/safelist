package safelist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSafeListString(t *testing.T) {
	safeList := new(SafeList[string])

	// Test Add method
	safeList.Add("testing item")

	// Test GetAll method
	items := safeList.GetAll()

	if len(items) != 1 {
		t.Errorf("Expected 1 item in the list but got %d", len(items))
	}

	if items[0] != "testing item" {
		t.Errorf("Expected item to be 'testing item' but got '%v'", items[0])
	}
}

func routine(safeList *SafeList[string], wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		safeList.Add(fmt.Sprintf("lap: %d", i))
	}
}

func TestSafeListInteger(t *testing.T) {
	wg := new(sync.WaitGroup)
	safeList := new(SafeList[string])

	wg.Add(1)
	go routine(safeList, wg)
	wg.Wait()

	assert.Equal(t, 100, len(safeList.GetAll()))

	safeList.Clear()
}
