package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 声明方法的方式很不一样, 需要声明为Node接受者的
// func Print(node Node) { print(node) 这样调用也是一样的, 
// 这里的node也是值传递, 也可以使用node *Node
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

// 工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value} //返回局部变量的地址, 不需要关心分配在堆还是栈上, go有垃圾回收机制
}
