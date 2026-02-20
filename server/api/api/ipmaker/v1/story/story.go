package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
)

// StoryCreateReq is the request for creating a new story.
type StoryCreateReq struct {
	g.Meta `path:"/story/create" method:"post" tags:"Story前台管理" summary:"创建Story"`
	*appin.StoryCreateInput
}

// StoryUpdateReq is the request for updating an existing story.
type StoryUpdateReq struct {
	g.Meta `path:"/story/update" method:"post" tags:"Story前台管理" summary:"更新Story"`
	*appin.StoryUpdateInput
}

// StoryDeleteReq is the request for deleting an existing story.
type StoryDeleteReq struct {
	g.Meta `path:"/story/delete" method:"post" tags:"Story前台管理" summary:"删除Story"`
	*appin.StoryDeleteInput
}

// StoryViewReq is the request for viewing an existing story.
type StoryViewReq struct {
	g.Meta `path:"/story/view" method:"get" tags:"Story前台管理" summary:"获取单个Story信息"`
	*appin.StoryViewInput
}

// StoryListReq is the request for listing existing stories.
type StoryListReq struct {
	g.Meta `path:"/story/list" method:"get" tags:"Story前台管理" summary:"获取Story列表"`
	*appin.StoryListInput
}

// StoryListRes is the response for listing existing stories.
type StoryListRes struct {
	*appin.StoryListOutput
}

// StoryViewRes is the response for viewing an existing story.
type StoryViewRes struct {
	*entity.Story
}
