package usecase

func setIfNotNil[T any](dest *T, src *T) {
	if src != nil {
		*dest = *src
	}
}
