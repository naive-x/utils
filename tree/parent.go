package tree

import "github.com/samber/lo"

type Node struct {
	Id       string
	Name     string
	Children []Node
}

type idGetter[T any] func(T) string
type nameGetter[T any] func(T) string

func BuildTree[T any](
	array []T,
	idGetter, parentIdGetter idGetter[T],
	nameGetter nameGetter[T],
) Node {
	grp := lo.GroupBy(array, func(item T) string {
		return parentIdGetter(item)
	})
	mp := lo.KeyBy(array, func(item T) string {
		return idGetter(item)
	})

	var buildNode func(parentId string) Node
	buildNode = func(parentId string) Node {
		tmp := mp[parentId]
		var top Node = Node{
			Id:       idGetter(tmp),
			Name:     nameGetter(tmp),
			Children: make([]Node, 0),
		}
		if children, ok := grp[parentId]; ok {
			for _, child := range children {
				childNode := buildNode(idGetter(child))
				top.Children = append(top.Children, childNode)
			}
		}
		return top
	}

	tree := buildNode("")
	return tree
}
