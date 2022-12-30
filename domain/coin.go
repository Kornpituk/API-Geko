package domain

type GeckoTrue struct {
	Id                    string       `json:"id"`
	Symbol                string       `json:"symbol"`
	Name                  string       `json:"name"`
	Asset_platform_id     string       `json:"asset_platform_id"`
	Platforms             Platforms    `json:"platforms"`
	Block_time_in_minutes int          `json:"block_time_in_minutes"`
	Hashing_algorithm     string       `json:"hashing_algorithm"`
	Categories            []string     `json:"categories"`
	Localization          Localization `json:"localization"`
	Description           Description  `json:"description"`
	PublicNotice          string       `json:"public_notice"`
	AdditionalNotices     []string     `json:"additional_notices"`
	Link                  Link         `json:"link"`
	Image                 Image        `json:"image"`
}
