package mal

import "context"

// GetUserInfo to get user info.
//
// Only `@me` in username param works for now.
//
// Need oauth2.
func (c *Client) GetUserInfo(username string, fields ...string) (*User, error) {
	return c.GetUserInfoWithContext(context.Background(), username, fields...)
}

// GetUserInfoWithContext to get user info with context.
//
// Only `@me` in username param works for now.
//
// Need oauth2.
func (c *Client) GetUserInfoWithContext(ctx context.Context, username string, fields ...string) (*User, error) {
	url := c.generateURL(map[string]interface{}{
		"fields": fields,
	}, "users", username)

	var user User
	if err := c.get(ctx, url, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
