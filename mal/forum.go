package mal

import "context"

// GetForumBoards to get forum board list.
func (c *Client) GetForumBoards() (*ForumBoardCategories, error) {
	return c.GetForumBoardsWithContext(context.Background())
}

// GetForumBoardsWithContext to get forum board list with context.
func (c *Client) GetForumBoardsWithContext(ctx context.Context) (*ForumBoardCategories, error) {
	url := c.generateURL(nil, "forum", "boards")

	var forumBoardCategories ForumBoardCategories
	if err := c.get(ctx, url, &forumBoardCategories); err != nil {
		return nil, err
	}

	return &forumBoardCategories, nil
}

// GetForumTopics to get forum topic list.
func (c *Client) GetForumTopics(param GetForumTopicsParam) (*ForumTopicPaging, error) {
	return c.GetForumTopicsWithContext(context.Background(), param)
}

// GetForumTopicsWithContext to get forum topic list with context.
func (c *Client) GetForumTopicsWithContext(ctx context.Context, param GetForumTopicsParam) (*ForumTopicPaging, error) {
	url := c.generateURL(map[string]interface{}{
		"board_id":        param.BoardID,
		"subboard_id":     param.SubboardID,
		"limit":           param.Limit,
		"offset":          param.Offset,
		"sort":            param.Sort,
		"q":               param.Query,
		"topic_user_name": param.TopicUsername,
		"user_name":       param.Username,
	}, "forum", "topics")

	var forumTopic ForumTopicPaging
	if err := c.get(ctx, url, &forumTopic); err != nil {
		return nil, err
	}

	return &forumTopic, nil
}

// GetForumTopicDetails to get forum topic details.
func (c *Client) GetForumTopicDetails(param GetForumTopicDetailsParam) (*ForumTopicDetailPaging, error) {
	return c.GetForumTopicDetailsWithContext(context.Background(), param)
}

// GetForumTopicDetailsWithContext to get forum topic details with context.
func (c *Client) GetForumTopicDetailsWithContext(ctx context.Context, param GetForumTopicDetailsParam) (*ForumTopicDetailPaging, error) {
	url := c.generateURL(map[string]interface{}{
		"limit":  param.Limit,
		"offset": param.Offset,
	}, "forum", "topic", param.ID)

	var forumTopic ForumTopicDetailPaging
	if err := c.get(ctx, url, &forumTopic); err != nil {
		return nil, err
	}

	return &forumTopic, nil
}
