// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateCalendarEvent
//
// 该接口用于以当前身份（应用 / 用户）在日历上创建一个日程。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份必须对日历有 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/create
func (r *CalendarService) CreateCalendarEvent(ctx context.Context, request *CreateCalendarEventReq, options ...MethodOptionFunc) (*CreateCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarCreateCalendarEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#CreateCalendarEvent mock enable")
		return r.cli.mock.mockCalendarCreateCalendarEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "CreateCalendarEvent",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockCalendarCreateCalendarEvent(f func(ctx context.Context, request *CreateCalendarEventReq, options ...MethodOptionFunc) (*CreateCalendarEventResp, *Response, error)) {
	r.mockCalendarCreateCalendarEvent = f
}

func (r *Mock) UnMockCalendarCreateCalendarEvent() {
	r.mockCalendarCreateCalendarEvent = nil
}

type CreateCalendarEventReq struct {
	CalendarID       string                            `path:"calendar_id" json:"-"`        // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	Summary          *string                           `json:"summary,omitempty"`           // 日程标题, 示例值："日程标题", 最大长度：`1000` 字符
	Description      *string                           `json:"description,omitempty"`       // 日程描述, 示例值："日程描述", 最大长度：`8192` 字符
	NeedNotification *bool                             `json:"need_notification,omitempty"` // 更新日程是否给日程参与人发送bot通知，默认为true, 示例值：false
	StartTime        *CreateCalendarEventReqStartTime  `json:"start_time,omitempty"`        // 日程开始时间
	EndTime          *CreateCalendarEventReqEndTime    `json:"end_time,omitempty"`          // 日程结束时间
	Vchat            *CreateCalendarEventReqVchat      `json:"vchat,omitempty"`             // 视频会议信息。
	Visibility       *string                           `json:"visibility,omitempty"`        // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 示例值："default", 可选值有: `default`：默认权限，跟随日历权限，默认仅向他人显示是否“忙碌”, `public`：公开，显示日程详情, `private`：私密，仅自己可见详情
	AttendeeAbility  *string                           `json:"attendee_ability,omitempty"`  // 参与人权限, 示例值："can_see_others", 可选值有: `none`：无法编辑日程、无法邀请其它参与人、无法查看参与人列表, `can_see_others`：无法编辑日程、无法邀请其它参与人、可以查看参与人列表, `can_invite_others`：无法编辑日程、可以邀请其它参与人、可以查看参与人列表, `can_modify_event`：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus   *string                           `json:"free_busy_status,omitempty"`  // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 示例值："busy", 可选值有: `busy`：忙碌, `free`：空闲
	Location         *CreateCalendarEventReqLocation   `json:"location,omitempty"`          // 日程地点
	Color            *int64                            `json:"color,omitempty"`             // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。, 示例值：-1
	Reminders        []*CreateCalendarEventReqReminder `json:"reminders,omitempty"`         // 日程提醒列表
	Recurrence       *string                           `json:"recurrence,omitempty"`        // 重复日程的重复性规则；参考[rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10)；, 不支持COUNT和UNTIL同时出现；, 预定会议室重复日程长度不得超过两年。, 示例值："FREQ=DAILY;INTERVAL=1", 最大长度：`2000` 字符
	Schemas          []*CreateCalendarEventReqSchema   `json:"schemas,omitempty"`           // 日程自定义信息
}

type CreateCalendarEventReqStartTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定, 示例值：" 2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区), 示例值："1605024000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai, 示例值："Asia/Shanghai"
}

type CreateCalendarEventReqEndTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定, 示例值：" 2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区), 示例值："1605024000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai, 示例值："Asia/Shanghai"
}

type CreateCalendarEventReqVchat struct {
	VCType      *string `json:"vc_type,omitempty"`     // 视频会议类型, 示例值："third_party", 可选值有: `vc`：飞书视频会议，取该类型时，其他字段无效。, `third_party`：第三方链接视频会议，取该类型时，icon_type、description、meeting_url字段生效。, `no_meeting`：无视频会议，取该类型时，其他字段无效。, `lark_live`：Lark直播，内部类型，只读。, `unknown`：未知类型，做兼容使用，只读。
	IconType    *string `json:"icon_type,omitempty"`   // 第三方视频会议icon类型；可以为空，为空展示默认icon。, 示例值："vc", 可选值有: `vc`：飞书视频会议icon, `live`：直播视频会议icon, `default`：默认icon
	Description *string `json:"description,omitempty"` // 第三方视频会议文案，可以为空，为空展示默认文案, 示例值："发起视频会议", 长度范围：`0` ～ `500` 字符
	MeetingURL  *string `json:"meeting_url,omitempty"` // 视频会议URL, 示例值："https://example.com", 长度范围：`1` ～ `2000` 字符
}

type CreateCalendarEventReqLocation struct {
	Name      *string  `json:"name,omitempty"`      // 地点名称, 示例值："地点名称", 长度范围：`1` ～ `512` 字符
	Address   *string  `json:"address,omitempty"`   // 地点地址, 示例值："地点地址", 长度范围：`1` ～ `255` 字符
	Latitude  *float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准, 示例值：xxxxx
	Longitude *float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准, 示例值：xxxxx
}

type CreateCalendarEventReqReminder struct {
	Minutes *int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒,新建或更新日程时传入该字段，仅对当前身份生效, 示例值：5, 取值范围：`-20160` ～ `20160`
}

