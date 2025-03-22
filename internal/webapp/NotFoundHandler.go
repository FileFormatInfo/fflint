package webapp

import (
	"html/template"
	"net/http"

	"github.com/FileFormatInfo/fflint/ui"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

	localizer := ui.GetLocalizer(r)

	ui.RunTemplate(w, r, "_404.tmpl", map[string]any{
		"http.status": http.StatusNotFound,
		"Message": template.HTML(localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "_404.message_html",
			TemplateData: map[string]string{
				"Path": r.URL.Path,
			},
		})),
	})

}
