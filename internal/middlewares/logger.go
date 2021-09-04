package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"github.com/sirupsen/logrus"
)

func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		logrus.Info(fmt.Sprintf("%q", x))

		rec := httptest.NewRecorder()
		fn(rec, r)

		logrus.Info(fmt.Sprintf("%q", rec.Body))

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		rec.Body.WriteTo(w)
	}
}
