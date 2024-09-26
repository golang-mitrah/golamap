package golamap

const (
	DirectionResponse = `
	{
	  "geocoded_waypoints": [
		{
		  "geocoder_status": "OK",
		  "place_id": "tdqcey303",
		  "types": []
		}
	  ],
	  "routes": [
		{
		  "bounds": "string",
		  "copyrights": "OLA Map data ©2024",
		  "legs": [
			{
			  "distance": 3836,
			  "readable_distance": "3 km 800 metres",
			  "duration": 698,
			  "readable_duration": "14 mins",
			  "end_address": "12.93397277.611342",
			  "end_location": {
				"lat": 12.933972,
				"lng": 77.611342
			  },
			  "start_address": "12.90934277.621689",
			  "start_location": {
				"lat": 12.909342,
				"lng": 77.62168
			  },
			  "steps": [
				{
				  "distance": 476,
				  "readable_distance": "0.5 km",
				  "duration": 99,
				  "readable_duration": "0 hours 2 min",
				  "start_location": {
					"lat": 12.90934,
					"lng": 77.62169
				  },
				  "end_location": {
					"lat": 12.908710000000001,
					"lng": 77.62158
				  },
				  "instructions": "Head west on NH48",
				  "maneuver": "turn-right",
				  "bearing_before": 0,
				  "bearing_after": 259
				}
			  ]
			}
		  ],
		  "travel_advisory": "0,1,0 | 1,3,15 | 3,4,10",
		  "summary": "",
		  "warnings": [],
		  "waypoint_order": [
			0,
			3,
			2,
			1
		  ]
		}
	  ],
	  "status": "OK",
	  "source_from": "Ola Maps"
	}`

	DistanceMatrixResponse = `{
		"rows": [
		  {
			"elements": [
			  {
				"duration": 719,
				"distance": 3224,
				"polyline": "su|mAyjvxM[GMCPmA?E@CJgADGDAxBTDJ[lCe@nDc@lDALAX^?d@WnAk@hBcAHEb@UVOJSHa@@CL_@Pc@FUDSJc@Lg@dBiGzFaTBIFUpBeGZaA|@iCpCiJ@CpBqGf@{AbByFx@_CFW@kBg@iFASUiBGu@y@kJAECMASKaAASCSSaCOaBUsCG}@E}@?YD]",
				"status": "OK"
			  },
			  {
				"duration": 1257,
				"distance": 5863,
				"status": "OK"
			  }
			]
		  }
		],
	  }`

	GeoCodeResponse = `{
		"status": "ok",
		"geocodingResults": [
		  {
			"formatted_address": "Mumbai, Mumbai Suburban, Mumbai Suburban District, Maharashtra, India",
			"types": [],
			"name": "Mumbai",
			"geometry": {
			  "viewport": {
				"southwest": {
				  "lng": 72.87232696702962,
				  "lat": 19.114096358602232
				},
				"northeast": {
				  "lng": 72.87498903297039,
				  "lat": 19.116611641397768
				}
			  },
			  "location": {
				"lng": 72.873658,
				"lat": 19.115354
			  },
			  "location_type": "geometric_center"
			},
			"address_components": [
			  {
				"types": [
				  "country"
				],
				"short_name": "India",
				"long_name": "India"
			  },
			  {
				"types": [
				  "administrative_area_level_1"
				],
				"short_name": "Maharashtra",
				"long_name": "Maharashtra"
			  },
			  {
				"types": [
				  "administrative_area_level_2"
				],
				"short_name": "Mumbai Suburban District",
				"long_name": "Mumbai Suburban District"
			  },
			  {
				"types": [
				  "administrative_area_level_3"
				],
				"short_name": "Mumbai Suburban",
				"long_name": "Mumbai Suburban"
			  },
			  {
				"types": [
				  "locality"
				],
				"short_name": "Mumbai",
				"long_name": "Mumbai"
			  },
			  {
				"types": [
				  "postal_code"
				],
				"short_name": "400059",
				"long_name": "400059"
			  },
			  {
				"types": [
				  "street_address"
				],
				"short_name": "Road Number 16",
				"long_name": "Road Number 16"
			  }
			],
			"plus_code": {
			  "compound_code": "NA",
			  "global_code": "NA"
			},
			"place_id": "ola-platform:1",
			"layer": [
			  "locality"
			]
		  }
		]
	  }`

	ReverseGeocodeResponse = `{
		"error_message": "",
		"info_messages": [],
		"results": [
		  {
			"formatted_address": "West Oidar Services, Central Tax, Bangalore, 560085, Phase 3, Banashankari, Bengaluru, Karnataka",
			"types": [
			  "point_of_interest",
			  "establishment"
			],
			"name": "West Oidar Services, Central Tax, Bangalore",
			"geometry": {
			  "viewport": {
				"southwest": {
				  "lng": 77.55141967245793,
				  "lat": 12.922552358602232
				},
				"northeast": {
				  "lng": 77.55400032754208,
				  "lat": 12.925067641397767
				}
			  },
			  "location": {
				"lng": 77.55271,
				"lat": 12.92381
			  },
			  "location_type": "rooftop"
			},
			"address_components": [
			  {
				"types": [
				  [
					"country"
				  ]
				],
				"short_name": "India",
				"long_name": "India"
			  }
			],
			"plus_code": {
			  "compound_code": "NA",
			  "global_code": "NA"
			},
			"place_id": "ola-platform:90be49b9e86d386a7076e25f71b402ca",
			"layer": [
			  "venue"
			]
		  }
		],
		"plus_code": {
		  "compound_code": "NA",
		  "global_code": "NA"
		},
		"status": "ok"
	  }`

	SnapToRoadResponse = `{
		"status": "SUCCESS",
		"snapped_points": [
		  {
			"location": {
			  "lat": 12.992597,
			  "lng": 77.659875
			},
			"original_index": 2,
			"snapped_type": "Nearest"
		  }
		]
	  }`

	NearestRoadResponse = `{
		"status": "SUCCESS",
		"results": [
		  {
			"lat": 12.97711937255269,
			"lng": 77.57183856264156,
			"distance": 34.15771496364373,
			"originalIndex": 0,
			"status": "SUCCESS"
		  }
		]
	  }`

	AutoCompleteResponse = `{
		"predictions": [
		  {
			"reference": "ChIJZWJEdf4crjsRjkEpoelwbCk",
			"types": [
			  "airport",
			  "point_of_interest",
			  "establishment"
			],
			"matched_substrings": [
			  {
				"offset": 0,
				"length": 5
			  }
			],
			"distance_meters": 31312,
			"terms": [
			  {
				"offset": 0,
				"value": "Kempegowda International Airport Bengaluru (BLR)"
			  }
			],
			"structured_formatting": {
			  "main_text_matched_substrings": [
				{
				  "offset": 0,
				  "length": 5
				}
			  ],
			  "secondary_text": "KIAL Rd, Devanahalli, Bengaluru, Karnataka, India",
			  "main_text": "Kempegowda International Airport Bengaluru (BLR)"
			},
			"description": "Kempegowda International Airport Bengaluru (BLR), KIAL Rd, Devanahalli, Bengaluru, Karnataka, India",
			"geometry": {
			  "location": {
				"lng": 77.61194,
				"lat": 13.05158
			  }
			},
			"place_id": "ola-platform:a1e99e3bd07b545284643f7bde183120"
		  }
		],
		"info_messages": [],
		"error_message": "",
		"status": "ok"
	  `

	PlaceDetailResponse = `{
		"html_attributions": [],
		"result": {
		  "address_components": [
			{
			  "long_name": "India",
			  "short_name": "India",
			  "types": [
				"country"
			  ]
			},
			{
			  "long_name": "KARNATAKA",
			  "short_name": "KARNATAKA",
			  "types": [
				"administrative_area_level_1"
			  ]
			},
			{
			  "long_name": "Bangalore",
			  "short_name": "Bangalore",
			  "types": [
				"administrative_area_level_2"
			  ]
			},
			{
			  "long_name": "Bengaluru South",
			  "short_name": "Bengaluru South",
			  "types": [
				"administrative_area_level_3"
			  ]
			},
			{
			  "long_name": "Bengaluru",
			  "short_name": "Bengaluru",
			  "types": [
				"locality"
			  ]
			},
			{
			  "long_name": "Koramangala",
			  "short_name": "Koramangala",
			  "types": [
				"postal_town"
			  ]
			},
			{
			  "long_name": "Koramangala Industrial Layout",
			  "short_name": "Koramangala Industrial Layout",
			  "types": [
				"sublocality"
			  ]
			},
			{
			  "long_name": "560095",
			  "short_name": "560095",
			  "types": [
				"postal_code"
			  ]
			}
		  ],
		  "formatted_address": "Ola Electric, 2, Hosur Rd, Koramangala Industrial Layout, Koramangala, Bengaluru, Karnataka, 560095, India",
		  "geometry": {
			"location": {
			  "lat": 12.93139,
			  "lng": 77.61647
			},
			"viewport": {
			  "northeast": {
				"lat": 12.932647641397768,
				"lng": 77.61776036672592
			  },
			  "southwest": {
				"lat": 12.930132358602233,
				"lng": 77.6151796332741
			  }
			}
		  },
		  "place_id": "ola-platform:a79ed32419962a11a588ea92b83ca78e",
		  "reference": "a79ed32419962a11a588ea92b83ca78e",
		  "business_status": "NA",
		  "formatted_phone_number": "NA",
		  "icon": "NA",
		  "icon_background_color": "#000000",
		  "icon_mask_base_uri": "https://placeicons.stg.olamaps.io/place_icons/default.png",
		  "international_phone_number": "NA",
		  "name": "Ola Electric",
		  "opening_hours": {
			"open_now": false,
			"periods": [
			  {
				"close": {
				  "day": 0,
				  "time": "NA"
				},
				"open": {
				  "day": 0,
				  "time": "NA"
				}
			  }
			],
			"weekday_text": []
		  },
		  "plus_code": {
			"compound_code": "NA",
			"global_code": "NA"
		  },
		  "rating": 0,
		  "reviews": [
			{
			  "author_name": "NA",
			  "author_url": "NA",
			  "language": "en",
			  "profile_photo_url": "NA",
			  "rating": 0,
			  "relative_time_description": "NA",
			  "text": "NA",
			  "time": 0
			}
		  ],
		  "types": [
			"point_of_interest",
			"establishment"
		  ],
		  "layer": [
			"venue"
		  ],
		  "url": "NA",
		  "user_ratings_total": 0,
		  "utc_offset": 0,
		  "vicinity": "NA",
		  "website": "NA",
		  "price_level": "NA",
		  "photos": [],
		  "adr_address": "NA"
		},
		"info_messages": [],
		"error_message": "",
		"status": "ok"
	  }`

	NearBySearchResponses = `
	  {
		"predictions": [
		  {
			"description": "Chaishai, 18th Cross Road, HSR Layout, Bengaluru, Karnataka, 560102, India",
			"matched_substrings": [],
			"place_id": "ola-platform:5000021987309",
			"reference": "node/5000021987309",
			"structured_formatting": {
			  "main_text": "Chaishai",
			  "main_text_matched_substrings": [],
			  "secondary_text": "18th Cross Road, HSR Layout, Bengaluru, Karnataka, 560102, India",
			  "secondary_text_matched_substrings": []
			},
			"terms": [
			  {
				"offset": 0,
				"value": "Chaishai"
			  },
			  {
				"offset": 10,
				"value": "18th Cross Road"
			  },
			  {
				"offset": 27,
				"value": "HSR Layout"
			  },
			  {
				"offset": 39,
				"value": "Bengaluru"
			  },
			  {
				"offset": 50,
				"value": "Karnataka"
			  },
			  {
				"offset": 61,
				"value": "560102"
			  },
			  {
				"offset": 69,
				"value": "India"
			  }
			],
			"types": [
			  "amenity:cafe"
			],
			"layer": [
			  "venue"
			],
			"distance_meters": 271
		  },
		  {
			"description": "McCafe By McDonald's, 18th Cross Road, HSR Layout, Bengaluru, Karnataka, 560102, India",
			"matched_substrings": [],
			"place_id": "ola-platform:5000015318986",
			"reference": "5000015318986",
			"structured_formatting": {
			  "main_text": "McCafe By McDonald's",
			  "main_text_matched_substrings": [],
			  "secondary_text": "18th Cross Road, HSR Layout, Bengaluru, Karnataka, 560102, India",
			  "secondary_text_matched_substrings": []
			},
			"terms": [
			  {
				"offset": 0,
				"value": "McCafe By McDonald's"
			  },
			  {
				"offset": 22,
				"value": "18th Cross Road"
			  },
			  {
				"offset": 39,
				"value": "HSR Layout"
			  },
			  {
				"offset": 51,
				"value": "Bengaluru"
			  },
			  {
				"offset": 62,
				"value": "Karnataka"
			  },
			  {
				"offset": 73,
				"value": "560102"
			  },
			  {
				"offset": 81,
				"value": "India"
			  }
			],
			"types": [
			  "amenity:cafe"
			],
			"layer": [
			  "venue"
			],
			"distance_meters": 363
		  },
		  {
			"description": "Cafe Irani, 19th Main Road, HSR Layout, Bengaluru, Karnataka, 560102, India",
			"matched_substrings": [],
			"place_id": "ola-platform:5000029415873",
			"reference": "5000029415873",
			"structured_formatting": {
			  "main_text": "Cafe Irani",
			  "main_text_matched_substrings": [],
			  "secondary_text": "19th Main Road, HSR Layout, Bengaluru, Karnataka, 560102, India",
			  "secondary_text_matched_substrings": []
			},
			"terms": [
			  {
				"offset": 0,
				"value": "Cafe Irani"
			  },
			  {
				"offset": 12,
				"value": "19th Main Road"
			  },
			  {
				"offset": 28,
				"value": "HSR Layout"
			  },
			  {
				"offset": 40,
				"value": "Bengaluru"
			  },
			  {
				"offset": 51,
				"value": "Karnataka"
			  },
			  {
				"offset": 62,
				"value": "560102"
			  },
			  {
				"offset": 70,
				"value": "India"
			  }
			],
			"types": [
			  "amenity:cafe"
			],
			"layer": [
			  "venue"
			],
			"distance_meters": 579
		  }
		],
		"info_messages": [],
		"error_message": "",
		"status": "ok"
	  }`

	TextSearchResponse = `{
		"predictions": [
		  {
			"formatted_address": "McDonald's, Unit No 10, AKS Plaza No 10, Jyoti Nivas College Rd, Opposite Forum Mall Musallah Masjid, KHB Colony, Koramangala Industrial Layout, Koramangala, Bengaluru, Karnataka, 560095, India",
			"geometry": {
			  "location": {
				"lat": 12.93538,
				"lng": 77.61545
			  }
			},
			"place_id": "ola-platform:3459f477f482a0431ceef4689f92b3dd",
			"name": "McDonald's",
			"types": [
			  "meal_delivery",
			  "cafe",
			  "restaurant",
			  "food",
			  "point_of_interest",
			  "store",
			  "establishment"
			]
		  },
		  {
			"formatted_address": "McDonald's, Mantri Avenue, KHB Games Village, Koramangala, Bengaluru, Karnataka, 560095, India",
			"geometry": {
			  "location": {
				"lat": 12.94502,
				"lng": 77.62006
			  }
			},
			"place_id": "ola-platform:1bb43a7cce8924b1ea0a0c280ee4e013",
			"name": "McDonald's",
			"types": [
			  "restaurant",
			  "food",
			  "point_of_interest",
			  "establishment"
			]
		  },
		  {
			"formatted_address": "McDonald's, Total Mall, Madivala Commercial Plaza Junction Of Hosur Road, Sarjapur - Marathahalli Rd, Sidharata Colony, Santhosapuram, 2nd Block, Koramangala, Bengaluru, Karnataka, 560034, India",
			"geometry": {
			  "location": {
				"lat": 12.92166,
				"lng": 77.62051
			  }
			},
			"place_id": "ola-platform:7cc658faa423f8fe2c85c90f99dd79fd",
			"name": "McDonald's",
			"types": [
			  "meal_delivery",
			  "cafe",
			  "restaurant",
			  "store",
			  "point_of_interest",
			  "food",
			  "establishment"
			]
		  },
		  {
			"formatted_address": "McCafe By McDonald's, 80 Feet Road, Koramangala, Bengaluru, Karnataka, 560095, India",
			"geometry": {
			  "location": {
				"lat": 12.94556,
				"lng": 77.6199
			  }
			},
			"place_id": "ola:openstreetmap:venue:node/5000014702218",
			"name": "McCafe By McDonald's",
			"types": [
			  "amenity:cafe"
			]
		  },
		  {
			"formatted_address": "McDonald's Head Office, 633, National Games Village Complex, KHB Games Village, Koramangala, Bengaluru, Karnataka, 560095, India",
			"geometry": {
			  "location": {
				"lat": 12.94493,
				"lng": 77.62086
			  }
			},
			"place_id": "ola-platform:04435c3a45c9bf457de1ba4cf7b0fc95",
			"name": "McDonald's Head Office",
			"types": [
			  "restaurant",
			  "food",
			  "point_of_interest",
			  "establishment"
			]
		  }
		],
		"info_messages": [],
		"error_message": "",
		"status": "ok"
	  }`

	MapStyleResponse = `[
		{
		  "version": 8,
		  "name": "Light",
		  "id": "light",
		  "url": "https://api.olamaps.io/tiles/vector/v1/styles/default-light-standard/style.json"
		}
	  ]`

	StyleDetailResponse = `{
		"id": "positron",
		"version": 8,
		"name": "Light",
		"sources": {
		  "openmaptiles": {
			"type": "vector",
			"url": "https://api.olamaps.io/tiles/vector/v1/data/planet.json"
		  }
		},
		"glyphs": "https://api.olamaps.io/tiles/vector/v1/fonts/{fontstack}/{range}.pbf",
		"layers": [
		  {
			"id": "water",
			"type": "fill",
			"source": "openmaptiles",
			"source-layer": "water",
			"filter": [
			  "all",
			  [
				"==",
				"$type",
				"Polygon"
			  ],
			  [
				"!=",
				"brunnel",
				"tunnel"
			  ]
			],
			"layout": {
			  "visibility": "visible"
			},
			"paint": {
			  "fill-antialias": true,
			  "fill-color": "rgb(194, 200, 202)"
			}
		  }
		]
	  }`

	ArrayOfDataResponse = `{
		"attribution": "basename:planet.mbtiles",
		"bounds": [-180, -85.05113, 180, 85.05113],
		"center": [0, 0, 0],
		"compression": "gzip",
		"data_version": "v2.5.8",
		"description": "",
		"format": "pbf",
		"id": "planet",
		"maxzoom": 14,
		"mbtiles_version": "v2.5.8.0",
		"minzoom": 0,
		"name": "Tiles",
		"planetiler": {
		  "buildtime": "2024-06-13T11:00:45.395Z",
		  "githash": "4982e2e133e5bee495aa8a5843bd420e33822bef",
		  "osm": {
			"osmosisreplicationseq": 0,
			"osmosisreplicationtime": "1970-01-01T00:00:00Z",
			"osmosisreplicationurl": ""
		  },
		  "version": "0.7-SNAPSHOT"
		},
		"tilejson": "2.0.0",
		"tiles": [
		  "https://api.olamaps.io/tiles/vector/v1/data/planet/{z}/{x}/{y}.pbf?key="
		],
		"type": "baselayer",
		"vector_layers": [
		  {
			"fields": {
			  "name": "String",
			  "status": "String"
			},
			"id": "aerodrome_label",
			"maxzoom": 14,
			"minzoom": 6
		  },
		  {
			"fields": {
			  "class": "String",
			  "ref": "String"
			},
			"id": "aeroway",
			"maxzoom": 14,
			"minzoom": 10
		  },
		  {
			"fields": {
			  "admin_level": "Number",
			  "disputed": "Number",
			  "maritime": "Number"
			},
			"id": "boundary",
			"maxzoom": 14,
			"minzoom": 0
		  },
		  {
			"fields": {
			  "colour": "String",
			  "hide_3d": "Boolean",
			  "name": "String",
			  "name:hi": "String",
			  "name:kn": "String",
			  "name:ml": "String",
			  "name:mr": "String",
			  "name:ta": "String",
			  "name:te": "String",
			  "render_height": "Number",
			  "render_min_height": "Number"
			},
			"id": "building",
			"maxzoom": 14,
			"minzoom": 13
		  },
		  {
			"fields": {
			  "class": "String",
			  "subclass": "String"
			},
			"id": "landcover",
			"maxzoom": 14,
			"minzoom": 3
		  },
		  {
			"fields": {
			  "class": "String"
			},
			"id": "landuse",
			"maxzoom": 14,
			"minzoom": 4
		  },
		  {
			"fields": {
			  "class": "String",
			  "ele": "Number",
			  "ele_ft": "Number",
			  "name": "String",
			  "name:hi": "String",
			  "name:kn": "String",
			  "name:ml": "String",
			  "name:mr": "String",
			  "name:ta": "String",
			  "name:te": "String",
			  "rank": "Number"
			},
			"id": "mountain_peak",
			"maxzoom": 14,
			"minzoom": 7
		  },
		  {
			"fields": {
			  "class": "String",
			  "name": "String",
			  "name:hi": "String",
			  "name:kn": "String",
			  "name:ml": "String",
			  "name:mr": "String",
			  "name:ta": "String",
			  "name:te": "String",
			  "rank": "Number"
			},
			"id": "park",
			"maxzoom": 14,
			"minzoom": 4
		  },
		  {
			"fields": {
			  "class": "String",
			  "name": "String",
			  "population": "Number"
			},
			"id": "place",
			"maxzoom": 14,
			"minzoom": 1
		  },
		  {
			"fields": {
			  "class": "String",
			  "name": "String",
			  "name:hi": "String",
			  "name:kn": "String",
			  "name:ml": "String",
			  "name:mr": "String",
			  "name:ta": "String",
			  "name:te": "String",
			  "place_id": "String",
			  "pop": "Number",
			  "popularity": "Number",
			  "shortName": "String",
			  "subclass": "Stri
	  `
)
