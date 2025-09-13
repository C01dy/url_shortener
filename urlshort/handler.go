package urlshort

import (
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if toRedirect, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, toRedirect, http.StatusFound)
			return 
		}
		fallback.ServeHTTP(w, r) 
		
	})
}

func DataHandler(data []byte, parser Parser, fallback http.Handler) (http.HandlerFunc, error) {
    pathUrls, err := parser.Parse(data)
    if err != nil {
        return nil, err
    }

    pathsToUrls := BuildPath(pathUrls) 
    return MapHandler(pathsToUrls, fallback), nil
}

