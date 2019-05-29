package wrapper

import "net/http"
import "io/ioutil"
import "encoding/json"
import "strconv"

// User : r/memeeconomybot user
type User struct {
	ID                   int
	Name                 string
	Balance              int
	CompletedInvestments int `json:"completed"`
	TimesBroke           int
	Firm                 int
	FirmRole             string `json:"firm_role"`
	Networth             int
	Ranking              int `json:"rank"`
}

// Firm : meme economy firm
type Firm struct {
	ID         int
	Name       string
	Balance    int
	Size       int
	Execs      int
	Tax        int
	Rank       int
	Private    bool
	LastPayout int `json:"last_payout"`
}
// Investment : user investment what else do u want from me golint
type Investment struct {
	ID            int
	Post          string
	Upvotes       int
	comment       string
	Name          string
	CoinsInvested int `json:"amount"`
	Time          int
	Completed     bool `json:"done"`
	response      string
	FinalUpvotes  int `json:"final_upvotes"`
	Success       bool
	Profit        int
}

// GetUser : get user info by reddit name
func GetUser(name string) (info User) {
	url := "https://meme.market/api/investor/" + name
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("referer", "https://meme.market/user.html?account="+name)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var userinfo User
	json.Unmarshal(body, &userinfo)
	return userinfo
}

// GetFirm : get's firm info from user
func (User *User) GetFirm() (firm Firm) {
	url := "https://meme.market/api/firm/" + strconv.Itoa(User.Firm)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("referer", "https://meme.market/user.html?account="+User.Name)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &firm)
	return firm
}

// GetInvestments : Get investments from user
func (User *User) GetInvestments() (investments []Investment) {
	url := "https://meme.market/api/investor/" + User.Name + "/investments?"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("referer", "https://meme.market/user.html?account="+User.Name)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &investments)
	return investments
}
