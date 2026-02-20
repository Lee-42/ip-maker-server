package story

import (
	"context"
	v1 "hotgo/api/api/ipmaker/v1/story"
	"hotgo/internal/service/ipmaker"
)

// CStory is the controller for story.
var CStory = cStory{}

type cStory struct{}

// Delete deletes a story.
func (c *cStory) Delete(ctx context.Context, req *v1.StoryDeleteReq) (res *v1.StoryDeleteReq, err error) {
	err = ipmaker.Story().Delete(ctx, req.StoryDeleteInput)
	return
}

// Edit edits a story.
func (c *cStory) Edit(ctx context.Context, req *v1.StoryUpdateReq) (res *v1.StoryUpdateReq, err error) {
	err = ipmaker.Story().Edit(ctx, req.StoryUpdateInput)
	return
}

// Add adds a new story.
func (c *cStory) Add(ctx context.Context, req *v1.StoryCreateReq) (res *v1.StoryCreateReq, err error) {
	err = ipmaker.Story().Add(ctx, req.StoryCreateInput)
	return
}

// List returns a list of stories.
func (c *cStory) List(ctx context.Context, req *v1.StoryListReq) (res *v1.StoryListRes, err error) {
	list, err := ipmaker.Story().List(ctx, req.StoryListInput)
	if err != nil {
		return
	}
	res = &v1.StoryListRes{
		StoryListOutput: list,
	}
	return
}

// View returns a story.
func (c *cStory) View(ctx context.Context, req *v1.StoryViewReq) (res *v1.StoryViewRes, err error) {
	view, err := ipmaker.Story().View(ctx, req.StoryViewInput)
	if err != nil {
		return
	}
	res = &v1.StoryViewRes{
		Story: view,
	}
	return
}
