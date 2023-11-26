package ani

type cat struct{}

func NewCat() Animal {
	return cat{}
}

func (c cat) Speak(str string) string {
	return str
}
