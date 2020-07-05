package host

import (
	"context"
	"time"

	"github.com/filanov/bm-inventory/internal/hardware"
	"github.com/filanov/bm-inventory/models"
	logutil "github.com/filanov/bm-inventory/pkg/log"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func NewDisconnectedState(log logrus.FieldLogger, db *gorm.DB, hwValidator hardware.Validator) *disconnectedState {
	return &disconnectedState{
		baseState: baseState{
			log: log,
			db:  db,
		},
		hwValidator: hwValidator,
	}
}

type disconnectedState struct {
	baseState
	hwValidator hardware.Validator
}

func (d *disconnectedState) UpdateHwInfo(ctx context.Context, h *models.Host, hwInfo string) (*UpdateReply, error) {
	h.HardwareInfo = hwInfo
	return updateHwInfo(logutil.FromContext(ctx, d.log), d.hwValidator, h, d.db)
}

func (d *disconnectedState) UpdateInventory(ctx context.Context, h *models.Host, inventory string) (*UpdateReply, error) {
	h.Inventory = inventory
	return updateInventory(logutil.FromContext(ctx, d.log), d.hwValidator, h, d.db)
}

func (d *disconnectedState) RefreshStatus(ctx context.Context, h *models.Host, db *gorm.DB) (*UpdateReply, error) {
	log := logutil.FromContext(ctx, d.log)
	if time.Since(time.Time(h.CheckedInAt)) < 3*time.Minute {
		return updateState(log, HostStatusDiscovering, statusInfoDiscovering, h, d.db)
	}
	// Stay in the same state
	return defaultReply(h)
}

func (d *disconnectedState) EnableHost(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	// State in the same state
	return &UpdateReply{
		State:     HostStatusDisconnected,
		IsChanged: false,
	}, nil
}

func (d *disconnectedState) DisableHost(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	return updateState(logutil.FromContext(ctx, d.log), HostStatusDisabled, statusInfoDisabled, h, d.db)
}
