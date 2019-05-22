package segment_tree

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

var (
	maxFunc = func(l, r interface{}) interface{} {
		if l.(int) > r.(int) {
			return l
		} else {
			return r
		}
	}
	minFunc = func(l, r interface{}) interface{} {
		if l.(int) > r.(int) {
			return r
		} else {
			return l
		}
	}
	sumFunc = func(l, r interface{}) interface{} {
		return l.(int) + r.(int)
	}
)

func TestNew(t *testing.T) {
	type args struct {
		functions []Function
		args      []interface{}
	}
	tests := []struct {
		name string
		args args
		want *SegmentTree
	}{
		// TODO: Add test cases.
		{
			"test new",
			args{
				[]Function{maxFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.functions, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegmentTree_Query(t *testing.T) {
	type fields struct {
		Functions []Function
		Args      []interface{}
	}
	type args struct {
		l int
		r int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		// TODO: Add test cases.
		{
			"test max1",
			fields{
				[]Function{maxFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			args{1, 2},
			[]interface{}{3},
		},

		{
			"test max2",
			fields{
				[]Function{maxFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			args{0, 5},
			[]interface{}{5},
		},
		{
			"test max3",
			fields{
				[]Function{maxFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			args{3, 4},
			[]interface{}{5},
		},

		{
			"test min1",
			fields{
				[]Function{minFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			args{3, 4},
			[]interface{}{4},
		},
		{
			"test min2",
			fields{
				[]Function{minFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			args{2, 4},
			[]interface{}{2},
		},
		{
			"test min3",
			fields{
				[]Function{minFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			args{0, 4},
			[]interface{}{1},
		},

		{
			"test max min,sum",
			fields{
				[]Function{maxFunc, minFunc, sumFunc},
				[]interface{}{1, 3, 2, 4, 5},
			},
			args{0, 4},
			[]interface{}{5, 1, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := New(tt.fields.Functions, tt.fields.Args)
			if got := st.Query(tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SegmentTree.Query() = %v, want %v", got, tt.want)
			} else {
				t.Logf("SegmentTree.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMooc(t *testing.T) {
	var want []interface{}
	{
		list := []int{rand.Int() % 1000}
		//list中最近20个数里最大的
		max20 := func() int {
			max := -math.MaxInt32
			for i := 0; i < 20; i++ {
				pos := len(list) - 1 - i
				if list[pos] > max {
					max = list[pos]
				}
			}
			return max
		}
		//list中最近20个数里最小的
		min20 := func() int {
			min:=math.MaxInt32
			for i:=0;i<20;i++{
				pos:=len(list)-1-i
				if list[pos]<min{
					min=list[pos]
				}
			}
			return min
		}
		//list最后一个数
		last := func() int {
			return list[len(list)-1]
		}
		//模拟期货的涨跌,构造假数据
		mooc := func(base int) int {
			pos := true
			//涨或者跌
			if rand.Int()%2 == 1 {
				pos = false
			}
			//每次不超过10%
			rate := float64(rand.Int()%10) / 100
			if pos {
				return int(float64(base) * (1 + rate))
			} else {
				return int(float64(base) * (1 - rate))
			}
		}
		//求区间最大值和最小值的线段树
		st := New([]Function{maxFunc,minFunc}, nil)
		//新增了19个数,共计20
		for i := 0; i < 19; i++ {
			tmp := mooc(last())
			st.Append(tmp)
			list = append(list, tmp)
		}

		t.Log(list)
		//模拟 新数据出现100次
		for i := 0; i < 100; i++ {

			tmp := mooc(last())
			st.Append(tmp)
			list = append(list, tmp)

			want = []interface{}{max20(),min20()}
			t.Logf("push data:%d ,pop data:%d", tmp, list[len(list)-21])

			if got := st.Query(st.Length-20, st.Length-1); !reflect.DeepEqual(got, want) {
				t.Errorf("get %v,want %v", got, want)
				t.Log(list)
			} else {
				t.Logf("get %v,want %v", got, want)
			}
		}
	}
}

func TestFree(t *testing.T) {
	var want []interface{}

	{
		st := New([]Function{maxFunc, minFunc}, []interface{}{1, 3, 2, 4, 5})

		want = []interface{}{5, 4}
		if got := st.Query(3, 4); !reflect.DeepEqual(got, want) {
			t.Errorf("get %v,want %v", got, want)
		} else {
			t.Logf("get %v,want %v", got, want)
		}
		st.Append(6)

		want = []interface{}{6, 1}
		if got := st.Query(0, 5); !reflect.DeepEqual(got, want) {
			t.Errorf("get %v,want %v", got, want)
		} else {
			t.Logf("get %v,want %v", got, want)
		}
	}

	{
		st := New([]Function{sumFunc}, []interface{}{5, 18, 13})
		want = []interface{}{36}
		if got := st.Query(0, 2); !reflect.DeepEqual(got, want) {
			t.Errorf("get %v,want %v", got, want)
		} else {
			t.Logf("get %v,want %v", got, want)
		}

		st.Update(1, -1)
		st.Update(2, 3)
		st.Update(0, 5)
		st.Update(0, -4)

		want = []interface{}{-2}
		if got := st.Query(0, 2); !reflect.DeepEqual(got, want) {
			t.Errorf("get %v,want %v", got, want)
		} else {
			t.Logf("get %v,want %v", got, want)
		}
	}
	{
		st := New([]Function{sumFunc}, nil)

		st.Append(2)
		st.Append(3)
		want = []interface{}{5}
		if got := st.Query(0, 1); !reflect.DeepEqual(got, want) {
			t.Errorf("get %v,want %v", got, want)
		} else {
			t.Logf("get %v,want %v", got, want)
		}
	}

}
