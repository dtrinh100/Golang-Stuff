package bsearch


// Search implements binary search to search a sorted array in O(logn). Returns the index of the searched key or -1 if it does not exists
func Search(arr []int, key, low, high int) int {

mid := (low+high)/2
cmid := arr[mid]

if (mid == low && low == high && mid == high) {
    if (cmid != key) {
        return -1
    }
} 

// If the mid pointer is less than the key, then we move the low pointer to mid+1, otherwise, we move the high pointer back by one from mid
if (cmid < key) {

low = mid+1
mid = Search(arr, key, low, high)

} else if (cmid > key) {

high = mid-1
mid = Search(arr, key, low, high) 

} 

return mid

}

