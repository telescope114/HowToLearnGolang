package main;
import (
	"fmt"
);

type ListNode struct {
	Val int64;
	Next *ListNode;
};

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var carry int64;
  var head, tail *ListNode;
  var sum int64;

  for l1 != nil || l2 != nil || carry != 0 {
	sum = 0

	if l1 != nil {
      sum += l1.Val;
	  l1 = l1.Next;
	}
	if l2 != nil {
	  sum += l2.Val;
	  l2 = l2.Next;
	}

	sum += carry;
	carry = sum / 10;
	newNode := &ListNode{Val: sum % 10};

	if head == nil {
	  head = newNode;
	  tail = newNode;
	} else {
	  tail.Next = newNode;
	  tail = newNode;
	}
  }

	return head;
}

func setListNode(intArr []int64) *ListNode {
  var out *ListNode
  var cur *ListNode

  // 初始化头节点
  for i, v := range intArr {
  	newNode := &ListNode{Val: v}
  	if i == 0 {
  		out = newNode
  		cur = out
  	} else {
  		cur.Next = newNode
  		cur = cur.Next
  	}
  }

  return out
}

func printlnListNode(listNode *ListNode) {
	if listNode == nil {
		return;
	}
	fmt.Println(listNode.Val);
	printlnListNode(listNode.Next);
}

func main () {
  var l1 *ListNode = setListNode([]int64{9,9,9,9,9,9,9});
  var l2 *ListNode = setListNode([]int64{9,9,9,9});
  printlnListNode(addTwoNumbers(l1, l2));
}