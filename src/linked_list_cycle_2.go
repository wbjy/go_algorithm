/**
 *环形链表 II
 *
 * 解题思路：一般这种判断有环的问题都可以用快慢指针来解决，
 * 有一个计算公式：当快慢指针相遇时，慢指针走过长度等于环外长a+环内距离b,快指针距离等于N*(b+c)+a+b,b+c是环长，
 * 此时快慢指针的路径关系为快指针是慢指针的路径的两倍：2(a+b) = N*(b+c)+a+b ==> 2(a+b) = a + (N+1)b+N*c ==> a = c + (N-1)*(b+c)
 * 所以可以得出，从相遇点到入环点的距离再加上n-1圈就等于出发点到入环点，所以当快慢指针相遇时，再找一个指针P从头出发，P和慢指针会在入环点相遇
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一：哈希表（效率不高）
//  var m map[*ListNode]int
// func detectCycle(head *ListNode) *ListNode {
//     m = make(map[*ListNode]int)
//     return dfs(head)
// }

// func dfs(head *ListNode) *ListNode {
//     if head == nil{
//         return nil
//     }

//     if _, ok := m[head];ok{
//         return head
//     }
//     m[head] = 1
//     lastNode := dfs(head.Next)
//     m[head] = 0
//     return lastNode

// }

// 解法二：快慢指针

func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil {
        slow = slow.Next
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next

        if fast == slow {
            p := head
            for p != slow {
                p = p.Next
                slow = slow.Next
            }
            return p
        }
    }

    return nil 
}