package redirect

import (
	"context"
	"net/http"
	"url_shortener/app/services"
)

func redirectController(ctx context.Context, key string) (interface{}, error) {

	// return cached value
	val, err := services.GetValue(ctx, key)
	if err != nil {
		return nil, err
	}

	return val, nil
	// return db value
	// filter := make(map[string]interface{})
	// val, err := services.FindData(ctx, models.AdAnalyticsCollectionName, filter)
}

// Redirect handler responsible for url redirection
func Redirect(w http.ResponseWriter, r *http.Request) error {
	// need to set default redirection.
	// var rURL string
	// var err error
	// ctx := r.Context()

	// http.Redirect(w, r, rURL, http.StatusSeeOther)
	return nil
}
