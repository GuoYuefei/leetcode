// 整数转罗马数字

func intToRoman(num int) string {
    var result string = ""
    mp := map[int]string{
        1: "I",
        5: "V",
        10: "X",
        50: "L",
        100: "C",
        500: "D",
        1000: "M",
    }
    temp := num
    for i := 0; temp > 0; i++ {
        pow10 := powInt10(i)
        str := ""
        t := temp % 10
        temp = temp / 10
        // 将t分析成5 1 的结合
        switch t {
            case 1,2,3:
                for t > 0 {
                    t--
                    str = str + mp[1 * pow10]
                }
            case 4:
                str = mp[pow10] + mp[5 * pow10]
            case 5,6,7,8:
                str = mp[5 * pow10]
                for t > 5 {
                    t--
                    str = str + mp[1 * pow10]
                }
            case 9:
                str = mp[1 * pow10] + mp[10 * pow10]
        }
        result = str + result
    }
    return result
}

func powInt10(x int) int {
    result := 1
    for x > 0 {
        x--
        result = result * 10
    }
    return result
}
