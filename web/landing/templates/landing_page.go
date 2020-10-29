package landing

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/sirupsen/logrus"
)

var templates = template.Must(template.ParseGlob("web/landing/templates/*"))

// PageHandler displays the landing page code
func PageHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("templates: %+v", templates)
	err := templates.ExecuteTemplate(w, "landing.html", nil)
	if err != nil {
		logrus.Infof("web/landing/PageHandler: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("landing page handler %s", err)))
		return
	}
}
