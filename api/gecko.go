package api

import "myapp-project-gecko/domain"

type GeckoRequest struct {
}

type GeckoResponse struct {
	Id                 string              `json:"id"`
	Symbol             string              `json:"symbol"`
	Name               string              `json:"name"`
	AssetPlatformId    string              `json:"asset_platform_id"`
	Platforms          domain.Platforms    `json:"platforms"`
	BlockTimeInMinutes int                 `json:"block_time_in_minutes"`
	HashingAlgorithm   string              `json:"hashing_algorithm"`
	Categories         []string            `json:"categories"`
	PublicNotice       string              `json:"public_notice"`
	AdditionalNotices  []string            `json:"additional_notices"`
	Localization       domain.Localization `json:"localization"`
	Description        domain.Description  `json:"description"`
	Link               domain.Link         `json:"link"`
	Image              domain.Image        `json:"image"`
}
