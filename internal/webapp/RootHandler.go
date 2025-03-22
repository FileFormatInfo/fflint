package webapp

import (
	"net/http"
	"os"

	"github.com/FileFormatInfo/fflint/internal/common"
	"github.com/FileFormatInfo/fflint/ui"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	fromLogout := r.URL.Query().Get("logout")
	if fromLogout != "" {
		T := ui.GetT(r)
		sess, _ := common.GetSession(r)
		sess.AddFlash("success", T("index.logged_out"))
		sess.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ui.RunTemplate(w, r, "index.tmpl", map[string]any{
		"WaitlistEnabled": os.Getenv("WAITLIST_URL") != "",
	})

	//http.Redirect(w, r, "/domains/", http.StatusSeeOther)
}
