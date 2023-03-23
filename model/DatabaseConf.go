package model

import _ "github.com/jinzhu/gorm"

type Emplike struct {
	Etype string `grom:"varchar(16)" json:"Etype"`
	Eid   string `grom:"varchar(4)" json:"Eid"`
	Elike string `grom:"varchar(30)" json:"Elike"`
}

type Employee struct {
	Eid    string `grom:"varchar(4)" json:"Eid"`
	Ename  string `grom:"varchar(8)" json:"Ename"`
	Eemail string `grom:"varchar(16)" json:"Eemail"`
	Epos   string `grom:"varchar(8)" json:"Epos"`
	Shid   string `grom:"varchar(4)" json:"Shid"`
}

type PaiBan struct {
	Planid string `grom:"varchar(4)" json:"Planid"`
	Eid    string `grom:"varchar(4)" json:"Eid"`
	Epos   string `grom:"varchar(8)" json:"Epos"`
	W1     string `grom:"varchar(60)" json:"W1"`
	W2     string `grom:"varchar(60)" json:"W2"`
	W3     string `grom:"varchar(60)" json:"W3"`
	W4     string `grom:"varchar(60)" json:"W4"`
	W5     string `grom:"varchar(60)" json:"W5"`
	W6     string `grom:"varchar(60)" json:"W6"`
	W7     string `grom:"varchar(60)" json:"W7"`
}

type Shop struct {
	Shid      string `grom:"varchar(4)" json:"Shid"`
	Shname    string `grom:"varchar(12)" json:"Shname"`
	Shaddress string `grom:"varchar(20)" json:"Shaddress"`
	Shsize    string `grom:"float" json:"Shsize"`
}

type Shop_rule struct {
	Ruletype string `grom:"char(8)" json:"Ruletype"`
	Shid     string `grom:"varchar(4)" json:"Shid"`
	Shrule   string `grom:"varchar(50)" json:"Shrule"`
}

type User_manager struct {
	Userid   string `grom:"varchar(16)" json:"Userid"`
	Eid      string `grom:"varchar(4)" json:"Eid"`
	Epos     string `grom:"varchar(8)" json:"Epos"`
	Username string `grom:"varchar(16)" json:"Username"`
	Pwd      int    `grom:"int" json:"Pwd"`
}
