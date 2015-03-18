package InversionCount

import "testing"

type testpair struct {
   values []int
   count int
}

var tests = []testpair {
{  []int{1,2}, 0 },
{  []int{3,2,1}, 3},
{  []int{0,0,0,0}, 0},
{  []int{1,2,3,5,4}, 1},
{  []int{3,1,6,5,2,4}, 7},
{  []int{5,2,10,8,1,9,4,3,6,7}, 22},

}

func TestInversionCount(t *testing.T) {
  for _, pair := range tests {
    c := InversionCount(pair.values)
    if c != pair.count {
        t.Error(
        "Expected:", pair.count,
        "Got:", c,
     )
   }
  }
}
