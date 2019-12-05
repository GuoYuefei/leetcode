// Z 字形变换
func convert(s string, numRows int) string {
    if numRows == 1 {
        // 1 就原样输出啊
        return s
    }
    var result = ""
    var row int = 0
    var flag bool = true       // true row+=1 flase row-=1
    // 初始化矩阵
    temp := make([][]byte, numRows)
    for i := 0; i < numRows; i++ {
        temp[i] = make([]byte, 0, 16)
    }
    // 形成矩阵
    for i := 0; i < len(s); i++ {
        temp[row] = append(temp[row], s[i])

        if row == numRows - 1 {
            flag = false
        } else if row == 0 {
            flag = true
        }
        if flag {
            row++
        } else {
            row--
        }
    }
    // 重新读取
    for _, vs := range temp {
        result = result + string(vs)
    }
    return result
}


/**
    1. 想到最简的方法就是用二维数组重新存储，矩阵可压缩，
    2. 重新读取
*/
