package session_s

const capacityCache = 7

type sessionCache struct {
	lastRequestIndex  int
	lastResponseIndex int

	//Dictionary  {Request: Response}
	Cache map[string]string
}
