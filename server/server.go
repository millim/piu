package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Server struct {
	Host string
	Port int
}

var values map[string]string

func init() {
	values = map[string]string{}
}

func (s *Server) InitArgs() {
	flag.StringVar(&s.Host, "host", "127.0.0.1", "server host")
	flag.IntVar(&s.Port, "port", 8080, "server port")
}

func (s *Server) Run() {
	http.HandleFunc("/values", valuesRouter)
	fmt.Println(fmt.Sprintf("server start at %s:%v", s.Host, s.Port))
	http.ListenAndServe(fmt.Sprintf("%s:%v", s.Host, s.Port), nil)
}

func valuesRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("key is empty"))
			return
		}

		value := getValue(key)
		if value == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("key is empty"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(value))
	}

	if r.Method == http.MethodPost {
		deco := json.NewDecoder(r.Body)
		resp := map[string]string{}
		err := deco.Decode(&resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if resp["key"] == "" || resp["value"] == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("key or value is empty"))
			return
		}
		setValue(resp["key"], resp["value"])
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

func setValue(key, value string) {
	values[key] = value
}

func getValue(key string) string {
	value := values[key]
	return value
}
