package gate

type (
    CreateProductReq struct {
        ProductItem
        Desc string `json:"desc"`
    }
    CreateProductResp struct{}

    QueryPopularProductReq  struct{}
    QueryPopularProductResp struct {
        ProductList []*ProductItem `json:"product_list"`
    }

    ProductItem struct {
        Pid      string  `json:"pid"`
        Title    string  `json:"title"`
        Category string  `json:"category"`
        Price    float64 `json:"price"`
        CoverUrl string  `json:"cover_url"`
    }
)
