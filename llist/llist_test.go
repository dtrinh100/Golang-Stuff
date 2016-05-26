package llist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	l := New()
	l.Add(1)
	l.Add(3)
	assert.True(t, l.Contains(3), "LList Get(3) should return true")
	assert.Equal(t, 2, l.Size(), "Size should be 3")
}

func TestContains(t *testing.T) {
	l := New()
	l.Add(0)
	l.Add(10)
	l.Add(2)
	l.Add(7)
	assert.False(t, l.Contains(5), "Contains(5) Should return false")
	assert.True(t, l.Contains(10), "Contains(10) Should return true")
	assert.True(t, l.Contains(0), "Contains(0) Should return true")
}

func TestDelete(t *testing.T) {
	l := New()
	l.Add(0)
	l.Add(10)
	l.Add(2)
	l.Add(7)
	l.Delete(10)
	assert.True(t, l.Contains(0), "0 should exist")
	assert.False(t, l.Contains(10), "10 should no longer exist")
	assert.True(t, l.Contains(2), "2 should exist")
	assert.True(t, l.Contains(7), "7 should exist")
	assert.Equal(t, 3, l.Size(), "Size should now decrease to 3")

	l.Delete(2)
	assert.False(t, l.Contains(2), "2 should no longer exist")
}

func TestDeleteBack(t *testing.T) {
	l := New()
	l.Add(0)
	l.Add(10)
	l.Add(2)
	l.Add(7)
	l.Delete(7)

	assert.False(t, l.Contains(7), "7 should no longer exist")

}

func TestDeleteFront(t *testing.T) {
	l := New()
	l.Add(0)
	l.Add(10)
	l.Add(2)
	l.Add(7)
	l.Delete(0)

	assert.False(t, l.Contains(0), "0 should no longer exist")

}

func TestPrint(t *testing.T) {
	l := New()
	l.Add(0)
	l.Add(10)
	l.Add(2)
	l.Add(7)

	assert.Equal(t, "0 10 2 7", l.String())
}
