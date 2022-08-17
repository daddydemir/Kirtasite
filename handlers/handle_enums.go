package handlers

// http methods
const (
	AccessMethods string = "Access-Control-Allow-Methods"
	GET           string = "GET"
	POST          string = "POST"
	PUT           string = "PUT"
	DELETE        string = "DELETE"
	PATHC         string = "PATCH"
)

// origins
const (
	AccessOrigin string = "Access-Control-Allow-Origin"
	ORIGIN       string = "*"
)

// content type
const (
	ContentType string = "Content-Type"
	JSON        string = "application/json"
)
