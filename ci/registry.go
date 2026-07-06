package ci

type Registory struct {
	providers map[string]Provider
}

func NewRegistory() *Registory {
	return &Registory{
		providers: make(map[string]Provider),
	}
}

func (r *Registory) Registor(name string, p Provider) {
	r.providers[name] = p
}

func (r *Registory) Get(name string) (Provider, bool) {
	p, ok := r.providers[name]
	return p, ok
}
