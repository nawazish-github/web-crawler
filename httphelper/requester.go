package httphelper

//Requester requests the given url and
//returns the response and/or error
type Requester interface {
	Request() ([]byte, error)
}
