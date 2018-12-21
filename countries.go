package ipapi

var capitals = map[string]string{
	"BD": "Dhaka",
	"BE": "Brussels",
	"BF": "Ouagadougou",
	"BG": "Sofia",
	"BA": "Sarajevo",
	"BB": "Bridgetown",
	"WF": "Mata Utu",
	"BL": "Gustavia",
	"BM": "Hamilton",
	"BN": "Bandar Seri Begawan",
	"BO": "Sucre",
	"BH": "Manama",
	"BI": "Bujumbura",
	"BJ": "Porto-Novo",
	"BT": "Thimphu",
	"JM": "Kingston",
	"BV": "",
	"BW": "Gaborone",
	"WS": "Apia",
	"BQ": "",
	"BR": "Brasilia",
	"BS": "Nassau",
	"JE": "Saint Helier",
	"BY": "Minsk",
	"BZ": "Belmopan",
	"RU": "Moscow",
	"RW": "Kigali",
	"RS": "Belgrade",
	"TL": "Dili",
	"RE": "Saint-Denis",
	"TM": "Ashgabat",
	"TJ": "Dushanbe",
	"RO": "Bucharest",
	"TK": "",
	"GW": "Bissau",
	"GU": "Hagatna",
	"GT": "Guatemala City",
	"GS": "Grytviken",
	"GR": "Athens",
	"GQ": "Malabo",
	"GP": "Basse-Terre",
	"JP": "Tokyo",
	"GY": "Georgetown",
	"GG": "St Peter Port",
	"GF": "Cayenne",
	"GE": "Tbilisi",
	"GD": "St. George's",
	"GB": "London",
	"GA": "Libreville",
	"SV": "San Salvador",
	"GN": "Conakry",
	"GM": "Banjul",
	"GL": "Nuuk",
	"GI": "Gibraltar",
	"GH": "Accra",
	"OM": "Muscat",
	"TN": "Tunis",
	"JO": "Amman",
	"HR": "Zagreb",
	"HT": "Port-au-Prince",
	"HU": "Budapest",
	"HK": "Hong Kong",
	"HN": "Tegucigalpa",
	"HM": "",
	"VE": "Caracas",
	"PR": "San Juan",
	"PS": "East Jerusalem",
	"PW": "Melekeok",
	"PT": "Lisbon",
	"SJ": "Longyearbyen",
	"PY": "Asuncion",
	"IQ": "Baghdad",
	"PA": "Panama City",
	"PF": "Papeete",
	"PG": "Port Moresby",
	"PE": "Lima",
	"PK": "Islamabad",
	"PH": "Manila",
	"PN": "Adamstown",
	"PL": "Warsaw",
	"PM": "Saint-Pierre",
	"ZM": "Lusaka",
	"EH": "El-Aaiun",
	"EE": "Tallinn",
	"EG": "Cairo",
	"ZA": "Pretoria",
	"EC": "Quito",
	"IT": "Rome",
	"VN": "Hanoi",
	"SB": "Honiara",
	"ET": "Addis Ababa",
	"SO": "Mogadishu",
	"ZW": "Harare",
	"SA": "Riyadh",
	"ES": "Madrid",
	"ER": "Asmara",
	"ME": "Podgorica",
	"MD": "Chisinau",
	"MG": "Antananarivo",
	"MF": "Marigot",
	"MA": "Rabat",
	"MC": "Monaco",
	"UZ": "Tashkent",
	"MM": "Nay Pyi Taw",
	"ML": "Bamako",
	"MO": "Macao",
	"MN": "Ulan Bator",
	"MH": "Majuro",
	"MK": "Skopje",
	"MU": "Port Louis",
	"MT": "Valletta",
	"MW": "Lilongwe",
	"MV": "Male",
	"MQ": "Fort-de-France",
	"MP": "Saipan",
	"MS": "Plymouth",
	"MR": "Nouakchott",
	"IM": "Douglas, Isle of Man",
	"UG": "Kampala",
	"TZ": "Dodoma",
	"MY": "Kuala Lumpur",
	"MX": "Mexico City",
	"IL": "Jerusalem",
	"FR": "Paris",
	"IO": "Diego Garcia",
	"SH": "Jamestown",
	"FI": "Helsinki",
	"FJ": "Suva",
	"FK": "Stanley",
	"FM": "Palikir",
	"FO": "Torshavn",
	"NI": "Managua",
	"NL": "Amsterdam",
	"NO": "Oslo",
	"NA": "Windhoek",
	"VU": "Port Vila",
	"NC": "Noumea",
	"NE": "Niamey",
	"NF": "Kingston",
	"NG": "Abuja",
	"NZ": "Wellington",
	"NP": "Kathmandu",
	"NR": "Yaren",
	"NU": "Alofi",
	"CK": "Avarua",
	"XK": "Pristina",
	"CI": "Yamoussoukro",
	"CH": "Berne",
	"CO": "Bogota",
	"CN": "Beijing",
	"CM": "Yaounde",
	"CL": "Santiago",
	"CC": "West Island",
	"CA": "Ottawa",
	"CG": "Brazzaville",
	"CF": "Bangui",
	"CD": "Kinshasa",
	"CZ": "Prague",
	"CY": "Nicosia",
	"CX": "Flying Fish Cove",
	"CR": "San Jose",
	"CW": " Willemstad",
	"CV": "Praia",
	"CU": "Havana",
	"SZ": "Mbabane",
	"SY": "Damascus",
	"SX": "Philipsburg",
	"KG": "Bishkek",
	"KE": "Nairobi",
	"SS": "Juba",
	"SR": "Paramaribo",
	"KI": "Tarawa",
	"KH": "Phnom Penh",
	"KN": "Basseterre",
	"KM": "Moroni",
	"ST": "Sao Tome",
	"SK": "Bratislava",
	"KR": "Seoul",
	"SI": "Ljubljana",
	"KP": "Pyongyang",
	"KW": "Kuwait City",
	"SN": "Dakar",
	"SM": "San Marino",
	"SL": "Freetown",
	"SC": "Victoria",
	"KZ": "Astana",
	"KY": "George Town",
	"SG": "Singapur",
	"SE": "Stockholm",
	"SD": "Khartoum",
	"DO": "Santo Domingo",
	"DM": "Roseau",
	"DJ": "Djibouti",
	"DK": "Copenhagen",
	"VG": "Road Town",
	"DE": "Berlin",
	"YE": "Sanaa",
	"DZ": "Algiers",
	"US": "Washington",
	"UY": "Montevideo",
	"YT": "Mamoudzou",
	"UM": "",
	"LB": "Beirut",
	"LC": "Castries",
	"LA": "Vientiane",
	"TV": "Funafuti",
	"TW": "Taipei",
	"TT": "Port of Spain",
	"TR": "Ankara",
	"LK": "Colombo",
	"LI": "Vaduz",
	"LV": "Riga",
	"TO": "Nuku'alofa",
	"LT": "Vilnius",
	"LU": "Luxembourg",
	"LR": "Monrovia",
	"LS": "Maseru",
	"TH": "Bangkok",
	"TF": "Port-aux-Francais",
	"TG": "Lome",
	"TD": "N'Djamena",
	"TC": "Cockburn Town",
	"LY": "Tripolis",
	"VA": "Vatican City",
	"VC": "Kingstown",
	"AE": "Abu Dhabi",
	"AD": "Andorra la Vella",
	"AG": "St. John's",
	"AF": "Kabul",
	"AI": "The Valley",
	"VI": "Charlotte Amalie",
	"IS": "Reykjavik",
	"IR": "Tehran",
	"AM": "Yerevan",
	"AL": "Tirana",
	"AO": "Luanda",
	"AQ": "",
	"AS": "Pago Pago",
	"AR": "Buenos Aires",
	"AU": "Canberra",
	"AT": "Vienna",
	"AW": "Oranjestad",
	"IN": "New Delhi",
	"AX": "Mariehamn",
	"AZ": "Baku",
	"IE": "Dublin",
	"ID": "Jakarta",
	"UA": "Kiev",
	"QA": "Doha",
	"MZ": "Maputo",
}

