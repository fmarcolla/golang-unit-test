package main

func VerificaPar(value int) bool {

	if value < 0 {
		return false;
	}

	if value % 2 == 0 {
		return true
	}

	return false;
}