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

// BatchCreateSearchDataSourceItem 为提高索引数据记录的速度, 特提供批量索引数据记录的接口
//
// 注意: 一个batch中所有数据项的datasourceID需要一致, tenantID也需要一致
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/batch_create
func (r *SearchService) BatchCreateSearchDataSourceItem(ctx context.Context, request *BatchCreateSearchDataSourceItemReq, options ...MethodOptionFunc) (*BatchCreateSearchDataSourceItemResp, *Response, error) {
	if r.cli.mock.mockSearchBatchCreateSearchDataSourceItem != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#BatchCreateSearchDataSourceItem mock enable")
		return r.cli.mock.mockSearchBatchCreateSearchDataSourceItem(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "BatchCreateSearchDataSourceItem",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources/:data_source_id/items/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchCreateSearchDataSourceItemResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchBatchCreateSearchDataSourceItem mock SearchBatchCreateSearchDataSourceItem method
func (r *Mock) MockSearchBatchCreateSearchDataSourceItem(f func(ctx context.Context, request *BatchCreateSearchDataSourceItemReq, options ...MethodOptionFunc) (*BatchCreateSearchDataSourceItemResp, *Response, error)) {
	r.mockSearchBatchCreateSearchDataSourceItem = f
}

// UnMockSearchBatchCreateSearchDataSourceItem un-mock SearchBatchCreateSearchDataSourceItem method
func (r *Mock) UnMockSearchBatchCreateSearchDataSourceItem() {
	r.mockSearchBatchCreateSearchDataSourceItem = nil
}

// BatchCreateSearchDataSourceItemReq ...
type BatchCreateSearchDataSourceItemReq struct {
	DataSourceID string                                    `path:"data_source_id" json:"-"` // 数据源的ID, 示例值: "service_ticket"
	Items        []*BatchCreateSearchDataSourceItemReqItem `json:"items,omitempty"`         // 待进入索引的item列表
}

// BatchCreateSearchDataSourceItemReqItem ...
type BatchCreateSearchDataSourceItemReqItem struct {
	ID             string                                          `json:"id,omitempty"`              // item 在 datasource 中的唯一标识, 示例值: "01010111"
	ACL            []*BatchCreateSearchDataSourceItemReqItemACL    `json:"acl,omitempty"`             // item 的访问权限控制。 acl 字段为空数组, 则默认数据不可见。如果数据是全员可见, 需要设置 access="allow"; type="user"; value="everyone"
	Metadata       *BatchCreateSearchDataSourceItemReqItemMetadata `json:"metadata,omitempty"`        // item 的元信息
	StructuredData string                                          `json:"structured_data,omitempty"` // 结构化数据（以 json 字符串传递）, 这些字段是搜索结果的展示字段（title字段无须在此另外指定）；, 示例值: "{"key":"value"}"
	Content        *BatchCreateSearchDataSourceItemReqItemContent  `json:"content,omitempty"`         // 非结构化数据, 如文档文本, 飞书搜索会用来做召回
}

// BatchCreateSearchDataSourceItemReqItemACL ...
type BatchCreateSearchDataSourceItemReqItemACL struct {
	Access *string `json:"access,omitempty"` // 权限类型, 优先级: Deny > Allow, 示例值: "allow", 可选值有: `allow`: 允许访问, `deny`: 禁止访问
	Value  *string `json:"value,omitempty"`  // 设置的权限值, 例如 userID, 依赖 type 描述, 注: 在 type 为 user 且 access 为 allow 时, 可填 "everyone" 来表示该数据项对全员可见；, 示例值: "d35e3c23"
	Type   *string `json:"type,omitempty"`   // 权限值类型, 示例值: "user", 可选值有: `user`: 访问权限控制中指定“用户”可以访问或拒绝访问该条数据
}

// BatchCreateSearchDataSourceItemReqItemContent ...
type BatchCreateSearchDataSourceItemReqItemContent struct {
	Format      *string `json:"format,omitempty"`       // 内容的格式, 示例值: "html", 可选值有: `html`: html格式, `plaintext`: 纯文本格式
	ContentData *string `json:"content_data,omitempty"` // 全文数据, 示例值: "这是一个很长的文本"
}

// BatchCreateSearchDataSourceItemReqItemMetadata ...
type BatchCreateSearchDataSourceItemReqItemMetadata struct {
	Title      string `json:"title,omitempty"`       // 该条数据记录对应的标题, 示例值: "工单: 无法创建文章"
	SourceURL  string `json:"source_url,omitempty"`  // 该条数据记录对应的跳转url, 示例值: "http://www.abc.com.cn"
	CreateTime *int64 `json:"create_time,omitempty"` // 数据项的创建时间。Unix 时间, 单位为秒, 示例值: 1618831236
	UpdateTime int64  `json:"update_time,omitempty"` // 数据项的更新时间。Unix 时间, 单位为秒, 示例值: 1618831236
}

// BatchCreateSearchDataSourceItemResp ...
type BatchCreateSearchDataSourceItemResp struct {
	Result *BatchCreateSearchDataSourceItemRespResult `json:"result,omitempty"` // 返回信息列表
}

// BatchCreateSearchDataSourceItemRespResult ...
type BatchCreateSearchDataSourceItemRespResult struct {
	ItemID    string `json:"item_id,omitempty"`    // itemID
	IsSuccess bool   `json:"is_success,omitempty"` // 该数据项是否成功发往索引, 注意: 存在极端情况该字段为True, 但并未进入索引
	Err       string `json:"err,omitempty"`        // 若该数据项索引失败, 则表示该数据的失败信息
}

// batchCreateSearchDataSourceItemResp ...
type batchCreateSearchDataSourceItemResp struct {
	Code int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                               `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateSearchDataSourceItemResp `json:"data,omitempty"`
}
