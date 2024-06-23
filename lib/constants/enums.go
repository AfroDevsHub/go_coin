package constants

var Statuses = struct {
	NEW      Status
	ACTIVE   Status
	INACTIVE Status
	DISABLED Status
	DELETED  Status
}{
	NEW:      NEW,
	ACTIVE:   ACTIVE,
	INACTIVE: INACTIVE,
	DISABLED: DISABLED,
	DELETED:  DELETED,
}

var Roles = struct {
	ADMIN     Role
	USER      Role
	DEVELOPER Role
}{
	ADMIN:     ADMIN,
	USER:      USER,
	DEVELOPER: DEVELOPER,
}

var CardTypes = struct {
	CHEQUE  CardType
	SAVINGS CardType
	CREDIT  CardType
}{
	CHEQUE:  CHEQUE,
	SAVINGS: SAVINGS,
	CREDIT:  CREDIT,
}
