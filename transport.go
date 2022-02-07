package main

import (
	"context"
	"encoding/json"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetRequest
	vars := mux.Vars(r)
	packageName := vars["PackageName"]
	var err error
	 if req.PackageName = packageName;err != nil{
		 return nil, err
	 }
	return req, nil
}

func makeGetPEndpoint(svc WebService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{},error) {
		req := request.(GetRequest)
		d, err:=svc.GetP(req)
		if err!= nil{
			return nil, err
		}
		return d, nil
	}
}

func MakeHandler(svc WebService) http.Handler {
	r := mux.NewRouter()

	GetHandler := httptransport.NewServer(
		makeGetPEndpoint(svc),
		decodeGetRequest,
		encodeResponse,
	)
	
	r.Methods(http.MethodGet).Path("/scrap/{PackageName}").Handler(GetHandler)
	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if err:=json.NewEncoder(w).Encode(response); err != nil{
		return err
	}
	return nil
}