var currencies = map[string]string{
	"BD": "BDT",
	"BE": "EUR",
	"BF": "XOF",
	"BG": "BGN",
	"BA": "BAM",
	"BB": "BBD",
	"WF": "XPF",
	"BL": "EUR",
	"BM": "BMD",
	"BN": "BND",
	"BO": "BOB",
	"BH": "BHD",
	"BI": "BIF",
	"BJ": "XOF",
	"BT": "BTN",
	"JM": "JMD",
	"BV": "NOK",
	"BW": "BWP",
	"WS": "WST",
	"BQ": "USD",
	"BR": "BRL",
	"BS": "BSD",
	"JE": "GBP",
	"BY": "BYR",
	"BZ": "BZD",
	"RU": "RUB",
	"RW": "RWF",
	"RS": "RSD",
	"TL": "USD",
	"RE": "EUR",
	"TM": "TMT",
	"TJ": "TJS",
	"RO": "RON",
	"TK": "NZD",
	"GW": "XOF",
	"GU": "USD",
	"GT": "GTQ",
	"GS": "GBP",
	"GR": "EUR",
	"GQ": "XAF",
	"GP": "EUR",
	"JP": "JPY",
	"GY": "GYD",
	"GG": "GBP",
	"GF": "EUR",
	"GE": "GEL",
	"GD": "XCD",
	"GB": "GBP",
	"GA": "XAF",
	"SV": "USD",
	"GN": "GNF",
	"GM": "GMD",
	"GL": "DKK",
	"GI": "GIP",
	"GH": "GHS",
	"OM": "OMR",
	"TN": "TND",
	"JO": "JOD",
	"HR": "HRK",
	"HT": "HTG",
	"HU": "HUF",
	"HK": "HKD",
	"HN": "HNL",
	"HM": "AUD",
	"VE": "VEF",
	"PR": "USD",
	"PS": "ILS",
	"PW": "USD",
	"PT": "EUR",
	"SJ": "NOK",
	"PY": "PYG",
	"IQ": "IQD",
	"PA": "PAB",
	"PF": "XPF",
	"PG": "PGK",
	"PE": "PEN",
	"PK": "PKR",
	"PH": "PHP",
	"PN": "NZD",
	"PL": "PLN",
	"PM": "EUR",
	"ZM": "ZMK",
	"EH": "MAD",
	"EE": "EUR",
	"EG": "EGP",
	"ZA": "ZAR",
	"EC": "USD",
	"IT": "EUR",
	"VN": "VND",
	"SB": "SBD",
	"ET": "ETB",
	"SO": "SOS",
	"ZW": "ZWL",
	"SA": "SAR",
	"ES": "EUR",
	"ER": "ERN",
	"ME": "EUR",
	"MD": "MDL",
	"MG": "MGA",
	"MF": "EUR",
	"MA": "MAD",
	"MC": "EUR",
	"UZ": "UZS",
	"MM": "MMK",
	"ML": "XOF",
	"MO": "MOP",
	"MN": "MNT",
	"MH": "USD",
	"MK": "MKD",
	"MU": "MUR",
	"MT": "EUR",
	"MW": "MWK",
	"MV": "MVR",
	"MQ": "EUR",
	"MP": "USD",
	"MS": "XCD",
	"MR": "MRO",
	"IM": "GBP",
	"UG": "UGX",
	"TZ": "TZS",
	"MY": "MYR",
	"MX": "MXN",
	"IL": "ILS",
	"FR": "EUR",
	"IO": "USD",
	"SH": "SHP",
	"FI": "EUR",
	"FJ": "FJD",
	"FK": "FKP",
	"FM": "USD",
	"FO": "DKK",
	"NI": "NIO",
	"NL": "EUR",
	"NO": "NOK",
	"NA": "NAD",
	"VU": "VUV",
	"NC": "XPF",
	"NE": "XOF",
	"NF": "AUD",
	"NG": "NGN",
	"NZ": "NZD",
	"NP": "NPR",
	"NR": "AUD",
	"NU": "NZD",
	"CK": "NZD",
	"XK": "EUR",
	"CI": "XOF",
	"CH": "CHF",
	"CO": "COP",
	"CN": "CNY",
	"CM": "XAF",
	"CL": "CLP",
	"CC": "AUD",
	"CA": "CAD",
	"CG": "XAF",
	"CF": "XAF",
	"CD": "CDF",
	"CZ": "CZK",
	"CY": "EUR",
	"CX": "AUD",
	"CR": "CRC",
	"CW": "ANG",
	"CV": "CVE",
	"CU": "CUP",
	"SZ": "SZL",
	"SY": "SYP",
	"SX": "ANG",
	"KG": "KGS",
	"KE": "KES",
	"SS": "SSP",
	"SR": "SRD",
	"KI": "AUD",
	"KH": "KHR",
	"KN": "XCD",
	"KM": "KMF",
	"ST": "STD",
	"SK": "EUR",
	"KR": "KRW",
	"SI": "EUR",
	"KP": "KPW",
	"KW": "KWD",
	"SN": "XOF",
	"SM": "EUR",
	"SL": "SLL",
	"SC": "SCR",
	"KZ": "KZT",
	"KY": "KYD",
	"SG": "SGD",
	"SE": "SEK",
	"SD": "SDG",
	"DO": "DOP",
	"DM": "XCD",
	"DJ": "DJF",
	"DK": "DKK",
	"VG": "USD",
	"DE": "EUR",
	"YE": "YER",
	"DZ": "DZD",
	"US": "USD",
	"UY": "UYU",
	"YT": "EUR",
	"UM": "USD",
	"LB": "LBP",
	"LC": "XCD",
	"LA": "LAK",
	"TV": "AUD",
	"TW": "TWD",
	"TT": "TTD",
	"TR": "TRY",
	"LK": "LKR",
	"LI": "CHF",
	"LV": "EUR",
	"TO": "TOP",
	"LT": "LTL",
	"LU": "EUR",
	"LR": "LRD",
	"LS": "LSL",
	"TH": "THB",
	"TF": "EUR",
	"TG": "XOF",
	"TD": "XAF",
	"TC": "USD",
	"LY": "LYD",
	"VA": "EUR",
	"VC": "XCD",
	"AE": "AED",
	"AD": "EUR",
	"AG": "XCD",
	"AF": "AFN",
	"AI": "XCD",
	"VI": "USD",
	"IS": "ISK",
	"IR": "IRR",
	"AM": "AMD",
	"AL": "ALL",
	"AO": "AOA",
	"AQ": "",
	"AS": "USD",
	"AR": "ARS",
	"AU": "AUD",
	"AT": "EUR",
	"AW": "AWG",
	"IN": "INR",
	"AX": "EUR",
	"AZ": "AZN",
	"IE": "EUR",
	"ID": "IDR",
	"UA": "UAH",
	"QA": "QAR",
	"MZ": "MZN",
}

