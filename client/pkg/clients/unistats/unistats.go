package unistats

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/supperdoggy/diploma_university_statistics_tool/models/rest"
	"go.uber.org/zap"
)

const (
	post  = "POST"
	get   = "GET"
	apiv1 = "/api/v1"
)

type IUniStats interface {
	// schools
	SchoolList() (*rest.ListSchoolsResponse, error)
	TopCompanies(school string) (*rest.ListSchoolsTopCompaniesResponse, error)
	TopHiredDegrees(school, company string) (*rest.TopHiredDegreesResponse, error)

	TopSchools(company string) (*rest.ListCompaniesTopSchoolsResponse, error)
}

type uniStats struct {
	url string

	log *zap.Logger
}

func NewUniStats(log *zap.Logger, url string) IUniStats {
	return &uniStats{
		url: url,

		log: log,
	}
}

func (u *uniStats) makeApiV1Req(path, method string, body, response interface{}) error {
	var raw io.Reader
	if method != get {
		var err error
		data, err := json.Marshal(body)
		if err != nil {
			u.log.Error("Error while marshaling body", zap.Error(err))
			return err
		}

		raw = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, u.url+apiv1+path, raw)
	if err != nil {
		u.log.Error("Error while creating request", zap.Error(err))
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		u.log.Error("Error while making request", zap.Error(err))
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Error while reading response", zap.Error(err))
		return err
	}

	if err := json.Unmarshal(data, response); err != nil {
		u.log.Error("Error while unmarshaling response", zap.Error(err))
		return err
	}

	return nil
}

func (u *uniStats) SchoolList() (*rest.ListSchoolsResponse, error) {
	var resp rest.ListSchoolsResponse
	if err := u.makeApiV1Req("/schools/list", get, nil, &resp); err != nil {
		u.log.Error("Error while making request", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (u *uniStats) TopCompanies(school string) (*rest.ListSchoolsTopCompaniesResponse, error) {
	var resp rest.ListSchoolsTopCompaniesResponse
	if err := u.makeApiV1Req("/schools/top_companies", post, rest.ListSchoolsTopCompaniesRequest{School: school}, &resp); err != nil {
		u.log.Error("Error while making request", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (u *uniStats) TopHiredDegrees(school, company string) (*rest.TopHiredDegreesResponse, error) {
	var resp rest.TopHiredDegreesResponse
	if err := u.makeApiV1Req("/companies/top_degrees_hired", post,
		rest.TopHiredDegreesRequest{School: school, Company: company}, &resp); err != nil {
		u.log.Error("Error while making request", zap.Error(err))
		return nil, err
	}

	for k, v := range resp.Degrees {
		if v.Name == "" {
			resp.Degrees[k].Name = "no data"
		}
	}

	return &resp, nil
}

func (u *uniStats) TopSchools(company string) (*rest.ListCompaniesTopSchoolsResponse, error) {
	var resp rest.ListCompaniesTopSchoolsResponse
	if err := u.makeApiV1Req("/companies/top_schools", post,
		rest.ListCompaniesTopSchoolsRequest{Company: company}, &resp); err != nil {
		u.log.Error("Error while making request", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}
