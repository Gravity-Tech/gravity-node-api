package model


// swagger:parameters getAllNebulas getAllNodes getNodeRewards
type QueryParam struct {
	// a QueryParam acts like a search string for fields.
	// Use as a regexp expression.
	//
	// in: query
	Query string `json:"q"`
}

// swagger:parameters getAllNebulas getAllNodes getNodeRewards
type CurrentPageParam struct {
	// a CurrentPageParam represents current page
	// Only positive integers allowed. Default is 1
	//
	// in: query
	Page uint64 `json:"page"`
}

// swagger:parameters getAllNebulas getAllNodes getNodeRewards
type ItemsPerPageParam struct {
	// an ItemsPerPageParam represents items count per page
	// Only positive integers allowed. Default is 12
	//
	// in: query
	Items uint64 `json:"items"`
}