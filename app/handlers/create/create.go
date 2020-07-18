package create

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"url_shortener/app/helpers"
	"url_shortener/app/models"
	"url_shortener/app/services"
)

func shortURLController(doc models.URLData) error {

	ctx := context.

	// updating data into mongo db.
	_, err := services.InsertData(ctx, models.URLDataCollectionName, doc)
	if err != nil {
		return err
	}

	// updating data in redis
	// key := strconv.FormatUint(uint64(doc.URLHash), 10)
	// exp := time.Duration(doc.ExpireAt)
	// if err := services.SetValue(ctx, key, doc, exp.Second()); err != nil {
	// 	return err
	// }
	// return nil
}

// ShortURL handler responsible for creating short url
func ShortURL(w http.ResponseWriter, r *http.Request) error {
	var doc models.URLData
	err := r.ParseForm()
	if err != nil {
		return err
	}
	doc.URLOriginal = r.Form.Get("url")
	doc.URLHash = helpers.GetHash()
	doc.CreatedAt = time.Now()
	t, err := helpers.ParseDate(r.Form.Get("expireAt"))
	if err != nil {
		return err
	}
	doc.ExpireAt = t
	doc.IsWebToMob, err = strconv.ParseBool(r.Form.Get("isWebToWeb"))
	if err != nil {
		return err
	}
	doc.IsWebToMob, err = strconv.ParseBool(r.Form.Get("isWebToMob"))
	if err != nil {
		return nil
	}
	doc.IsMobToApp, err = strconv.ParseBool(r.Form.Get("isMobToApp"))
	if err != nil {
		return nil
	}
	doc.IsAppToApp, err = strconv.ParseBool(r.Form.Get("isAppToApp"))
	if err != nil {
		return err
	}

	// ctx := r.Context()

	if err = shortURLController(doc); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doc)

	return nil
}
