package api

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"bitbucket.org/ajmal_azm/scraperP/model"
	"bitbucket.org/ajmal_azm/scraperP/store"
	//"bitbucket.org/ajmal_azm/scraperP/web"
)

type MockRepository struct {
	mock.Mock
	store.WebStore
}

func (mock *MockRepository) GetPackageNameDetails(model.GetRequest) (*model.Model, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Model), args.Error(1)
}

func (mock *MockRepository) GetChangeLogDetails(model.GetRequest) (*[]model.Changelog, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*[]model.Changelog), args.Error(1)
}

func TestGetPackageNameDetails(t *testing.T) {
	mockrepo := new(MockRepository)
	get := model.GetRequest{PackageName: "net.androgames.level"}
	post := model.Model{
		Url:                 "https://play.google.com/store/apps/details?id=net.androgames.level",
		PackageName:         "net.androgames.level",
		AppName:             "Bubble level",
		DeveloperName:       "PixelProse SARL",
		ImageUrl:            "https://play-lh.googleusercontent.com/ip7eoggNFyAQrhSuuF2qXSASkF0jr5wapaCOuTQne4420OvjqEgQfCbYcAHVIXXowg=s180",
		Rating:              "4.6",
		RatedPeopleCount:    "260,979",
		InAppPurchase:       true,
		ContainAds:          true,
		Updated:             "January 28, 2022",
		Size:                "9.2M",
		Installs:            "10,000,000+",
		CurrentVersion:      "10.0.3",
		AndroidVersion:      "4.1 and up",
		ContentRating:       "Rated for 3+Learn more",
		InteractiveElements: "",
		InAppProducts:       "₹199.00 - ₹240.00 per item",
		OfferedBy:           "PixelProse SARL",
		Developer:           "Visit websitelevel.app@pixelprose.frPrivacy Policy466 route de Ferrières\n74350 Cuvat\nFrance",
	}
	var err error
	mockrepo.On("GetPackageNameDetails").Return(&post, err)
	testservice := NewWebService(mockrepo)
	result, err := testservice.GetPackageNameDetails(get)
	log.Println(result)
	if err != nil {
		log.Println(err)
	}
	mockrepo.AssertExpectations(t)
	assert.Equal(t, "https://play.google.com/store/apps/details?id=net.androgames.level", result.Url)
	assert.Equal(t, "net.androgames.level", result.PackageName)
	assert.Equal(t, "Bubble level", result.AppName)
	assert.Equal(t, "PixelProse SARL", result.DeveloperName)
	assert.Equal(t, "https://play-lh.googleusercontent.com/ip7eoggNFyAQrhSuuF2qXSASkF0jr5wapaCOuTQne4420OvjqEgQfCbYcAHVIXXowg=s180", result.ImageUrl)
	assert.Equal(t, "4.6", result.Rating)
	assert.Equal(t, "260,979", result.RatedPeopleCount)
	assert.Equal(t, true, result.InAppPurchase)
	assert.Equal(t, true, result.ContainAds)
	assert.Equal(t, "January 28, 2022", result.Updated)
	assert.Equal(t, "9.2M", result.Size)
	assert.Equal(t, "10,000,000+", result.Installs)
	assert.Equal(t, "10.0.3", result.CurrentVersion)
	assert.Equal(t, "4.1 and up", result.AndroidVersion)
	assert.Equal(t, "Rated for 3+Learn more", result.ContentRating)
	assert.Equal(t, "", result.InteractiveElements)
	assert.Equal(t, "₹199.00 - ₹240.00 per item", result.InAppProducts)
	assert.Equal(t, "PixelProse SARL", result.OfferedBy)
	assert.Equal(t, "Visit websitelevel.app@pixelprose.frPrivacy Policy466 route de Ferrières\n74350 Cuvat\nFrance", result.Developer)
	assert.Nil(t, err)
}

func TestGetChangeLogDetails(t *testing.T) {
	mockrepo := new(MockRepository)
	get := model.GetRequest{PackageName: "com.taxsee.taxsee"}
	post := []model.Changelog{
		{
			PackageName: "com.taxsee.taxsee",
			Field:       "rating",
			Old:         "4.4",
			New:         "4.3",
			Updated:     "2022-02-10T00:00:00Z",
		}, {
			PackageName: "com.taxsee.taxsee",
			Field:       "rated_people",
			Old:         "1,751,471",
			New:         "1,751,909",
			Updated:     "2022-02-10T00:00:00Z",
		},
	}
	var err error
	mockrepo.On("GetChangeLogDetails").Return(&post, err)
	testservice := NewWebService(mockrepo)
	resultGroup, err := testservice.GetChangeLogDetails(get)
	var result []model.Changelog
	result = append(result, *resultGroup...)
	if err != nil {
		log.Println(err)
	}

	mockrepo.AssertExpectations(t)
	assert.Equal(t, "com.taxsee.taxsee", result[0].PackageName)
	assert.Equal(t, "rating", result[0].Field)
	assert.Equal(t, "4.4", result[0].Old)
	assert.Equal(t, "4.3", result[0].New)
	assert.Equal(t, "2022-02-10T00:00:00Z", result[0].Updated)
	assert.Equal(t, "com.taxsee.taxsee", result[1].PackageName)
	assert.Equal(t, "rated_people", result[1].Field)
	assert.Equal(t, "1,751,471", result[1].Old)
	assert.Equal(t, "1,751,909", result[1].New)
	assert.Equal(t, "2022-02-10T00:00:00Z", result[1].Updated)
	assert.Nil(t, err)
}
