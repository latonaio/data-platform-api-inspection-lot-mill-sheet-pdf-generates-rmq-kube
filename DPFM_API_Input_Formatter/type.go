package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey         string                  `json:"connection_key"`
	Result                bool                    `json:"result"`
	RedisKey              string                  `json:"redis_key"`
	Filepath              string                  `json:"filepath"`
	APIStatusCode         int                     `json:"api_status_code"`
	RuntimeSessionID      string                  `json:"runtime_session_id"`
	BusinessPartnerID     *int                    `json:"business_partner"`
	ServiceLabel          string                  `json:"service_label"`
	Product               Product                 `json:"Product"`
	Header                []Header                `json:"Header"`
	SpecDetails           []SpecDetails           `json:"SpecDetail"`
	ComponentCompositions []ComponentCompositions `json:"ComponentComposition"`
	Inspections           []Inspections           `json:"Inspection"`
	APISchema             string                  `json:"api_schema"`
	Accepter              []string                `json:"accepter"`
	Deleted               bool                    `json:"deleted"`
	DocData               string                  `json:"doc_data"`
	APIProcessingResult   *bool                   `json:"api_processing_result"`
	APIProcessingError    string                  `json:"api_processing_error"`
}

type Header struct {
	InspectionLot                      int    `json:"InspectionLot"`
	InspectionPlantBusinessPartnerName string `json:"InspectionPlantBusinessPartnerName"`
	InspectionLotDate                  string `json:"InspectionLotDate"`
	InspectionSpecification            string `json:"InspectionSpecification"`
	Product                            string `json:"Product"`
	SizeOrDimensionText                string `json:"SizeOrDimensionText"`
	ProductSpecification               string `json:"ProductSpecification"`
	MarkingOfMaterial                  string `json:"MarkingOfMaterial"`
	ProductionVersion                  *int   `json:"ProductionVersion"`
	ProductionVersionItem              *int   `json:"ProductionVersionItem"`
	ProductionOrder                    *int   `json:"ProductionOrder"`
	ProductionOrderItem                *int   `json:"ProductionOrderItem"`
	HeatNumber                         string `json:"HeatNumber"`
	DrawingNo                          string `json:"DrawingNo"`
	Remarks                            string `json:"Remarks"`
	ChiefOfInspectionSection           string `json:"ChiefOfInspectionSection"`
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

type Product struct {
	Product    string     `json:"Product"`
	GeneralDoc GeneralDoc `json:"GeneralDoc"`
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
