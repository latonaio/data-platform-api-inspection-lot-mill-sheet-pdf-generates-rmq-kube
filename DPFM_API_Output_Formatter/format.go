package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-inspection-lot-mill-sheet-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	"fmt"
)

func SetToPdfData(
	headerData *dpfm_api_input_reader.Header,
	inputSpecDetails []dpfm_api_input_reader.SpecDetails,
	inputComponentCompositions []dpfm_api_input_reader.ComponentCompositions,
	inputInspections []dpfm_api_input_reader.Inspections,
) *Header {
	pm := &Header{}

	var specDetails []SpecDetails
	for _, dataSDs := range inputSpecDetails {
		if headerData.InspectionLot == dataSDs.InspectionLot {
			specDetails = append(specDetails, SpecDetails{
				InspectionLot:    dataSDs.InspectionLot,
				SpecType:         dataSDs.SpecType,
				UpperLimitValue:  dataSDs.UpperLimitValue,
				LowerLimitValue:  dataSDs.LowerLimitValue,
				StandardValue:    dataSDs.StandardValue,
				SpecTypeUnit:     dataSDs.SpecTypeUnit,
				HeatNumber:       dataSDs.HeatNumber,
				SpecTypeTextInJA: dataSDs.SpecTypeTextInJA,
			})
		}
	}

	var componentCompositions []ComponentCompositions
	for _, dataCCs := range inputComponentCompositions {
		if headerData.InspectionLot == dataCCs.InspectionLot {
			componentCompositions = append(componentCompositions, ComponentCompositions{
				InspectionLot:                              dataCCs.InspectionLot,
				ComponentCompositionType:                   dataCCs.ComponentCompositionType,
				ComponentCompositionUpperLimitInPercent:    dataCCs.ComponentCompositionUpperLimitInPercent,
				ComponentCompositionLowerLimitInPercent:    dataCCs.ComponentCompositionLowerLimitInPercent,
				ComponentCompositionStandardValueInPercent: dataCCs.ComponentCompositionStandardValueInPercent,
				HeatNumber: dataCCs.HeatNumber,
			})
		}
	}

	var inspections []Inspections
	for _, dataIPTs := range inputInspections {
		if headerData.InspectionLot == dataIPTs.InspectionLot {
			inspections = append(inspections, Inspections{
				InspectionLot:                            dataIPTs.InspectionLot,
				Inspection:                               dataIPTs.Inspection,
				InspectionType:                           dataIPTs.InspectionType,
				InspectionTypeCertificateValueInText:     dataIPTs.InspectionTypeCertificateValueInText,
				InspectionTypeCertificateValueInQuantity: dataIPTs.InspectionTypeCertificateValueInQuantity,
				InspectionTypeValueUnit:                  dataIPTs.InspectionTypeValueUnit,
				InspectionTypeTextInJA:                   dataIPTs.InspectionTypeTextInJA,
			})
		}
	}

	pm = &Header{
		InspectionLot:                      fmt.Sprintf("%d", headerData.InspectionLot),
		InspectionPlantBusinessPartnerName: headerData.InspectionPlantBusinessPartnerName,
		InspectionLotDate:                  headerData.InspectionLotDate,
		InspectionSpecification:            headerData.InspectionSpecification,
		Product:                            headerData.Product,
		SizeOrDimensionText:                headerData.SizeOrDimensionText,
		ProductSpecification:               headerData.ProductSpecification,
		MarkingOfMaterial:                  headerData.MarkingOfMaterial,
		ProductionOrder:                    headerData.ProductionOrder,
		ProductionOrderItem:                headerData.ProductionOrderItem,
		HeatNumber:                         headerData.HeatNumber,
		DrawingNo:                          headerData.DrawingNo,
		SpecDetails:                        specDetails,
		ComponentCompositions:              componentCompositions,
		Inspections:                        inspections,
		Remarks:                            headerData.Remarks,
		ChiefOfInspectionSection:           headerData.ChiefOfInspectionSection,
	}

	return pm
}
