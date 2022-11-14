package erratum

func Use(opener ResourceOpener, input string) (err error) {
	res, err := opener()
	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(opener, input)
		default:
			return err
		}
	}

	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case FrobError:
				res.Defrob(t.defrobTag)
				err = t.inner
			case error:
				err = t
			}
		}
		defer res.Close()
	}()
	res.Frob(input)
	return
}
