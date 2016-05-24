package account

import (
	"database/sql"
	"net/http"
	"time"
)

type Cookie struct {
	ID        int `sql:"id"`
	AccountID int `sql:"account_id"`

	Name  string `sql:"name"`
	Value string `sql:"value"`

	Path       string    `sql:"path"`        // optional
	Domain     string    `sql:"domain"`      // optional
	Expires    time.Time `sql:"time"`        // optional
	RawExpires string    `sql:"raw_expires"` // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int      `sql:"max_age"`
	Secure   bool     `sql:"secure"`
	HttpOnly bool     `sql:"http_only"`
	Raw      string   `sql:"raw"`
	Unparsed []string `sql:"unparsed"` // Raw text of unparsed attribute-value pairs
}

func GetCookies(db *sql.DB, id int) ([]*Cookie, error) {
	rows, err := db.Query("SELECT * FROM `cookies` WHERE `account_id`=?", id)
	if err != nil {
		return nil, err
	}

	var cookies []*Cookie
	for rows.Next() {
		var c Cookie
		if err := rows.Scan(
			&c.ID,
			&c.AccountID,
			&c.Name,
			&c.Value,
			&c.Path,
			&c.Domain,
			&c.Expires,
			&c.RawExpires,
			&c.MaxAge,
			&c.Secure,
			&c.HttpOnly,
			&c.Raw,
			&c.Unparsed); err != nil {
			return nil, err
		}

		cookies = append(cookies, &c)
	}

	return cookies, nil
}

func Save(db *sql.DB, cookies []Cookie) error {
	return nil
}

func From(c *http.Cookie) Cookie {
	return Cookie{
		Name:  c.Name,
		Value: c.Value,

		Path:       c.Path,
		Domain:     c.Domain,
		Expires:    c.Expires,
		RawExpires: c.RawExpires,

		MaxAge:   c.MaxAge,
		Secure:   c.Secure,
		HttpOnly: c.HttpOnly,
		Raw:      c.Raw,
		Unparsed: c.Unparsed,
	}
}

func To(c *Cookie) *http.Cookie {
	return &http.Cookie{
		Name:  c.Name,
		Value: c.Value,

		Path:       c.Path,
		Domain:     c.Domain,
		Expires:    c.Expires,
		RawExpires: c.RawExpires,

		MaxAge:   c.MaxAge,
		Secure:   c.Secure,
		HttpOnly: c.HttpOnly,
		Raw:      c.Raw,
		Unparsed: c.Unparsed,
	}
}
