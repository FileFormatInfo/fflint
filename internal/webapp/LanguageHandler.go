package webapp

import (
	"cmp"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/FileFormatInfo/fflint/internal/common"
	"github.com/FileFormatInfo/fflint/ui"
)

type languageItem struct {
	Code      string
	Name      string
	IsCurrent bool
}

func LanguageGetHandler(w http.ResponseWriter, r *http.Request) {

	next := r.URL.Query().Get("next")

	T := ui.GetT(r)
	currentLanguage := T("language.current")
	codes := ui.GetLanguages()
	var languages []languageItem = make([]languageItem, len(codes))
	for i, code := range codes {
		languages[i] = languageItem{
			Code:      code,
			Name:      T(fmt.Sprintf("language.%s", code)),
			IsCurrent: code == currentLanguage,
		}
	}
	slices.SortFunc(languages, func(a, b languageItem) int {
		return cmp.Compare(a.Name, b.Name)
	})

	ui.RunTemplate(w, r, "language.tmpl", map[string]any{
		"Languages": languages,
		"Next":      next,
	})

}

func LanguagePostHandler(w http.ResponseWriter, r *http.Request) {

	sess, _ := common.GetSession(r)
	lang := r.FormValue("lang")
	next := safeNext(r.FormValue("next"), "/")

	sess.Set("Language", lang)
	sess.Save(r, w)

	//LATER: also save to user profile if logged in

	T := ui.GetT(r)
	sess.AddFlash("success", T("language.changed"))
	sess.Save(r, w)

	http.Redirect(w, r, next, http.StatusSeeOther)
}

func safeNext(next string, defaultValue string) string {

	if next == "" {
		return defaultValue
	}

	parsedNext, parseErr := url.Parse(next)
	if parseErr != nil {
		common.Logger.Error("Unable to parse next URL", "error", parseErr, "next", next)
		return defaultValue
	}

	if parsedNext.Host != "" {
		common.Logger.Error("Invalid next URL", "next", next, "host", parsedNext.Host)
		return defaultValue
	}

	return next
}
