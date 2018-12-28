package importer

import (
	"encoding/json"
	"highloadcup/app/structs"
	"io/ioutil"
	"os"
	"sort"
)

type ById []structs.Account

func (a ById) Len() int           { return len(a) }
func (a ById) Less(i, j int) bool { return a[i].ID < a[j].ID }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func ImportAccounts() []structs.Account {
	jsonFile, err := os.Open("/tmp/accounts.json")

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	accounts := structs.Accounts{}
	json.Unmarshal(bytes, &accounts)

	sort.Sort(ById(accounts.Accounts))

	return accounts.Accounts
}
