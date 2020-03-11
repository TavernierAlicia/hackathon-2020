package main

type EventForm struct {
	Mail          string
	Name          string
	Surname       string
	Birth         string
	ObjectName0   string
	ObjectType0   string
	ObjectState0  string
	ObjectSize0   int
	ObjectWeight0 int
	ObjectName1   string
	ObjectType1   string
	ObjectState1  string
	ObjectSize1   int
	ObjectWeight1 int
	ObjectName2   string
	ObjectType2   string
	ObjectState2  string
	ObjectSize2   int
	ObjectWeight2 int
	EventId       int
}

type ItemsForm struct {
	Mail       string
	Name       string
	Surname    string
	Birth      string
	PickupDate string
	ItemId     int
}

type WorkshopForm struct {
	Mail       string
	Name       string
	Surname    string
	Birth      string
	WorkshopId int
}

/*
type Events struct {
	TeamName       string  `db:"name"`
	NbCallswi      int     `db:"inbcallsw"`
	Idurationwi    float64 `db:"iiduration_minutesw"`
	Odurationwi    float64 `db:"ioduration_minutesw"`
	TMCwi          float64 `db:"iweekly_TMC"`
	Ocostwi        float64 `db:"iocost_eurosw"`
	Creditwi       float64 `db:"icreditw"`
	CreditLimitwi  float64 `db:"icredit_limitw"`
	Userswi        int64   `db:"iusersw"`
	Numberswi      string  `db:"inumbersw"`
	BilledReccurwi float64 `db:"ibilled_reccurw"`
	CreationDatewi string  `db:"iCREATION_DATEw"`
}

type Event struct {
	TeamName       string  `db:"name"`
	NbCallswi      int     `db:"inbcallsw"`
	Idurationwi    float64 `db:"iiduration_minutesw"`
	Odurationwi    float64 `db:"ioduration_minutesw"`
	TMCwi          float64 `db:"iweekly_TMC"`
	Ocostwi        float64 `db:"iocost_eurosw"`
	Creditwi       float64 `db:"icreditw"`
	CreditLimitwi  float64 `db:"icredit_limitw"`
	Userswi        int64   `db:"iusersw"`
	Numberswi      string  `db:"inumbersw"`
	BilledReccurwi float64 `db:"ibilled_reccurw"`
	CreationDatewi string  `db:"iCREATION_DATEw"`
}

type Items struct {
	TeamName       string  `db:"name"`
	NbCallswi      int     `db:"inbcallsw"`
	Idurationwi    float64 `db:"iiduration_minutesw"`
	Odurationwi    float64 `db:"ioduration_minutesw"`
	TMCwi          float64 `db:"iweekly_TMC"`
	Ocostwi        float64 `db:"iocost_eurosw"`
	Creditwi       float64 `db:"icreditw"`
	CreditLimitwi  float64 `db:"icredit_limitw"`
	Userswi        int64   `db:"iusersw"`
	Numberswi      string  `db:"inumbersw"`
	BilledReccurwi float64 `db:"ibilled_reccurw"`
	CreationDatewi string  `db:"iCREATION_DATEw"`
}

type Item struct {
	TeamName       string  `db:"name"`
	NbCallswi      int     `db:"inbcallsw"`
	Idurationwi    float64 `db:"iiduration_minutesw"`
	Odurationwi    float64 `db:"ioduration_minutesw"`
	TMCwi          float64 `db:"iweekly_TMC"`
	Ocostwi        float64 `db:"iocost_eurosw"`
	Creditwi       float64 `db:"icreditw"`
	CreditLimitwi  float64 `db:"icredit_limitw"`
	Userswi        int64   `db:"iusersw"`
	Numberswi      string  `db:"inumbersw"`
	BilledReccurwi float64 `db:"ibilled_reccurw"`
	CreationDatewi string  `db:"iCREATION_DATEw"`
}

type Workshops struct {
	TeamName       string  `db:"name"`
	NbCallswi      int     `db:"inbcallsw"`
	Idurationwi    float64 `db:"iiduration_minutesw"`
	Odurationwi    float64 `db:"ioduration_minutesw"`
	TMCwi          float64 `db:"iweekly_TMC"`
	Ocostwi        float64 `db:"iocost_eurosw"`
	Creditwi       float64 `db:"icreditw"`
	CreditLimitwi  float64 `db:"icredit_limitw"`
	Userswi        int64   `db:"iusersw"`
	Numberswi      string  `db:"inumbersw"`
	BilledReccurwi float64 `db:"ibilled_reccurw"`
	CreationDatewi string  `db:"iCREATION_DATEw"`
}
*/
