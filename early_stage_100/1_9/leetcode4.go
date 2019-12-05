// 寻找两个有序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    total := len(nums1) + len(nums2)

    var f int = total / 2 + 1

    result := make([]int, f)
    for i, j := 0, 0; (i < len(nums1) || j < len(nums2)) && i + j < f; {
        if i == len(nums1) {
            result[i+j] = nums2[j]
            j++
            continue
        }
        if j == len(nums2) {
            result[i+j] = nums1[i]
            i++
            continue
        }
        if nums1[i] < nums2[j] {
            result[i+j] = nums1[i]
            i++
        } else {
            result[i+j] = nums2[j]
            j++
        }
    }
    if total % 2 == 0 {
        return float64(result[f-1] + result[f-2])/2
    }

    return float64(result[f-1])
}


/**
    归并排序
    完成一半即可

*/
