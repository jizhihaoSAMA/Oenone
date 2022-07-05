package crud

//func ESSearchWithBody(body gin.H) (map[string]interface{}, error) {
//	es := base.GLOBAL_RESOURCE[base.ESClient].(*elastic.Client)
//
//	query
//
//	//result, err := es.Search(
//	//	es.Search.WithIndex(base.NeighborhoodInfo),
//	//	es.Search.WithBody(&buf),
//	//	es.Search.WithPretty(),
//	//)
//	//
//	//if err != nil {
//	//	return nil, err
//	//}
//	//defer result.Body.Close()
//	//
//	//var r map[string]interface{}
//	//if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
//	//	return nil, err
//	//}
//	//NewSearchResponse(result.Body)
//	//if result.IsError() {
//	//	return nil, errors.New(r["error"].(map[string]interface{})["reason"].(string))
//	//} else {
//	//	return r, nil
//	//}
//}
