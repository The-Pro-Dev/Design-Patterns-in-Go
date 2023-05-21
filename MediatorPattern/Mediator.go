package MediatorPattern

type Mediator interface {
	perform()
}

type ChoreoGrapher struct {
	Artists []Artist
}

func (ch ChoreoGrapher) perform() {
	for _, a := range ch.Artists {
		a.perform()
	}
}

func (ch *ChoreoGrapher) add(a Artist)  {
	ch.Artists = append(ch.Artists, a)
}
