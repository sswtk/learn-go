/**
 * @Author: roninsswang
 * @Date: 2022/2/23 9:37 下午
 * @File: main
 */

package main

import "net/http"

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello, world"))
	})
	http.ListenAndServe(":8080", nil)
}