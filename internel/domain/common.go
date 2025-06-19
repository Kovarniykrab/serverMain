package domain

import (
	"github.com/cadyrov/godict/v2"
	"github.com/cadyrov/goerr/v2"
	"net/http"
)

type SearchForm struct {
	Query       string  `json:"query"`
	IDs         []int64 `json:"ids"`
	ExcludedIDs []int64 `json:"excludedIds"`
	Limit       int     `json:"limit"`
	Page        int     `json:"page"`
}

func (sf SearchForm) Pagination(total int) godict.Pagination {
	return godict.Pagination{Page: sf.Page, Limit: sf.Limit, Total: total}
}

func BadErr(e error) goerr.IError {
	return err(http.StatusBadRequest, e)
}

func AuthErr(e error) goerr.IError {
	return err(http.StatusUnauthorized, e)
}

func IntErr(e error) goerr.IError {
	return err(http.StatusInternalServerError, e)
}

func ConfErr(e error) goerr.IError {
	return err(http.StatusConflict, e)
}

func ForbidErr(e error) goerr.IError {
	return err(http.StatusForbidden, e)
}

func NotFoundErr(e error) goerr.IError {
	return err(http.StatusNotFound, e)
}

func err(code int, err error) goerr.IError {
	return goerr.New(code, err)
}
