package main

import (
	"controllers/productcontroller"
	"net/http"
)

func main() {

	http.HandleFunc("/", productcontroller.Index)
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/index", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/processadd", productcontroller.ProcessAdd)
	http.HandleFunc("/product/delete", productcontroller.Delete)
	http.HandleFunc("/product/update", productcontroller.Update)
	http.HandleFunc("/product/edit", productcontroller.Edit)

	http.ListenAndServe(":3000", nil)

}
