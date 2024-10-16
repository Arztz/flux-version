package gitlab

type Interface interface {
	LoadRepo() (err error)
	DeleteRepo() (err error)
}
