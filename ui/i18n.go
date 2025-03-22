package ui

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"

	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/FileFormatInfo/fflint/internal/common"
)

//go:embed messages
var messagesFS embed.FS

var bundle = initTranslations()

func initTranslations() *i18n.Bundle {

	// Initialize i18n with English (default) and French languages
	theBundle := i18n.NewBundle(language.English)           // Default language
	theBundle.RegisterUnmarshalFunc("json", json.Unmarshal) // Register JSON unmarshal function
	walkErr := fs.WalkDir(messagesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			common.Logger.Debug("Loading messages", "path", path)
			theBundle.LoadMessageFileFS(messagesFS, path)
		}
		return nil
	})
	if walkErr != nil {
		common.Logger.Error("unable to walk messages", "error", walkErr)
	}

	common.Logger.Info("Loaded messages", "languages", theBundle.LanguageTags())

	return theBundle
}

func GetLanguages() []string {

	tags := bundle.LanguageTags()
	languages := make([]string, len(tags))
	for i, tag := range tags {
		languages[i] = tag.String()
	}
	return languages
}

func GetLocalizer(r *http.Request) *i18n.Localizer {

	var lang string
	sess, sessErr := common.GetSession(r)
	if sessErr != nil {
		lang = ""
	} else {
		lang = sess.GetString("Language")
		common.Logger.Trace("Language from session", "lang", lang)
	}

	if lang == "" {
		cookie, cookieErr := r.Cookie("Language")
		if cookieErr == nil {
			lang = cookie.Value
			common.Logger.Trace("Language from cookie", "lang", lang)
		}
	}

	accept := r.Header.Get("Accept-Language")

	localizer := i18n.NewLocalizer(bundle, lang, accept) // Initialize localizer with detected language

	return localizer
}

func GetT(r *http.Request) func(string) string {

	localizer := GetLocalizer(r)

	return func(msg string) string {
		xlat, xlatErr := localizer.Localize(&i18n.LocalizeConfig{MessageID: msg})
		if xlatErr != nil {
			common.Logger.Error("Translation error", "error", xlatErr, "message", msg)
			return msg
		}
		return xlat
	}
}

func GetMaybeT(r *http.Request) func(string) string {

	localizer := GetLocalizer(r)

	return func(msg string) string {
		xlat, xlatErr := localizer.Localize(&i18n.LocalizeConfig{MessageID: msg})
		if xlatErr != nil {
			return ""
		}
		return xlat
	}
}
