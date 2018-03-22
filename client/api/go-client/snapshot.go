//
// Copyright (c) 2018 The heketi Authors
//
// This file is licensed to you under your choice of the GNU Lesser
// General Public License, version 3 or any later version (LGPLv3 or
// later), as published by the Free Software Foundation,
// or under the Apache License, Version 2.0 <LICENSE-APACHE2 or
// http://www.apache.org/licenses/LICENSE-2.0>.
//
// You may not use this file except in compliance with those terms.
//

package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/heketi/heketi/pkg/utils"
)

func (c *Client) SnapshotClone(id string, scr *api.SnapshotCloneRequest) (*api.SnapshotInfoResponse, error) {
	// Marshal request to JSON
	buffer, err := json.Marshal(scr)
	if err != nil {
		return nil, err
	}

	// Create a request
	req, err := http.NewRequest("POST", c.host+"/snapshots/"+id+"/clone", bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Set token
	err = c.setToken(req)
	if err != nil {
		return nil, err
	}

	// Send request
	r, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusAccepted {
		return nil, utils.GetErrorFromResponse(r)
	}

	// Wait for response
	r, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, utils.GetErrorFromResponse(r)
	}

	// Read JSON response
	var snapshot api.SnapshotInfoResponse
	err = utils.GetJsonFromResponse(r, &snapshot)
	if err != nil {
		return nil, err
	}

	return &snapshot, nil

}

func (c *Client) SnapshotList() (*api.SnapshotListResponse, error) {

	// Create request
	req, err := http.NewRequest("GET", c.host+"/snapshots", nil)
	if err != nil {
		return nil, err
	}

	// Set token
	err = c.setToken(req)
	if err != nil {
		return nil, err
	}

	// Get info
	r, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, utils.GetErrorFromResponse(r)
	}

	// Read JSON response
	var snapshots api.SnapshotListResponse
	err = utils.GetJsonFromResponse(r, &snapshots)
	if err != nil {
		return nil, err
	}

	return &snapshots, nil
}

func (c *Client) SnapshotInfo(id string) (*api.SnapshotInfoResponse, error) {

	// Create request
	req, err := http.NewRequest("GET", c.host+"/snapshots/"+id, nil)
	if err != nil {
		return nil, err
	}

	// Set token
	err = c.setToken(req)
	if err != nil {
		return nil, err
	}

	// Get info
	r, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, utils.GetErrorFromResponse(r)
	}

	// Read JSON response
	var snapshot api.SnapshotInfoResponse
	err = utils.GetJsonFromResponse(r, &snapshot)
	if err != nil {
		return nil, err
	}

	return &snapshot, nil
}

func (c *Client) SnapshotDelete(snapshot_uuid string) error {

	// Create a request
	req, err := http.NewRequest("DELETE", c.host+"/snapshots/"+snapshot_uuid, nil)
	if err != nil {
		return err
	}

	// Set token
	err = c.setToken(req)
	if err != nil {
		return err
	}

	// Send request
	r, err := c.do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusAccepted {
		return utils.GetErrorFromResponse(r)
	}

	// Wait for response
	r, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusNoContent {
		return utils.GetErrorFromResponse(r)
	}

	return nil
}
