package sqlstore

import (
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetFullUrlByUid)
	bus.AddHandler("sql", CreateShortUrl)
}

func GetFullUrlByUid(query *models.GetFullUrlQuery) error {
	var shortUrl models.ShortUrl
	exists, err := x.Where("uid=?", query.Uid).Get(&shortUrl)
	if err != nil {
		return err
	}

	if !exists {
		return models.ErrShortUrlNotFound
	}

	query.Result = &shortUrl
	return nil
}

func CreateShortUrl(command *models.CreateShortUrlCommand) error {
	return inTransaction(func(sess *DBSession) error {

		shortUrl := models.ShortUrl{
			Uid:       command.Uid,
			Path:      command.Path,
			CreatedBy: command.CreatedBy,
			CreatedAt: command.CreatedAt,
		}

		if _, err := sess.Table("shortUrl").Insert(&shortUrl); err != nil {
			return err
		}

		return nil
	})
}
