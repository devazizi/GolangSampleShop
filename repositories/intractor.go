package repositories

type Intractor struct {
}

func NewSqlQuery(db any) *Intractor {
	return &Intractor{}
}
