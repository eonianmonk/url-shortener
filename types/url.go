package types

import "strings"

type Url string
type ShortUrl string
type ID int32

func (u *Url) Verify() {
	u.verifyScheme()
}

func (u *Url) verifyScheme() {
	if strings.HasPrefix(string(*u), "http://") || strings.HasPrefix(string(*u), "https://") {
		return
	}
	*u = "http://" + *u

}
