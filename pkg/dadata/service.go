package dadata

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"mayak/config"
	"mayak/internal/domain"
	"net/http"
)

const (
	MinLenQueryDadata = 10
	MaxResultCount    = 10
	suggestAddress    = "suggest/address?"
	suggestAddressIp  = "iplocate/address?"

	findById = "findById/party?"
)

type Response struct {
	Suggestions `json:"suggestions"`
}
type ResponseInn struct {
	SuggestionsInns `json:"suggestions"`
}

type Location struct {
	Location LocationData `json:"location"`
}
type GeoCoordinates struct {
	Latitude  string `json:"geo_lat"`
	Longitude string `json:"geo_lon"`
}

type LocationData struct {
	GeoCoordinates GeoCoordinates `json:"data"`
}

type Service struct {
	logger *zerolog.Logger
	cfg    config.Dadata
	client *http.Client
}

// nolint
func New(conf config.Dadata, log *zerolog.Logger) *Service {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Service{
		logger: log,
		cfg:    conf,
		client: &http.Client{Transport: tr},
	}
}

var ErrInvalidResponce = errors.New("invalid response")

func (s *Service) SuggestAddress(ctx context.Context, queryWithCount string) (Suggestions, error) {
	url := fmt.Sprintf("%s%s%s", s.cfg.APIURI, suggestAddress, queryWithCount)

	resp, err := s.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if s.cfg.DebugMode {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			s.logger.Error().Str("check response status code", string(body))
		}

		return nil, fmt.Errorf("check response status code: %w", ErrInvalidResponce)
	}

	var dadataResp Response

	if err := json.NewDecoder(resp.Body).Decode(&dadataResp); err != nil {
		return nil, err
	}

	return dadataResp.Suggestions, nil
}

func (s *Service) doRequest(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token "+s.cfg.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// nolint
func (s *Service) SuggestInnAddress(ctx context.Context, queryWithCount string) (domain.CompanyDadats, error) {
	url := fmt.Sprintf("%s%s%s", s.cfg.APIURI, findById, queryWithCount)

	resp, err := s.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if s.cfg.DebugMode {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			s.logger.Error().Str("check response status code", string(body))
		}

		return nil, fmt.Errorf("check response status code: %w", ErrInvalidResponce)
	}

	var dadataRespInn ResponseInn

	if err := json.NewDecoder(resp.Body).Decode(&dadataRespInn); err != nil {
		return nil, err
	}

	var companies []domain.CompanyDadata

	for _, item := range dadataRespInn.SuggestionsInns {
		company := domain.CompanyDadata{
			Name:          item.Value,
			Address:       item.Data.Address.Value,
			Inn:           item.Data.INN,
			Kpp:           item.Data.KPP,
			ActualAddress: item.Data.Address.UnrestrictedValue,
			Ogrn:          item.Data.OGRN,
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (s *Service) SuggestAddressIp(ctx context.Context, queryWithCount string) (LocationData, error) {
	url := fmt.Sprintf("%s%s%s", s.cfg.APIURI, suggestAddressIp, queryWithCount)

	resp, err := s.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return LocationData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if s.cfg.DebugMode {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return LocationData{}, err
			}

			s.logger.Error().Str("check response status code", string(body))
		}

		return LocationData{}, fmt.Errorf("check response status code: %w", ErrInvalidResponce)
	}

	var data Location

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return LocationData{}, err
	}

	return data.Location, nil
}
