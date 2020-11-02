package main

var oidToType = map[uint32]string{
	16:   "bool",
	17:   "bytea",
	18:   "char",
	20:   "int8",
	21:   "int2",
	23:   "int4",
	25:   "text",
	114:  "json",
	142:  "xml",
	600:  "point",
	601:  "lseg",
	602:  "path",
	603:  "box",
	604:  "polygon",
	628:  "line",
	700:  "float4",
	701:  "float8",
	718:  "circle",
	790:  "money",
	829:  "macaddr",
	869:  "inet",
	650:  "cidr",
	774:  "macaddr8",
	1042: "bpchar",
	1043: "varchar",
	1082: "date",
	1083: "time",
	1114: "timestamp",
	1184: "timestamptz",
	1186: "interval",
	1266: "timetz",
	1560: "bit",
	1562: "varbit",
	1700: "numeric",
	1790: "refcursor",
	2950: "uuid",
	3220: "pg_lsn",
	3614: "tsvector",
	3642: "gtsvector",
	3615: "tsquery",
	3802: "jsonb",
	4072: "jsonpath",
	3904: "int4range",
	3906: "numrange",
	3908: "tsrange",
	3910: "tstzrange",
	3912: "daterange",
	3926: "int8range",
	2249: "record",
	2287: "_record",
	2275: "cstring",
	2276: "any",
	2277: "anyarray",
	2278: "void",
	2280: "language_handler",
	2282: "opaque",
	2283: "anyelement",
	2776: "anynonarray",
	3500: "anyenum",
	3115: "fdw_handler",
	325:  "index_am_handler",
	3310: "tsm_handler",
	269:  "table_am_handler",
	3831: "anyrange",
	1000: "_bool",
	1001: "_bytea",
	1002: "_char",
	1003: "_name",
	1016: "_int8",
	1005: "_int2",
	1006: "_int2vector",
	1007: "_int4",
	1008: "_regproc",
	1009: "_text",
	1028: "_oid",
	1010: "_tid",
	1011: "_xid",
	1012: "_cid",
	1013: "_oidvector",
	199:  "_json",
	143:  "_xml",
	1017: "_point",
	1018: "_lseg",
	1019: "_path",
	1020: "_box",
	1027: "_polygon",
	629:  "_line",
	1021: "_float4",
	1022: "_float8",
	719:  "_circle",
	791:  "_money",
	1040: "_macaddr",
	1041: "_inet",
	651:  "_cidr",
	775:  "_macaddr8",
	1034: "_aclitem",
	1014: "_bpchar",
	1015: "_varchar",
	1182: "_date",
	1183: "_time",
	1115: "_timestamp",
	1185: "_timestamptz",
	1187: "_interval",
	1270: "_timetz",
	1561: "_bit",
	1563: "_varbit",
	1231: "_numeric",
	2201: "_refcursor",
	2207: "_regprocedure",
	2208: "_regoper",
	2209: "_regoperator",
	2210: "_regclass",
	2211: "_regtype",
	4097: "_regrole",
	4090: "_regnamespace",
	2951: "_uuid",
	3221: "_pg_lsn",
	3643: "_tsvector",
	3644: "_gtsvector",
	3645: "_tsquery",
	3735: "_regconfig",
	3770: "_regdictionary",
	3807: "_jsonb",
	4073: "_jsonpath",
	2949: "_txid_snapshot",
	3905: "_int4range",
	3907: "_numrange",
	3909: "_tsrange",
	3911: "_tstzrange",
	3913: "_daterange",
	3927: "_int8range",
	1263: "_cstring",
}