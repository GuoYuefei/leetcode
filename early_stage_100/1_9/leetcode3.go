// 无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
    result := 0
    // 可能的最大长度
    max := checke(s)

    for i := 0; i < len(s); {
        for j := i; j <= len(s); j++ {
            l := checke(s[i:j])
            if l < j-i {
                // 有重复
                iplus := strings.IndexByte(s[i:j], s[j-1])
                if iplus == -1 || iplus == j-i {
                    // 应该不会存在这种情况
                    iplus = 1
                }
                // 因为发生重复最外成循环开始跳跃
                i += iplus
                break
            } else {
                if l == max {
                    return max
                }
                if result < l {
                    result = l
                }
            }
        }
        i++
    }
    return result
}

// 检查字符串中最大不相同个数
func checke(s string) int {
    mp := make(map[rune]int)
    for _, v := range s {
        mp[v] = 1
    }

    return len(mp)
}


/**
本解法只考虑了ascii， 由于golang string的特殊性，中文不包含在本解法中
先**找最大可能性@0**， 然后**按顺序找@1**，如果出现最大可能性长度，那么就认定是这个长度

0. 找出最大可能性
    简单方式：确定所有不同的字符是几个

1. 按顺序找的方案：
    像abcb这种的，就可以直接从cb。。。开始找了
*/

