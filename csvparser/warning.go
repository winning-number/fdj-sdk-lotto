package csvparser

type Warning map[string][]string

func NewWarning() Warning {
	return make(map[string][]string)
}

func (w *Warning) addUnusedValues(key, value string) {
	if _, ok := (*w)[key]; !ok {
		(*w)[key] = []string{value}

		return
	}
	(*w)[key] = append((*w)[key], value)
}
