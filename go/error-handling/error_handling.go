package erratum

import "time"

func Use(o ResourceOpener, input string) (ret_err error) {
	resource, err := o()
	for {

		if err != nil {
			if _, ok := err.(TransientError); ok {
				time.Sleep(1e9)
				resource, err = o()
			} else {

				ret_err = err
				return
			}
		} else {
			break
		}

	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			if ferr, ok := r.(FrobError); ok {
				resource.Defrob(ferr.defrobTag)
				ret_err = ferr.inner
			}
			ret_err = r.(error)
		} else {
		  ret_err = nil
		}
	}()
	resource.Frob("hello")
	return
}
