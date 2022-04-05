package iJSON

type AvlNode struct {
	hash        uint
	name        string
	value       interface{}
	height      int
	left, right *AvlNode
}

func NewAvlNode(name string, value interface{}) *AvlNode {
	an := new(AvlNode)
	an.hash = HashToUInt(name)
	an.name = name
	an.value = value
	an.height = 0
	an.left = nil
	an.right = nil
	return an
}

func (an *AvlNode) Name() string {
	return an.name
}

func (an *AvlNode) Value() interface{} {
	return an.value
}

func (an *AvlNode) getHeight() int {
	if an == nil {
		return 0
	}
	return an.height
}

func (an *AvlNode) setHeight() {

	if an.left.getHeight() > an.right.getHeight() {
		an.height = an.left.getHeight() + 1
	} else {
		an.height = an.right.getHeight() + 1
	}
}

func (an *AvlNode) Height() int {
	if an == nil {
		return 0
	}
	hL := an.left.Height()
	hR := an.right.Height()
	if hL > hR {
		return hL + 1
	}
	return hR + 1
}

/*
LeftLeftRotate
左左旋转 使根结点的左结点变为新的根结点，其右结点变为老的根结点的左结点，老的根结点变为其右结点，
返回旋转后的根结点

           a                                b
	     /   \                            /   \
        b     c                          d     a
      /   \                               \    / \
      d    e                               f  e   c
       \
		f
*/
func (an *AvlNode) LeftLeftRotate() (rt *AvlNode) {
	rt = an.left
	an.left = rt.right
	rt.right = an
	an.setHeight()
	rt.setHeight()
	return rt
}

/*
RightRightRotate
右右旋转 使根结点的右结点变为新的根结点，其左结点变为老的根结点的右结点，老的根结点变为其左结点，
返回旋转后的根结点

           a                                c
	     /   \                            /   \
        b     c                          a     e
            /   \                       /  \    \
            d    e                      b   d    f
                /
		        f
*/
func (an *AvlNode) RightRightRotate() (rt *AvlNode) {
	rt = an.right
	an.right = rt.left
	rt.left = an
	an.setHeight()
	rt.setHeight()
	return rt
}

/*
LeftRightRotate
左右旋转 先对根结点的左结点右右旋转，再对根结点左左旋转
返回旋转后的根结点
*/
func (an *AvlNode) LeftRightRotate() (rt *AvlNode) {
	an.left = an.left.RightRightRotate()
	return an.LeftLeftRotate()
}

/*
RightLeftRotate
右左旋转 先对根结点的右结点左左旋转，再对根结点右右旋转
返回旋转后的根结点
*/
func (an *AvlNode) RightLeftRotate() (rt *AvlNode) {
	an.right = an.right.LeftLeftRotate()
	return an.RightRightRotate()
}

//保持平衡
func (an *AvlNode) balance() (rt *AvlNode) {
	rt = an
	if (rt.left.getHeight() - rt.right.getHeight()) == 2 {
		if rt.left.left.getHeight() > rt.left.right.getHeight() {
			rt = rt.LeftLeftRotate()
		} else {
			rt = rt.LeftRightRotate()
		}
	} else if (rt.right.getHeight() - rt.left.getHeight()) == 2 {
		if rt.right.right.getHeight() > rt.right.left.getHeight() {
			rt = rt.RightRightRotate()
		} else {
			rt = rt.RightLeftRotate()
		}
	}
	return rt
}

//获取最小节点
func (an *AvlNode) GetMinNode() (rt *AvlNode) {
	rt = an
	if rt.left != nil {
		rt = rt.left.GetMinNode()
	} else {
		rt = an
	}
	return rt
}

//获取最大节点
func (an *AvlNode) GetMaxNode() (rt *AvlNode) {
	rt = an
	if rt.right != nil {
		rt = rt.right.GetMinNode()
	} else {
		rt = an
	}
	return rt
}

