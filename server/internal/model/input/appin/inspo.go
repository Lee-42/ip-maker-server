package appin

// InspoCreateInput 创建灵感输入参数
type InspoCreateInput struct {
	Title    string
	Type     string
	Content  string
	Operator int64
}

// InspoUpdateInput 更新灵感输入参数
type InspoUpdateInput struct {
	Id       int64
	Title    *string
	Type     *string
	Content  *string
	Status   *int
	Operator int64
}

// InspoListInput 获取列表输入参数
type InspoListInput struct {
	Page   int
	Size   int
	Status *int
}
