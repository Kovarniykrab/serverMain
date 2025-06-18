package dadata

type Suggestion struct {
	Value             string      `json:"value"`
	UnrestrictedValue string      `json:"unrestricted_value"`
	Data              AddressData `json:"data"`
}

type SuggestionInn struct {
	Value             string           `json:"value"`
	UnrestrictedValue string           `json:"unrestricted_value"`
	Data              OrganizationData `json:"data"`
}

type Suggestions []Suggestion
type SuggestionsInns []SuggestionInn

type AddressInn struct {
	Value             string         `json:"value"`
	UnrestrictedValue string         `json:"unrestricted_value"`
	Data              AddressDataInn `json:"data"`
}

type AddressDataInn struct {
	PostalCode     string `json:"postal_code"`
	Country        string `json:"country"`
	CountryISOCode string `json:"country_iso_code"`
	City           string `json:"city"`
	CityDistrict   string `json:"city_district"`
	Street         string `json:"street"`
	House          string `json:"house"`
}

type OrganizationData struct {
	INN     string     `json:"inn"`
	KPP     string     `json:"kpp"`
	OGRN    string     `json:"ogrn"`
	Name    string     `json:"full_with_opf"`
	Address AddressInn `json:"address"`
	State   struct {
		Status           string `json:"status"`
		RegistrationDate int64  `json:"registration_date"`
	} `json:"state"`
}

type Companys []Company

type Company struct {
	INN               string `json:"inn"`
	KPP               string `json:"kpp"`
	OGRN              string `json:"ogrn"`
	Name              string `json:"full_with_opf"`
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	PostalCode        string `json:"postal_code"`
	Country           string `json:"country"`
	CountryISOCode    string `json:"country_iso_code"`
	City              string `json:"city"`
	CityDistrict      string `json:"city_district"`
	Street            string `json:"street"`
	House             string `json:"house"`
	Status            string `json:"status"`
	RegistrationDate  int64  `json:"registration_date"`
}

