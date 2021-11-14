/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 12:40
 * @Desc:
 */

package vmMacModel

type VmMac struct {
	Id   uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id"`
	VmId uint64 `gorm:"Column:vm_id;NOT NULL"`
	Mac  string `gorm:"type:char(17);Column:mac;NOT NULL"`
}
