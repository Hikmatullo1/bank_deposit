package deposit

func Calculate(amount int, currency string) (min int, max int) {

	minPercent := 0
	maxPercent := 0

	if currency == "TJS" {
		minPercent = 4
		maxPercent = 6
	} else if currency == "USD" {
		minPercent = 1
		maxPercent = 2
	}

	min = amount * minPercent / 100 / 12
	max = amount * maxPercent / 100 / 12

	return min, max
}
