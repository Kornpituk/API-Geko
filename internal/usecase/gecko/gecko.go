package gecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myapp-project-gecko/api"
	"myapp-project-gecko/domain"
	"net/http"
	"strings"
)

func (s *Service) Gecko() (*api.GeckoResponse, *api.ErrorResponse) {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, strings.ReplaceAll(s.config.Api.Coin, ":id", "binancecoin"), nil)

	if err != nil {
		fmt.Println(err)
		return nil, api.ErrInternalError
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, api.ErrInternalError
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, api.ErrInternalError
	}

	var gecko domain.GeckoTrue

	err = json.Unmarshal(body, &gecko)
	if err != nil {
		log.Fatal(err)
	}

	result := &api.GeckoResponse{
		Id:                 gecko.Id,
		Symbol:             gecko.Symbol,
		Name:               gecko.Name,
		AssetPlatformId:    "",
		Platforms:          domain.Platforms{},
		BlockTimeInMinutes: 0,
		HashingAlgorithm:   "",
		Categories:         gecko.Categories,
		PublicNotice:       gecko.PublicNotice,
		AdditionalNotices:  gecko.AdditionalNotices,
		Localization:       gecko.Localization,
		Description:        gecko.Description,
		Link:               domain.Link{},
		Image:              domain.Image{},
	}

	return result, nil
}
