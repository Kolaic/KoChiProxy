package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Run(port int) {
	r := router()
	_ = http.ListenAndServe(":8080", r)
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/random", GetProxy)
	r.HandleFunc("/count", GetCount)
	r.HandleFunc("/limit/{limit}", GetLimitProxy)
	return r
}
// 获取随机代理
func GetProxy(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(666)
}

// 获取代理数量
func GetCount(w http.ResponseWriter, r *http.Request)  {
	
}

// 一次获取多个代理

func GetLimitProxy(w http.ResponseWriter, r *http.Request)  {

}