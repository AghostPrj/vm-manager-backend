/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 12:39
 * @Desc:
 */

package vmListModel

import (
	"time"
)

type VmList struct {
	Id          uint64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id"`
	Name        string     `gorm:"type:varchar(128);Column:name;NOT NULL"`
	Cpu         int32      `gorm:"Column:cpu;NOT NULL"`
	Mem         int32      `gorm:"Column:mem;NOT NULL"`
	AutoStartup int32      `gorm:"Column:auto_startup;NOT NULL"`
	Status      int32      `gorm:"Column:status;NOT NULL"`
	CreatedAt   *time.Time `gorm:"Column:create_time"`
	UpdatedAt   *time.Time `gorm:"Column:update_time"`
}
