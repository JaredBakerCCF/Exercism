package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	t := []int{2, 6, 9}
    return t
    
	panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < len(slice) && index >= 0 {
        return slice[index]
    } else {
    	return -1
    }
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index < len(slice) && index <= 0 {
    	slice[index] = value
        return slice
    } else {
    	return append(slice, index, value)
    }
	panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(slice, values ...)
	panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if (isIndexOutOfBounds(slice, index)) {
		return slice
	}
	return append(slice[:index], slice[index + 1:]...)
}

func isIndexOutOfBounds (slice []int, index int) bool {
	return index < 0 || index >= len(slice)
}
