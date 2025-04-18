package dto

type AddRoutingRulesInput struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
	Protocol string `json:"protocol"`
}

type UpdateRoutingRulesInput struct {
	ID       string `json:"id"`
	Source   string `json:"source"`
	Target   string `json:"target"`
	Protocol string `json:"protocol"`
}

type DeleteRoutingRulesInput struct {
	ID string `json:"id"`
}

type GetRoutingRulesByIDInput struct {
	ID string `json:"id"`
}
