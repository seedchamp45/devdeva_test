package model

type Data struct {
	ActivePower int `json:"active_power"`
	InputPower  int `json:"input_power"`
}

var allDataList []Data

func SetData(dataList []Data) []Data {
	allDataList = dataList
	return dataList
}

func GetData() []Data {
	return allDataList
}
