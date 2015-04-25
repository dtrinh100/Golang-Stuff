package bsearch


import (
"github.com/stretchr/testify/assert"
"testing"
)

func TestSearch(t *testing.T) {
  arr := []int{1,2,3}

  assert.Equal(t, Search(arr, 2, 0, len(arr)-1), 1, "Should be equal")
  assert.Equal(t, Search(arr, 1, 0, len(arr)-1), 0, "Should be equal")
  assert.Equal(t, Search(arr, 3, 0, len(arr)-1), 2, "Should be equal")
  assert.Equal(t, Search(arr, 5, 0, len(arr)-1), -1, "Key should not be found")
 
  arr2 := []int{5, 12, 17, 23, 38, 44, 77, 84, 90}
  
  assert.Equal(t, Search(arr2, 44, 0, len(arr2)-1), 5, "Should be equal")
  assert.Equal(t, Search(arr2, 90, 0, len(arr2)-1), 8, "Should be equal")
  assert.Equal(t, Search(arr2, 1000, 0, len(arr2)-1), -1, "Key should not be found")

  arr3 := []int{1}
  assert.Equal(t, Search(arr3, 1, 0, len(arr3)-1), 0, "Should be equal")
  assert.Equal(t, Search(arr3, 3, 0, len(arr3)-1), -1, "Key should not be found")
  
}
