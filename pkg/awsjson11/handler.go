package awsjson11

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fernomac/aws-local/pkg/common"
)

func sendError(resp http.ResponseWriter, err error) {
	resp.Header().Add("Content-Type", "application/x-amz-json-1.1")
	resp.WriteHeader(400)

	msg := map[string]string{}

	ce, ok := err.(common.Error)
	if ok {
		msg["__type"] = ce.Code
		if ce.Message != "" {
			msg["message"] = ce.Message
		}
	} else {
		msg["__type"] = "InternalFailure"
		msg["message"] = err.Error()
	}

	body, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	resp.Write(body)
}

// HandlerFunc is the type of function the Handler uses to handle things.
type HandlerFunc func([]byte) (interface{}, error)

// Handler handles HTTP requests.
type Handler struct {
	prefix   string
	handlers map[string]HandlerFunc
}

// NewHandler creates a new handler.
func NewHandler(prefix string) *Handler {
	return &Handler{
		prefix:   prefix + ".",
		handlers: make(map[string]HandlerFunc),
	}
}

// HandleWith handles the given operation with the given handler function.
func (h *Handler) HandleWith(op string, handler HandlerFunc) {
	h.handlers[op] = handler
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	target := req.Header.Get("x-amz-target")
	if req.Method != "POST" || req.RequestURI != "/" || !strings.HasPrefix(target, h.prefix) {
		sendError(resp, common.NewError("UnknownOperationException"))
		return
	}

	target = target[len(h.prefix):]

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(500)
		return
	}

	handler, ok := h.handlers[target]
	if !ok {
		sendError(resp, common.NewError("UnknownOperationException"))
		return
	}

	out, err := handler(body)
	if err != nil {
		sendError(resp, err)
		return
	}

	var rbody []byte
	if out != nil {
		rbody, err = json.Marshal(out)
		if err != nil {
			sendError(resp, err)
		}
	}

	resp.Header().Add("Content-Type", "application/x-amz-json-1.1")
	resp.WriteHeader(200)

	if out != nil {
		resp.Write(rbody)
	}
}
