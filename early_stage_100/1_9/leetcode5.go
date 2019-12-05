// 最长回文子串
func longestPalindrome(s string) string {
    max := 0
    longest := ""
    for x := 0; x < len(s); x++ {
        son, l := centerSpread(s, x, x)
        if l > max {
           longest = son
           max = l
        }
        son, l = centerSpread(s, x, x+1)
        if l > max {
            longest = son
            max = l
        }
    }
    return longest
}

// @param s 就是辣个字符串
// @param i, j i与j确定中心位置  if i == j then s is 奇数子串
func centerSpread(s string, i, j int) (string, int) {
    var left, right int = i, j
    if i == j {
        left = i - 1
        right = i + 1
    }

    for x := 0; x < len(s) / 2; x++ {
        if left < 0 || right >= len(s) || s[left] != s[right] {
            break
        }
        left--
        right++
    }

    return s[left+1 : right], right-left-1
}
