package mal

import "context"

// GetMangaDetails to get manga details.
func (c *Client) GetMangaDetails(id int, fields ...string) (*Manga, error) {
	return c.GetMangaDetailsWithContext(context.Background(), id, fields...)
}

// GetMangaDetailsWithContext to get manga details with context.
func (c *Client) GetMangaDetailsWithContext(ctx context.Context, id int, fields ...string) (*Manga, error) {
	url := c.generateURL(map[string]interface{}{
		"fields": fields,
	}, "manga", id)

	var manga Manga
	if err := c.get(ctx, url, &manga); err != nil {
		return nil, err
	}

	return &manga, nil
}

// GetMangaList to get manga list.
func (c *Client) GetMangaList(param GetMangaListParam, fields ...string) (*MangaPaging, error) {
	return c.GetMangaListWithContext(context.Background(), param, fields...)
}

//  GetMangaListWithContext to get manga list with context.
func (c *Client) GetMangaListWithContext(ctx context.Context, param GetMangaListParam, fields ...string) (*MangaPaging, error) {
	url := c.generateURL(map[string]interface{}{
		"q":      param.Query,
		"nsfw":   param.Nsfw,
		"limit":  param.Limit,
		"offset": param.Offset,
		"fields": fields,
	}, "manga")

	var manga MangaPaging
	if err := c.get(ctx, url, &manga); err != nil {
		return nil, err
	}

	return &manga, nil
}

// GetMangaRanking to get manga ranking list.
func (c *Client) GetMangaRanking(param GetMangaRankingParam, fields ...string) (*MangaRankingPaging, error) {
	return c.GetMangaRankingWithContext(context.Background(), param, fields...)
}

// GetMangaRankingWithContext to get manga ranking list with context.
func (c *Client) GetMangaRankingWithContext(ctx context.Context, param GetMangaRankingParam, fields ...string) (*MangaRankingPaging, error) {
	url := c.generateURL(map[string]interface{}{
		"ranking_type": param.RankingType,
		"nsfw":         param.Nsfw,
		"limit":        param.Limit,
		"offset":       param.Offset,
		"fields":       fields,
	}, "manga", "ranking")

	var manga MangaRankingPaging
	if err := c.get(ctx, url, &manga); err != nil {
		return nil, err
	}

	return &manga, nil
}
