package base

import "github.com/olivere/elastic/v7"

func GetErrorDetailFromES(err error) string {
	e, ok := err.(*elastic.Error)
	if !ok {
		return "这不是ES错误"
	}
	return MustMarshal(e.Details)
}
