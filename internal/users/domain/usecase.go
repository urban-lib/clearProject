package domain

type UseCase interface {
	Authenticate()
	Logout()
}
