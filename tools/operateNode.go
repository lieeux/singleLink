package tools

import "fmt"

type Node struct {
	Id   int
	Name string
	next *Node //指向下一个节点
}

func InsertNode(head *Node, newNode *Node) { //在链表末尾添加节点
	temp := head //创建一个辅助节点
	for {        //找链表尾
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode //将链表末尾指向新节点
}

func InsertNodeByNum(head *Node, newNode *Node) { //按num顺序插入节点
	temp := head //创建一个辅助节点
	for {
		if temp.next == nil {
			break
		} else if temp.next.Id >= newNode.Id {
			break
		}
		temp = temp.next
	}
	newNode.next = temp.next //先把新节点指向下一个节点
	temp.next = newNode      //再把当前节点指向新节点，不然链就断了
}

func DelNode(head *Node, id int) { //删除节点
	temp := head //创建一个辅助节点
	for {
		if temp.next == nil {
			break
		} else if temp.next.Id == id {
			break
		}
		temp = temp.next
	}
	temp.next = temp.next.next //把当前节点指向下个节点的下个节点
}

func ListNode(head *Node) { //输出链表
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("[%d %s]->", temp.next.Id, temp.next.Name) //不输出链表头，所以输出的是下一个节点
		temp = temp.next                                      //先跳转
		if temp.next == nil {                                 //再判断
			break
		}
	}
	fmt.Println()
}
