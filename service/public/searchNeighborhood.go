package public

import (
	"Oenone/common/base"
	"Oenone/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/thoas/go-funk"
)

func SearchNeighborhood(neighborhood string, area string, size int) ([]model.NeighborhoodSearchHit, error) {
	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)

	boolQuery := elastic.NewBoolQuery()
	matchQuery := elastic.NewMatchQuery("name", neighborhood)
	termQuery := elastic.NewTermQuery("area", area)

	boolQuery.Must(matchQuery, termQuery)
	result, err := es.Search().Index(base.NeighborhoodInfo).
		Query(boolQuery).
		Pretty(true).
		From(0).
		Size(size).
		Highlight(elastic.NewHighlight().Field("name").PreTags("<b style='color: black;'>").PostTags("</b>")).
		FetchSourceContext(elastic.NewFetchSourceContext(true).Include("name")).
		Do(context.Background())

	if err != nil {
		return nil, err
	}
	res := funk.Map(result.Hits.Hits, func(hit *elastic.SearchHit) model.NeighborhoodSearchHit {
		var tempNeighborhood model.Neighborhood
		json.Unmarshal(hit.Source, &tempNeighborhood)
		return model.NeighborhoodSearchHit{
			Name:    tempNeighborhood.Name,
			Content: hit.Highlight["name"][0],
		}
	})

	return res.([]model.NeighborhoodSearchHit), nil
}
