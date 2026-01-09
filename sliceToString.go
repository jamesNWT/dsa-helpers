import "strconv"

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
