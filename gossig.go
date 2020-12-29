package gossig

import (
	"os"
	"os/signal"
)

var SignalProcessor *SignalsT

type SignalsT struct {
	handlers         map[os.Signal]interface{}
	procees          bool
	signalReceiverCh chan os.Signal
	cSignal          os.Signal
}

func (sls *SignalsT) HandlerAdd(signal os.Signal, handler interface{}) {
	handlerFunc, ok := handler.(func())
	if !ok {
		return
	}
	if sls.handlers == nil {
		sls.handlers = make(map[os.Signal]interface{})
	}
	sls.handlers[signal] = handlerFunc
}

func (sls *SignalsT) HandlerRemove(signal os.Signal) {
	delete(sls.handlers, signal)
}

func (sls *SignalsT) Stop() {
	sls.procees = false
	signal.Stop(sls.signalReceiverCh)
}

func (sls *SignalsT) processSignal() {
	if sls.handlers[sls.cSignal] != nil {
		sls.handlers[sls.cSignal].(func())()
	}
}

func (sls *SignalsT) Run() {
	if sls.procees == true {
		return
	}
	sls.procees = true
	sls.signalReceiverCh = make(chan os.Signal, 1)
	signal.Notify(sls.signalReceiverCh)

	go func() {
		for {
			sls.cSignal = <-sls.signalReceiverCh
			if sls.procees {
				sls.processSignal()
			}
		}
	}()
}

func init() {
	SignalProcessor = new(SignalsT)
}