type CreateCalendarEventReqSchema struct {
	UiName   *string `json:"ui_name,omitempty"`   // UI名称。取值范围如下： \,ForwardIcon: 日程转发按钮 \,MeetingChatIcon: 会议群聊按钮 \,MeetingMinutesIcon: 会议纪要按钮 \,MeetingVideo: 视频会议区域 \,RSVP: 接受/拒绝/待定区域 \,Attendee: 参与者区域 \,OrganizerOrCreator: 组织者/创建者区域, 示例值："ForwardIcon"
	UiStatus *string `json:"ui_status,omitempty"` // UI项自定义状态。目前只支持hide, 示例值："hide", 可选值有: `hide`：隐藏显示, `readonly`：只读, `editable`：可编辑, `unknown`：未知UI项自定义状态，仅用于读取时兼容
	AppLink  *string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接, 示例值："xxxxx", 最大长度：`2000` 字符
}

type createCalendarEventResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *CreateCalendarEventResp `json:"data,omitempty"`
}

type CreateCalendarEventResp struct {
	Event *CreateCalendarEventRespEvent `json:"event,omitempty"` // 新创建的日程实体
}

type CreateCalendarEventRespEvent struct {
	EventID          string                                  `json:"event_id,omitempty"`           // 日程ID
	Summary          string                                  `json:"summary,omitempty"`            // 日程标题
	Description      string                                  `json:"description,omitempty"`        // 日程描述
	NeedNotification bool                                    `json:"need_notification,omitempty"`  // 更新日程是否给日程参与人发送bot通知，默认为true
	StartTime        *CreateCalendarEventRespEventStartTime  `json:"start_time,omitempty"`         // 日程开始时间
	EndTime          *CreateCalendarEventRespEventEndTime    `json:"end_time,omitempty"`           // 日程结束时间
	Vchat            *CreateCalendarEventRespEventVchat      `json:"vchat,omitempty"`              // 视频会议信息。
	Visibility       string                                  `json:"visibility,omitempty"`         // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `default`：默认权限，跟随日历权限，默认仅向他人显示是否“忙碌”, `public`：公开，显示日程详情, `private`：私密，仅自己可见详情
	AttendeeAbility  string                                  `json:"attendee_ability,omitempty"`   // 参与人权限, 可选值有: `none`：无法编辑日程、无法邀请其它参与人、无法查看参与人列表, `can_see_others`：无法编辑日程、无法邀请其它参与人、可以查看参与人列表, `can_invite_others`：无法编辑日程、可以邀请其它参与人、可以查看参与人列表, `can_modify_event`：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus   string                                  `json:"free_busy_status,omitempty"`   // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `busy`：忙碌, `free`：空闲
	Location         *CreateCalendarEventRespEventLocation   `json:"location,omitempty"`           // 日程地点
	Color            int64                                   `json:"color,omitempty"`              // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders        []*CreateCalendarEventRespEventReminder `json:"reminders,omitempty"`          // 日程提醒列表
	Recurrence       string                                  `json:"recurrence,omitempty"`         // 重复日程的重复性规则；参考[rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10)；, 不支持COUNT和UNTIL同时出现；, 预定会议室重复日程长度不得超过两年。
	Status           string                                  `json:"status,omitempty"`             // 日程状态, 可选值有: `tentative`：未回应, `confirmed`：已确认, `cancelled`：日程已取消
	IsException      bool                                    `json:"is_exception,omitempty"`       // 日程是否是一个重复日程的例外日程
	RecurringEventID string                                  `json:"recurring_event_id,omitempty"` // 例外日程的原重复日程的event_id
	Schemas          []*CreateCalendarEventRespEventSchema   `json:"schemas,omitempty"`            // 日程自定义信息
}

type CreateCalendarEventRespEventStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type CreateCalendarEventRespEventEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type CreateCalendarEventRespEventVchat struct {
	VCType      string `json:"vc_type,omitempty"`     // 视频会议类型, 可选值有: `vc`：飞书视频会议，取该类型时，其他字段无效。, `third_party`：第三方链接视频会议，取该类型时，icon_type、description、meeting_url字段生效。, `no_meeting`：无视频会议，取该类型时，其他字段无效。, `lark_live`：Lark直播，内部类型，只读。, `unknown`：未知类型，做兼容使用，只读。
	IconType    string `json:"icon_type,omitempty"`   // 第三方视频会议icon类型；可以为空，为空展示默认icon。, 可选值有: `vc`：飞书视频会议icon, `live`：直播视频会议icon, `default`：默认icon
	Description string `json:"description,omitempty"` // 第三方视频会议文案，可以为空，为空展示默认文案
	MeetingURL  string `json:"meeting_url,omitempty"` // 视频会议URL
}

type CreateCalendarEventRespEventLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称
	Address   string  `json:"address,omitempty"`   // 地点地址
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
}

type CreateCalendarEventRespEventReminder struct {
	Minutes int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒,新建或更新日程时传入该字段，仅对当前身份生效
}

type CreateCalendarEventRespEventSchema struct {
	UiName   string `json:"ui_name,omitempty"`   // UI名称。取值范围如下： \,ForwardIcon: 日程转发按钮 \,MeetingChatIcon: 会议群聊按钮 \,MeetingMinutesIcon: 会议纪要按钮 \,MeetingVideo: 视频会议区域 \,RSVP: 接受/拒绝/待定区域 \,Attendee: 参与者区域 \,OrganizerOrCreator: 组织者/创建者区域
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态。目前只支持hide, 可选值有: `hide`：隐藏显示, `readonly`：只读, `editable`：可编辑, `unknown`：未知UI项自定义状态，仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接
}
