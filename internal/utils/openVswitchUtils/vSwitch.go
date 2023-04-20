/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2023/4/18 17:19
 * @Desc:
 */

package openVswitchUtils

import (
	"errors"
	"github.com/AghostPrj/vm-manager-backend/internal/constData/errorCode"
	"github.com/digitalocean/go-openvswitch/ovs"
	"os/exec"
	"regexp"
	"strings"
)

const (
	bridgeName         = "br-aghost-vm-manager"
	portPrefix         = "aghost-vm-port-"
	physicalPortPrefix = portPrefix + "phy-"
	virtualPortPrefix  = portPrefix + "virtual-"

	virtualPortPathPrefix = "/run/openvswitch/vm-port-"

	virtualPortNameRegularExpression = "aghost-vm-port-virtual-([a-f0-9]{2}_){5}[a-f0-9]{2}"
)

var (
	ovsClient                *ovs.Client
	virtualPortNameRegexp, _ = regexp.Compile(virtualPortNameRegularExpression)
)

func CheckOvsExists() bool {
	_, err := exec.Command("which", "ovs-vsctl").CombinedOutput()
	if err != nil {
		return false
	}

	return checkOvsIsDpdk()
}

func checkOvsIsDpdk() bool {
	output, err := exec.Command("which", "update-alternatives").CombinedOutput()
	if err != nil {
		return false
	}

	updateAlternativesPath := strings.TrimSpace(string(output))

	output, err = exec.Command("bash", "-c",
		updateAlternativesPath+" --query ovs-vswitchd | grep Value | awk '{print $2}'").CombinedOutput()
	if err != nil {
		return false
	}

	ovsLinkPath := strings.TrimSpace(string(output))

	if !strings.HasSuffix(ovsLinkPath, "-dpdk") {
		return false
	}

	output, err = exec.Command("bash", "-c",
		"ovs-vsctl get Open_vSwitch . other_config").CombinedOutput()
	if err != nil {
		return false
	}

	ovsDpdkConfStr := strings.TrimSpace(string(output))

	contains := strings.Contains(ovsDpdkConfStr, "dpdk-init=\"true\"")

	return contains

}

func Init() bool {
	exists := CheckOvsExists()
	if exists {
		ovsClient = ovs.New(
			ovs.Sudo(),
		)

	}

	return exists
}

func CheckBridgeExists() (bool, error) {
	list, err := GetBridgeList()
	if err != nil {
		return false, err
	}

	for _, name := range *list {
		if strings.TrimSpace(name) == bridgeName {
			return true, nil
		}
	}

	return false, nil
}

func CreateBridge() error {
	if ovsClient == nil {
		return errors.New(errorCode.NoOvsSupportError)
	}
	cmdStr := "sudo ovs-vsctl add-br " + bridgeName + " -- set bridge " + bridgeName + " datapath_type=netdev"
	_, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()

	return err
}

func GetBridgeList() (*[]string, error) {
	if ovsClient == nil {
		return nil, errors.New(errorCode.NoOvsSupportError)
	}

	bridges, err := ovsClient.VSwitch.ListBridges()
	if err != nil {
		return nil, err
	}

	return &bridges, nil
}

func getPhysicalPortName(deviceDomain string) string {
	return physicalPortPrefix +
		strings.ReplaceAll(strings.ReplaceAll(deviceDomain, ":", "_"), ".", "_")
}

func CreatePhysicalPort(deviceDomain string) error {
	if ovsClient == nil {
		return errors.New(errorCode.NoOvsSupportError)
	}

	deviceDomain = strings.TrimSpace(deviceDomain)

	portName := getPhysicalPortName(deviceDomain)

	cmdStr := "sudo ovs-vsctl add-port " + bridgeName + " " +
		portName + " -- set Interface " + portName +
		" type=dpdk  \"options:dpdk-devargs=" + deviceDomain + "\""
	_, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()

	return err
}

func CheckPhysicalPortExists(deviceDomain string) (bool, error) {
	if ovsClient == nil {
		return false, errors.New(errorCode.NoOvsSupportError)
	}

	deviceDomain = strings.TrimSpace(deviceDomain)

	portName := getPhysicalPortName(deviceDomain)

	list, err := GetPortList(bridgeName)
	if err != nil {
		return false, err
	}

	for _, p := range *list {
		if strings.TrimSpace(p) == portName {
			return true, nil
		}
	}

	return false, nil
}

func DelPhysicalPort(deviceDomain string) error {
	exists, err := CheckPhysicalPortExists(deviceDomain)
	if err != nil {
		return err
	}

	if !exists {
		return nil
	}
	deviceDomain = strings.TrimSpace(deviceDomain)
	portName := getPhysicalPortName(deviceDomain)

	cmdStr := "sudo ovs-vsctl del-port " + bridgeName + " " + portName
	_, err = exec.Command("bash", "-c", cmdStr).CombinedOutput()

	return err

}

func GetPortList(bridge string) (*[]string, error) {
	if ovsClient == nil {
		return nil, errors.New(errorCode.NoOvsSupportError)
	}

	bridges, err := ovsClient.VSwitch.ListPorts(bridge)
	if err != nil {
		return nil, err
	}

	return &bridges, nil
}

func getVirtualPortName(macAddress string) string {
	return virtualPortPrefix + strings.ReplaceAll(macAddress, ":", "_")
}
func CheckVirtualPortExists(macAddress string) (bool, error) {
	if ovsClient == nil {
		return false, errors.New(errorCode.NoOvsSupportError)
	}

	macAddress = strings.TrimSpace(macAddress)

	portName := getVirtualPortName(macAddress)

	list, err := GetPortList(bridgeName)
	if err != nil {
		return false, err
	}

	for _, p := range *list {
		if strings.TrimSpace(p) == portName {
			return true, nil
		}
	}

	return false, nil
}

func CreateVirtualPort(macAddress string) error {
	if ovsClient == nil {
		return errors.New(errorCode.NoOvsSupportError)
	}

	macAddress = strings.TrimSpace(macAddress)

	portName := getVirtualPortName(macAddress)

	cmdStr := "sudo ovs-vsctl add-port " + bridgeName + " " +
		portName + " -- set Interface " + portName +
		" type=dpdkvhostuserclient  \"options:vhost-server-path=" + GetVirtualPortPath(macAddress) + "\""
	_, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()

	return err
}

func GetVirtualPortPath(macAddress string) string {
	return virtualPortPathPrefix + strings.ReplaceAll(macAddress, ":", "_")
}

func DelVirtualPort(macAddress string) error {
	exists, err := CheckVirtualPortExists(macAddress)
	if err != nil {
		return err
	}

	if !exists {
		return nil
	}
	macAddress = strings.TrimSpace(macAddress)
	portName := getVirtualPortName(macAddress)

	cmdStr := "sudo ovs-vsctl del-port " + bridgeName + " " + portName
	_, err = exec.Command("bash", "-c", cmdStr).CombinedOutput()

	return err
}

func GetVirtualPortList() (*[]string, error) {
	if ovsClient == nil {
		return nil, errors.New(errorCode.NoOvsSupportError)
	}

	list, err := GetPortList(bridgeName)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, name := range *list {
		name = strings.TrimSpace(name)
		if virtualPortNameRegexp.MatchString(name) {
			result = append(result, name)
		}
	}

	return &result, nil
}
