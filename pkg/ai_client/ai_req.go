package ai_client

import (
	"context"
	"errors"
	"log"
	"time"
)

func (c *ClientAIHttp) GetBrigV1(ctx context.Context, promt string) (GetBrigV1Response, error) {

	var (
		url      = "/analyze"
		err      error
		request  GetBrigV1Request = GetBrigV1Request{}
		response BrigV1Response   = BrigV1Response{}
	)
	log.Println(c.httpClient1.BaseURL)
	if promt == "" {
		return GetBrigV1Response{}, errors.New("sent empty data lol")
	}

	request.Text = promt

	ctx_timeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	resp, err := c.httpClient1.R().SetHeader("Content-Type", "application/json").SetContext(ctx_timeout).SetResult(&response).SetBody(request).Post(url)
	_ = resp
	log.Println("CLIENT ", response, resp)
	if err != nil {
		return GetBrigV1Response{}, err
	}
	return GetBrigV1Response{
		Stations: response.Stations,
		Period:   response.Period,
	}, nil
}

// TODO
func (c *ClientAIHttp) GetVendorAIRU(ctx context.Context, promt string) (GetVendorAIRUResponse, error) {

	// var (
	// 	url      = "/"
	// 	err      error
	// 	response *models.GetAVRGResp = new(models.GetAVRGResp)
	// )

	// ctx_timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	// params := &models.GetRateByRatingRequest{
	// 	Rating: in.Rating,
	// }

	// resp, err := c.httpClient.R().SetContext(ctx_timeout).SetResult(response).SetBody(params).Post(url)
	// _ = resp
	// if err != nil {
	// 	return 0, err
	// }

	// avrg, err := strconv.ParseFloat(response.Rate, 64)
	// if err != nil {
	// 	return 0, err
	// }
	return GetVendorAIRUResponse{}, nil
}

// TODO
func (c *ClientAIHttp) GetVendorAIEU(ctx context.Context, promt string) (GetVendorAIEUResponse, error) {

	// var (
	// 	url      = "/analyze"
	// 	err      error
	// 	response *models.GetAVRGResp = new(models.GetAVRGResp)
	// )

	// ctx_timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	// params := &models.GetRateByRatingRequest{
	// 	Rating: in.Rating,
	// }

	// resp, err := c.httpClient.R().SetContext(ctx_timeout).SetResult(response).SetBody(params).Post(url)
	// _ = resp
	// if err != nil {
	// 	return 0, err
	// }

	// avrg, err := strconv.ParseFloat(response.Rate, 64)
	// if err != nil {
	// 	return 0, err
	// }
	return GetVendorAIEUResponse{}, nil
}
