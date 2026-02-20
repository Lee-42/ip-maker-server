package story

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
	"hotgo/internal/service/ipmaker"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// SStory is the service for story.
type SStory struct{}

func init() {
	ipmaker.RegisterStory(NewStory())
}

// NewStory returns a new SStory.
func NewStory() *SStory {
	return &SStory{}
}

// Delete deletes a story.
func (s *SStory) Delete(ctx context.Context, in *appin.StoryDeleteInput) (err error) {
	_, err = g.Model("story").Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit edits a story.
func (s *SStory) Edit(ctx context.Context, in *appin.StoryUpdateInput) (err error) {
	if in.Id <= 0 {
		return gerror.New("ID is required")
	}
	_, err = g.Model("story").Ctx(ctx).Where("id", in.Id).Data(in.Story).Update()
	return
}

// Add adds a new story.
func (s *SStory) Add(ctx context.Context, in *appin.StoryCreateInput) (err error) {
	_, err = g.Model("story").Ctx(ctx).Data(in.Story).Insert()
	return
}

// List returns a list of stories.
func (s *SStory) List(ctx context.Context, in *appin.StoryListInput) (res *appin.StoryListOutput, err error) {
	res = new(appin.StoryListOutput)
	mod := g.Model("story").Ctx(ctx)
	if in.IpId > 0 {
		mod = mod.Where("ip_id", in.IpId)
	}
	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
	}
	res.Total, err = mod.Count()
	if err != nil {
		return
	}
	if res.Total == 0 {
		res.List = []*entity.Story{}
		return
	}
	var list []*entity.Story
	err = mod.Page(in.PageReq.Page, in.PageReq.PerPage).Scan(&list)
	if err != nil {
		return
	}
	res.List = list
	return
}

// View returns a story.
func (s *SStory) View(ctx context.Context, in *appin.StoryViewInput) (res *entity.Story, err error) {
	err = g.Model("story").Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}
