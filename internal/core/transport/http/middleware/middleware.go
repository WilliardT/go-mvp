package core_http_middleware

type Middleware func(http.Handler) http.Handler