package statuscodes

type Codes struct {
	list []int
}

func (c *Codes) IsInList(code int) bool {
	for _, b := range c.list {
		if b == code {
			return true
		}
	}
	return false
}

func GetAllowedCodes() *Codes {
	return &Codes{
		list: []int{100, 101, 102, 200, 201, 202, 203, 204, 205, 206, 300, 301, 302, 303, 304, 305, 306, 307, 308, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 422, 428, 429, 431, 451, 500, 501, 502, 503, 504, 505, 511, 520, 522, 524},
	}
}
