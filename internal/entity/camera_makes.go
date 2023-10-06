package entity

var CameraMakes = map[string]string{
	"acer":                        "Acer",
	"ACER":                        "Acer",
	"asus":                        "ASUS",
	"Asus":                        "ASUS",
	"ASUS_AI2302":                 "ASUS",
	"apple":                       "Apple",
	"Casio":                       "CASIO",
	"CASIO COMPUTER":              "CASIO",
	"CASIO COMPUTER CO":           "CASIO",
	"CASIO COMPUTER CO.":          "CASIO",
	"CASIO COMPUTER CO.,LTD":      "CASIO",
	"CASIO COMPUTER CO.,LTD.":     "CASIO",
	"CASIO CORPORATION":           "CASIO",
	"Fujifilm":                    "FUJIFILM",
	"FUJIFILM CORPORATION":        "FUJIFILM",
	"Garmin-Asus":                 "Garmin",
	"google":                      "Google",
	"GOOGLE":                      "Google",
	"Hewlett-Packard":             "HP",
	"htc":                         "HTC",
	"Kodak":                       "KODAK",
	"EASTMAN KODAK":               "KODAK",
	"EASTMAN KODAK COMPANY":       "KODAK",
	"Leica Camera AG":             "Leica",
	"LEICA":                       "Leica",
	"LG Electronics":              "LG",
	"LGE":                         "LG",
	"lge":                         "LG",
	"Minolta Co., Ltd.":           "Minolta",
	"MINOLTA CO.,LTD":             "Minolta",
	"KONICA MINOLTA":              "Konica Minolta",
	"Motorol":                     "Motorola",
	"Motorola Mobility":           "Motorola",
	"motorola":                    "Motorola",
	"samsung":                     "Samsung",
	"SAMSUNG":                     "Samsung",
	"Samsung Electronics":         "Samsung",
	"SAMSUNG TECHWIN Co.":         "Samsung",
	"SAMSUNG TECHWIN":             "Samsung",
	"Samsung Techwin":             "Samsung",
	"sharp":                       "SHARP",
	"Sharp":                       "SHARP",
	"sigma":                       "SIGMA",
	"Sigma":                       "SIGMA",
	"OLYMPUS":                     "Olympus",
	"OLYMPUS CORPORATION":         "Olympus",
	"OLYMPUS DIGITAL CAMERA":      "Olympus",
	"OLYMPUS IMAGING CORP.":       "Olympus",
	"OLYMPUS OPTICAL CO.,LTD":     "Olympus",
	"Nikon":                       "NIKON",
	"NIKON CORPORATION":           "NIKON",
	"Huawei":                      "HUAWEI",
	"RaspberryPi":                 "Raspberry Pi",
	"RICOH IMAGING COMPANY, LTD.": "RICOH",
	"Ricoh":                       "RICOH",
	"Pentax":                      "PENTAX",
	"PENTAX Corporation":          "PENTAX",
	"PENTAX CORPORATION":          "PENTAX",
	"Blackberry":                  "BlackBerry",
	"Research In Motion":          "BlackBerry",
	"Sony":                        "SONY",
	"VenTrade GmbH, Germany":      "VenTrade",
}

