package list

type List struct {
	Head []interface{}
	Tail []interface{}
}

func New()*List {
	dq:=&List{
		Head:make([]interface{},0,16),
		Tail:make([]interface{},0,16),
	}
	return dq
}

func(this*List) PushFront(v interface{}){
	this.push_front(v)
}
func(this *List)PushBack(v interface{}){
	this.push_back(v)
}
func (this *List)GetPopFront()interface{}{
	res:=this.front()
	this.pop_front()
	return res
}
func(this*List)PopFront(){
	this.pop_front()
}
func(this*List)PopBack(){
	this.pop_back()
}
func(this *List)GetPopBack()interface{}{
	res:=this.back()
	this.pop_back()
	return res
}
func (this*List)Front()interface{}{
	return this.front()
}
func(this*List)Back()interface{}{
	return this.back()
}
func(this*List)Size()int{
	return this.size()
}
func(this*List)Empty()bool{
	return this.empty()
}

//not safe function
func (this *List)push_front(v interface{}) {
	this.Head=append(this.Head,v)
}
func (this *List)push_back(v interface{}){
	this.Tail=append(this.Tail,v)
}
func (this *List)pop_front(){
	if len(this.Head)!=0{
		this.Head=this.Head[:len(this.Head)-1]
	}else{
		this.Tail=this.Tail[1:]
	}
}
func (this *List)pop_back(){
	if len(this.Tail)!=0{
		this.Tail=this.Tail[:len(this.Tail)-1]
	}else{
		this.Head=this.Head[1:]
	}
}

func (this*List)front()interface{}{
	if len(this.Head)!=0{
		return this.Head[len(this.Head)-1]
	}else{
		return this.Tail[0]
	}
}
func (this*List)back()interface{}{
	if len(this.Tail)!=0{
		return this.Tail[len(this.Tail)-1]
	}else{
		return this.Head[0]
	}
}

func (this *List)size()int{
	return len(this.Tail)+len(this.Head)
}
func (this *List)empty()bool {
	return this.size()==0
}

