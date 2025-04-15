package domain

type Backend struct {
	ID           string
	IPv4         string
	Hostname     *string
	Port         int
	Protocol     string
	Weight       int
	InitialState bool
	PoolID       string
}