var CameraModels = map[string]string{
	"AI2302":                     "Zenfone 10",
	"_AI2302":                    "Zenfone 10",
	"ASUS_AI2302":                "Zenfone 10",
	"Blackberry Q5":              "BlackBerry Q5",
	"Blackberry Q10":             "BlackBerry Q10",
	"Blackberry Z10":             "BlackBerry Z10",
	"Blackberry Z3":              "BlackBerry Z3",
	"Blackberry Z30":             "BlackBerry Z30",
	"Blackberry Leap":            "BlackBerry Leap",
	"Blackberry Classic":         "BlackBerry Classic",
	"Blackberry Passport":        "BlackBerry Passport",
	"iPhone SE (1st generation)": "iPhone SE",
	"iPhone SE (2nd generation)": "iPhone SE",
	"iPhone SE (3rd generation)": "iPhone SE",
	"iPhone SE (4th generation)": "iPhone SE",
	"iPhone SE (5th generation)": "iPhone SE",
	"iPad (1st generation)":      "iPad",
	"iPad (2nd generation)":      "iPad",
	"iPad (3rd generation)":      "iPad",
	"iPad (4th generation)":      "iPad",
	"iPad (5th generation)":      "iPad",
	"iPad Air (1st generation)":  "iPad Air",
	"iPad Air (2nd generation)":  "iPad Air",
	"iPad Air (3rd generation)":  "iPad Air",
	"iPad Air (4th generation)":  "iPad Air",
	"iPad Air (5th generation)":  "iPad Air",
	"iPad Pro (1st generation)":  "iPad Pro",
	"iPad Pro (2nd generation)":  "iPad Pro",
	"iPad Pro (3rd generation)":  "iPad Pro",
	"iPad Pro (4th generation)":  "iPad Pro",
	"iPad Pro (5th generation)":  "iPad Pro",
	"SM-G780F":                   "Galaxy S20",
	"SM-G781B":                   "Galaxy S20 FE",
	"SM-G991A":                   "Galaxy S21",
	"SM-G991B":                   "Galaxy S21",
	"SM-G990A":                   "Galaxy S21 FE",
	"SM-G990B":                   "Galaxy S21 FE",
	"SM-G996A":                   "Galaxy S21+",
	"SM-G996B":                   "Galaxy S21+",
	"SM-G998A":                   "Galaxy S21 Ultra",
	"SM-G998B":                   "Galaxy S21 Ultra",
	"SM-S911A":                   "Galaxy S23",
	"SM-S911B":                   "Galaxy S23",
	"SM-S916A":                   "Galaxy S23+",
	"SM-S916B":                   "Galaxy S23+",
	"SM-S918A":                   "Galaxy S23 Ultra",
	"SM-S918B":                   "Galaxy S23 Ultra",
	"WAS-LX1":                    "P10 lite",
	"WAS-LX2":                    "P10 lite",
	"WAS-LX3":                    "P10 lite",
	"WAS-LX1A":                   "P10 lite",
	"WAS-LX2J":                   "P10 lite",
	"WAS-L03T":                   "P10 lite",
	"WAS-AL00":                   "P10 lite",
	"WAS-TL10":                   "P10 lite",
	"VTR-L29":                    "P10",
	"VTR-AL00":                   "P10",
	"VTR-TL00":                   "P10",
	"VTR-L09":                    "P10",
	"EML-AL00":                   "P20",
	"EML-L09":                    "P20",
	"EML-L09C":                   "P20",
	"EML-L29":                    "P20",
	"EML-L29C":                   "P20",
	"CLT-AL00":                   "P20 Pro",
	"CLT-AL01":                   "P20 Pro",
	"CLT-TL01":                   "P20 Pro",
	"CLT-L09":                    "P20 Pro",
	"CLT-L29":                    "P20 Pro",
	"ELE-L29":                    "P30",
	"ELE-AL00":                   "P30",
	"ELE-L04":                    "P30",
	"ELE-L09":                    "P30",
	"ELE-TL00":                   "P30",
	"VOG-L29":                    "P30 Pro",
	"VOG-L09":                    "P30 Pro",
	"VOG-L04":                    "P30 Pro",
	"VOG-AL00":                   "P30 Pro",
	"VOG-AL10":                   "P30 Pro",
	"VOG-TL00":                   "P30 Pro",
	"MAR-L01A":                   "P30 lite",
	"MAR-L21A":                   "P30 lite",
	"MAR-LX1A":                   "P30 lite",
	"MAR-LX1M":                   "P30 lite",
	"MAR-LX2":                    "P30 lite",
	"MAR-L21MEA":                 "P30 lite",
	"MAR-L22A":                   "P30 lite",
	"MAR-L22B":                   "P30 lite",
	"MAR-LX3A":                   "P30 lite",
	"ANA-AN00":                   "P40",
	"ANA-TN00":                   "P40",
	"ELS-AN00":                   "P40 Pro",
	"ELS-TN00":                   "P40 Pro",
	"ELS-NX9":                    "P40 Pro",
	"ELS-N04":                    "P40 Pro",
	"JNY-L21A":                   "P40 lite",
	"JNY-L01A":                   "P40 lite",
	"JNY-L21B":                   "P40 lite",
	"JNY-L22A":                   "P40 lite",
	"JNY-L02A":                   "P40 lite",
	"JNY-L22B":                   "P40 lite",
	"STK-LX1":                    "Honor 9X",
	"HLK-AL00":                   "Honor 9X",
	"HLK-TL00":                   "Honor 9X",
	"SNE-AL00":                   "Mate 20 lite",
	"SNE-LX1":                    "Mate 20 lite",
	"SNE-LX2":                    "Mate 20 lite",
	"SNE-LX3":                    "Mate 20 lite",
	"INE-LX2":                    "Mate 20 lite",
	"HMA-L29":                    "Mate 20",
	"HMA-L09":                    "Mate 20",
	"HMA-LX9":                    "Mate 20",
	"HMA-AL00":                   "Mate 20",
	"HMA-TL00":                   "Mate 20",
	"LYA-L09":                    "Mate 20 Pro",
	"LYA-L29":                    "Mate 20 Pro",
	"LYA-AL00":                   "Mate 20 Pro",
	"LYA-AL10":                   "Mate 20 Pro",
	"LYA-TL00":                   "Mate 20 Pro",
	"LYA-L0C":                    "Mate 20 Pro",
	"TAS-L09":                    "Mate 30",
	"TAS-L29":                    "Mate 30",
	"TAS-AL00":                   "Mate 30",
	"TAS-TL00":                   "Mate 30",
	"LIO-L09":                    "Mate 30 Pro",
	"LIO-L29":                    "Mate 30 Pro",
	"LIO-AL00":                   "Mate 30 Pro",
	"LIO-TL00":                   "Mate 30 Pro",
}
