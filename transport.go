package main

import (
	// "context"
	// "encoding/json"
	// "net/http"

	// "github.com/go-kit/kit/endpoint"
)

// type DetailsRequest struct {
// 	S 			string `json:"s"`
// }

type DetailsResponse struct {
	Url		           		string		`gorm:"column:urlid"`
	PackageName				string		`gorm:"column:package_name"`
	APPNAME       			string		`gorm:"column:app_name"`
	DEVELOPERNAME 			string		`gorm:"column:company"`
	ImageUrl				string		`gorm:"column:image_url"`
	Rating        			string		`gorm:"column:rating"`
	RatedPeopleCount		string		`gorm:"column:rated_people"`
	InAppPurchase 			bool		`gorm:"column:purchase_ads"`
	ContainAds    			bool		`gorm:"column:contain_ads"`
	Updated		  			string		`gorm:"column:updated"`
	Size		  			string		`gorm:"column:size"`
	Installs	  			string		`gorm:"column:installs"`
	CurrentVersion  		string		`gorm:"column:current_version"`
	AndroidVersion			string		`gorm:"column:android_version"`
	ContentRating			string		`gorm:"column:content_rating"`
	InteractiveElements		string		`gorm:"column:interactive_ele"`
	InAppProducts			string		`gorm:"column:in_app_products"`
	OfferedBy				string		`gorm:"column:offered_by"`
	Developer				string		`gorm:"column:developer"`
}


// func makeDetailsEndpoint(svc WebService) endpoint.Endpoint {
// 	return func(_ context.Context, request interface{}) (interface{}, error) {
// 		req := request.(DetailsRequest)

// 		v, err := svc.details(req.S)

// 		if err != nil {
// 			return v, nil
// 		}

// 		return v, nil

// 	}
// }

// func decodeDetailsRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var request DetailsRequest
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		return nil, err
// 	}
// 	return request, nil
// }

// func encodeDetailsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
// 	return json.NewEncoder(w).Encode(response)

// }