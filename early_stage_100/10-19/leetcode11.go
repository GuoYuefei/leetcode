// 盛最多水的容器
func maxArea(height []int) int {
    result := 0
    b, e := 0, len(height) - 1

    for b < e {
        m := area(height, b, e)
        if m > result {
            result = m
        }
        if height[b] < height[e] {
            b++  
        } else {
            e--
        }
    }

    return result
}

// 计算面积函数
func area(height []int, b, e int) int {
    return min(height[b], height[e]) * (e - b)
}

// 其实本题如果考虑a b相等时的一种情况效率会更高，但是大多数情况增加的微乎其微
func min(a, b int) int {
    if a < b {
        return a  
    }
    return b
}

