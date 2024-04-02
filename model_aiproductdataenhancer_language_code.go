/*
aiproductdataenhancer/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aiproductdataenhancer

import (
	"encoding/json"
	"fmt"
)

// AiproductdataenhancerLanguageCode the model 'AiproductdataenhancerLanguageCode'
type AiproductdataenhancerLanguageCode string

// List of aiproductdataenhancerLanguageCode
const (
	AIPRODUCTDATAENHANCERLANGUAGECODE_UNKNOWN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_UNKNOWN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AB AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AB"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AF AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AF"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AK AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AK"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AM AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AM"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_AZ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_AZ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BH AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BH"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BM AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BM"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_BS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_BS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CH AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CH"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_CY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_CY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_DA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_DA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_DE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_DE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_DV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_DV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_DZ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_DZ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_EE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_EE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_EL AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_EL"
	AIPRODUCTDATAENHANCERLANGUAGECODE_EN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_EN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_EO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_EO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ES AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ES"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ET AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ET"
	AIPRODUCTDATAENHANCERLANGUAGECODE_EU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_EU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_FA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_FA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_FF AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_FF"
	AIPRODUCTDATAENHANCERLANGUAGECODE_FI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_FI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_FJ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_FJ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_FO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_FO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_FR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_FR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_FY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_FY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_GA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_GA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_GD AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_GD"
	AIPRODUCTDATAENHANCERLANGUAGECODE_GL AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_GL"
	AIPRODUCTDATAENHANCERLANGUAGECODE_GN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_GN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_GU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_GU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_GV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_GV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HT AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HT"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_HZ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_HZ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ID AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ID"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_II AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_II"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IK AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IK"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IT AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IT"
	AIPRODUCTDATAENHANCERLANGUAGECODE_IU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_IU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_JA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_JA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_JV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_JV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KJ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KJ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KK AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KK"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KL AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KL"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KM AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KM"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KW AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KW"
	AIPRODUCTDATAENHANCERLANGUAGECODE_KY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_KY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LB AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LB"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LT AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LT"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_LV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_LV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MH AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MH"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MK AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MK"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ML AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ML"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MT AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MT"
	AIPRODUCTDATAENHANCERLANGUAGECODE_MY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_MY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NB AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NB"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ND AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ND"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NL AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NL"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_NY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_NY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_OC AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_OC"
	AIPRODUCTDATAENHANCERLANGUAGECODE_OJ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_OJ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_OM AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_OM"
	AIPRODUCTDATAENHANCERLANGUAGECODE_OR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_OR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_OS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_OS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_PA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_PA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_PI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_PI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_PL AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_PL"
	AIPRODUCTDATAENHANCERLANGUAGECODE_PS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_PS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_PT AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_PT"
	AIPRODUCTDATAENHANCERLANGUAGECODE_QU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_QU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_RM AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_RM"
	AIPRODUCTDATAENHANCERLANGUAGECODE_RN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_RN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_RO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_RO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_RU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_RU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_RW AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_RW"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SC AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SC"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SD AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SD"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SK AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SK"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SL AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SL"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SM AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SM"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SQ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SQ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ST AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ST"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SU"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SV AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SV"
	AIPRODUCTDATAENHANCERLANGUAGECODE_SW AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_SW"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TH AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TH"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TK AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TK"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TL AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TL"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TN AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TN"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TS AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TS"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TT AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TT"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TW AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TW"
	AIPRODUCTDATAENHANCERLANGUAGECODE_TY AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_TY"
	AIPRODUCTDATAENHANCERLANGUAGECODE_UG AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_UG"
	AIPRODUCTDATAENHANCERLANGUAGECODE_UK AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_UK"
	AIPRODUCTDATAENHANCERLANGUAGECODE_UR AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_UR"
	AIPRODUCTDATAENHANCERLANGUAGECODE_UZ AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_UZ"
	AIPRODUCTDATAENHANCERLANGUAGECODE_VE AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_VE"
	AIPRODUCTDATAENHANCERLANGUAGECODE_VI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_VI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_VO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_VO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_WA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_WA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_WO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_WO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_XH AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_XH"
	AIPRODUCTDATAENHANCERLANGUAGECODE_YI AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_YI"
	AIPRODUCTDATAENHANCERLANGUAGECODE_YO AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_YO"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ZA AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ZA"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ZH AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ZH"
	AIPRODUCTDATAENHANCERLANGUAGECODE_ZU AiproductdataenhancerLanguageCode = "LANGUAGE_CODE_ZU"
)

// All allowed values of AiproductdataenhancerLanguageCode enum
var AllowedAiproductdataenhancerLanguageCodeEnumValues = []AiproductdataenhancerLanguageCode{
	"LANGUAGE_CODE_UNKNOWN",
	"LANGUAGE_CODE_AA",
	"LANGUAGE_CODE_AB",
	"LANGUAGE_CODE_AE",
	"LANGUAGE_CODE_AF",
	"LANGUAGE_CODE_AK",
	"LANGUAGE_CODE_AM",
	"LANGUAGE_CODE_AN",
	"LANGUAGE_CODE_AR",
	"LANGUAGE_CODE_AS",
	"LANGUAGE_CODE_AV",
	"LANGUAGE_CODE_AY",
	"LANGUAGE_CODE_AZ",
	"LANGUAGE_CODE_BA",
	"LANGUAGE_CODE_BE",
	"LANGUAGE_CODE_BG",
	"LANGUAGE_CODE_BH",
	"LANGUAGE_CODE_BM",
	"LANGUAGE_CODE_BI",
	"LANGUAGE_CODE_BN",
	"LANGUAGE_CODE_BO",
	"LANGUAGE_CODE_BR",
	"LANGUAGE_CODE_BS",
	"LANGUAGE_CODE_CA",
	"LANGUAGE_CODE_CE",
	"LANGUAGE_CODE_CH",
	"LANGUAGE_CODE_CO",
	"LANGUAGE_CODE_CR",
	"LANGUAGE_CODE_CS",
	"LANGUAGE_CODE_CU",
	"LANGUAGE_CODE_CV",
	"LANGUAGE_CODE_CY",
	"LANGUAGE_CODE_DA",
	"LANGUAGE_CODE_DE",
	"LANGUAGE_CODE_DV",
	"LANGUAGE_CODE_DZ",
	"LANGUAGE_CODE_EE",
	"LANGUAGE_CODE_EL",
	"LANGUAGE_CODE_EN",
	"LANGUAGE_CODE_EO",
	"LANGUAGE_CODE_ES",
	"LANGUAGE_CODE_ET",
	"LANGUAGE_CODE_EU",
	"LANGUAGE_CODE_FA",
	"LANGUAGE_CODE_FF",
	"LANGUAGE_CODE_FI",
	"LANGUAGE_CODE_FJ",
	"LANGUAGE_CODE_FO",
	"LANGUAGE_CODE_FR",
	"LANGUAGE_CODE_FY",
	"LANGUAGE_CODE_GA",
	"LANGUAGE_CODE_GD",
	"LANGUAGE_CODE_GL",
	"LANGUAGE_CODE_GN",
	"LANGUAGE_CODE_GU",
	"LANGUAGE_CODE_GV",
	"LANGUAGE_CODE_HA",
	"LANGUAGE_CODE_HE",
	"LANGUAGE_CODE_HI",
	"LANGUAGE_CODE_HO",
	"LANGUAGE_CODE_HR",
	"LANGUAGE_CODE_HT",
	"LANGUAGE_CODE_HU",
	"LANGUAGE_CODE_HY",
	"LANGUAGE_CODE_HZ",
	"LANGUAGE_CODE_IA",
	"LANGUAGE_CODE_ID",
	"LANGUAGE_CODE_IE",
	"LANGUAGE_CODE_IG",
	"LANGUAGE_CODE_II",
	"LANGUAGE_CODE_IK",
	"LANGUAGE_CODE_IO",
	"LANGUAGE_CODE_IS",
	"LANGUAGE_CODE_IT",
	"LANGUAGE_CODE_IU",
	"LANGUAGE_CODE_JA",
	"LANGUAGE_CODE_JV",
	"LANGUAGE_CODE_KA",
	"LANGUAGE_CODE_KG",
	"LANGUAGE_CODE_KI",
	"LANGUAGE_CODE_KJ",
	"LANGUAGE_CODE_KK",
	"LANGUAGE_CODE_KL",
	"LANGUAGE_CODE_KM",
	"LANGUAGE_CODE_KN",
	"LANGUAGE_CODE_KO",
	"LANGUAGE_CODE_KR",
	"LANGUAGE_CODE_KS",
	"LANGUAGE_CODE_KU",
	"LANGUAGE_CODE_KV",
	"LANGUAGE_CODE_KW",
	"LANGUAGE_CODE_KY",
	"LANGUAGE_CODE_LA",
	"LANGUAGE_CODE_LB",
	"LANGUAGE_CODE_LG",
	"LANGUAGE_CODE_LI",
	"LANGUAGE_CODE_LN",
	"LANGUAGE_CODE_LO",
	"LANGUAGE_CODE_LT",
	"LANGUAGE_CODE_LU",
	"LANGUAGE_CODE_LV",
	"LANGUAGE_CODE_MG",
	"LANGUAGE_CODE_MH",
	"LANGUAGE_CODE_MI",
	"LANGUAGE_CODE_MK",
	"LANGUAGE_CODE_ML",
	"LANGUAGE_CODE_MN",
	"LANGUAGE_CODE_MR",
	"LANGUAGE_CODE_MS",
	"LANGUAGE_CODE_MT",
	"LANGUAGE_CODE_MY",
	"LANGUAGE_CODE_NA",
	"LANGUAGE_CODE_NB",
	"LANGUAGE_CODE_ND",
	"LANGUAGE_CODE_NE",
	"LANGUAGE_CODE_NG",
	"LANGUAGE_CODE_NL",
	"LANGUAGE_CODE_NN",
	"LANGUAGE_CODE_NO",
	"LANGUAGE_CODE_NR",
	"LANGUAGE_CODE_NV",
	"LANGUAGE_CODE_NY",
	"LANGUAGE_CODE_OC",
	"LANGUAGE_CODE_OJ",
	"LANGUAGE_CODE_OM",
	"LANGUAGE_CODE_OR",
	"LANGUAGE_CODE_OS",
	"LANGUAGE_CODE_PA",
	"LANGUAGE_CODE_PI",
	"LANGUAGE_CODE_PL",
	"LANGUAGE_CODE_PS",
	"LANGUAGE_CODE_PT",
	"LANGUAGE_CODE_QU",
	"LANGUAGE_CODE_RM",
	"LANGUAGE_CODE_RN",
	"LANGUAGE_CODE_RO",
	"LANGUAGE_CODE_RU",
	"LANGUAGE_CODE_RW",
	"LANGUAGE_CODE_SA",
	"LANGUAGE_CODE_SC",
	"LANGUAGE_CODE_SD",
	"LANGUAGE_CODE_SE",
	"LANGUAGE_CODE_SG",
	"LANGUAGE_CODE_SI",
	"LANGUAGE_CODE_SK",
	"LANGUAGE_CODE_SL",
	"LANGUAGE_CODE_SM",
	"LANGUAGE_CODE_SN",
	"LANGUAGE_CODE_SO",
	"LANGUAGE_CODE_SQ",
	"LANGUAGE_CODE_SR",
	"LANGUAGE_CODE_SS",
	"LANGUAGE_CODE_ST",
	"LANGUAGE_CODE_SU",
	"LANGUAGE_CODE_SV",
	"LANGUAGE_CODE_SW",
	"LANGUAGE_CODE_TA",
	"LANGUAGE_CODE_TE",
	"LANGUAGE_CODE_TG",
	"LANGUAGE_CODE_TH",
	"LANGUAGE_CODE_TI",
	"LANGUAGE_CODE_TK",
	"LANGUAGE_CODE_TL",
	"LANGUAGE_CODE_TN",
	"LANGUAGE_CODE_TO",
	"LANGUAGE_CODE_TR",
	"LANGUAGE_CODE_TS",
	"LANGUAGE_CODE_TT",
	"LANGUAGE_CODE_TW",
	"LANGUAGE_CODE_TY",
	"LANGUAGE_CODE_UG",
	"LANGUAGE_CODE_UK",
	"LANGUAGE_CODE_UR",
	"LANGUAGE_CODE_UZ",
	"LANGUAGE_CODE_VE",
	"LANGUAGE_CODE_VI",
	"LANGUAGE_CODE_VO",
	"LANGUAGE_CODE_WA",
	"LANGUAGE_CODE_WO",
	"LANGUAGE_CODE_XH",
	"LANGUAGE_CODE_YI",
	"LANGUAGE_CODE_YO",
	"LANGUAGE_CODE_ZA",
	"LANGUAGE_CODE_ZH",
	"LANGUAGE_CODE_ZU",
}

func (v *AiproductdataenhancerLanguageCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiproductdataenhancerLanguageCode(value)
	for _, existing := range AllowedAiproductdataenhancerLanguageCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiproductdataenhancerLanguageCode", value)
}

// NewAiproductdataenhancerLanguageCodeFromValue returns a pointer to a valid AiproductdataenhancerLanguageCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiproductdataenhancerLanguageCodeFromValue(v string) (*AiproductdataenhancerLanguageCode, error) {
	ev := AiproductdataenhancerLanguageCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiproductdataenhancerLanguageCode: valid values are %v", v, AllowedAiproductdataenhancerLanguageCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiproductdataenhancerLanguageCode) IsValid() bool {
	for _, existing := range AllowedAiproductdataenhancerLanguageCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to aiproductdataenhancerLanguageCode value
func (v AiproductdataenhancerLanguageCode) Ptr() *AiproductdataenhancerLanguageCode {
	return &v
}

type NullableAiproductdataenhancerLanguageCode struct {
	value *AiproductdataenhancerLanguageCode
	isSet bool
}

func (v NullableAiproductdataenhancerLanguageCode) Get() *AiproductdataenhancerLanguageCode {
	return v.value
}

func (v *NullableAiproductdataenhancerLanguageCode) Set(val *AiproductdataenhancerLanguageCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAiproductdataenhancerLanguageCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAiproductdataenhancerLanguageCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiproductdataenhancerLanguageCode(val *AiproductdataenhancerLanguageCode) *NullableAiproductdataenhancerLanguageCode {
	return &NullableAiproductdataenhancerLanguageCode{value: val, isSet: true}
}

func (v NullableAiproductdataenhancerLanguageCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiproductdataenhancerLanguageCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

