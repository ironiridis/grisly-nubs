package main

import "net/http"
import "io"

//import "html/template"

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

}

func HandleSlotPOSTIdx(s int, w http.ResponseWriter, r *http.Request) {
	var err error
	err = r.ParseMultipartForm(157286400) // about or exactly 150mB
	if err != nil {
		http.Error(w, "Unable to parse upload", http.StatusInternalServerError)
		return
	}

	rfp, _, err := r.FormFile("f")
	if err != nil {
		http.Error(w, "Unable to obtain uploaded file", http.StatusInternalServerError)
		return
	}
	defer rfp.Close()

	if v, err := ValidImage(rfp); v == false {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	fn := fmt.Sprintf("/slots/n%d", s)
	// TODO update configuration

	CmdRemountRW.Run()
	lfp, err := os.Create(fn)
	if err != nil {
		http.Error(w, "Unable to store file locally", http.StatusInternalServerError)
		return
	}
	defer lfp.Close()
	io.Copy(rfp, lfp)
	CmdRemountRO.Run()

	// TODO redirect back to index page
}
