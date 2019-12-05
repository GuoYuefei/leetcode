// 回文数
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    tmp := x
    a := make([]int, 0, 20)
    for tmp != 0 {
        a = append(a, tmp % 10)
        tmp = tmp / 10
    }
    for i := 0; i < len(a) / 2; i++ {
        if a[i] != a[len(a) - 1 - i] {
            return false
        }
    }
    return true
}
