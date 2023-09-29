package response

import "time"

type RetornoEndPointPacientesQmatic struct {
	ID                       int    `json:"id"`
	PublicID                 string `json:"publicId"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	CardNumber               string `json:"cardNumber"`
	ConsentIdentifier        string `json:"consentIdentifier"`
	ConsentTimestamp         string `json:"consentTimestamp"`
	RetentionPolicy          string `json:"retentionPolicy"`
	LastInteractionTimestamp string `json:"lastInteractionTimestamp"`
	DeletionTimestamp        string `json:"deletionTimestamp"`
	Properties               struct {
		CustomField5    string `json:"customField5"`
		LastName        string `json:"lastName"`
		Gender          string `json:"gender"`
		PhoneMobile     string `json:"phoneMobile"`
		PhoneHome       string `json:"phoneHome"`
		DateOfBirth     string `json:"dateOfBirth"`
		AccountNumber   string `json:"accountNumber"`
		PhoneWork       string `json:"phoneWork"`
		PhoneNumber     string `json:"phoneNumber"`
		AddressLine1    string `json:"addressLine1"`
		AddressPostCode string `json:"addressPostCode"`
		AddressLine2    string `json:"addressLine2"`
		AddressLine3    string `json:"addressLine3"`
		AddressLine4    string `json:"addressLine4"`
		AddressLine5    string `json:"addressLine5"`
		Email           string `json:"email"`
	} `json:"properties"`
}

type RetornoAgendamentosQmatic struct {
	Id                int         `json:"id,omitempty"`
	ExternalId        string      `json:"externalId,omitempty"`
	QpCalendarId      interface{} `json:"qpCalendarId,omitempty"`
	BranchId          int         `json:"branchId,omitempty"`
	ResourceId        int         `json:"resourceId,omitempty"`
	ResourceName      string      `json:"resourceName,omitempty"`
	ResourceGroupId   interface{} `json:"resourceGroupId,omitempty"`
	ResourceGroupName interface{} `json:"resourceGroupName,omitempty"`
	Services          []struct {
		Id          int    `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Duration    int    `json:"duration,omitempty"`
	} `json:"services,omitempty"`
	ResourceServiceDecorators []interface{} `json:"resourceServiceDecorators,omitempty"`
	Customers                 []struct {
		Id                       int         `json:"id,omitempty"`
		PublicId                 string      `json:"publicId,omitempty"`
		FirstName                string      `json:"firstName,omitempty"`
		LastName                 string      `json:"lastName,omitempty"`
		CardNumber               string      `json:"cardNumber,omitempty"`
		ConsentIdentifier        interface{} `json:"consentIdentifier,omitempty"`
		ConsentTimestamp         interface{} `json:"consentTimestamp,omitempty"`
		RetentionPolicy          string      `json:"retentionPolicy,omitempty"`
		LastInteractionTimestamp string      `json:"lastInteractionTimestamp,omitempty"`
		DeletionTimestamp        string      `json:"deletionTimestamp,omitempty"`
		Properties               struct {
			CustomField5    string      `json:"customField5,omitempty"`
			LastName        string      `json:"lastName,omitempty"`
			Gender          string      `json:"gender,omitempty"`
			PhoneMobile     string      `json:"phoneMobile,omitempty"`
			PhoneHome       string      `json:"phoneHome,omitempty"`
			DateOfBirth     time.Time   `json:"dateOfBirth,omitempty"`
			AccountNumber   string      `json:"accountNumber,omitempty"`
			PhoneWork       string      `json:"phoneWork,omitempty"`
			PhoneNumber     interface{} `json:"phoneNumber,omitempty"`
			AddressLine1    string      `json:"addressLine1,omitempty"`
			AddressPostCode string      `json:"addressPostCode,omitempty"`
			AddressLine2    string      `json:"addressLine2,omitempty"`
			AddressLine3    string      `json:"addressLine3,omitempty"`
			AddressLine4    string      `json:"addressLine4,omitempty"`
			AddressLine5    string      `json:"addressLine5,omitempty"`
			Email           string      `json:"email,omitempty"`
		} `json:"properties,omitempty"`
	} `json:"customers,omitempty"`
	StartTime  string      `json:"startTime,omitempty"`
	EndTime    string      `json:"endTime,omitempty"`
	UpdateTime string      `json:"updateTime,omitempty"`
	Status     string      `json:"status,omitempty"`
	Title      interface{} `json:"title,omitempty"`
	Properties struct {
		Authorization string `json:"authorization,omitempty"`
		OtrosDatos    string `json:"otros_datos,omitempty"`
		Notes         string `json:"notes,omitempty"`
	} `json:"properties"`
	UserName interface{} `json:"userName,omitempty"`
}
