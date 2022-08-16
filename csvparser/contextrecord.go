package csvparser

type ContextRecord interface {
	Get(key string) (value string, found bool)
	Set(key string, value string)
}

type contextRecord struct {
	params map[string]string
}

func NewContextRecord() ContextRecord {
	return &contextRecord{params: make(map[string]string)}
}

// Get retreive the associated value from the key parameter and return it
// If no value is associated, empty string is return and found should be set at false
func (c *contextRecord) Get(key string) (string, bool) {
	if v, ok := c.params[key]; ok {
		return v, true
	}

	return "", false
}

// Set associate the value string to the key parameter
// If a value already exist for the key, it should be replace
func (c *contextRecord) Set(key, value string) {
	c.params[key] = value
}
