package entity

type Backend struct {
	ID       string
	IPv4     string
	Hostname *string
	Port     int
	Protocol string
	Weight   *int
	State    bool
	PoolID   *string
}

func (b *Backend) RemoveFromPool() {
	b.PoolID = nil
	b.Inactivate()
}

func (b *Backend) Activate() {
	b.State = true
}

func (b *Backend) Inactivate() {
	b.State = false
}