var phonePrefixes = map[string]string{
	"BD": "+880",
	"BE": "+32",
	"BF": "+226",
	"BG": "+359",
	"BA": "+387",
	"BB": "+1-246",
	"WF": "+681",
	"BL": "+590",
	"BM": "+1-441",
	"BN": "+673",
	"BO": "+591",
	"BH": "+973",
	"BI": "+257",
	"BJ": "+229",
	"BT": "+975",
	"JM": "+1-876",
	"BV": "",
	"BW": "+267",
	"WS": "+685",
	"BQ": "+599",
	"BR": "+55",
	"BS": "+1-242",
	"JE": "+44-1534",
	"BY": "+375",
	"BZ": "+501",
	"RU": "+7",
	"RW": "+250",
	"RS": "+381",
	"TL": "+670",
	"RE": "+262",
	"TM": "+993",
	"TJ": "+992",
	"RO": "+40",
	"TK": "+690",
	"GW": "+245",
	"GU": "+1-671",
	"GT": "+502",
	"GS": "",
	"GR": "+30",
	"GQ": "+240",
	"GP": "+590",
	"JP": "+81",
	"GY": "+592",
	"GG": "+44-1481",
	"GF": "+594",
	"GE": "+995",
	"GD": "+1-473",
	"GB": "+44",
	"GA": "+241",
	"SV": "+503",
	"GN": "+224",
	"GM": "+220",
	"GL": "+299",
	"GI": "+350",
	"GH": "+233",
	"OM": "+968",
	"TN": "+216",
	"JO": "+962",
	"HR": "+385",
	"HT": "+509",
	"HU": "+36",
	"HK": "+852",
	"HN": "+504",
	"HM": "+ ",
	"VE": "+58",
	"PR": "+1-787 and 1-939",
	"PS": "+970",
	"PW": "+680",
	"PT": "+351",
	"SJ": "+47",
	"PY": "+595",
	"IQ": "+964",
	"PA": "+507",
	"PF": "+689",
	"PG": "+675",
	"PE": "+51",
	"PK": "+92",
	"PH": "+63",
	"PN": "+870",
	"PL": "+48",
	"PM": "+508",
	"ZM": "+260",
	"EH": "+212",
	"EE": "+372",
	"EG": "+20",
	"ZA": "+27",
	"EC": "+593",
	"IT": "+39",
	"VN": "+84",
	"SB": "+677",
	"ET": "+251",
	"SO": "+252",
	"ZW": "+263",
	"SA": "+966",
	"ES": "+34",
	"ER": "+291",
	"ME": "+382",
	"MD": "+373",
	"MG": "+261",
	"MF": "+590",
	"MA": "+212",
	"MC": "+377",
	"UZ": "+998",
	"MM": "+95",
	"ML": "+223",
	"MO": "+853",
	"MN": "+976",
	"MH": "+692",
	"MK": "+389",
	"MU": "+230",
	"MT": "+356",
	"MW": "+265",
	"MV": "+960",
	"MQ": "+596",
	"MP": "+1-670",
	"MS": "+1-664",
	"MR": "+222",
	"IM": "+44-1624",
	"UG": "+256",
	"TZ": "+255",
	"MY": "+60",
	"MX": "+52",
	"IL": "+972",
	"FR": "+33",
	"IO": "+246",
	"SH": "+290",
	"FI": "+358",
	"FJ": "+679",
	"FK": "+500",
	"FM": "+691",
	"FO": "+298",
	"NI": "+505",
	"NL": "+31",
	"NO": "+47",
	"NA": "+264",
	"VU": "+678",
	"NC": "+687",
	"NE": "+227",
	"NF": "+672",
	"NG": "+234",
	"NZ": "+64",
	"NP": "+977",
	"NR": "+674",
	"NU": "+683",
	"CK": "+682",
	"XK": "",
	"CI": "+225",
	"CH": "+41",
	"CO": "+57",
	"CN": "+86",
	"CM": "+237",
	"CL": "+56",
	"CC": "+61",
	"CA": "+1",
	"CG": "+242",
	"CF": "+236",
	"CD": "+243",
	"CZ": "+420",
	"CY": "+357",
	"CX": "+61",
	"CR": "+506",
	"CW": "+599",
	"CV": "+238",
	"CU": "+53",
	"SZ": "+268",
	"SY": "+963",
	"SX": "+599",
	"KG": "+996",
	"KE": "+254",
	"SS": "+211",
	"SR": "+597",
	"KI": "+686",
	"KH": "+855",
	"KN": "+1-869",
	"KM": "+269",
	"ST": "+239",
	"SK": "+421",
	"KR": "+82",
	"SI": "+386",
	"KP": "+850",
	"KW": "+965",
	"SN": "+221",
	"SM": "+378",
	"SL": "+232",
	"SC": "+248",
	"KZ": "+7",
	"KY": "+1-345",
	"SG": "+65",
	"SE": "+46",
	"SD": "+249",
	"DO": "+1-809 and 1-829",
	"DM": "+1-767",
	"DJ": "+253",
	"DK": "+45",
	"VG": "+1-284",
	"DE": "+49",
	"YE": "+967",
	"DZ": "+213",
	"US": "+1",
	"UY": "+598",
	"YT": "+262",
	"UM": "+1",
	"LB": "+961",
	"LC": "+1-758",
	"LA": "+856",
	"TV": "+688",
	"TW": "+886",
	"TT": "+1-868",
	"TR": "+90",
	"LK": "+94",
	"LI": "+423",
	"LV": "+371",
	"TO": "+676",
	"LT": "+370",
	"LU": "+352",
	"LR": "+231",
	"LS": "+266",
	"TH": "+66",
	"TF": "",
	"TG": "+228",
	"TD": "+235",
	"TC": "+1-649",
	"LY": "+218",
	"VA": "+379",
	"VC": "+1-784",
	"AE": "+971",
	"AD": "+376",
	"AG": "+1-268",
	"AF": "+93",
	"AI": "+1-264",
	"VI": "+1-340",
	"IS": "+354",
	"IR": "+98",
	"AM": "+374",
	"AL": "+355",
	"AO": "+244",
	"AQ": "",
	"AS": "+1-684",
	"AR": "+54",
	"AU": "+61",
	"AT": "+43",
	"AW": "+297",
	"IN": "+91",
	"AX": "+358-18",
	"AZ": "+994",
	"IE": "+353",
	"ID": "+62",
	"UA": "+380",
	"QA": "+974",
	"MZ": "+258",
}