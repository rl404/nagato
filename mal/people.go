package mal

import "context"

// GetPeopleDetails to get people details.
//
// Undocumented.
func (c *Client) GetPeopleDetails(id int, fields ...string) (*People, error) {
	return c.GetPeopleDetailsWithContext(context.Background(), id, fields...)
}

// GetPeopleDetailsWithContext to get people details with context.
//
// Undocumented.
func (c *Client) GetPeopleDetailsWithContext(ctx context.Context, id int, fields ...string) (*People, error) {
	url := c.generateURL(map[string]interface{}{
		"fields": fields,
	}, "people", id)

	var people People
	if err := c.get(ctx, url, &people); err != nil {
		return nil, err
	}

	return &people, nil
}
