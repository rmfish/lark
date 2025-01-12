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

// GetChat 获取群名称、群描述、群头像、群主 ID 等群基本信息。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 机器人或授权用户必须在群里（否则只会返回群名称、群头像等基本信息）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get
func (r *ChatService) GetChat(ctx context.Context, request *GetChatReq, options ...MethodOptionFunc) (*GetChatResp, *Response, error) {
	if r.cli.mock.mockChatGetChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#GetChat mock enable")
		return r.cli.mock.mockChatGetChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChat",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatGetChat mock ChatGetChat method
func (r *Mock) MockChatGetChat(f func(ctx context.Context, request *GetChatReq, options ...MethodOptionFunc) (*GetChatResp, *Response, error)) {
	r.mockChatGetChat = f
}

// UnMockChatGetChat un-mock ChatGetChat method
func (r *Mock) UnMockChatGetChat() {
	r.mockChatGetChat = nil
}

// GetChatReq ...
type GetChatReq struct {
	ChatID     string  `path:"chat_id" json:"-"`       // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_a0553eda9014c201e6969b478895c230"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetChatResp ...
type GetChatResp struct {
	Avatar                 string               `json:"avatar,omitempty"`                   // 群头像 URL
	Name                   string               `json:"name,omitempty"`                     // 群名称
	Description            string               `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames           `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    AddMemberPermission  `json:"add_member_permission,omitempty"`    // 群成员添加权限(all_members/only_owner)
	ShareCardPermission    ShareCardPermission  `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed)
	AtAllPermission        AtAllPermission      `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner)
	EditPermission         EditPermission       `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner)
	OwnerIDType            IDType               `json:"owner_id_type,omitempty"`            // 群主 ID 的类型(open_id/user_id/union_id), 群主是机器人时, 不返回该字段。
	OwnerID                string               `json:"owner_id,omitempty"`                 // 群主 ID, 群主是机器人时, 不返回该字段。
	ChatMode               ChatMode             `json:"chat_mode,omitempty"`                // 群模式(group/topic/p2p)
	ChatType               ChatType             `json:"chat_type,omitempty"`                // 群类型(private/public)
	ChatTag                string               `json:"chat_tag,omitempty"`                 // 优先级最高的一个群tag(inner/tenant/department/edu/meeting/customer_service)
	JoinMessageVisibility  MessageVisibility    `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility MessageVisibility    `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	MembershipApproval     MembershipApproval   `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	ModerationPermission   ModerationPermission `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner/moderator_list)
	External               bool                 `json:"external,omitempty"`                 // 是否是外部群
	TenantKey              string               `json:"tenant_key,omitempty"`               // tenant key
	UserCount              string               `json:"user_count,omitempty"`               // 群成员人数
	BotCount               string               `json:"bot_count,omitempty"`                // 群机器人数
}

// getChatResp ...
type getChatResp struct {
	Code int64        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string       `json:"msg,omitempty"`  // 错误描述
	Data *GetChatResp `json:"data,omitempty"`
}
