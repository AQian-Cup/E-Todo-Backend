package response

type Response struct {
	HTTP int
	Result
}

type Result interface {
	Result()
}

type ErrorResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type OkResult map[string]interface{}

type OkResultList []map[string]interface{}

func (e ErrorResult) Result() {

}

func (o OkResult) Result() {

}

func (o OkResultList) Result() {

}
