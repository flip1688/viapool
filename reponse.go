package viapool

type DefaultResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type DefaultListResponse struct {
	Code        int           `json:"code"`
	Count       int           `json:"count"`
	CurrentPage int           `json:"curr_page"`
	Data        []interface{} `json:"data"`
	HasNext     bool          `json:"has_next"`
	Total       int           `json:"total"`
	TotalPage   int           `json:"total_page"`
	Message     string        `json:"message"`
}

type PagingDataResponse struct {
	Count       int           `json:"count"`
	CurrentPage int           `json:"curr_page"`
	Data        []interface{} `json:"data"`
	HasNext     bool          `json:"has_next"`
	Total       int           `json:"total"`
	TotalPage   int           `json:"total_page"`
}

type AccountInfoResponse struct {
	Account           AccountInfo        `json:"account"`
	Observers         []*Observer        `json:"observer"`
	WihtdrawAddresses []*WithdrawAddress `json:"withdraw_address"`
	BalanceWallets    []*WalletBalance   `json:"balance"`
}

type AccountInfo struct {
	ID           int    `json:"id"`
	ParentUserId *int   `json:"parent_user_id"`
	CreateTime   int64  `json:"create_time"`
	AccountName  string `json:"account"`
	Email        string `json:"email"`
	CountryCode  string `json:"country_code"`
	Mobile       string `json:"mobile"`
	Country      string `json:"country"`
}

type SubAccount struct {
	Account   string `json:"account"`
	ApiKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
	ID        int    `json:"id"`
	CreatTime int64  `json:"create_time"`
}

type Observer struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	AccessKey string `json:"access_key"`
	CreatTime int64  `json:"create_time"`
}

type WithdrawAddress struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
}

type WalletBalance struct {
	Coin   string  `json:"coin"`
	Amount float64 `json:"amount,string"`
}
