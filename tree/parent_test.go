package tree

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/samber/lo"
)

func TestYyy(t *testing.T) {
	s := BuildTree[Company](companies, func(c Company) string {
		return c.Id
	}, func(c Company) string { return c.ParentId },
		func(c Company) string { return c.Name })

	b, err := json.Marshal(s)
	fmt.Println(string(b), err)
}

func TestXxx(t *testing.T) {

	grp := lo.GroupBy(companies, func(item Company) string {
		return item.ParentId
	})
	mp := lo.KeyBy(companies, func(item Company) string {
		return item.Id
	})

	var buildNode func(parentId string) Node
	buildNode = func(parentId string) Node {
		fmt.Println("cc", companies)
		tmp := mp[parentId]
		var top Node = Node{
			Id:       tmp.Id,
			Name:     tmp.Name,
			Children: make([]Node, 0),
		}
		if children, ok := grp[parentId]; ok {
			for _, child := range children {
				childNode := buildNode(child.Id)
				top.Children = append(top.Children, childNode)
			}
		}
		return top
	}

	fmt.Println("xx", grp)
	fmt.Println("mp", mp)

	tre := buildNode("")
	b, err := json.Marshal(tre)
	fmt.Println("tree", string(b), err)
}
