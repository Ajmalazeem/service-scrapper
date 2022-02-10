package scrap

import (

	"testing"

	"bitbucket.org/ajmal_azm/scraperP/model"
)


func TestScrap(t *testing.T) {

	want := model.Model{
		Url:                 "https://play.google.com/store/apps/details?id=com.pg.oralb.oralbapp",
		PackageName:         "com.pg.oralb.oralbapp",
		AppName:             "Oral-B",
		DeveloperName:       "Procter & Gamble Productions",
		ImageUrl:            "https://play-lh.googleusercontent.com/fO1sD60QXN1_EUZgAg7e6WfsAu0GyzlgkJ0O8ONja0a8bp7AACu8VLh-7pho212cKchL=s180",
		Rating:              "4.3",
		//RatedPeopleCount:    "64,921",
		InAppPurchase:       false,
		ContainAds:          false,
		Updated:             "January 4, 2022",
		Size:                "83M",
		Installs:            "1,000,000+",
		CurrentVersion:      "8.7.1",
		AndroidVersion:      "7.0 and up",
		ContentRating:       "Rated for 3+Learn more",
		InteractiveElements: "",
		InAppProducts:       "",
		OfferedBy:           "Procter & Gamble Productions",
		Developer:           "Visit websiteproctergambleproductions@gmail.comPrivacy Policy",
	}

	got := Scraper("https://play.google.com/store/apps/details?id=com.pg.oralb.oralbapp")

	if got.Url != want.Url {
		t.Errorf("url fail got := %s , want := %s",got.AppName,want.AppName)
	}
	if got.PackageName != want.PackageName {
		t.Errorf("package name fail got := %s , want := %s",got.PackageName,want.PackageName)
	}
	if got.AppName != want.AppName {
		t.Errorf("appname fail got := %s , want := %s",got.AppName,want.AppName)
	}
	if got.DeveloperName != want.DeveloperName {
		t.Errorf("developername fail got := %s , want := %s",got.DeveloperName,want.DeveloperName)
	}
	if got.ImageUrl != want.ImageUrl {
		t.Errorf("image url fail got := %s , want := %s",got.ImageUrl,want.ImageUrl)
	}
	if got.Rating != want.Rating {
		t.Errorf("rating fail got := %s , want := %s",got.Rating,want.Rating)
	}
	// if got.RatedPeopleCount != want.RatedPeopleCount {
	// 	t.Errorf("rated people count fail got := %s , want := %s",got.RatedPeopleCount,want.RatedPeopleCount)
	// }
	if got.InAppPurchase != want.InAppPurchase {
		t.Errorf("inapppurchase fail got := %t , want := %t",got.InAppPurchase,want.InAppPurchase)
	}
	if got.ContainAds != want.ContainAds {
		t.Errorf("contain ads fail got := %t , want := %t",got.ContainAds,want.ContainAds)
	}
	if got.Updated != want.Updated {
		t.Errorf("updated fail got := %s , want := %s",got.Updated,want.Updated)
	}
	if got.Size != want.Size {
		t.Errorf("size fail got := %s , want := %s",got.Size,want.Size)
	}
	if got.Installs != want.Installs {
		t.Errorf("installs fail got := %s , want := %s",got.Installs,want.Installs)
	}
	if got.CurrentVersion != want.CurrentVersion {
		t.Errorf("current version fail got := %s , want := %s",got.CurrentVersion,want.CurrentVersion)
	}
	if got.AndroidVersion != want.AndroidVersion {
		t.Errorf("Android version fail got := %s , want := %s",got.AndroidVersion,want.AndroidVersion)
	}
	if got.ContentRating != want.ContentRating {
		t.Errorf("content rating fail got := %s , want := %s",got.ContentRating,want.ContentRating)
	}
	if got.InteractiveElements != want.InteractiveElements {
		t.Errorf("interactive elemnts fail got := %s , want := %s",got.InteractiveElements,want.InteractiveElements)
	}
	if got.InAppProducts != want.InAppProducts {
		t.Errorf("ina app products fail got := %s , want := %s",got.InAppProducts,want.InAppProducts)
	}
	if got.OfferedBy != want.OfferedBy {
		t.Errorf("offered by fail got := %s , want := %s",got.OfferedBy,want.OfferedBy)
	}
	if got.Developer != want.Developer {
		t.Errorf("Developer fail got := %s , want := %s",got.Developer,want.Developer)
	}

}
