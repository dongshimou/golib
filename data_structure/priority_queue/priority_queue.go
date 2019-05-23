package priority_queue

type PriorityQueue struct {
	Compare func(interface{}, interface{}) bool
	heap    []interface{}
	count int
}

func NewPriorityQueue(compare func(interface{},interface{})bool)*PriorityQueue{
	pq:=PriorityQueue{
		Compare: compare,
		heap:    make([]interface{}, 0, 16),
		count:   0,
	}
	return &pq
}
func (this*PriorityQueue)Push(v interface{}){
	this.push(v)
}
func (this*PriorityQueue)Top()interface{}{
	return this.top()
}
func (this*PriorityQueue)Pop()interface{}{
	res:=this.top()
	this.pop()
	return res
}
func (this*PriorityQueue)Size()interface{}{
	return this.size()
}
func (this *PriorityQueue)Empty()bool{
	return this.empty()
}
//----------------------------------------------------------------------------------------------------------------------
func(this*PriorityQueue)push(v interface{}) {
	if this.count >= len(this.heap) {
		this.heap = append(this.heap, v)
	} else {
		this.heap[this.count] = v
	}
	this.siftUp(this.count)
	this.count++
}
func (this*PriorityQueue)top()interface{}{
	return this.heap[0]
}
func (this*PriorityQueue)pop(){
	this.count--
	this.heap[0]=this.heap[this.count]
	this.siftDown(0)
}
func (this*PriorityQueue)size()int{
	return this.count
}
func (this *PriorityQueue)empty()bool{
	return this.size()==0
}
func (this* PriorityQueue)siftUp(n int){
	v:=this.heap[n]
	for n2 := n / 2; n > 0 && this.Compare(v, this.heap[n2]); {
		this.heap[n]=this.heap[n2]
		n=n2
		n2/=2
	}
	this.heap[n]=v
}
func (this* PriorityQueue)siftDown(n int){
	v:=this.heap[n]
	if this.count<=1{
		return
	}
	for n2:=n*2;n2<this.count;{
		if n2+1<this.count&&this.Compare(this.heap[n2+1],this.heap[n2]){
			n2++
		}
		if this.Compare(v,this.heap[n2]) {
			break
		}
		this.heap[n]=this.heap[n2]
		n=n2
		n2*=2
	}
	this.heap[n]=v
}


