package response

type Result struct {
	Code    string
	Message string
}

type InterfaceResponse interface {
	error() InterfaceErrorResponse
	data() InterfaceDataResponse
}

type StructErrorResponse struct {
	HTTP int
	Result
}

func (E *StructErrorResponse) getHTTP() int {
	return E.HTTP
}

func (E *StructErrorResponse) getResult() Result {
	return E.Result
}

type StructDataResponse struct {
	HTTP int
	Result
}

func (D *StructDataResponse) getHTTP() int {
	return D.HTTP
}

func (D *StructDataResponse) getResult() Result {
	return D.Result
}

type InterfaceErrorResponse interface {
	getHTTP() int
	getResult() Result
}

type InterfaceDataResponse interface {
	getHTTP() int
	getResult() Result
}

func (E *StructErrorResponse) error() InterfaceErrorResponse {
	return E
}

func (E *StructErrorResponse) data() InterfaceDataResponse {
	return nil
}

func (D *StructDataResponse) error() InterfaceErrorResponse {
	return nil
}

func (D *StructDataResponse) data() InterfaceDataResponse {
	return D
}
