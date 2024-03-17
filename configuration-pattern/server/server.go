package server

type OptionFunc func(*Option)

type Option struct {
	maxCon int
	tls    bool
	id     string
}

type Server struct {
	Option
}

func defaultOption() *Option {
	return &Option{
		maxCon: 10,
		tls:    false,
		id:     "default",
	}
}

func WithMaxCon(maxCon int) OptionFunc {
	return func(option *Option) {
		option.maxCon = maxCon
	}
}

func WithTLS(option *Option) {
	option.tls = true
}

func WithId(id string) OptionFunc {
	return func(option *Option) {
		option.id = id
	}
}

func New(options ...OptionFunc) *Server {
	option := defaultOption()

	for _, optFunc := range options {
		optFunc(option)
	}

	return &Server{Option: *option}
}
