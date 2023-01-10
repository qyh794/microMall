package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// 获取 mysql 的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}

//type MysqlConfig struct {
//	Host     string `json:"host"`
//	User     string `json:"user"`
//	Pwd      string `json:"pwd"`
//	Database string `json:"database"`
//	Port     string `json:"port"`
//}
//
//// get mysql config
////func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
////	mysqlConfig := &MysqlConfig{}
////	config.Get(path...).Scan(mysqlConfig)
////	fmt.Println("zheli" + mysqlConfig.Password)
////	return mysqlConfig
////}
//
//func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
//	mysqlConfig := &MysqlConfig{}
//	// big problem
//	value := config.Get(path...)
//	bytes, err := json.Marshal(value)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("shidahwiodhawd" + string(bytes))
//	value.Scan(mysqlConfig)
//	return mysqlConfig
//}
