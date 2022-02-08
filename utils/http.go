package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SuccessResponse(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	resp := apiResponse{
		Status:  true,
		Message: "",
		Data:    data,
	}
	marshalledResp, _ := json.Marshal(resp)
	writer.Write(marshalledResp)
}
func SuccessResponseHTML(writer http.ResponseWriter, message string) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusOK)

	html := fmt.Sprintf(htmlSuccess, message)
	writer.Write([]byte(html))
}
func ErrorResponseHTML(writer http.ResponseWriter, message string) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusOK)
	html := fmt.Sprintf(htmlError, message)
	writer.Write([]byte(html))
}

func ErrorResponse(writer http.ResponseWriter, message string, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	data := apiResponse{Status: false, Message: message}
	marshalledData, _ := json.Marshal(data)
	writer.Write(marshalledData)
}

type apiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var htmlError = `<html>
<link href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
<script src="//code.jquery.com/jquery-1.11.1.min.js"></script>
<!------ Include the above in your HEAD tag ---------->

<!-- http://www.jquery2dotnet.com -->
<div class="container" text-align="center">
    <div class="row">
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
               <span class="glyphicon glyphicon-ok"></span> <strong>Something is wrong</strong>
                <hr class="message-inner-separator">
                <p>
                    %s</p>
            </div>
    </div>
</div>
</html>`

var htmlSuccess = `
<html>
<link href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
<script src="//code.jquery.com/jquery-1.11.1.min.js"></script>
<!------ Include the above in your HEAD tag ---------->

<!-- http://www.jquery2dotnet.com -->
<div class="container" text-align="center">
    <div class="row">
            <div class="alert alert-success">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
               <span class="glyphicon glyphicon-ok"></span> <strong>Success Message</strong>
                <hr class="message-inner-separator">
                <p>
                   %s </p>
            </div>
    </div>
</div>
</html>
`
