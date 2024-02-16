class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        carry = 0
        dummy = head = ListNode(0)
        
        while l1 or l2 or carry:
            a = l1.val if l1 else 0
            b = l2.val if l2 else 0
            
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
            
            carry, digit = divmod(a + b + carry, 10)
            head.next = ListNode(digit)
            head = head.next
        
        return dummy.next