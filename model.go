package main


type Model struct {
	Url                 string `gorm:"column:urlid"`
	PackageName         string `gorm:"column:package_name"`
	AppName             string `gorm:"column:app_name"`
	DeveloperName       string `gorm:"column:company"`
	ImageUrl            string `gorm:"column:image_url"`
	Rating              string `gorm:"column:rating"`
	RatedPeopleCount    string `gorm:"column:rated_people"`
	InAppPurchase       bool   `gorm:"column:purchase_ads"`
	ContainAds          bool   `gorm:"column:contain_ads"`
	Updated             string `gorm:"column:updated"`
	Size                string `gorm:"column:size"`
	Installs            string `gorm:"column:installs"`
	CurrentVersion      string `gorm:"column:current_version"`
	AndroidVersion      string `gorm:"column:android_version"`
	ContentRating       string `gorm:"column:content_rating"`
	InteractiveElements string `gorm:"column:interactive_ele"`
	InAppProducts       string `gorm:"column:in_app_products"`
	OfferedBy           string `gorm:"column:offered_by"`
	Developer           string `gorm:"column:developer"`
}


