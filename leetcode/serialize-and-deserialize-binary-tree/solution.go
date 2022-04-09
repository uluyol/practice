package leetcode

import (
	"encoding/binary"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var s encoder
	var sb strings.Builder
	s.encode(root, &sb)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	d := decoder{nodes: []*TreeNode{nil}}
	d.decode(data)
	if len(d.nodes) == 0 {
		return nil
	}
	return d.nodes[len(d.nodes)-1]
}

type decoder struct {
	nodes []*TreeNode
}

type encoder struct {
	nextID int32
}

func (enc *encoder) encode(n *TreeNode, sb *strings.Builder) int32 {
	if n == nil {
		return 0
	}
	rightID := enc.encode(n.Right, sb)
	leftID := enc.encode(n.Left, sb)

	enc.nextID++
	myID := enc.nextID
	var buf [16]byte
	binary.LittleEndian.PutUint32(buf[0:4], uint32(myID))
	binary.LittleEndian.PutUint32(buf[4:8], uint32(int32(n.Val)))
	binary.LittleEndian.PutUint32(buf[8:12], uint32(rightID))
	binary.LittleEndian.PutUint32(buf[12:16], uint32(leftID))

	sb.Write(buf[:])
	return myID
}

func (dec *decoder) decode(s string) {
	for s != "" {
		if len(s) < 16 {
			panic("invalid data")
		}
		myID := int(int32(binary.LittleEndian.Uint32([]byte(s[0:4]))))
		val := int(int32(binary.LittleEndian.Uint32([]byte(s[4:8]))))
		rightID := int(int32(binary.LittleEndian.Uint32([]byte(s[8:12]))))
		leftID := int(int32(binary.LittleEndian.Uint32([]byte(s[12:16]))))

		if leftID >= len(dec.nodes) || rightID >= len(dec.nodes) {
			panic("left or right ID too large")
		}

		if myID != len(dec.nodes) {
			panic("ids not in order")
		}
		dec.nodes = append(dec.nodes, &TreeNode{
			Val:   val,
			Right: dec.nodes[rightID],
			Left:  dec.nodes[leftID],
		})
		s = s[16:]
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
