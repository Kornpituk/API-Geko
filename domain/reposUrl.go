package domain

type ReposUrl struct {
	Github []string `json:"github"`
	Bitbucket []string `json:"bitbucket"`
}