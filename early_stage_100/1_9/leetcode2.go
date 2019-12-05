// 两数相加
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    CF := 0
    null := &ListNode{
        0,
        nil,
    }
    var result *ListNode = &ListNode{
        0,
        nil,
    }
    var l3 = result
    for  {
        // 如果不等于nil就说明还有
        l3.Val = l1.Val + l2.Val + CF
        if l3.Val > 9 {
            CF = 1
            l3.Val = l3.Val - 10
        } else {
            CF = 0
        }

        if l1.Next != nil || l2.Next != nil || CF == 1 {
            if l1.Next != nil {
                l1 = l1.Next
            } else {
                l1 = null
            }
            if l2.Next != nil {
                l2 = l2.Next
            } else {
                l2 = null
            }
            temp := &ListNode{
                0,
                nil,
            }
            l3.Next = temp
            l3 = temp
        } else {
            break
        }
    }

    return result
}


/**
直接从低位开始相加，进位记录，下次相加如有进位信号就+1，直到最长链表结束且CF==1
*/
