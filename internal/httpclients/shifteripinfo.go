package httpclients

/*
{
  "ip": "209.50.27.37",
  "city": "Lincoln (Downtown)",
  "city_geoname_id": 5072006,
  "region": "Nebraska",
  "region_code": "NE",
  "region_geoname_id": 5073708,
  "district": "Lancaster",
  "district_code": null,
  "district_geoname_id": 5071741,
  "country": "US",
  "country_name": "United States",
  "country_geoname_id": 6252001,
  "is_eu": false,
  "continent": "NA",
  "continent_name": "North America",
  "continent_geoname_id": 6255149,
  "loc": "40.8106,-96.6903",
  "latitude": 40.8106,
  "longitude": -96.6903,
  "postal": "68510",
  "timezone": "America/Chicago",
  "weather_code": "USKS0314",
  "asn": 15108,
  "as_name": "Allo Communications LLC",
  "isp": "Allo Communications LLC",
  "org": "Corporate - Lincoln",
  "connection_type": "Corporate",
  "user_type": "business",
  "is_anycast": false
}
*/

type ShifterIpInfo struct {
	Ip string `json:"ip"`
}
