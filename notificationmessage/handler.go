package notificationmessage

import "fmt"

type RecordHandler func(*Record)

var NopRecordHandler = func(r *Record) {
	fmt.Println(fmt.Sprintf("message record created: %s", r.String()))
}
