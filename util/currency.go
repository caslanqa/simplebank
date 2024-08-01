package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	JPY = "JPY"
	GBP = "GBP"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, JPY, GBP:
		return true
	}
	return false
}
