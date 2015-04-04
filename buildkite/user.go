// Copyright 2014 Mark Wolfe. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildkite

import "fmt"

// UserService handles communication with the user related
// methods of the buildkite API.
//
// buildkite API docs: https://buildkite.com/docs/api
type UserService struct {
	client *Client
}

// User represents a buildkite user.
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// Get the current user.
//
// buildkite API docs: https://buildkite.com/docs/api
func (os *OrganizationsService) Get() (*User, *Response, error) {
	var u string

	u = fmt.Sprintf("v1/user")

	req, err := os.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := os.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}