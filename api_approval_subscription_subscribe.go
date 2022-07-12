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

// SubscribeApprovalSubscription 应用订阅 approval_code 后, 该应用就可以收到该审批定义对应实例的事件通知。同一应用只需要订阅一次, 无需重复订阅。
//
// 当应用不希望再收到审批事件时, 可以使用取消订阅接口进行取消, 取消后将不再给应用推送消息。
// 订阅和取消订阅都是应用维度的, 多个应用可以同时订阅同一个 approval_code, 每个应用都能收到审批事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/subscribe
func (r *ApprovalService) SubscribeApprovalSubscription(ctx context.Context, request *SubscribeApprovalSubscriptionReq, options ...MethodOptionFunc) (*SubscribeApprovalSubscriptionResp, *Response, error) {
	if r.cli.mock.mockApprovalSubscribeApprovalSubscription != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#SubscribeApprovalSubscription mock enable")
		return r.cli.mock.mockApprovalSubscribeApprovalSubscription(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "SubscribeApprovalSubscription",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/approvals/:approval_code/subscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(subscribeApprovalSubscriptionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalSubscribeApprovalSubscription mock ApprovalSubscribeApprovalSubscription method
func (r *Mock) MockApprovalSubscribeApprovalSubscription(f func(ctx context.Context, request *SubscribeApprovalSubscriptionReq, options ...MethodOptionFunc) (*SubscribeApprovalSubscriptionResp, *Response, error)) {
	r.mockApprovalSubscribeApprovalSubscription = f
}

// UnMockApprovalSubscribeApprovalSubscription un-mock ApprovalSubscribeApprovalSubscription method
func (r *Mock) UnMockApprovalSubscribeApprovalSubscription() {
	r.mockApprovalSubscribeApprovalSubscription = nil
}

// SubscribeApprovalSubscriptionReq ...
type SubscribeApprovalSubscriptionReq struct {
	ApprovalCode string `path:"approval_code" json:"-"` // 审批定义唯一标识, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
}

// SubscribeApprovalSubscriptionResp ...
type SubscribeApprovalSubscriptionResp struct {
}

// subscribeApprovalSubscriptionResp ...
type subscribeApprovalSubscriptionResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *SubscribeApprovalSubscriptionResp `json:"data,omitempty"`
}
