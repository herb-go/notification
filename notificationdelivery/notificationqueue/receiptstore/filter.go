package receiptstore

import (
	"github.com/herb-go/notification"
	"github.com/herb-go/notification/notificationdelivery/notificationqueue"
)

type Filter interface {
	//FilterReceipt filter receipt with given context
	//Return if Receipt is valid
	FilterReceipt(r *notificationqueue.Receipt, ctx *notification.ConditionContext) (bool, error)
	//ApplyCondition apply search condition to filter
	//ErrConditionNotSupported should be returned if condition keyword is not supported
	ApplyCondition(cond *notification.Condition) error
}

//ApplyToFilter apply condiitons to filter.
func ApplyToFilter(f Filter, conds []*notification.Condition) error {
	for k := range conds {
		err := f.ApplyCondition(conds[k])
		if err != nil {
			return err
		}
	}
	return nil
}
