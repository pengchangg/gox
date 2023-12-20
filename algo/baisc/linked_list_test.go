package basic

import (
	"reflect"
	"testing"
)

func TestListNode_Delete(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	var tests []struct {
		name   string
		fields fields
		want   *ListNode
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name   string
			fields fields
			want   *ListNode
		}{name: "", fields: fields{Val: i}, want: NewListNode(i)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := node.Delete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_Find(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	type args struct {
		val int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   *ListNode
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name   string
			fields fields
			args   args
			want   *ListNode
		}{name: "", fields: fields{Val: i}, args: args{val: i / 2}, want: NewListNode(i)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := node.Find(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_Insert(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	type args struct {
		val int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   *ListNode
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name   string
			fields fields
			args   args
			want   *ListNode
		}{name: "", fields: fields{Val: i}, args: args{val: i / 2}, want: NewListNode(i)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := node.Insert(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_Reverse(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	var tests []struct {
		name   string
		fields fields
		want   *ListNode
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name   string
			fields fields
			want   *ListNode
		}{name: "", fields: fields{Val: i}, want: NewListNode(i)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := node.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_Traverse(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	var tests []struct {
		name   string
		fields fields
		want   []int
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name   string
			fields fields
			want   []int
		}{name: "", fields: fields{Val: i}, want: []int{i}})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := node.Traverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewListNode(t *testing.T) {
	type args struct {
		val int
	}
	var tests []struct {
		name string
		args args
		want *ListNode
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			want *ListNode
		}{name: "", args: args{val: i}, want: NewListNode(i)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
