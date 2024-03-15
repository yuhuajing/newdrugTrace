package tabletypes

type Identity string

var (
	User  Identity = "user"
	Admin Identity = "admin"
	//Prod  Identity = "prod"
	//Logis Identity = "logis"
	//Store Identity = "store"
	//Sale  Identity = "sale"
)

type UserInfo struct {
	Id       string   `json:"id" gorm:"primary_key"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Identity Identity `json:"identity"`
	Approved int64    `json:"approved"`
}

type Info struct {
	Username string `json:"username"`
}

type DealInfo struct {
	Dealdata string `json:"dealdata"`
}
type Tousu struct {
	Id       string `json:"id" gorm:"primary_key"`
	User     string `json:"user"`
	Tousu    string `json:"tousu"`
	Dealdata string `json:"dealdata"`
	Dealed   int64  `json:"dealed"`
}

type TraceInfo struct {
	TraceID       string `json:"traceID"`
	ProdID        string `json:"prodID" `
	ProdBaseInfo  string `json:"prodBaseInfo"`
	ProdTxHash    string `json:"prodTxHash"`
	StoreID       string `json:"storeID" `
	StoreBaseInfo string `json:"storeBaseInfo"`
	StoreTxHash   string `json:"storeTxHash"`
	LogisID       string `json:"logisID" `
	LogisBaseInfo string `json:"logisBaseInfo"`
	LogisTxHash   string `json:"logisTxHash"`
	SaleID        string `json:"saleID" `
	SaleBaseInfo  string `json:"saleBaseInfo"`
	SaleTxHash    string `json:"saleTxHash"`
}

type ProduceInfo struct {
	Id           string `json:"id" gorm:"primary_key"`
	Prodcompany  string `json:"prodcompany"`
	Name         string `json:"name"`
	Prodbatchnum string `json:"prodbatchnum"`
	Prodtraceid  string `json:"prodtraceid"`
	Baseinfo     string `json:"baseinfo"`
	TxHash       string `json:"txHash"`
}

type EcoResData struct {
	Id          string `json:"id"`
	Prodcompany string `json:"prodcompany"`
	Prodtraceid string `json:"prodtraceid"`
	Baseinfo    string `json:"baseinfo"`
	Identity    string `json:"identity"`
	Hash        string `json:"hash"`
}

type NewUserInfo struct {
	Username    string `json:"username"`
	Newusername string `json:"newusername"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
type NewAdminUserInfo struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
