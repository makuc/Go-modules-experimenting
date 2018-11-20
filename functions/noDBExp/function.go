package noDBExp

import (
	"github.com/makuc/go-mod-exp/modules/mytest"
	"net/http"
)

func BrezBaze(w http.ResponseWriter, r *http.Request) {
	w.Write(mytest.Message())
}
