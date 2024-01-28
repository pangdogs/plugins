package gtp_gate

import (
	"git.golaxy.org/core/service"
)

// GetSession 查询会话
func GetSession(servCtx service.Context, sessionId string) (ISession, bool) {
	return Using(servCtx).GetSession(sessionId)
}

// RangeSessions 遍历所有会话
func RangeSessions(servCtx service.Context, fun func(session ISession) bool) {
	Using(servCtx).RangeSessions(fun)
}

// CountSessions 统计所有会话数量
func CountSessions(servCtx service.Context) int {
	return Using(servCtx).CountSessions()
}