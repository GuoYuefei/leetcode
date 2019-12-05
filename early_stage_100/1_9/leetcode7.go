// 整数反转

func reverse(x int) int {
    var result int = 0
    v := x
    for v != 0 {
        a := v % 10
        v = v / 10
        // 虽然可以使用int64， 但是这边题目的意思是在此之前判溢出
        if result > (math.MaxInt32 - a) / 10.0 ||
            result < (math.MinInt32 - a) / 10.0 {
                return 0
        }
        result = result * 10 + a
    }
    return result
}
