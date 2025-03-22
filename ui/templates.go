package ui

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"github.com/FileFormatInfo/fflint/internal/common"
)

//go:embed partials
var partialsFS embed.FS

//go:embed all:views
var viewsFS embed.FS

type TemplateFunc func(data any) (string, error)

type TemplateData map[string]any

var templateCache = initTemplates()

func initTemplates() map[string]TemplateFunc {

	funcMap := template.FuncMap{
		"dec": func(i int) int {
			return i - 1
		},
		"inc": func(i int) int {
			return i + 1
		},
		"loop": func(from, to int) []int {
			result := []int{}
			for i := from; i < to; i++ {
				result = append(result, i)
			}
			return result
		},
		"toJson": func(data any) string {
			jsonStr, jsonErr := json.MarshalIndent(data, "", "    ")
			if jsonErr != nil {
				return jsonErr.Error()
			}
			return string(jsonStr)
		},
		"toString": func(data any) any {
			if data == nil {
				return template.HTML("<i>(not set)</i>")
			}

			return data.(string)
		},
	}

	theCache := make(map[string]TemplateFunc)

	var partials bytes.Buffer

	partialErr := fs.WalkDir(partialsFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			// read the partials file and append to partials
			content, readErr := fs.ReadFile(partialsFS, path)
			if readErr != nil {
				common.Logger.Error("unable to read partials file", "err", readErr, "filename", path)
				return err
			}
			partials.Write(content)
		}
		return nil
	})
	if partialErr != nil {
		common.Logger.Error("unable to register partials", "err", partialErr)
	}

	viewErr := fs.WalkDir(viewsFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			common.Logger.Error("walkdir error", "err", err)
			return err
		}
		if !d.IsDir() {
			common.Logger.Trace("registering view", "filename", path)
			content, readErr := fs.ReadFile(viewsFS, path)
			if readErr != nil {
				common.Logger.Error("unable to read view file", "err", readErr, "filename", path)
				return err
			}
			name := path[len("views/"):]

			var templateBuffer bytes.Buffer
			templateBuffer.Write(content)
			templateBuffer.Write(partials.Bytes())
			t := template.New(name).Funcs(funcMap)
			template, parseErr := t.Parse(templateBuffer.String())
			if parseErr != nil {
				common.Logger.Error("unable to parse template", "err", parseErr, "filename", path, "content", string(content))
				return parseErr
			}
			theCache[name] = func(data any) (string, error) {
				var buf bytes.Buffer
				err := template.Execute(&buf, data)
				if err != nil {
					common.Logger.Error("unable to execute template", "err", err, "filename", path, "content", string(content))
					return "", err
				}
				return buf.String(), nil
			}
		}
		return nil
	})
	if viewErr != nil {
		common.Logger.Error("unable to register views", "err", viewErr)
	}

	return theCache
}

type CrumbtrailEntry struct {
	Text string
	URL  string
}

func makeCrumbtrail(r *http.Request, data TemplateData) []CrumbtrailEntry {
	crumbtrail := []CrumbtrailEntry{}

	// Add additional entries based on the request path
	path := r.URL.Path
	segments := strings.Split(path, "/")
	for i := 1; i < len(segments); i++ {
		entry := CrumbtrailEntry{
			Text: segments[i],
			URL:  strings.Join(segments[0:i+1], "/"),
		}
		crumbtrail = append(crumbtrail, entry)
	}

	if data["Title"] != nil {
		crumbtrail[len(crumbtrail)-1].Text = data["Title"].(string)
	}

	return crumbtrail
}

func RunTemplate(w http.ResponseWriter, r *http.Request, templateName string, data TemplateData) {

	fn := templateCache[templateName]
	if fn == nil {
		common.Logger.Error("template not found", "template", templateName)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if data == nil {
		data = make(map[string]any)
	}
	T := GetT(r)
	MaybeT := GetMaybeT(r)
	key := strings.ReplaceAll(templateName[:len(templateName)-5], "/", ".")
	data["Title"] = T(fmt.Sprintf("%s.title", key))
	_, hasH1 := data["H1"]
	if !hasH1 {
		h1 := MaybeT(fmt.Sprintf("%s.h1", key))
		if h1 != "" {
			data["H1"] = h1
		}
	}
	data["T"] = T
	data["TH"] = func(msg string) template.HTML { return template.HTML(T(msg)) }
	data["crumbtrail"] = makeCrumbtrail(r, data)

	sess, sessErr := common.GetSession(r)
	if sessErr != nil {
		//LATER: flash
		common.Logger.Error("unable to get session", "error", sessErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data["Host"] = r.Host
	data["Path"] = r.URL.Path

	data["flashes"] = sess.ConsumeFlashes()
	sess.Save(r, w)

	rawStatus := data["http.status"]
	if rawStatus != nil {
		if status, ok := rawStatus.(int); !ok {
			common.Logger.Error("unable to convert http.status to int", "status", rawStatus)
		} else {
			w.WriteHeader(status)
		}
	}

	result, execErr := fn(data)
	if execErr != nil {
		common.Logger.Error("template failed", "err", execErr, "template", templateName)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(result))
}
