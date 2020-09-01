package bootstrap

import "github.com/inhere/go-web-skeleton/app/clog"

type Bootstrapper interface {
	// Name() string
	Boot() error
}

type BootFunc func() error

func (bf BootFunc) Boot() error {
	return bf()
}

type Launcher struct {
	Boots []Bootstrapper
}

func (l *Launcher) Add(boots ...Bootstrapper)  {
	l.Boots = append(l.Boots, boots...)
}

func (l *Launcher) Run() {
	for _, boot := range l.Boots {
		err := boot.Boot()
		if err != nil {
			clog.Fatalf(err.Error())
		}
	}
}
