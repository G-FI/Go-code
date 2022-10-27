package Observer

type Publisher struct {
	listeners []Listener
}

func (p *Publisher) AddListener(listener Listener) {
	p.listeners = append(p.listeners, listener)
}
func (p *Publisher) Notify() {
	for _, l := range p.listeners {
		l.Update()
	}
}
