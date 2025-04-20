package dto

type AddBackendInput struct {
	IPv4     string  `json:"ipv4"`
	Hostname string `json:"hostname"`
	Port     int     `json:"port"`
	Protocol string  `json:"protocol"`
	Weight   int    `json:"weight"`
	PoolID   string  `json:"poolId"`
}

type RemoveBackendFromPoolInput struct {
	ID     string `json:"id"`
	PoolID string `json:"poolId"`
}

type UpdateBackendInput struct {
	ID       string  `json:"id"`
	IPv4     string  `json:"ipv4"`
	Hostname string `json:"hostname"`
	Port     int     `json:"port"`
	Protocol string  `json:"protocol"`
	Weight   int    `json:"weight"`
	PoolID   string  `json:"poolId"`
}

type ActivateBackendInput struct {
	ID string `json:"id"`
}

type InactivateBackendInput struct {
	ID string `json:"id"`
}

type DeleteBackendInput struct {
	ID string `json:"id"`
}

type GetBackendByIDInput struct {
	ID string `json:"id"`
}

type GetBackendsByPoolIDInput struct {
	PoolID string `json:"poolId"`
}
