package session

import (
	"mainpkg/util"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type T struct {
	Data map[string]*SingleSession

	quit           chan struct{}
	cookieDuration time.Duration
}

type SingleSession struct {
	ExpirationDate time.Time
	UserId         primitive.ObjectID
	Data           string
}

func newSingleSession(expirationDate time.Time) *SingleSession {
	return &SingleSession{ExpirationDate: expirationDate}
}

func (sessions T) clearSessions() {
	for key, val := range sessions.Data {
		if val.ExpirationDate.Before(time.Now()) {
			delete(sessions.Data, key)
		}
	}
}

type HandlerFunc func(http.ResponseWriter, *http.Request)
type SessionHandlerFunc func(http.ResponseWriter, *http.Request, SingleSession)

func (sessions T) DeleteSession(writer *http.ResponseWriter) {
	http.SetCookie(*writer, &http.Cookie{
		Name:    "sessionId",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
}

func (sessions T) NewSession(writer *http.ResponseWriter) *SingleSession {
	sessionId := util.RandomString(10)
	newSession := newSingleSession(time.Now().Add(sessions.cookieDuration * time.Second))
	sessions.Data[sessionId] = newSession

	http.SetCookie(*writer, &http.Cookie{
		Name:    "sessionId",
		Value:   sessionId,
		Expires: newSession.ExpirationDate,
		Path:    "/",
	})

	return newSession
}

func (sessions T) GetSession(r *http.Request) *SingleSession {
	sessionCookie, _ := r.Cookie("sessionId")
	if sessionCookie == nil {
		return nil
	}

	session, ok := sessions.Data[sessionCookie.Value]
	if !ok {
		return nil
	}

	return session
}

func (sessions T) LoggedIn(handler HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := sessions.GetSession(r)
		if session == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		handler(w, r)
	}
}

func Init(cookieDuration int) *T {
	var result T
	result.Data = make(map[string]*SingleSession)
	result.quit = util.RunEvery(result.clearSessions, 10*time.Second)
	result.cookieDuration = time.Duration(cookieDuration)
	return &result
}

func (sessions T) Close() {
	close(sessions.quit)
}
