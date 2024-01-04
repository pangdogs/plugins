package gtp_gate

import (
	"github.com/elliotchance/pie/v2"
	"golang.org/x/net/context"
)

func (s *_Session) newDataWatcher(ctx context.Context, handler RecvDataHandler) *_DataWatcher {
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithCancel(ctx)

	watcher := &_DataWatcher{
		Context:     ctx,
		cancel:      cancel,
		stoppedChan: make(chan struct{}),
		session:     s,
		handler:     handler,
	}
	s.dataWatchers.Append(watcher)

	s.gate.wg.Add(1)
	go watcher.mainLoop()

	return watcher
}

type _DataWatcher struct {
	context.Context
	cancel      context.CancelFunc
	stoppedChan chan struct{}
	session     *_Session
	handler     RecvDataHandler
}

func (w *_DataWatcher) Stop() <-chan struct{} {
	w.cancel()
	return w.stoppedChan
}

func (w *_DataWatcher) mainLoop() {
	defer func() {
		w.cancel()
		w.session.gate.wg.Done()
		close(w.stoppedChan)
	}()

	select {
	case <-w.Done():
	case <-w.session.Done():
	}

	w.session.dataWatchers.AutoLock(func(watchers *[]*_DataWatcher) {
		*watchers = pie.DropWhile(*watchers, func(other *_DataWatcher) bool {
			return other == w
		})
	})
}