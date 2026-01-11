import "strconv"

// Note: sometimes it will be faster to weave this logic into a loop that 
// already exists in your algorithm.
func intSliceToString(nums []int) string {
    if len(nums) == 0 {
        return ""
    }
    var strBuilder strings.Builder
    for i, n := range nums {
        if i > 0 {
            strBuilder.WriteByte(',')
        }
        b.WriteString(strconv.Itoa(n))
    }
    return strBuilder.String()
}
