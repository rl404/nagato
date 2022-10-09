package nagato

import (
	"context"

	"github.com/rl404/nagato/mal"
)

// GetUserMangaList to get user manga list.
func (c *Client) GetUserMangaList(param GetUserMangaListParam, fields ...MangaField) ([]UserManga, error) {
	return c.GetUserMangaListWithContext(context.Background(), param, fields...)
}

// GetUserMangaListWithContext to get user manga list with context.
func (c *Client) GetUserMangaListWithContext(ctx context.Context, param GetUserMangaListParam, fields ...MangaField) ([]UserManga, error) {
	if err := c.validate(&param); err != nil {
		return nil, err
	}

	manga, err := c.mal.GetUserMangaListWithContext(ctx, mal.GetUserMangaListParam{
		Username: param.Username,
		Status:   string(param.Status),
		Nsfw:     param.NSFW,
		Sort:     string(param.Sort),
		Limit:    param.Limit,
		Offset:   param.Offset,
	}, c.mangaFieldsToStrs(fields...)...)
	if err != nil {
		return nil, err
	}

	return c.userMangaPagingToUserMangaList(manga), nil
}
