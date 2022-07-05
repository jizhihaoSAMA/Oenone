package public

import (
    "Oenone/common/base"
    "Oenone/model"
    "context"
    "encoding/json"
    "github.com/olivere/elastic/v7"
    "github.com/thoas/go-funk"
    "log"
)

func SearchHouse(name string, area string, bc string, neighborhood string) (interface{}, error) {
    es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)

    query := elastic.NewBoolQuery().Must(elastic.NewMatchQuery("house_name", name))
    if bc != "" {
        query.Must(elastic.NewTermQuery("house_loc_bc.keyword", bc))
    }
    if area != "" {
        query.Must(elastic.NewTermQuery("house_loc_area.keyword", area))
    }
    if neighborhood != "" {
        query.Must(elastic.NewTermQuery("house_neighborhood.keyword", neighborhood))
    }

    query.Must(elastic.NewTermQuery("status", model.Online))
    queryResult, err := es.Search().
        Index(base.HouseInfo).
        Query(query).
        Size(5).
        Highlight(elastic.NewHighlight().Field("house_name").PreTags("<b style='color: black;'>").PostTags("</b>")).
        FetchSourceContext(elastic.NewFetchSourceContext(true).Include("house_loc_name", "house_loc_area", "house_loc_bc", "house_neighborhood")).
        Do(context.Background())
    if err != nil {
        log.Println("[SearchHouse] 查找房屋出错，err: ", err)
        return nil, err
    }

    res := funk.Map(queryResult.Hits.Hits, func(hit *elastic.SearchHit) model.HouseSearchHit {
        var tempHouse model.House
        json.Unmarshal(hit.Source, &tempHouse)
        return model.HouseSearchHit{
            ID:                hit.Id,
            HouseName:         tempHouse.HouseName,
            HouseLocArea:      tempHouse.HouseLocArea,
            HouseLocBC:        tempHouse.HouseLocBC,
            HouseNeighborhood: tempHouse.HouseNeighborhood,
            Content:           hit.Highlight["house_name"][0],
        }
    })
    return res, nil

}
