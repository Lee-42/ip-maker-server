package ipmaker

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
)

type (
	IStory interface {
		Delete(ctx context.Context, in *appin.StoryDeleteInput) (err error)
		Edit(ctx context.Context, in *appin.StoryUpdateInput) (err error)
		Add(ctx context.Context, in *appin.StoryCreateInput) (err error)
		List(ctx context.Context, in *appin.StoryListInput) (res *appin.StoryListOutput, err error)
		View(ctx context.Context, in *appin.StoryViewInput) (res *entity.Story, err error)
	}
)

var (
	localStory IStory
)

func Story() IStory {
	if localStory == nil {
		panic("implement not found for interface IStory, forgot register?")
	}
	return localStory
}

func RegisterStory(i IStory) {
	localStory = i
}
