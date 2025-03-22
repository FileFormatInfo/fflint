package common

import (
	"encoding/gob"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var SessionName = "_jsessionid"

type Session struct {
	base      *sessions.Session
	needsSave bool
}

type FlashMessage struct {
	Msg  string
	Type string
}

func init() {
	gob.Register(&FlashMessage{})
}

func GetSession(r *http.Request) (*Session, error) {
	baseSession, err := SessionStore.Get(r, SessionName)
	if err != nil {
		return nil, err
	}
	return &Session{base: baseSession}, nil
}

func (s *Session) Get(key string) any {
	return s.base.Values[key]
}

func (s *Session) GetString(key string) string {
	value := s.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func (s *Session) Reset(r *http.Request, w http.ResponseWriter) error {
	s.base.Options.MaxAge = -1
	return s.base.Save(r, w)
}

func (s *Session) Save(r *http.Request, w http.ResponseWriter) error {
	if s.needsSave {
		saveErr := s.base.Save(r, w)
		if saveErr != nil {
			Logger.Error("unable to save session", "error", saveErr)
			return saveErr
		}
		s.needsSave = false
	}
	return nil
}

func (s *Session) Set(key string, value any) {
	s.needsSave = true
	s.base.Values[key] = value
}

func (s *Session) NeedsSave() bool {
	return s.needsSave
}

func (s *Session) AddFlash(flashType string, msg string) {
	s.needsSave = true
	flashMsg := FlashMessage{Msg: msg, Type: flashType}
	flashStr, jsonErr := json.Marshal(flashMsg)
	if jsonErr != nil {
		// should never occur!
		Logger.Error("unable to marshal flash message", "error", jsonErr)
		return
	}
	s.base.AddFlash(flashStr)
}

func (s *Session) ConsumeFlashes() []FlashMessage {
	rawFlashes := s.base.Flashes()
	if len(rawFlashes) == 0 {
		return []FlashMessage{}
	} else {
		s.needsSave = true
	}
	Logger.Info("consuming flashes", "flashes", rawFlashes)
	flashes := make([]FlashMessage, len(rawFlashes))
	for i, rawFlash := range rawFlashes {
		jsonErr := json.Unmarshal(rawFlash.([]byte), &flashes[i])
		if jsonErr != nil {
			Logger.Error("unable to unmarshal flash message", "error", jsonErr)
		}
	}
	return flashes
}
