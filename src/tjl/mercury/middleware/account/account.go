package account

import (
	"errors"
	"example.com/m/src/tjl/mercury/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func processRequest(ctx *gin.Context) {
	var userSession session.Session
	var err error
	defer func() {
		if userSession == nil {
			userSession, _ = session.CreateSession()
		}
		ctx.Set(MercurySessionName, userSession)
	}()
	cookie, err := ctx.Request.Cookie(CookieSessionId)
	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}

	sessionId := cookie.Value
	if len(sessionId) == 0 {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	userSession, err = session.Get(sessionId)
	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))

		return
	}
	tmpUserId, err := userSession.Get(MercuryUserId)
	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	userId, ok := tmpUserId.(int64)
	if !ok || userId == 0 {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	ctx.Set(MercuryUserId, int64(userId))
	ctx.Set(MercuryUserLoginStatus, int64(1))
	return
}
func GetUserId(ctx *gin.Context) (userId int64, err error) {
	tempUserId, exists := ctx.Get(MercuryUserId)
	if !exists {
		err = errors.New("user id not exists")
		return
	}
	userId, ok := tempUserId.(int64)
	if !ok {
		err = errors.New("user id convert to int64 failed")
		return
	}
	return
}
func IsLogin(ctx *gin.Context) (login bool) {
	tempLoginStatus, exists := ctx.Get(MercuryUserLoginStatus)
	if !exists {
		return
	}
	loginstatus, ok := tempLoginStatus.(int64)
	if !ok {
		return
	}
	if loginstatus == 0 {
		return
	}
	login = true
	return
}
func processResponse(ctx *gin.Context) {
	var userSession session.Session
	tempSession, exits := ctx.Get(MercurySessionName)
	if !exits {
		return
	}
	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}
	if userSession.IsModify() == false {
		return
	}
	err := userSession.Save()
	if err != nil {
		return
	}
	sessionId := userSession.Id()
	cookie := &http.Cookie{
		Name:     CookieSessionId,
		Value:    sessionId,
		Path:     "/",
		MaxAge:   CookieMaxAge,
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(ctx.Writer, cookie)
	return
}
