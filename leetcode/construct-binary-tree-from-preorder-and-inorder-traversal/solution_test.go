package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		want     []int
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     []int{3, 9, 20, -666, -666, 15, 7},
		},
	}

	for _, test := range tests {
		t.Logf("preorder = %v", test.preorder)
		t.Logf("inorder = %v", test.inorder)
		got := buildTree(test.preorder, test.inorder).ToList()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
