package domain

type Link struct {
	Homepage                    []string `json:"homepage"`
	BlockchainSite              []string `json:"blockchain_site"`
	OfficialForumUrl            []string `json:"official_forum_url"`
	ChatUrl                     []string `json:"chat_url"`
	AnnouncementUrl             []string `json:"announcement_url"`
	TwitterScreenName           []string `json:"twitter_screen_name"`
	FacebookUsername            []string `json:"facebookUsername"`
	BitcointalkThreadIdentifier []string `json:"bitcointalk_thread_identifier"`
	TelegramChannelIdentifier   []string `json:"telegram_channel_identifier"`
	SubredditUrl                []string `json:"subreddit_url"`
	ReposUrl                    ReposUrl `json:"repos_url"`
}
