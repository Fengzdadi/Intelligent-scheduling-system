package model

type Emplike struct {
	Etype string `xorm:"varchar(16)" json:"Etype"`
	Eid   string `xorm:"varchar(4)" json:"Eid"`
	Elike string `xorm:"varchar(30)" json:"Elike"`
}

type Employee struct {
	Eid    string `xorm:"varchar(4)" json:"Eid"`
	Ename  string `xorm:"varchar(8)" json:"Ename"`
	Eemail string `xorm:"varchar(16)" json:"Eemail"`
	Epos   string `xorm:"varchar(8)" json:"Epos"`
	Shid   string `xorm:"varchar(4)" json:"Shid"`
}

type PaiBan struct {
	Planid string `xorm:"varchar(4)" json:"Planid"`
	Eid    string `xorm:"varchar(4)" json:"Eid"`
	Epos   string `xorm:"varchar(8)" json:"Epos"`
	W1     string `xorm:"varchar(60)" json:"W1"`
	W2     string `xorm:"varchar(60)" json:"W2"`
	W3     string `xorm:"varchar(60)" json:"W3"`
	W4     string `xorm:"varchar(60)" json:"W4"`
	W5     string `xorm:"varchar(60)" json:"W5"`
	W6     string `xorm:"varchar(60)" json:"W6"`
	W7     string `xorm:"varchar(60)" json:"W7"`
}

type Shop struct {
	Shid      string `xorm:"varchar(4)" json:"Shid"`
	Shname    string `xorm:"varchar(12)" json:"Shname"`
	Shaddress string `xorm:"varchar(20)" json:"Shaddress"`
	Shsize    string `xorm:"float" json:"Shsize"`
}

type Shop_rule struct {
	Ruletype string `xorm:"char(8)" json:"Ruletype"`
	Shid     string `xorm:"varchar(4)" json:"Shid"`
	Shrule   string `xorm:"varchar(50)" json:"Shrule"`
}

type User_manager struct {
	Userid   string `xorm:"varchar(16)" json:"Userid"`
	Eid      string `xorm:"varchar(4)" json:"Eid"`
	Epos     string `xorm:"varchar(8)" json:"Epos"`
	Username string `xorm:"varchar(16)" json:"Username"`
	Pwd      int    `xorm:"int" json:"Pwd"`
}
