package create

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"url_shortener/app/helpers"
)

func shortURLController(ctx context.Context) {
}

// ShortURL handler responsible for creating short url
func ShortURL(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	filter := make(map[string]interface{})
	filter["urlOriginal"] = r.Form.Get("url")

	update := make(map[string]interface{})
	for _, val := range [4]string{"isWebToWeb", "isWebToMob", "isMobToApp", "isAppToApp"} {
		update[val], err = strconv.ParseBool(r.Form.Get(val))
		if err != nil {
			return err
		}
	}

	t, err := helpers.ParseDate(r.Form.Get("expireAt"))
	if err != nil {
		return err
	}
	update["expireAt"] = t

	update["expireAt"] = time.Now()

	// ctx := r.Context()
	// filter := make(map[string]interface{})

	return nil
}
