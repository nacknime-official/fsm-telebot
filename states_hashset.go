package fsm

type statesHashset map[State]struct{}

func newHashset() statesHashset {
	return make(statesHashset)
}

func newHashsetFromSlice(states []State) statesHashset {
	h := newHashset()
	for _, state := range states {
		h.Add(state)
	}
	return h
}

func (h statesHashset) Add(s State) {
	h[s] = struct{}{}
}

func (h statesHashset) Has(state State) bool {
	_, ok := h[state]
	return ok
}

func (h statesHashset) Delete(state State) {
	delete(h, state)
}

func (h statesHashset) Size() int {
	return len(h)
}
