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

// SubscribeCalendarChangeEvent 该接口用于以用户身份订阅当前身份下日历列表中的所有日历变更。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/subscription
func (r *CalendarService) SubscribeCalendarChangeEvent(ctx context.Context, request *SubscribeCalendarChangeEventReq, options ...MethodOptionFunc) (*SubscribeCalendarChangeEventResp, *Response, error) {
	if r.cli.mock.mockCalendarSubscribeCalendarChangeEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#SubscribeCalendarChangeEvent mock enable")
		return r.cli.mock.mockCalendarSubscribeCalendarChangeEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Calendar",
		API:                 "SubscribeCalendarChangeEvent",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/subscription",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(subscribeCalendarChangeEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarSubscribeCalendarChangeEvent mock CalendarSubscribeCalendarChangeEvent method
func (r *Mock) MockCalendarSubscribeCalendarChangeEvent(f func(ctx context.Context, request *SubscribeCalendarChangeEventReq, options ...MethodOptionFunc) (*SubscribeCalendarChangeEventResp, *Response, error)) {
	r.mockCalendarSubscribeCalendarChangeEvent = f
}

// UnMockCalendarSubscribeCalendarChangeEvent un-mock CalendarSubscribeCalendarChangeEvent method
func (r *Mock) UnMockCalendarSubscribeCalendarChangeEvent() {
	r.mockCalendarSubscribeCalendarChangeEvent = nil
}

// SubscribeCalendarChangeEventReq ...
type SubscribeCalendarChangeEventReq struct {
}

// subscribeCalendarChangeEventResp ...
type subscribeCalendarChangeEventResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *SubscribeCalendarChangeEventResp `json:"data,omitempty"`
}

// SubscribeCalendarChangeEventResp ...
type SubscribeCalendarChangeEventResp struct {
}