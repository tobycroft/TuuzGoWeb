package Net

func (h Curl) SetHeaderJson() Curl {
	h.request.SetHeaders(map[string]string{"Content-Type": "application/json"})
	return h
}

func (h Curl) SetHeaderUrlEncode() Curl {
	h.request.SetHeaders(map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	return h
}

func (h Curl) SetHeaderFormData() Curl {
	h.request.SetHeaders(map[string]string{"Content-Type": "multipart/form-data"})
	return h
}

func (h Curl) SetHeaderTextPlain() Curl {
	h.request.SetHeaders(map[string]string{"Content-Type": "text/plain"})
	return h
}
