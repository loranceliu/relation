package vo

type UserInfoVO struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	IsAdmin  bool   `json:"isAdmin"`
	Exp      int    `json:"exp"`
}

type UserMineInfoVO struct {
	UserId    uint    `json:"userId"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Time      int     `json:"time"`
	IncomeNum int     `json:"incomeNum"`
	ExpendNum int     `json:"expendNum"`
	Revenue   float32 `json:"revenue"`
}
