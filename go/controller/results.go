// results.go implements handler and helper functions
// for the results and resultslist tabs
package controller

import (
	"github.com/geobe/gostip/go/model"
	"github.com/geobe/gostip/go/view"
	"html/template"
	"net/http"
	"time"
)

func ShowResults(w http.ResponseWriter, r *http.Request) {
	if checkMethodAllowed(http.MethodPost, w, r) != nil {
		return
	}
	app, err := fetchApplicant(w, r, "appid")
	//fmt.Printf("got applicant %s\n", app.Data.FirstName)
	if err != nil {
		return
	}
	values := viewmodel{
		"disabled": template.HTMLAttr("disabled='true'"),
		"oblasts":  model.Oblasts(),
	}
	setViewModel(app, values)
	addResultsConfig(time.Now().Year(), app, values)
	view.Views().ExecuteTemplate(w, "work_results", values)

}

func SubmitResults(w http.ResponseWriter, r *http.Request) {
	if err := checkMethodAllowed(http.MethodPost, w, r); err == nil {
		saveApplicantSubmission(w, r)
	}

}

// add a slice of possible test results for the given year to the viewmodel
func addResultsConfig(y int, app model.Applicant, data viewmodel) {
	var exref model.ExamReference
	model.Db().First(&exref, "year = ?", y)
	if exref.ID == 0 {
		return
	}
	var nq int
	for i, v := range exref.Results {
		if v == 0 || i == model.NQESTION-1 {
			nq = i
			break
		}
	}

	var results [model.NQESTION]map[string]float32
	for i := 0; i <= nq; i++ {
		results[i] = map[string]float32{
			"val": float32(app.Data.Results[i])/10.,
			"max": float32(exref.Results[i])/10.,
			"no":  float32(i + 1),
		}
	}
	data["results"] = results[:nq]
	data["languageresult"] = app.Data.LanguageResult
	data["language"] = app.Data.Language
	data["languages"] = model.InitialLanguages
}
