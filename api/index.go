package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
<html>
    <h1>https://go-plot.now.sh/api/plot?expr=sin(-x)*pow(1.5,-r)</h1>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
}
