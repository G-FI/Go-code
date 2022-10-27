package builder

type Manager struct {
}

func (m *Manager) Make(builder Builder) {
	builder.Reset()
	builder.SetSeats()
	builder.SetEngin()
}

func NewManager() *Manager {
	return &Manager{}
}
