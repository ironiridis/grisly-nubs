package main

import "net/http"

//import "html/template"

func EPHandleUpload(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err = r.ParseMultipartForm(157286400) // about or exactly 150mB
	if err != nil {
		http.Error(w, "Unable to parse upload", http.StatusInternalServerError)
		return
	}

	fv, _, err := r.FormFile("f")
	if err != nil {
		http.Error(w, "Unable to obtain uploaded file", http.StatusInternalServerError)
		return
	}

	if v, err := ValidImage(fv); v == false {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

}
