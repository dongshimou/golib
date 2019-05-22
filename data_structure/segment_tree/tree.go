package segment_tree

//线段树

//节点
type Node struct {
	L     int //左区间index
	R     int //右区间index
	Deep  int //深度
	Index int //下标
	Value []interface{} //数量跟Functions一致
}

//函数 如 max,min,sum
//需要保证与线段树内的元素能够调用
type Function func(l, r interface{}) interface{}

//线段树结构
type SegmentTree struct {
	Functions []Function //需要计算的Function
	Element   [][]*Node  //所有的节点
	Root      *Node      //根节点
	Length    int        //长度 == len(Element[0])
	Capacity  int        //容量
}
//构造线段树
//functions 所有的函数
//args 初始化的区间,可以为nil
func New(functions []Function, args []interface{}) *SegmentTree {
	st := &SegmentTree{
		Functions: functions,
		Element:   [][]*Node{},
		Root:      nil,
		Length:    0,
		Capacity:  1,
	}
	st.Element = append(st.Element, []*Node{})
	for _, v := range args {
		st.Append(v)
	}
	return st
}

//查询线段树区间的func值
//l,r 范围为 [0,len)
func (st *SegmentTree) Query(l, r int) []interface{} {
	return query(l, r, st.Root, st)
}
func query(l, r int, root *Node, st *SegmentTree) []interface{} {
	if l <= root.L && r >= root.R {
		return root.Value
	}
	m := (root.L + root.R) >> 1
	var lres []interface{}
	var rres []interface{}
	if l <= m {
		if r < m {
			lres = query(l, r, st.Element[root.Deep-1][root.Index*2], st)
		} else {
			lres = query(l, m, st.Element[root.Deep-1][root.Index*2], st)
		}
	}
	if r > m {
		if m+1 < l {
			rres = query(l, r, st.Element[root.Deep-1][root.Index*2+1], st)
		} else {
			rres = query(m+1, r, st.Element[root.Deep-1][root.Index*2+1], st)
		}
	}
	if len(lres) != 0 && len(rres) != 0 {
		res := make([]interface{}, len(st.Functions))
		for i, f := range st.Functions {
			res[i] = f(lres[i], rres[i])
		}
		return res
	}
	if len(lres) != 0 {
		return lres
	} else {
		return rres
	}
}

//更新节点 pos范围 [0,len)
//大于等于 len 时,等同于append
func (st *SegmentTree) Update(pos int, value interface{}) {
	if pos >= st.Length {
		st.Append(value)
	} else {
		deep := 0
		index := pos
		//更新叶节点的值
		tmp := st.Element[deep][index]
		for i, _ := range tmp.Value {
			tmp.Value[i] = value
		}
		//更新值到根节点
		for {
			if deep == len(st.Element)-1 {
				break
			}
			//向上
			index >>= 1
			deep++
			//父节点
			ele := st.Element[deep][index]
			//左孩子
			left := st.Element[deep-1][index*2]
			//右孩子
			//可能没有右孩子
			if len(st.Element[deep-1]) <= index*2+1 {
				for i, _ := range st.Functions {
					ele.Value[i] = left.Value[i]
				}
			} else {
				right := st.Element[deep-1][index*2+1]
				for i, f := range st.Functions {
					ele.Value[i] = f(left.Value[i], right.Value[i])
				}
			}
		}
	}
}
func (st *SegmentTree) newNode(l, r, deep, index int, value interface{}) *Node {
	tmp := &Node{
		l,
		r,
		deep,
		index,
		make([]interface{}, len(st.Functions)),
	}
	for i, _ := range st.Functions {
		tmp.Value[i] = value
	}
	return tmp
}
func powLeft(s, d int) int {
	for {
		if d == 0 {
			break
		}
		d--
		s <<= 1
	}
	return s
}
func powRight(s, d int) int {
	for {
		if d == 0 {
			break
		}
		d--
		s <<= 1
		s += 1
	}
	return s
}

//新增节点到尾部
func (st *SegmentTree) Append(value interface{}) {
	deep := 0
	index := 0
	index = st.Length
	if st.Length >= st.Capacity {
		st.Capacity *= 2
	}

	tmp := st.newNode(index, index, deep, index, value)
	st.Element[deep] = append(st.Element[deep], tmp)
	st.Length++

	for {
		if index == 0 {
			break
		}
		//向上
		index >>= 1
		deep++
		//扩容
		if len(st.Element) <= deep {
			st.Element = append(st.Element, []*Node{})
		}
		//扩容
		if len(st.Element[deep]) <= index {
			st.Element[deep] = append(st.Element[deep], st.newNode(powLeft(index, deep), powRight(index, deep), deep, index, 0))
		}
		//当前父节点
		ele := st.Element[deep][index]
		//左孩子
		left := st.Element[deep-1][index*2]
		//当前节点
		//可能没有右孩子,当前节点就是左节点
		if len(st.Element[deep-1]) <= index*2+1 {
			for i, _ := range st.Functions {
				ele.Value[i] = left.Value[i]
			}
		} else {
			right := st.Element[deep-1][index*2+1]
			for i, f := range st.Functions {
				ele.Value[i] = f(left.Value[i], right.Value[i])
			}
		}
		tmp = ele
	}
	st.Root = tmp
}
