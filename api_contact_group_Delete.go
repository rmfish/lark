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

// DeleteContactGroup 通过该接口可删除企业中的用户组，请注意删除用户组时应用的通讯录权限范围需为“全部员工”，否则会删除失败，[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/delete
func (r *ContactService) DeleteContactGroup(ctx context.Context, request *DeleteContactGroupReq, options ...MethodOptionFunc) (*DeleteContactGroupResp, *Response, error) {
	if r.cli.mock.mockContactDeleteContactGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteContactGroup mock enable")
		return r.cli.mock.mockContactDeleteContactGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "DeleteContactGroup",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/:group_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteContactGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactDeleteContactGroup mock ContactDeleteContactGroup method
func (r *Mock) MockContactDeleteContactGroup(f func(ctx context.Context, request *DeleteContactGroupReq, options ...MethodOptionFunc) (*DeleteContactGroupResp, *Response, error)) {
	r.mockContactDeleteContactGroup = f
}

// UnMockContactDeleteContactGroup un-mock ContactDeleteContactGroup method
func (r *Mock) UnMockContactDeleteContactGroup() {
	r.mockContactDeleteContactGroup = nil
}

// DeleteContactGroupReq ...
type DeleteContactGroupReq struct {
	GroupID string `path:"group_id" json:"-"` // 需删除的用户组ID, 示例值："g1837191"
}

// deleteContactGroupResp ...
type deleteContactGroupResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteContactGroupResp `json:"data,omitempty"`
}

// DeleteContactGroupResp ...
type DeleteContactGroupResp struct {
}