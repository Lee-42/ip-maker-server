package appin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// StoryCreateInput is the input for creating a new story.
type StoryCreateInput struct {
	entity.Story
}

// StoryUpdateInput is the input for updating an existing story.
type StoryUpdateInput struct {
	entity.Story
}

// StoryDeleteInput is the input for deleting an existing story.
type StoryDeleteInput struct {
	Id int64 `json:"id" description:""`
}

// StoryViewInput is the input for viewing an existing story.
type StoryViewInput struct {
	Id int64 `json:"id" description:""`
}

// StoryListInput is the input for listing existing stories.
type StoryListInput struct {
	form.PageReq
	IpId   int64  `json:"ipId" description:"关联的IP ID"`
	ChatId string `json:"chatId" description:"关联的会话ID"`
	Title  string `json:"title" description:"故事标题"`
}

// StoryListOutput is the output for listing existing stories.
type StoryListOutput struct {
	List  []*entity.Story `json:"list"  description:"列表"`
	Total int             `json:"total" description:"总数"`
}
