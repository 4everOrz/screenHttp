package models

type ScreenData struct {
	S_pv_info          PV_info         `json:"pv_info,omitempty"`
	S_environment_info Enviroment_info `json:"environment_info,omitempty"`
	S_device_info      Device_info     `json:"device_info,omitempty"`
	S_sence_info       Sence_info      `json:"sence_info,omitempty"`
	S_leader_info      Leader_info     `json:"leader_info,omitempty"`
	S_visitor_info     Visitor_info    `json:"visitor_info,omitempty"`
	S_Light_info       Light_info      `json:"light_info,omitempty"`
}
type PV_info struct {
	//光伏
	P01 string
	P02 string
	P03 string
	P04 string
	P05 string
	P06 string
	P07 string
	P08 string
	P09 string
	P10 string
	P11 string
	P12 string
	P13 string
	P14 string
	P15 string
	P16 string
	P17 string
	P18 string
	P19 string
	P20 string
	P21 string
	P22 string
	P23 string
	P24 string
	P25 string
	P26 string
	P27 string
}
type Enviroment_info struct {
	//环境
	E01 string
	E02 string
	E03 string
	E04 string
	E05 string
	E06 string
	E07 string
	E08 string
	E09 string
	E10 string
	E11 string
	E12 string
	E13 string
	E14 string
	E15 string
	E16 string
	E17 string
	E18 string
	E19 string
	E20 string
	E21 string
	E22 string
}
type Device_info struct {
	//设备
	K01 string
	K02 string
	K03 string
	K04 string
	K05 string
	K06 string
	K07 string
	K08 string
	K09 string
	K10 string
	K11 string
	K12 string
	K13 string
	K14 string
	K15 string
	K16 string
	K17 string
	// this.device_info.K18 = "0";
	// this.device_info.K19 = "1";
	// this.device_info.K20 = "3";

	C01 string
	C02 string
	C03 string
	C04 string
	C05 string
	C06 string
	C07 string
	C08 string
	C09 string
	C10 string
	C11 string
	C12 string
	C13 string
	C14 string
	C15 string
	C16 string
	C17 string
	C18 string
	C19 string
	C20 string
	C21 string
	C22 string
	C23 string
	C24 string
	C25 string
	C26 string

	H01 string
	H02 string
	H03 string
	H04 string
	H05 string
	H06 string
	H07 string
	H08 string
	H09 string
	H10 string
	H11 string

	//  this.device_info.C27 = "0";
	/*
				 * “H01”:”0”// 0三天以前 1前天 2昨天 3今天 上一次洗衣
		“H02”:”0”// 0三天以前 1前天 2昨天 3今天 上一次洗碗
		“H03”:”30” //自来水DTS值
		“H04”:”10” //净水值DTS
		“H05”:”55”  //热水温度
		“H06”:”香菇焖鸡” //推荐菜谱
		“H07”:”新鲜” //冰箱
		“H08”:”0” //0撤防    1 布防
		“H09”:”0” //0正常 1报警 烟雾传感器
		“H10”:”0” //0正常1报警水浸传感器
		“H11”:”0” //0正常1报警紧急按钮

	*/
}
type Visitor_info struct {
	V01 string
	V02 string
	V03 string
	V04 string
	V05 string
	V06 string
}
type Leader_info struct {
	L01 string
	L02 string
	L03 string
	L04 string
	L05 string
}
type Sence_info struct {

	//情景模式 gohome,leaveHome,wakeUp,movie,sport,sleep
	S01 string
	S02 string
	S03 string
	S04 string
}
type Light_info struct {
	G01 string
	G02 string
	G03 string
	G04 string
	G05 string
	G06 string
	G07 string
	G08 string
}
