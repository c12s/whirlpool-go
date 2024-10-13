package whirlpool

import "fmt"

type WhirlPool struct {
	config []ProbeSettings
}

type ProbeSettings struct {
	nodeUrl string
	path    string
	probes  []Probe
}

type ProbeType int

const (
	Liveness ProbeType = iota
	Readiness
	Startup
)

type ProbeMethod interface {
	Perform() error
}

type HttpProbe struct {
	path string
	port int32
}

func (http *HttpProbe) Perform() error {
	return nil
}

type RPCProbe struct {
	port int32
}

func (rpc *RPCProbe) Perform() error {
	return nil
}

type Probe struct {
	probeType           ProbeType
	intervalSeconds     float32
	startupDelaySeconds float32
}

func New(settings []ProbeSettings) *WhirlPool {
	return &WhirlPool{
		config: settings,
	}
}

func (w *WhirlPool) StartProbing() error {
	go w.startActor()

	return nil
}

func (w *WhirlPool) startActor() error {
	fmt.Printf("Started probing.")

	return nil
}