func (an *AvlNode) GetNode(name string) (rt *AvlNode) {
	if an == nil {
		return nil
	}

	switch CompareName(name, an.Name()) {
	case -1:
		rt = an.left.GetNode(name)
	case 1:
		rt = an.right.GetNode(name)
	case 0:
		return an
	}
	return rt
}

func (an *AvlNode) Insert(in *AvlNode) (rt *AvlNode, code int) {
	if an == nil {
		in.setHeight()
		return in, 1
	}
	rt = an
	switch CompareAvlNode(in, an) {
	case +1:
		rt.right, code = rt.right.Insert(in)
	case -1:
		rt.left, code = rt.left.Insert(in)
	case 0:
		//数据已存在，覆盖
		rt.value = in.value
		return rt, 0
	}
	rt = rt.balance()
	rt.setHeight()
	return
}

/*
Remove
删除元素
*1、如果被删除结点只有一个子结点，就直接将A的子结点连至A的父结点上，并将A删除
*2、如果被删除结点有两个子结点，将该结点右子数内的最小结点取代A。
*3、平衡二叉树，重新设置高度
*/
func (an *AvlNode) Remove(name string) (rt *AvlNode, code int) {
	if an == nil {
		return nil, 0
	}
	rt = an
	switch CompareName(name, an.Name()) {
	case +1:
		rt.right, code = rt.right.Remove(name)
	case -1:
		rt.left, code = rt.left.Remove(name)
	case 0:
		if rt.left != nil && rt.right != nil {
			//结点有左孩子和右孩子
			rt.name = rt.right.GetMinNode().name
			rt.right, code = an.right.Remove(rt.name)
		} else {
			code = 1
			if rt.left != nil {
				//结点只有左孩子，无右孩子
				rt = rt.left
			} else {
				//结点只有右孩子或者无孩子
				rt = rt.right
			}
		}
	}

	if rt != nil {
		rt = rt.balance()
		rt.setHeight()
	}
	return
}

func (an *AvlNode) GetKeys() (rsArr []string) {
	if an == nil {
		return
	}
	arrLeft := an.left.GetKeys()
	arrRight := an.right.GetKeys()
	rsArr = append(rsArr, arrLeft...)
	rsArr = append(rsArr, arrRight...)
	rsArr = append(rsArr, an.name)
	return
}

func (an *AvlNode) GetValues() (rsArr []interface{}) {
	if an == nil {
		return
	}
	arrLeft := an.left.GetValues()
	arrRight := an.right.GetValues()
	rsArr = append(rsArr, arrLeft...)
	rsArr = append(rsArr, arrRight...)
	rsArr = append(rsArr, an.value)
	return
}

type AvlTree struct {
	root  *AvlNode
	count int
}

func NewAvlTree() *AvlTree {
	return &AvlTree{
		root:  nil,
		count: 0,
	}
}

func (at *AvlTree) Count() int {
	return at.count
}

func (at *AvlTree) Clear() {
	at.root = nil
	at.count = 0
}

func (at *AvlTree) IsEmpty() bool {
	return at.root == nil
}

func (at *AvlTree) Get(name string) *AvlNode {
	return at.root.GetNode(name)
}

func (at *AvlTree) Insert(in *AvlNode) {
	var code int
	at.root, code = at.root.Insert(in)
	at.count += code
}

func (at *AvlTree) Set(name string, value interface{}) {
	an := NewAvlNode(name, value)
	at.Insert(an)
}

func (at *AvlTree) Remove(name string) *AvlNode {
	var code int
	an := at.Get(name)
	if an == nil {
		return nil
	}
	at.root, code = at.root.Remove(name)
	at.count -= code
	return an
}

func (at *AvlTree) GetKeys() (rsArr []string) {
	return at.root.GetKeys()
}

func (at *AvlTree) GetValues() (rsArr []interface{}) {
	return at.root.GetValues()
}
