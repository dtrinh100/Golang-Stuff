/* Project Euler Problem 1: Find the sum of all numbers that are divisible by 3 and 5 that is less than 1000 */
package main


import "fmt"

func main() {
  sum := 0 
  for x := 0; x < 1000; x += 1 {
       /* If the number is divisible by both 3 AND 5, then only count it once */
       if (x % 5 == 0 && x % 3 == 0) {
          sum += x
       /* Otherwise, it just checks to see if it's divisible by 3 or 5, then it adds that number to current sum */
       }  else if x % 5 == 0 {
         sum += x
       }  else if x % 3 == 0 {
         sum += x
      }
    }
   fmt.Println(sum)
  


}
