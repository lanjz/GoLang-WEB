package odserver

type staticUrlMap []string

type StaticPathConfig struct {
	staticUrlMap
}

func newStaticPathConfig() *StaticPathConfig {
	return &StaticPathConfig{
		staticUrlMap: make([]string, 0),
	}
}

func (s *StaticPathConfig) setStaticPath(url string) {
	s.staticUrlMap = append(s.staticUrlMap, url)
}
