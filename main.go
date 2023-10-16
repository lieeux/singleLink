package main

import "singleLink/tools"

func main() {
	head := &tools.Node{} //初始化头结点

	hero1 := &tools.Node{
		Id:   1,
		Name: "穆",
	}
	hero2 := &tools.Node{
		Id:   2,
		Name: "阿鲁迪巴",
	}
	hero3 := &tools.Node{
		Id:   3,
		Name: "撒加",
	}
	hero4 := &tools.Node{
		Id:   4,
		Name: "加隆",
	}

	tools.InsertNodeByNum(head, hero4)
	tools.InsertNodeByNum(head, hero3)
	tools.InsertNodeByNum(head, hero2)
	tools.InsertNodeByNum(head, hero1)

	tools.ListNode(head)

	tools.DelNode(head, 3)

	tools.ListNode(head)
}
