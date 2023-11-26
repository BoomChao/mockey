package ani

type Animal interface {
	Speak(string) string
}

type Zoo struct {
	Ani Animal
}

func (zoo *Zoo) AniSpeak(str string) string {
	return zoo.Ani.Speak(str)
}

func (zoo *Zoo) aniWalk(str string) string {
	return str
}