type AddressData struct {
	PostalCode           string   `json:"postal_code"`             // Почтовый индекс.
	Country              string   `json:"country"`                 // Страна.
	CountryISOCode       string   `json:"country_iso_code"`        // ISO код страны.
	FederalDistrict      string   `json:"federal_district"`        // Федеральный округ.
	RegionFiasID         string   `json:"region_fias_id"`          // ID региона в ФИАС.
	RegionKladrID        string   `json:"region_kladr_id"`         // ID региона в КЛАДР.
	RegionISOCode        string   `json:"region_iso_code"`         // ISO код региона.
	RegionWithType       string   `json:"region_with_type"`        // Регион с типом.
	RegionType           string   `json:"region_type"`             // Тип региона.
	RegionTypeFull       string   `json:"region_type_full"`        // Полное значения региона.
	Region               string   `json:"region"`                  //  Региона.
	AreaFiasID           *string  `json:"area_fias_id"`            // ID района в ФИАС.
	AreaKladrID          *string  `json:"area_kladr_id" `          // ID района в КЛАДР.
	AreaWithType         *string  `json:"area_with_type"`          // Район с типом.
	AreaType             *string  `json:"area_type"`               // Тип района.
	AreaTypeFull         *string  `json:"area_type_full"`          // Полное название типа района.
	Area                 *string  `json:"area"`                    // Название района.
	CityFiasID           string   `json:"city_fias_id"`            // ID города в ФИАС.
	CityKladrID          string   `json:"city_kladr_id"`           // ID города в КЛАДР
	CityWithType         string   `json:"city_with_type"`          // Город с типом.
	CityType             string   `json:"city_type"`               // Тип города.
	CityTypeFull         string   `json:"city_type_full"`          // Полное название типа города.
	City                 string   `json:"city"`                    // Название города.
	CityArea             *string  `json:"city_area"`               // Район города.
	CityDistrictFiasID   *string  `json:"city_district_fias_id"`   // ID района города в ФИАС.
	CityDistrictKladrID  *string  `json:"city_district_kladr_id"`  // ID района города в КЛАДР.
	CityDistrictWithType string   `json:"city_district_with_type"` // Район города с типом.
	CityDistrictType     string   `json:"city_district_type"`      // Тип района города.
	CityDistrictTypeFull string   `json:"city_district_type_full"` // Полное название типа района города.
	CityDistrict         string   `json:"city_district"`           // Название района города.
	SettlementFiasID     *string  `json:"settlement_fias_id"`      // ID населенного пункта в ФИАС.
	SettlementKladrID    *string  `json:"settlement_kladr_id"`     // ID населенного пункта в КЛАДР.
	SettlementWithType   *string  `json:"settlement_with_type"`    // Населенный пункт с типом.
	SettlementType       *string  `json:"settlement_type"`         // Тип населенного пункта.
	SettlementTypeFull   *string  `json:"settlement_type_full"`    // Полное название типа населенного пункта.
	Settlement           *string  `json:"settlement"`              // Название населенного пункта.
	StreetFiasID         string   `json:"street_fias_id"`          // ID улицы в ФИАС.
	StreetKladrID        string   `json:"street_kladr_id"`         // ID улицы в КЛАДР.
	StreetWithType       string   `json:"street_with_type"`        // Улица с типом.
	StreetType           string   `json:"street_type"`             // Тип улицы.
	StreetTypeFull       string   `json:"street_type_full"`        // Полное названия улицы.
	Street               string   `json:"street"`                  //  Названия улицы.
	SteadFiasID          *string  `json:"stead_fias_id"`           // ID улицы в ФИАС.
	SteadCadnum          *string  `json:"stead_cadnum" `           // Кадастровый номер земельного участка.
	SteadType            *string  `json:"stead_type"`              // Тип земельного участка.
	SteadTypeFull        *string  `json:"stead_type_full"`         // Полное название типа земельного участка.
	Stead                *string  `json:"stead"`                   // Название земельного участка.
	HouseFiasID          string   `json:"house_fias_id"`           // ID дома в ФИАС.
	HouseKladrID         string   `json:"house_kladr_id"`          // ID дома в КЛАДР.
	HouseCadnum          *string  `json:"house_cadnum"`            // Кадастровый номер дома.
	HouseFlatCount       *int64   `json:"house_flat_count"`        // Количество квартир в доме.
	HouseType            string   `json:"house_type"`              // Тип дома.
	HouseTypeFull        string   `json:"house_type_full"`         // Полное название типа дома.
	House                string   `json:"house"`                   // Название дома.
	BlockType            *string  `json:"block_type"`              // Тип блока.
	BlockTypeFull        *string  `json:"block_type_full" `        // Полное название типа блока
	Block                *string  `json:"block"`                   // Корпус/строение.
	Entrance             *string  `json:"entrance"`                // Подъезд.
	Floor                *int64   `json:"floor"`                   // Этаж.
	FlatFiasID           *string  `json:"flat_fias_id"`            // ФИАС-код квартиры.
	FlatCadnum           *string  `json:"flat_cadnum"`             // Кадастровый номер квартиры.
	FlatType             *string  `json:"flat_type"`               // Тип квартиры (сокращенный).
	FlatTypeFull         *string  `json:"flat_type_full"`          // Тип квартиры.
	Flat                 *string  `json:"flat"`                    // Квартира.
	FlatArea             *float64 `json:"flat_area"`               // Площадь квартиры.
	SquareMeterPrice     *float64 `json:"square_meter_price"`      // Рыночная стоимость м².
	FlatPrice            *float64 `json:"flat_price"`              // Рыночная стоимость квартиры.
	RoomFiasID           *string  `json:"room_fias_id"`            // ФИАС-код адреса (идентификатор ФИАС)
	RoomCadnum           *string  `json:"room_cadnum"`             // Если квартира найдена в ФИАС.
	RoomType             *string  `json:"room_type"`               // Тип комнаты.
	RoomTypeFull         *string  `json:"room_type_full"`          // Полное название типа комнат.
	Room                 *string  `json:"room"`                    // Название комнаты.
	PostalBox            *string  `json:"postal_box"`              // Почтовый ящик.
	FiasID               string   `json:"fias_id"`                 // Общий ID в ФИАС.
	FiasCode             *string  `json:"fias_code"`               // Код ФИАС.
	FiasLevel            string   `json:"fias_level"`              // Уровень в иерархии ФИАС.
	FiasActualityState   string   `json:"fias_actuality_state"`    // Состояние актуальности записи ФИАС.
	KladrID              string   `json:"kladr_id"`                // КЛАДР-код адреса.
	GeonameID            string   `json:"geoname_id"`              // имя гео ID
	CapitalMarker        string   `json:"capital_marker"`          // Признак центра района или региона.
	OKATO                string   `json:"okato"`                   // Код ОКАТ.
	OKTMO                string   `json:"oktmo"`                   // Код ОКТМО.
	TaxOffice            string   `json:"tax_office"`              // Код ИФНС для физических лиц
	TaxOfficeLegal       string   `json:"tax_office_legal"`        // Код ИФНС для организаций.
	Timezone             *string  `json:"timezone"`                // Часовой пояс города.
	GeoLat               string   `json:"geo_lat"`                 // Координаты: широта.
	GeoLon               string   `json:"geo_lon"`                 // Координаты: широта
	UnparsedParts        string   `json:"unparsed_parts"`          // Нераспознанная часть адреса
}
