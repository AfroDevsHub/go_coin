package constants

import "fmt"

type DevicePermission int

const (
	CAMERA DevicePermission = iota
	STORAGE
	CONTACTS
)

var devicepermissions = map[DevicePermission]string{
	CAMERA:   "camera",
	STORAGE:  "STORAGE",
	CONTACTS: "contacts",
}

func (d DevicePermission) String() (string, error) {
	if value, ok := devicepermissions[d]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid device permission")
}
