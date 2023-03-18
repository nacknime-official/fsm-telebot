package fsm

import (
	"github.com/pkg/errors"
	tele "gopkg.in/telebot.v3"
)

// Filter object. Needs for graceful works with state filters.
type Filter struct {
	Endpoint interface{}
	States   []State
}

// F returns new Filter object.
func F(endpoint interface{}, states ...State) Filter {
	if len(states) == 0 {
		states = []State{DefaultState}
	}
	return Filter{Endpoint: endpoint, States: states}
}

// ForState creates handler with local filter for given state.
func (m *Manager) ForState(want State, handler Handler) tele.HandlerFunc {
	return m.ForStates(handler, want)
}

// ForStates  creates a handler with local filter
// for current state to check for presence in given states.
func (m *Manager) ForStates(h Handler, states ...State) tele.HandlerFunc {
	return m.HandlerAdapter(func(c tele.Context, state Context) error {
		s, err := state.State()
		if err != nil {
			return errors.Wrap(err, "fsm-telebot: get state in Manager.ForStates")
		}
		if ContainsState(s, states...) {
			return h(c, state)
		}
		return nil
	})
}

func (f Filter) CallbackUnique() string {
	switch end := f.Endpoint.(type) {
	case string:
		return end
	case tele.CallbackEndpoint:
		return end.CallbackUnique()
	default:
		panic("fsm: telebot: unsupported endpoint")
	}
}
