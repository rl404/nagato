package mal

import "context"

// GetCharacterDetails to get character details.
//
// Undocumented.
func (c *Client) GetCharacterDetails(id int, fields ...string) (*Character, error) {
	return c.GetCharacterDetailsWithContext(context.Background(), id, fields...)
}

// GetCharacterDetailsWithContext to get character details with context.
//
// Undocumented.
func (c *Client) GetCharacterDetailsWithContext(ctx context.Context, id int, fields ...string) (*Character, error) {
	url := c.generateURL(map[string]interface{}{
		"fields": fields,
	}, "characters", id)

	var character Character
	if err := c.get(ctx, url, &character); err != nil {
		return nil, err
	}

	return &character, nil
}
