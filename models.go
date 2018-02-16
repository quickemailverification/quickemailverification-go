package quickemailverification

// Response struct provides the data mappings for API responses
type Response struct {
	Result     string  `json:"result"`
	Reason     string  `json:"reason"`
	Disposable string  `json:"disposable"`
	AcceptAll  string  `json:"accept_all"`
	Role       string  `json:"role"`
	Free       string  `json:"free"`
	Email      string  `json:"email"`
	User       string  `json:"user"`
	Domain     string  `json:"domain"`
	SafeToSend string  `json:"safe_to_send"`
	Suggested  string  `json:"did_you_mean"`
	Success    string  `json:"success"`
	Message    string  `json:"message"`
}

