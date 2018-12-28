package structs

import "reflect"

// Like is a part of Account structure
type Like struct {
	ID int32 `json:"id"`
	Ts int32 `json:"ts"`
}

// PremiumPeriod is a part of Account structure
type PremiumPeriod struct {
	Start  int32 `json:"start"`
	Finish int32 `json:"finish"`
}

// ServiceQuery is a parsed http query
type ServiceQuery struct {
	Params map[string]map[string]string
	Limit  int32
	Order  int32
}

// NewServiceQuery inits ServiceQuery struct
func NewServiceQuery() *ServiceQuery {
	var query ServiceQuery
	query.Params = make(map[string]map[string]string)
	return &query
}

// Accounts collection
type Accounts struct {
	Accounts []Account `json:"accounts"`
}

// Account is a general service structure
type Account struct {
	ID        int32         `json:"id"`
	Email     string        `json:"email"`
	Fname     string        `json:"fname"`
	Sname     string        `json:"sname"`
	Phone     string        `json:"phone"`
	Birth     int32         `json:"birth"`
	Sex       string        `json:"sex"`
	Country   string        `json:"country"`
	City      string        `json:"city"`
	Joined    int32         `json:"joined"`
	Status    string        `json:"status"`
	Interests []string      `json:"interests"`
	Premium   PremiumPeriod `json:"premium"`
	Likes     []Like        `json:"likes"`
}

func (account *Account) GetFieldValue(name string) string {
	r := reflect.ValueOf(account)
	f := reflect.Indirect(r).FieldByName(name)
	return f.String()
}

func (account *Account) GetStringFieldValue(name string) string {
	r := reflect.ValueOf(account)
	f := reflect.Indirect(r).FieldByName(name)
	return f.String()
}

func (account *Account) GetInt32FieldValue(name string) int32 {
	r := reflect.ValueOf(account)
	f := reflect.Indirect(r).FieldByName(name)
	return int32(f.Int())
}
