package domain

// func MakeJSONResponse(httpCode int, message string, products []*Product) string {
func MakeJSONResponse(httpCode int, message string, products []*Product) map[string]any {
	return map[string]any{
		"meta": map[string]interface{}{
			"code":    httpCode,
			"message": message,
		},
		"products": products,
	}
	//_ = resp
	//j, err := json.Marshal(resp)
	//if err != nil {
	//	fmt.Println("json marshal error")
	//}
	//return j
}

/*
{
   "meta":{
       "code": 200, // http status code와 같은 code를 응답으로 전달
       "message":"ok" // 에러 발생시, 필요한 에러 메시지 전달
		},
		"data":{
       "products":[...]
		}
}


// 400 Bad Request Example
{
   "meta":{
       "code": 400,
       "message": "잘못된 상품 사이즈 입니다."
		},
		"data": null
}
*/
