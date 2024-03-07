package requests

type Header struct {
	InspectionLot			   			int		`json:"InspectionLot"`
	InspectionPlantBusinessPartnerName	string	`json:"InspectionPlantBusinessPartnerName"`
	InspectionLotDate					string	`json:"InspectionLotDate"`
	InspectionSpecification				string	`json:"InspectionSpecification"`
	Product                    			string	`json:"Product"`
	SizeOrDimensionText        			string	`json:"SizeOrDimensionText"`
	ProductSpecification	   			string	`json:"ProductSpecification"`
	MarkingOfMaterial					string	`json:"MarkingOfMaterial"`
	ProductionVersion				 	*int    `json:"ProductionVersion"`
	ProductionVersionItem			 	*int    `json:"ProductionVersionItem"`
	ProductionOrder			   			*int	`json:"ProductionOrder"`
	ProductionOrderItem		   			*int	`json:"ProductionOrderItem"`
	HeatNumber                 			string	`json:"HeatNumber"`
	DrawingNo                  			string	`json:"DrawingNo"`
	Header								[]Header	`json:"Header"`
	SpecDetails							[]SpecDetails	`json:"SpecDetails"`
	ComponentCompositions				[]ComponentCompositions	`json:"ComponentCompositions"`
	Inspections							[]Inspections	`json:"Inspections"`
	Remarks                    			string	`json:"Remarks"`
	ChiefOfInspectionSection   			string	`json:"ChiefOfInspectionSection"`
}

type SpecDetails struct {
	InspectionLot   	int      `json:"InspectionLot"`
	SpecType        	string   `json:"SpecType"`
	UpperLimitValue 	*float32 `json:"UpperLimitValue"`
	LowerLimitValue 	*float32 `json:"LowerLimitValue"`
	StandardValue   	*float32 `json:"StandardValue"`
	SpecTypeUnit    	*string  `json:"SpecTypeUnit"`
	HeatNumber			*string  `json:"HeatNumber"`
	SpecTypeTextInJA	string	 `json:"SpecTypeTextInJA"`
}

type ComponentCompositions struct {
	InspectionLot                              int		`json:"InspectionLot"`
	ComponentCompositionType                   string   `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    *float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    *float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent *float32 `json:"ComponentCompositionStandardValueInPercent"`
	HeatNumber								   *string  `json:"HeatNumber"`
}

type Inspections struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection								 int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypeTextInJA					 string   `json:"InspectionTypeTextInJA"`
}
