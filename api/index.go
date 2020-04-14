package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
<html>
	<h1>sin(-x)*pow(1.5,-r): <a href="plot?expr=sin(-x)*pow(1.5,-r)">Try it!</a></h1>
	<h1>sin(x*y/10)/10: <a href="plot?expr=sin(x*y/10)/10">Try it!</a></h1>
	<h1>pow(2,sin(y))*pow(2,sin(x))/12: <a href="plot?expr=pow(2,sin(y))*pow(2,sin(x))/12">Try it!</a></h1>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
	w.WriteHeader(http.StatusOK)
}
