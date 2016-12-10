// Package controller file enrol.go implements handler and helper functions for enrolment.
// I.e. for tabs enrol and edit.
package controller

import (
	"github.com/geobe/gostip/go/model"
	"github.com/geobe/gostip/go/view"
	"html"
	"html/template"
	"net/http"
)

// ShowEnrol is handler to show the selected applicant from the
// search select element for enrol and edit tabs. It returns
// an html page fragment that is inserted into the respective tab area.
func ShowCancel(w http.ResponseWriter, r *http.Request) {
	if checkMethodAllowed(http.MethodPost, w, r) != nil {
		return
	}
	app, err := fetchApplicant(w, r, "appid")
	if err != nil {
		return
	}
	action := html.EscapeString(r.PostFormValue("action"))
	enrol := action == "enrol"

	//data := app.Data
	values := viewmodel{
		//"enrolaction":  "/enrol/submit",
		//"displaystyle": "none",
		"oblasts": model.Oblasts(),
	}
	setViewModel(app, values)
	if enrol {
		values["disabled"] = template.HTMLAttr("disabled='true'")
		view.Views().ExecuteTemplate(w, "work_enrol", values)
	} else {
		view.Views().ExecuteTemplate(w, "work_edit", values)
	}
}

// SubmitCancel is handler that accepts form submissions from the cancellation tab.
// Only http POST method is accepted.
func SubmitCancel(w http.ResponseWriter, r *http.Request) {
	if err := checkMethodAllowed(http.MethodPost, w, r); err == nil {
		saveApplicantSubmission(w, r, true)
	}
}
