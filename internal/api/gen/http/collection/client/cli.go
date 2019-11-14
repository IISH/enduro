// Code generated by goa v3.0.7, DO NOT EDIT.
//
// collection HTTP client CLI support package
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"fmt"
	"strconv"

	collection "github.com/artefactual-labs/enduro/internal/api/gen/collection"
)

// BuildListPayload builds the payload for the collection list endpoint from
// CLI flags.
func BuildListPayload(collectionListOriginalID string, collectionListCursor string) (*collection.ListPayload, error) {
	var originalID *string
	{
		if collectionListOriginalID != "" {
			originalID = &collectionListOriginalID
		}
	}
	var cursor *string
	{
		if collectionListCursor != "" {
			cursor = &collectionListCursor
		}
	}
	payload := &collection.ListPayload{
		OriginalID: originalID,
		Cursor:     cursor,
	}
	return payload, nil
}

// BuildShowPayload builds the payload for the collection show endpoint from
// CLI flags.
func BuildShowPayload(collectionShowID string) (*collection.ShowPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionShowID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	payload := &collection.ShowPayload{
		ID: id,
	}
	return payload, nil
}

// BuildDeletePayload builds the payload for the collection delete endpoint
// from CLI flags.
func BuildDeletePayload(collectionDeleteID string) (*collection.DeletePayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionDeleteID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	payload := &collection.DeletePayload{
		ID: id,
	}
	return payload, nil
}

// BuildCancelPayload builds the payload for the collection cancel endpoint
// from CLI flags.
func BuildCancelPayload(collectionCancelID string) (*collection.CancelPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionCancelID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	payload := &collection.CancelPayload{
		ID: id,
	}
	return payload, nil
}

// BuildRetryPayload builds the payload for the collection retry endpoint from
// CLI flags.
func BuildRetryPayload(collectionRetryID string) (*collection.RetryPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionRetryID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	payload := &collection.RetryPayload{
		ID: id,
	}
	return payload, nil
}

// BuildWorkflowPayload builds the payload for the collection workflow endpoint
// from CLI flags.
func BuildWorkflowPayload(collectionWorkflowID string) (*collection.WorkflowPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionWorkflowID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	payload := &collection.WorkflowPayload{
		ID: id,
	}
	return payload, nil
}
