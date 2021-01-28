package env

type (
	DBOption struct {
		EnableYYDB bool
	}

	Option struct {
		DBOption
	}
)

func Configure(opt Option) error {

	for _, confFn := range []func() error{
		logger,
		timeLocation,
		databases,
	} {
		err := confFn()
		if err != nil {
			return err
		}
	}

	return nil
}
