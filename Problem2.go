/* Project Euler Problem #2: Get the sum of all even Fib numbers that don't exceed 4 million */
package main
import "fmt"
import "math/big"


func Fib(N uint) *big.Int {
  x := uint(2)
  temp := new(big.Int).SetInt64(0)
  tempZ := new(big.Int).SetInt64(0)
  temp1 := new(big.Int).SetInt64(1)
  even := new(big.Int).SetInt64(2)
  sum := new(big.Int).SetInt64(0)
  if (N == 0) {
     sum = tempZ
   } else {
    for x <= N {
    /* Makes sure that terms don't exceed 4 million */
    res := temp.Cmp(new(big.Int).SetInt64(4000000)) 
    
    /* Exits the loop if it does exceed 4 million */
    if (res == 1) {
       break
    }
    /* Checks to make sure that if the terms are even, then add it to previous sum value */
    switch (new(big.Int).Mod(temp, even).Cmp(new(big.Int).SetInt64(0))) {
    case 0:
         sum = new(big.Int).Add(sum, temp)
    }
    /* Starts by adding the previous 2 terms */
    temp = new(big.Int).Add(tempZ, temp1)
    /* Assigns the first term the second term, so now second term becomes the first term for the next iteration */
    tempZ = temp1
    /* Assigns the second term the current term, so the current term becomes the second term */
    temp1 = temp
    x += 1
  }

  }
   return sum
}


func main() {
  fmt.Println(Fib(100))
}
