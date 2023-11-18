package pack

func PFUNC() error {
	println("pi")
	var err error
	if err != nil {
		return err
	}

	PFUNC()

	return nil
}
