package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
	MillSheetPdf        string      `json:"mill_sheet_pdf"`
	MountPath           *string     `json:"mount_path"`
}

type Message struct {
	Header     []Header   `json:"Header"`
	GeneralDoc GeneralDoc `json:"GeneralDoc"`
}

type Header struct {
	InspectionLot                      string                  `json:"InspectionLot"`
	InspectionPlantBusinessPartnerName string                  `json:"InspectionPlantBusinessPartnerName"`
	InspectionLotDate                  string                  `json:"InspectionLotDate"`
	InspectionSpecification            string                  `json:"InspectionSpecification"`
	Product                            string                  `json:"Product"`
	SizeOrDimensionText                string                  `json:"SizeOrDimensionText"`
	ProductSpecification               string                  `json:"ProductSpecification"`
	MarkingOfMaterial                  string                  `json:"MarkingOfMaterial"`
	ProductionVersion                  *int                    `json:"ProductionVersion"`
	ProductionVersionItem              *int                    `json:"ProductionVersionItem"`
	ProductionOrder                    *int                    `json:"ProductionOrder"`
	ProductionOrderItem                *int                    `json:"ProductionOrderItem"`
	HeatNumber                         string                  `json:"HeatNumber"`
	DrawingNo                          string                  `json:"DrawingNo"`
	SpecDetails                        []SpecDetails           `json:"SpecDetails"`
	ComponentCompositions              []ComponentCompositions `json:"ComponentCompositions"`
	Inspections                        []Inspections           `json:"Inspections"`
	Remarks                            string                  `json:"Remarks"`
	ChiefOfInspectionSection           string                  `json:"ChiefOfInspectionSection"`
}

type SpecDetails struct {
	InspectionLot    int      `json:"InspectionLot"`
	SpecType         string   `json:"SpecType"`
	UpperLimitValue  *float32 `json:"UpperLimitValue"`
	LowerLimitValue  *float32 `json:"LowerLimitValue"`
	StandardValue    *float32 `json:"StandardValue"`
	SpecTypeUnit     *string  `json:"SpecTypeUnit"`
	HeatNumber       *string  `json:"HeatNumber"`
	SpecTypeTextInJA string   `json:"SpecTypeTextInJA"`
}

type ComponentCompositions struct {
	InspectionLot                              int      `json:"InspectionLot"`
	ComponentCompositionType                   string   `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    *float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    *float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent *float32 `json:"ComponentCompositionStandardValueInPercent"`
	HeatNumber                                 *string  `json:"HeatNumber"`
}

type Inspections struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypeTextInJA                   string   `json:"InspectionTypeTextInJA"`
}

type GeneralDoc struct {
	Product                  string  `json:"Product"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}
