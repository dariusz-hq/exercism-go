package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}
    return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice) < index + 1 || index < 0  {
        return -1
    }
    item := slice[index]
    return item
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice) < index + 1 || index < 0  {
        return append(slice, value)
    }
    result := append(slice[0:index], value)
    return append(result, slice[index+1:len(slice)]...)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	r := []int{}
    r = append(r, values...)
    return append(r, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) -1 {
        return slice
    }
    
    if index == len(slice) {
        return slice[:index]
    } else if index == 0 {
        return slice[index+1:]
    } else {
    	return append(slice[0:index], slice[index+1:len(slice)]...)    
    }
}
