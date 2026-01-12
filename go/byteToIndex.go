// Useful if you're using bytes as a key to map, and you know that there's 
// only going to be lowercase letters.
func getIndex(b byte) int{
    return int(b) - 97
}
