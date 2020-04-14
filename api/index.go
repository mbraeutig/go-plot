package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
<html>
	<a href="https://go-plot.now.sh/api/plot?expr=sin(-x)*pow(1.5,-r)">https://go-plot.now.sh/api/plot?expr=sin(-x)*pow(1.5,-r)</a>
	<p>A relative URL: <a href="api/plot?expr=sin(-x)*pow(1.5,-r)">The a tag</a></p>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
	w.WriteHeader(http.StatusOK)
}
