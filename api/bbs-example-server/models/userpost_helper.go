// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "bbs-example-server": Model Helpers
//
// Command:
// $ goagen
// --design=BBS-Example/api/bbs-example-server/design
// --out=$(GOPATH)/src/BBS-Example/api/bbs-example-server
// --version=v1.3.1

package models

import (
	"BBS-Example/api/bbs-example-server/app"
	"context"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"time"
)

// MediaType Retrieval Functions

// ListJSON returns an array of view: default.
func (m *UserPostDB) ListJSON(ctx context.Context) []*app.JSON {
	defer goa.MeasureSince([]string{"goa", "db", "json_", "listjson_"}, time.Now())

	var native []*UserPost
	var objs []*app.JSON
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing UserPost", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserPostToJSON())
	}

	return objs
}

// UserPostToJSON loads a UserPost and builds the default view of media type Json.
func (m *UserPost) UserPostToJSON() *app.JSON {
	userpost := &app.JSON{}
	userpost.Message = &m.Message

	return userpost
}

// OneJSON loads a UserPost and builds the default view of media type Json.
func (m *UserPostDB) OneJSON(ctx context.Context, id int) (*app.JSON, error) {
	defer goa.MeasureSince([]string{"goa", "db", "json_", "onejson_"}, time.Now())

	var native UserPost
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting UserPost", "error", err.Error())
		return nil, err
	}

	view := *native.UserPostToJSON()
	return &view, err
}