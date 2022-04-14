// Package delreply 撤回消息
package delreply

import (
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

// 插件主体
func init() {
	engine := control.Register("delreply", &control.Options{
		DisableOnDefault: false,
		Help:             "",
	})
	engine.OnRegex(`^\[CQ:reply,id=(.*)](.*)`, zero.AdminPermission, zero.OnlyToMe, zero.KeywordRule("撤回")).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			// 获取消息id
			mid := ctx.State["regex_matched"].([]string)[1]
			// 撤回消息
			ctx.DeleteMessage(message.NewMessageIDFromString(mid))
		})
}
