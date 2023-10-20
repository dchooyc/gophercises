package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLs []pathURL
	err := yaml.Unmarshal(yml, &pathURLs)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)

	for _, pu := range pathURLs {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}
