package whirlpool

import (
	"fmt"
	"net/http"
)

type WhirlPool struct {
	configs []ProbeMethod
}

func New(configs []ProbeMethod) *WhirlPool {
	return &WhirlPool{
		configs: configs,
	}
}

func (w *WhirlPool) StartProbing() (*http.Response, error) {
	config := w.configs[0]

	return config.Perform()
}

type ProbeType int

const (
	Liveness ProbeType = iota
	Readiness
	Startup
)

type ProbeMethod interface {
	Perform() (*http.Response, error)
}

type RestProbe struct {
	probes []Probe
	url    string
	path   string
	port   int32
}

func NewRestProbe(url, path string, port int32, probes []Probe) RestProbe {
	return RestProbe{
		probes: probes,
		url:    url,
		path:   path,
		port:   port,
	}
}

func (probe *RestProbe) formatedUrl() string {
	return fmt.Sprintf("%s:%d/%s", probe.url, probe.port, probe.path)
}

func (probe *RestProbe) Perform() (*http.Response, error) {
	return http.Get(probe.formatedUrl())
}

type Probe struct {
	probeType           ProbeType
	intervalSeconds     float32
	startupDelaySeconds float32
}

func NewLivenessProbe() Probe {
	return Probe{
		probeType:           Liveness,
		intervalSeconds:     1,
		startupDelaySeconds: 1,
	}
}
