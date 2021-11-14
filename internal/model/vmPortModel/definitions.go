/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 12:40
 * @Desc:
 */

package vmPortModel

type VmPort struct {
	Id   uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id"`
	VmId uint64 `gorm:"Column:vm_id;NOT NULL"`
	Type int32  `gorm:"Column:type;NOT NULL"`
	Port int32  `gorm:"Column:port;NOT NULL"`
}
