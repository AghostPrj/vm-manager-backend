/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2023/4/18 15:27
 * @Desc:
 */

package dpdkUtils

import (
	"errors"
	"github.com/AghostPrj/vm-manager-backend/internal/constData/errorCode"
	"github.com/AghostPrj/vm-manager-backend/internal/object/exchange/nicPayload"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"regexp"
	"strings"
)

const (
	DomainRegularExpression      = "\\d{4}\\:\\d{2}\\:\\d{2}\\.\\d+?"
	DescriptionRegularExpression = "'.*?'"
)

var (
	dpdkPath = ""

	domainRegexp, _      = regexp.Compile(DomainRegularExpression)
	descriptionRegexp, _ = regexp.Compile(DescriptionRegularExpression)
)

func CheckDpdkDevbindExists() bool {
	output, err := exec.Command("which", "dpdk-devbind.py").CombinedOutput()
	if err != nil {
		return false
	}

	dpdkPath = strings.TrimSpace(string(output))
	return true
}

func GetDpdkNetworkInterfaces() (*[]nicPayload.NetworkInterfaceController, error) {
	if len(dpdkPath) < 1 {
		return nil, errors.New(errorCode.NoDpdkSupportError)
	}

	output, err := exec.Command(dpdkPath, "--status-dev", "net").CombinedOutput()
	if err != nil {
		return nil, err
	}

	outputStr := string(output)
	lines := strings.Split(outputStr, "\n")

	filteredLines := make([]string, 0)

	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "==", "")
		lines[i] = strings.TrimSpace(lines[i])

		if lines[i] == "=" {
			lines[i] = ""
		}

		if len(lines[i]) > 0 {
			filteredLines = append(filteredLines, lines[i])
		}
	}

	typeStr := ""

	result := make([]nicPayload.NetworkInterfaceController, 0)

	for _, line := range filteredLines {

		if line == "Network devices using DPDK-compatible driver" {
			typeStr = nicPayload.NicTypeDpdk
			continue
		} else if line == "Network devices using kernel driver" ||
			line == "Other Network devices" {
			typeStr = nicPayload.NicTypeSystem
			continue
		}

		if len(typeStr) < 1 {
			continue
		}

		domain := domainRegexp.FindString(line)
		if len(domain) < 1 {
			continue
		}

		line = strings.ReplaceAll(line, domain, "")

		description := descriptionRegexp.FindString(line)
		if len(description) > 0 {
			line = strings.ReplaceAll(line, description, "")
		}

		line = strings.ReplaceAll(line, "*Active*", "")

		line = strings.TrimSpace(line)

		tmpPayload := nicPayload.NetworkInterfaceController{
			Type:        typeStr,
			Domain:      domain,
			Description: description,
		}

		params := strings.Split(line, " ")
		for _, param := range params {
			param = strings.TrimSpace(param)

			splitParam := strings.Split(param, "=")

			if len(splitParam) == 2 {
				switch strings.TrimSpace(splitParam[0]) {
				case "if":
					tmpPayload.InterfaceName = strings.TrimSpace(splitParam[1])
					break
				case "drv":
					tmpPayload.Driver = strings.TrimSpace(splitParam[1])
					break
				case "unused":
					drivers := strings.Split(splitParam[1], ",")
					tmpPayload.AvailableDriver = make([]string, 0)
					for _, driver := range drivers {
						tmpPayload.AvailableDriver = append(tmpPayload.AvailableDriver, strings.TrimSpace(driver))
					}
					break
				default:
					continue
				}
			}
		}

		result = append(result, tmpPayload)

	}

	log.WithFields(log.Fields{
		"op":   "GetDpdkNetworkInterfaces",
		"data": result,
	}).Debug()

	return &result, err
}
func BindDriver(domain, driver string) error {
	if len(dpdkPath) < 1 {
		return errors.New(errorCode.NoDpdkSupportError)
	}

	domain = strings.TrimSpace(domain)
	driver = strings.TrimSpace(driver)

	if len(domain) < 6 || len(driver) < 1 {
		return errors.New(errorCode.DataError)
	}

	interfaces, err := GetDpdkNetworkInterfaces()
	if err != nil {
		return err
	}

	if len(*interfaces) < 1 {
		return errors.New(errorCode.DeviceNotFoundError)
	}

	var nic *nicPayload.NetworkInterfaceController
	nic = nil

	for _, d := range *interfaces {
		if d.Domain == domain {
			nic = &d
			break
		}
	}

	if nic == nil || nic.Domain != domain {
		return errors.New(errorCode.DeviceNotFoundError)
	}

	driverFound := false

	if nic.Driver == driver {
		return nil
	}

	for _, d := range nic.AvailableDriver {
		if d == driver {
			driverFound = true
			break
		}
	}

	if !driverFound {
		return errors.New(errorCode.DeviceDriverNotSupportError)
	}

	output, err := exec.Command(dpdkPath, "-b", driver, domain).CombinedOutput()
	log.WithFields(log.Fields{
		"op":     "BindDpdkNicDriver",
		"driver": driver,
		"domain": domain,
		"err":    err,
		"output": string(output),
	}).Debug()

	if err != nil || strings.HasPrefix(strings.TrimSpace(string(output)), "Error") {
		return errors.New(errorCode.DeviceDriverNotSupportError)
	} else {
		return nil
	}

}
