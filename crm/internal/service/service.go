package service

import (
	"log"
	"sync"

	"github.com/crusttech/crust/internal/store"
)

type (
	db interface {
		Transaction(callback func() error) error
	}
)

var (
	o                   sync.Once
	DefaultRecord       RecordService
	DefaultModule       ModuleService
	DefaultTrigger      TriggerService
	DefaultChart        ChartService
	DefaultPage         PageService
	DefaultNotification NotificationService
	DefaultPermissions  PermissionsService
	DefaultAttachment   AttachmentService
)

func init() {
	o.Do(func() {
		fs, err := store.New("var/store")
		if err != nil {
			log.Fatalf("Failed to initialize store: %v", err)
		}

		DefaultPermissions = Permissions()
		DefaultRecord = Record()
		DefaultModule = Module()
		DefaultTrigger = Trigger()
		DefaultPage = Page()
		DefaultChart = Chart()
		DefaultNotification = Notification()
		DefaultAttachment = Attachment(fs)
	})
}
