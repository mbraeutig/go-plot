package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
<html>
	<h1>sin(-x)*pow(1.5,-r): <a href="plot?expr=sin(-x)*pow(1.5,-r)">Try it!</a></h1>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
	w.WriteHeader(http.StatusOK)
}
