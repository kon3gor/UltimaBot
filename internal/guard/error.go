package guard

type GuardErr struct {
	Msg string
}

func NewGuardErr(msg string) *GuardErr {
	return &GuardErr{msg}
}

func (self *GuardErr) Error() string {
	return self.Msg
}
