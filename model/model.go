package model

type Model struct {
	Url                 string `json:"urlid" gorm:"column:urlid"`
	PackageName         string `json:"package_name" gorm:"column:package_name"`
	AppName             string `json:"app_name" gorm:"column:app_name"`
	DeveloperName       string `json:"company" gorm:"column:company"`
	ImageUrl            string `json:"image_url" gorm:"column:image_url"`
	Rating              string `json:"rating" gorm:"column:rating"`
	RatedPeopleCount    string `json:"rated_people" gorm:"column:rated_people"`
	InAppPurchase       bool   `json:"purchase_ads" gorm:"column:purchase_ads"`
	ContainAds          bool   `json:"contain_ads" gorm:"column:contain_ads"`
	Updated             string `json:"updated" gorm:"column:updated"`
	Size                string `json:"size" gorm:"column:size"`
	Installs            string `json:"installs" gorm:"column:installs"`
	CurrentVersion      string `json:"current_version" gorm:"column:current_version"`
	AndroidVersion      string `json:"android_version" gorm:"column:android_version"`
	ContentRating       string `json:"content_rating" gorm:"column:content_rating"`
	InteractiveElements string `json:"interactive_ele" gorm:"column:interactive_ele"`
	InAppProducts       string `json:"in_app_products" gorm:"column:in_app_products"`
	OfferedBy           string `json:"offered_by" gorm:"column:offered_by"`
	Developer           string `json:"developer" gorm:"column:developer"`
}

type GetRequest struct {
	PackageName string
}

type Changelog struct {
	PackageName string `json:"package_name" gorm:"column:package_name"`
	Field       string `json:"field" gorm:"column:field"`
	Old         string `json:"old" gorm:"column:old"`
	New         string `json:"new" gorm:"column:new"`
	Updated     string `json:"updated" gorm:"column:updated"`
}
