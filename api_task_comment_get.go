// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetTaskComment 该接口用于通过评论ID获取评论详情
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/get
func (r *TaskService) GetTaskComment(ctx context.Context, request *GetTaskCommentReq, options ...MethodOptionFunc) (*GetTaskCommentResp, *Response, error) {
	if r.cli.mock.mockTaskGetTaskComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#GetTaskComment mock enable")
		return r.cli.mock.mockTaskGetTaskComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "GetTaskComment",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/comments/:comment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getTaskCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskGetTaskComment mock TaskGetTaskComment method
func (r *Mock) MockTaskGetTaskComment(f func(ctx context.Context, request *GetTaskCommentReq, options ...MethodOptionFunc) (*GetTaskCommentResp, *Response, error)) {
	r.mockTaskGetTaskComment = f
}

// UnMockTaskGetTaskComment un-mock TaskGetTaskComment method
func (r *Mock) UnMockTaskGetTaskComment() {
	r.mockTaskGetTaskComment = nil
}

// GetTaskCommentReq ...
type GetTaskCommentReq struct {
	TaskID    string `path:"task_id" json:"-"`    // 任务ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	CommentID string `path:"comment_id" json:"-"` // 评论ID, 示例值: "6937231762296684564"
}

// GetTaskCommentResp ...
type GetTaskCommentResp struct {
	Comment *GetTaskCommentRespComment `json:"comment,omitempty"` // 返回新的任务评论详情
}

// GetTaskCommentRespComment ...
type GetTaskCommentRespComment struct {
	Content         string `json:"content,omitempty"`           // 评论内容
	ParentID        string `json:"parent_id,omitempty"`         // 评论的父ID, 创建评论时若不为空则为某条评论的回复, 若为空则不是回复
	ID              string `json:"id,omitempty"`                // 评论ID, 由飞书服务器发号
	CreateMilliTime string `json:"create_milli_time,omitempty"` // 评论创建的时间戳, 单位为毫秒, 用于展示, 创建时不用填写
}

// getTaskCommentResp ...
type getTaskCommentResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetTaskCommentResp `json:"data,omitempty"`
}
