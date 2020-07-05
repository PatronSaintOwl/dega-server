package claimant

import "github.com/go-chi/chi"

// claimant model
type claimant struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	TagLine     string `json:"tag_line"`
	MediumID    uint   `json:"medium_id"`
}

// Router - Group of claimant router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Post("/", create)

	r.Route("/{claimant_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r

}
