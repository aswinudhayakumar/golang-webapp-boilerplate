package model

//Output is a type for rest api response
type Output struct {
	Code  int64
	Data  interface{}
	Error error
}
