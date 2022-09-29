package mal

import (
	"context"
)

// GetAnimeDetails to get anime details.
func (c *Client) GetAnimeDetails(id int, fields ...string) (*Anime, error) {
	return c.GetAnimeDetailsWithContext(context.Background(), id, fields...)
}

// GetAnimeDetailsWithContext to get anime details with context.
func (c *Client) GetAnimeDetailsWithContext(ctx context.Context, id int, fields ...string) (*Anime, error) {
	url := c.generateURL(map[string]interface{}{
		"fields": fields,
	}, "anime", id)

	var anime Anime
	if err := c.get(ctx, url, &anime); err != nil {
		return nil, err
	}

	return &anime, nil
}

// GetAnimeList to get anime list.
func (c *Client) GetAnimeList(param GetAnimeListParam, fields ...string) (*AnimePaging, error) {
	return c.GetAnimeListWithContext(context.Background(), param, fields...)
}

//  GetAnimeListWithContext to get anime list with context.
func (c *Client) GetAnimeListWithContext(ctx context.Context, param GetAnimeListParam, fields ...string) (*AnimePaging, error) {
	url := c.generateURL(map[string]interface{}{
		"q":      param.Query,
		"nsfw":   param.Nsfw,
		"limit":  param.Limit,
		"offset": param.Offset,
		"fields": fields,
	}, "anime")

	var anime AnimePaging
	if err := c.get(ctx, url, &anime); err != nil {
		return nil, err
	}

	return &anime, nil
}

// GetAnimeRanking to get anime ranking list.
func (c *Client) GetAnimeRanking(param GetAnimeRankingParam, fields ...string) (*AnimeRankingPaging, error) {
	return c.GetAnimeRankingWithContext(context.Background(), param, fields...)
}

// GetAnimeRankingWithContext to get anime ranking list with context.
func (c *Client) GetAnimeRankingWithContext(ctx context.Context, param GetAnimeRankingParam, fields ...string) (*AnimeRankingPaging, error) {
	url := c.generateURL(map[string]interface{}{
		"ranking_type": param.RankingType,
		"nsfw":         param.Nsfw,
		"limit":        param.Limit,
		"offset":       param.Offset,
		"fields":       fields,
	}, "anime", "ranking")

	var anime AnimeRankingPaging
	if err := c.get(ctx, url, &anime); err != nil {
		return nil, err
	}

	return &anime, nil
}

// GetSeasonalAnime to get seasonal anime list.
func (c *Client) GetSeasonalAnime(param GetSeasonalAnimeParam, fields ...string) (*SeasonalAnimePaging, error) {
	return c.GetSeasonalAnimeWithContext(context.Background(), param, fields...)
}

// GetSeasonalAnimeWithContext to get seasonal anime list with context.
func (c *Client) GetSeasonalAnimeWithContext(ctx context.Context, param GetSeasonalAnimeParam, fields ...string) (*SeasonalAnimePaging, error) {
	url := c.generateURL(map[string]interface{}{
		"nsfw":   param.Nsfw,
		"sort":   param.Sort,
		"limit":  param.Limit,
		"offset": param.Offset,
		"fields": fields,
	}, "anime", "season", param.Year, param.Season)

	var anime SeasonalAnimePaging
	if err := c.get(ctx, url, &anime); err != nil {
		return nil, err
	}

	return &anime, nil
}

// GetSuggestedAnime to get suggested anime list.
//
// Need oauth2.
func (c *Client) GetSuggestedAnime(param GetSuggestedAnimeParam, fields ...string) (*SuggestedAnimePaging, error) {
	return c.GetSuggestedAnimeWithContext(context.Background(), param, fields...)
}

// GetSuggestedAnimeWithContext to get suggested anime list with context.
//
// Need oauth2.
func (c *Client) GetSuggestedAnimeWithContext(ctx context.Context, param GetSuggestedAnimeParam, fields ...string) (*SuggestedAnimePaging, error) {
	url := c.generateURL(map[string]interface{}{
		"nsfw":   param.Nsfw,
		"limit":  param.Limit,
		"offset": param.Offset,
		"fields": fields,
	}, "anime", "suggestions")

	var anime SuggestedAnimePaging
	if err := c.get(ctx, url, &anime); err != nil {
		return nil, err
	}

	return &anime, nil
}
