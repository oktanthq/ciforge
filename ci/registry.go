package ci

type Registory struct {
	providers map[string]Provider
}

func NewRegistory() *Registory

func (r *Registory) Registor(name string, p Provider)
func (r *Registory) Get(name string) (Provider, bool)
