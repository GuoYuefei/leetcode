// 罗马数字转整数
func romanToInt(s string) int {
    var result int = 0
    l := len(s)
    sign := 1
    mp := map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    // 前后出现逆序就是4*pow10 9*pow10
    for i := 0; i < l; i++ {
        sign = 1
        t := mp[s[i]]
        if i+1 != l && t < mp[s[i+1]] {
            sign = -1
        }
        result = result + sign * t
    }
    return result
}
