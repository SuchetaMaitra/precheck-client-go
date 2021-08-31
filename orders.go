package precheck

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetOrder - Returns a specifc order
func (c *Client) GetOrder(orderID string) (*Order, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/users?id=35", "http://localhost:3000"), nil)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", c.HostURL, orderID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
