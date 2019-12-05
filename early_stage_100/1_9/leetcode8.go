// 字符串转换整数 (atoi)
func myAtoi(str string) int {
    var result int = 0
    begin, end := -1, -1
    sign := 1           // 1代表正数， -1为负数 最后*sign就行

    var flag = true         // 判定是不是第一个数字

    for i := 0; i < len(str); i++ {
        // 是数字就进行相关记录啊
        if str[i] >= '0' && str[i] <= '9' {
            if flag {
                flag = false
                begin = i
                if i != 0 && (str[i-1] == '-') {
                    // 前面有正负号的话就begin
                    sign = -1
                }
            }
        } else {
            if begin != -1 {
                //当出现过数字后，出现一次不是数字就记录并跳出循环
                end = i
                break
            }
        }

        // 检测不符合条件的情况
        if flag && (str[i] != ' ' ) {
            if str[i] == '+' || str[i] == '-' {
                if i < len(str) - 1 && str[i+1] >= '0' && str[i+1] <= '9' {
                    continue
                }
            }
            return 0
        }
    }

    // 没发现数字就返回0哈
    if flag {
        return 0
    }
    // 到最后还是数字
    if end == -1 {
        end = len(str)
    }

    nums := str[begin:end]
    // 整理数字字符串
    for i := 0; i<len(nums); i++ {
        if result > (math.MaxInt32 - int(nums[i] - '0')*sign) / 10.0 {
            return math.MaxInt32
        }
        if  result < (math.MinInt32 - int(nums[i] - '0')*sign) / 10.0 {
            return math.MinInt32
        }
        result = result * 10 + int(nums[i] - '0') * sign
    }

    return result
}



/**
找第一个数字
如果前面存在过非空格的字符且该字符不是与数字直接相邻的+、-
之后找出这个数字之后的所有连续的数字，直至无数字或字符为止
取出数字判溢出 与上题同

没数字就返回0

*/
