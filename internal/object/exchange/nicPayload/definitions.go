/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2023/4/18 15:16
 * @Desc:
 */

package nicPayload

const (
	NicTypeDpdk   = "dpdk"
	NicTypeSystem = "system"
)

type NetworkInterfaceController struct {
	Type            string   `json:"type"`
	Domain          string   `json:"domain"`
	Description     string   `json:"description"`
	InterfaceName   string   `json:"interface_name"`
	Driver          string   `json:"driver"`
	AvailableDriver []string `json:"available_driver"`
}
