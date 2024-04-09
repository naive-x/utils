package tree

type Company struct {
	Id       string
	Name     string
	ParentId string
}

var companies []Company = []Company{
	{
		Id:       "1-1",
		Name:     "一级-1",
		ParentId: "",
	},
	{
		Id:       "1-2",
		Name:     "一级-2",
		ParentId: "",
	},
	{
		Id:       "2-1",
		Name:     "二级-1",
		ParentId: "1-1",
	},
	{
		Id:       "2-2",
		Name:     "二级-2",
		ParentId: "1-1",
	},
}
