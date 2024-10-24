package connector

type HTTPClient interface {
	Get(url string) ([]byte, error)
}
