package shotner

type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
