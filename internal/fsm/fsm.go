package fsm

type State struct {
	cmd  string
	step int
}

type stateStore map[int64]*State

var StateStore = make(stateStore, 0)

func (self stateStore) GetOrCreateState(chatId int64) *State {
	if val, ok := self[chatId]; ok {
		return val
	} else {
		self[chatId] = &State{}
		return self[chatId]
	}
}

func (state *State) StartFlow(cmd string) {
	if state.cmd != "" {
		state.FinishFlow()
	}
	state.cmd = cmd
	state.step = 0
}

func (state *State) FinishFlow() {
	state.cmd = ""
	state.step = 0
}

func (state *State) Next() {
	state.step++
}

func (state *State) CurrentStep() int {
	return state.step
}

func (state *State) CurrentCmd() string {
	return state.cmd
}
