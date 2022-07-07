package u



// 将字符串切割到指定长度
// TruncatedString returns a truncated string.
func TruncatedString(src string, limit int) string {
	var (
		end = "..."
	)
	if len(src) <= limit {
		return src
	}
	return src[:limit-len(end)] + end
}
