package Utils

import (
	"net/http"
	"strings"
)

func CloneHeader(src http.Header, dest *http.Header)  {
	for k,v := range src{
		dest.Set(k,strings.Join(v, ","))
	}
}
