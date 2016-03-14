package main

import "net/http"
import "io"
import "fmt"
import "os"
import "html/template"

const indexTemplateRaw = `
<!DOCTYPE html><html><body>
{{range $i, $e := .Slots}}
<div class="slot{{- $i -}}">
<form action="/slot{{- $i -}}" method="post" enctype="multipart/form-data">
Slot {{ $i }}:
<a href="/slot{{- $i -}}">Activate</a>
<input type="text" size="40" placeholder="Label" value="{{- .Label -}}" name="label" />
<input type="file" name="f" />
<input type="submit" value="Update">
</form>
</div>
{{end}}
</body></html>`

var indexTemplate = template.Must(template.New("index").Parse(indexTemplateRaw))

func HandleSlot(w http.ResponseWriter, r *http.Request) {
	var s int
	switch r.URL.Path {
	case "/slot0":
		s = 0
	case "/slot1":
		s = 1
	case "/slot2":
		s = 2
	case "/slot3":
		s = 3
	case "/slot4":
		s = 4
	case "/slot5":
		s = 5
	case "/slot6":
		s = 6
	case "/slot7":
		s = 7
	case "/slot8":
		s = 8
	case "/slot9":
		s = 9
	default:
		http.Error(w, "Unable to determine requested slot", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		HandleSlotGETIdx(s, w, r)
	case "POST":
		HandleSlotPOSTIdx(s, w, r)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleSlotGETIdx(s int, w http.ResponseWriter, r *http.Request) {
	err := RenderSlotToFramebuffer(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleSlotPOSTIdx(s int, w http.ResponseWriter, r *http.Request) {
	var err error
	err = r.ParseMultipartForm(157286400) // about or exactly 150mB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rfp, _, err := r.FormFile("f")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rfp.Close()

	if v, err := ValidImage(rfp); v == false {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	fn := fmt.Sprintf("/slots/n%d", s)
	FSRemountRW()
	lfp, err := os.Create(fn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rfp.Seek(0, 0) // ValidImage above will read a ways into the file
	io.Copy(lfp, rfp)
	lfp.Close()
	FSRemountRO()

	conf.WriteStart()
	conf.Slots[s].Filename = fn
	conf.Slots[s].Label = r.FormValue("label")
	l := conf.LastRecalled
	conf.WriteDone()

	LoadSlot(s, fn) // Pull it into the cache immediately
	
	// if the last recalled slot is the one we just uploaded, re-render it
	if s == l {
		RenderSlotToFramebuffer(s)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	conf.ReadStart()
	err := indexTemplate.Execute(w, conf)
	conf.ReadDone()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/slot0", HandleSlot)
	http.HandleFunc("/slot1", HandleSlot)
	http.HandleFunc("/slot2", HandleSlot)
	http.HandleFunc("/slot3", HandleSlot)
	http.HandleFunc("/slot4", HandleSlot)
	http.HandleFunc("/slot5", HandleSlot)
	http.HandleFunc("/slot6", HandleSlot)
	http.HandleFunc("/slot7", HandleSlot)
	http.HandleFunc("/slot8", HandleSlot)
	http.HandleFunc("/slot9", HandleSlot)

	http.HandleFunc("/", HandleIndexPage)
}

func StartHTTP() {
	http.ListenAndServe(":8080", nil)
}
