package repository

import (
	"github.com/microcmsio/microcms-go-sdk"

	"github.com/Mckee404/microcms_search_api/config"
	"github.com/Mckee404/microcms_search_api/models"
)

func SearchBlog(q string) (models.BlogList, error) {
	serviceDomain := config.GetServiceDomain()
	apiKey := config.GetApiKey()

	client := microcms.New(serviceDomain, apiKey)

	var searchResult models.BlogList

	fields := []string{"id", "updatedAt", "title", "thumbnail", "tags.id", "tags.tag", "category.id", "category.category", "overview"}

	err := client.List(
		microcms.ListParams{
			Endpoint: "blog",
			Fields:   fields,
			Q:        q,
		},
		&searchResult,
	)

	if err != nil {
		return models.BlogList{}, err
	}

	return searchResult, nil
}
