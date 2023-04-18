/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2023/4/18 14:18
 * @Desc:
 */

package networkService

import (
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/AghostPrj/vm-manager-backend/internal/object/exchange/nicPayload"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/dpdkUtils"
	"net"
)

func GetSystemNic() (*[]nicPayload.NetworkInterfaceController, error) {

	if global.HavingDpdk {
		return dpdkUtils.GetDpdkNetworkInterfaces()
	} else {
		return GetNicListBySystem()
	}

}

func GetNicListBySystem() (*[]nicPayload.NetworkInterfaceController, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	result := make([]nicPayload.NetworkInterfaceController, 0)

	for _, i := range interfaces {
		result = append(result, nicPayload.NetworkInterfaceController{
			Type:          nicPayload.NicTypeSystem,
			InterfaceName: i.Name,
		})
	}

	return &result, nil
}

func BindNicDriver(domain, driver string) error {
	return dpdkUtils.BindDriver(domain, driver)
}
