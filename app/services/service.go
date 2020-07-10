package services

// Service generic service handler
type Service func(config interface{}) error

// ErrorService generic service error handler
func (fn Service) ErrorService(c interface{}) {
	if err := fn(c); err != nil {
		// TODO: need to add logging here
	}
}
