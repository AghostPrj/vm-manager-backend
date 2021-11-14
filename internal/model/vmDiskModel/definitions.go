/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 12:40
 * @Desc:
 */

package vmDiskModel

type VmDisk struct {
	Id       uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id"`
	VmId     uint64 `gorm:"Column:vm_id;NOT NULL"`
	DiskPath string `gorm:"type:varchar(1024);Column:disk_path;NOT NULL"`
}
