package main
import "fmt"

/* This is the partition code that I wrote while practicing my algorithm skills on an algorithms practice website, this code partitions an array for Quicksort. Quicksort is not included in this code */ 

func partition(arr []int, l, h int) int  {
    pivot := arr[0] // Uses the first element as a pivot
    i := l+1 // The left pointer is one element after the pivot
    j := h
   // For as long as the pointers don't meet 
   for i < j { 
    // For as long as the left pointer is less than the pivot and not out of bounds increment one 
    for arr[i] < pivot && i < len(arr)-1 { 
        
        i += 1
    }
    // For as long as the right pointer is more than the pivot and not out of bounds increment one 
    for arr[j] >= pivot && j > 0 {
       
        j--
       
    }
    // In some cases, the left and right pointer will cross in the loop, in this case, just exit the loop
    if i  > j {
        break 
    }
    
    // Swap the elements  
    tmp := arr[i]
   
    arr[i] = arr[j]
   
    arr[j] = tmp
    // After the swap is done, move the pointers one position
    i += 1
    j -= 1
   
    
   }
    // Once the pointers have met or cross, swap the pivot with the i-1 element, finishing the partition 
    if (j <= i) {
        tmp := arr[0]
        
        arr[0] = arr[i-1]
        
        
        arr[i-1] = tmp
        
       
    }
    
    return i-1 // returns the splitPoint 
}

func main() {
    var k int
    fmt.Scan(&k)
    test := k
    var s int
    var arr []int
    for i := 0; i < test; i++ {
    fmt.Scan(&s)
    arr = append(arr, s)
}
    fmt.Println(partition(arr, 0, len(arr)-1))
}

