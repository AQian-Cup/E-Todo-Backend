package response

type Response struct {
	HTTP int
	Result
}

type Result interface {
	Result()
}

type ErrorResult struct {
	Code    string
	Message string
}

type OkResult map[string]interface{}

func (e ErrorResult) Result() {

}

func (o OkResult) Result() {

}
