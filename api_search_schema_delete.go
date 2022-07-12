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

// DeleteSearchSchema 删除已存在的数据范式
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/schema/delete
func (r *SearchService) DeleteSearchSchema(ctx context.Context, request *DeleteSearchSchemaReq, options ...MethodOptionFunc) (*DeleteSearchSchemaResp, *Response, error) {
	if r.cli.mock.mockSearchDeleteSearchSchema != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#DeleteSearchSchema mock enable")
		return r.cli.mock.mockSearchDeleteSearchSchema(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "DeleteSearchSchema",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/schemas/:schema_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteSearchSchemaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchDeleteSearchSchema mock SearchDeleteSearchSchema method
func (r *Mock) MockSearchDeleteSearchSchema(f func(ctx context.Context, request *DeleteSearchSchemaReq, options ...MethodOptionFunc) (*DeleteSearchSchemaResp, *Response, error)) {
	r.mockSearchDeleteSearchSchema = f
}

// UnMockSearchDeleteSearchSchema un-mock SearchDeleteSearchSchema method
func (r *Mock) UnMockSearchDeleteSearchSchema() {
	r.mockSearchDeleteSearchSchema = nil
}

// DeleteSearchSchemaReq ...
type DeleteSearchSchemaReq struct {
	SchemaID string `path:"schema_id" json:"-"` // 用户自定义数据范式的唯一标识, 示例值: "custom_schema_id", 最大长度: `40` 字符, 正则校验: `^[a-zA-Z][a-zA-Z0-9-_].*$`
}

// DeleteSearchSchemaResp ...
type DeleteSearchSchemaResp struct {
}

// deleteSearchSchemaResp ...
type deleteSearchSchemaResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteSearchSchemaResp `json:"data,omitempty"`
}
