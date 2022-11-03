package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	eaiConf := SQSConf{
		Queue: "retekqueue",
	}

	conf := SQSs{
		RetekCreate: eaiConf,
	}

	rValue := reflect.ValueOf(conf)
	rType := reflect.TypeOf(conf)

	// Traverse through all the fields of a struct.
	if rType.Kind() == reflect.Struct {
		for i := 0; i < rType.NumField(); i++ {
			if strings.EqualFold(rType.Field(i).Tag.Get("mapstructure"), "retek_create") {
				prop := rValue.Field(i)
				cfg := prop.Interface().(SQSConf)
				fmt.Println(cfg.Queue)
			}
		}
		// fieldValue := rValue.FieldByName("eai")

		// if !fieldValue.IsValid() {
		// 	fmt.Println("invalid field")
		// 	os.Exit(0)
		// }
		// x := fieldValue.Interface().(SQSConf)
		// fmt.Println(x.Queue)
	}

}

type SQSs struct {
	Env              string  `mapstructure:"env"`
	Canonical        SQSConf `mapstructure:"canonical"`
	StoreOrderEvents SQSConf `mapstructure:"store_order_events"`
	OrderArchiver    SQSConf `mapstructure:"order_archiver"`
	RequestsArchiver SQSConf `mapstructure:"requests_archiver"`
	OrderIntents     SQSConf `mapstructure:"order_intents"`
	Mirakl           SQSConf `mapstructure:"mirakl"`
	EAI              SQSConf `mapstructure:"eai"`
	Charon           SQSConf `mapstructure:"charon"`
	Sofia            SQSConf `mapstructure:"sofia"`
	Cancellation     SQSConf `mapstructure:"cancellation"`
	Availability     SQSConf `mapstructure:"availability"`
	Dicktracy        SQSConf `mapstructure:"dicktracy"`
	RetekCreate      SQSConf `mapstructure:"retek_create"`
	RetekUpdate      SQSConf `mapstructure:"retek_update"`
	Activation       SQSConf `mapstructure:"activation"`
	ScheduledJobs    SQSConf `mapstructure:"scheduled_jobs"`
	JobScheduler     SQSConf `mapstructure:"job_scheduler"`
	ICP              SQSConf `mapstructure:"icp"`
}

type SQSConf struct {
	Region              string `mapstructure:"region"`
	Queue               string `mapstructure:"queue"`
	Endpoint            string `mapstructure:"endpoint"`
	MaxNumberOfMessages int    `mapstructure:"max_number_of_messages"`
	VisibilityTimeout   int    `mapstructure:"visibility_timeout"`
	WaitTimeSeconds     int    `mapstructure:"wait_time_seconds"`
	Retries             int    `mapstructure:"retries"`
}
