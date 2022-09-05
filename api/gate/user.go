package gate

type (
    CreateUserReq struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
        Avatar   string `json:"avatar"`
    }
    CreateUserResp struct{}
)

type (
    QueryCartReq struct {
        Uid string `form:"uid"`
    }
    QueryCartResp struct {
        Items []*CartItem `json:"items"`
    }
    CartItem struct {
        Pid      string  `json:"pid"`
        Title    string  `json:"title"`
        Type     string  `json:"type"`
        Count    int32   `json:"count"`
        Price    float64 `json:"price"`
        CoverUrl string  `json:"cover_url"`
    }
)
