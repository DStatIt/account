package account

import (
	"database/sql"
	"net/http"
	"time"
)

type Cookie struct {
	ID        int `json:"id"`
	AccountID int `json:"account_id"`

	Name  string `json:"name"`
	Value string `json:"value"`

	Path       string    `json:"path"`        // optional
	Domain     string    `json:"domain"`      // optional
	Expires    time.Time `json:"time"`        // optional
	RawExpires string    `json:"raw_expires"` // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int      `json:"max_age"`
	Secure   bool     `json:"secure"`
	HttpOnly bool     `json:"http_only"`
	Raw      string   `json:"raw"`
	Unparsed []string `json:"unparsed"` // Raw text of unparsed attribute-value pairs
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
