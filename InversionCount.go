package InversionCount

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
)

func InversionCount(a []int) int {
    if len(a) <= 1 {
        return 0
    }
    mid := len(a) / 2
    left := a[:mid]
    right := a[mid:]
    leftCount := InversionCount(left) 
    rightCount := InversionCount(right) 

    res := make([]int, 0, len(right)+len(left)) //temp slice to hold the sorted left side and right side

    iCount := mergeCount(left, right, &res)

    copy(a, res)        //assigns the original slice with the temp slice values
    
    return iCount + leftCount + rightCount
}

    func mergeCount(left, right []int, res *[]int) int {
        count := 0

        for len(left) > 0 || len(right) > 0 {
            if len(left) == 0 {
                *res = append(*res, right...)
                break
            }
            if len(right) == 0 {
                *res = append(*res, left...)
                break
            }
        if left[0] <= right[0] {
            *res = append(*res, left[0])
            left = left[1:]
        } else { //Inversion has been found
            count += len(left)
            *res = append(*res, right[0])
            right = right[1:]
        }
    }

    return count
}

func main() {
    file, err := os.Open(os.Args[1]) //Takes file from command line 
    if err != nil {
	log.Fatal(err)
    }
    defer file.Close()
    var lines []int //slice used to store integers read from file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      val, err := strconv.Atoi(scanner.Text()) //Converts str into int
        if err != nil {
             log.Fatal(err)
        }
        lines = append(lines, val) //Adds the values into the slice one by one 
    }
    fmt.Println(InversionCount(lines))

}
