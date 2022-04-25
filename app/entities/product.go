package entities

type Product struct {
	ID   *int64  `json:"id"`
	Type *string `json:"type"`
	Name *string `json:"name"`
}
