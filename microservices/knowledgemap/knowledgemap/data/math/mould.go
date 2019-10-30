package math

func GetMathKnowledgeMap() map[string]interface{} {
	return map[string]interface{}{
		"@context": map[string]interface{}{
			"Class":         "owl:Class",
			"Collection":    "skos:Collection",
			"Concept":       "skos:Concept",
			"NodeShape":     "sh:NodeShape",
			"Relation":      "rdf:Property",
			"affects":       "dcterms:isRequiredBy",
			"belongsTo":     "dcterms:isPartOf",
			"broader":       "skos:broader",
			"ckm":           "https://github.com/Atuno/CKM-Toolkit/blob/master/jsonld/",
			"contains":      "dcterms:hasPart",
			"datatype":      "sh:datatype",
			"dcterms":       "http://purl.org/dc/terms/",
			"derives":       "xkos:specializes",
			"equals":        "sh:equals",
			"id":            "@id",
			"inherits":      "xkos:generalizes",
			"isDependentOn": "dcterms:requires",
			"isSubjectOf":   "skos:isSubjectOf",
			"label":         "rdfs:label",
			"language":      "@language",
			"maxCount":      "sh:maxCount",
			"member":        "skos:member",
			"minCount":      "sh:minCount",
			"narrower":      "skos:narrower",
			"or":            "sh:or",
			"owl":           "http://www.w3.org/2002/07/owl#",
			"path":          "sh:path",
			"property":      "sh:property",
			"rdf":           "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
			"rdfs":          "http://www.w3.org/2000/01/rdf-schema#",
			"related":       "skos:related",
			"sh":            "http://www.w3.org/ns/shacl#",
			"skos":          "http://www.w3.org/2004/02/skos/core#",
			"targetClass":   "sh:targetClass",
			"type":          "@type",
			"value":         "@value",
			"weight":        "rdf:value",
			"xkos":          "http://rdf-vocabulary.ddialliance.org/xkos#",
			"xsd":           "http://www.w3.org/2001/XMLSchema#",
		},
		"@graph": []map[string]interface{}{
			// 三种联系的主体定义
			map[string]interface{}{
				"id": "dcterms:hasPart",
				"label": map[string]interface{}{
					"language": "zh",
					"value":    "包含",
				},
				"type": "Relation",
			},
			map[string]interface{}{
				"id": "dcterms:isRequiredBy",
				"label": map[string]interface{}{
					"language": "zh",
					"value":    "影响",
				},
				"type": "Relation",
			},
			map[string]interface{}{
				"id": "xkos:specializes",
				"label": map[string]interface{}{
					"language": "zh",
					"value":    "派生",
				},
				"type": "Relation",
			},
			// 三种联系的表现形式的属性定义，前段用
			map[string]interface{}{
				"id": "ckm:conceptShape",
				"property": []map[string]interface{}{
					map[string]interface{}{
						"minCount": 1,
						"path":     "@id",
					},
					map[string]interface{}{
						"minCount": 1,
						"path":     "@type",
					},
					map[string]interface{}{
						"datatype": "xsd:string",
						"label":    "名称",
						"minCount": 1,
						"path":     "rdfs:label",
					},
					map[string]interface{}{
						"datatype": "xsd:string",
						"label":    "资源",
						"minCount": 0,
						"path":     "skos:isSubjectOf",
					},
					map[string]interface{}{
						"datatype": "xsd:integer",
						"label":    "值",
						"minCount": 0,
						"path":     "rdf:value",
					},
				},
				"targetClass": "skos:Concept",
				"type":        "NodeShape",
			},
			map[string]interface{}{
				"id": "ckm:collectionShape",
				"property": []map[string]interface{}{
					map[string]interface{}{
						"minCount": 1,
						"path":     "@id",
					},
					map[string]interface{}{
						"minCount": 1,
						"path":     "@type",
					},
					map[string]interface{}{
						"datatype": "xsd:string",
						"label":    "名称",
						"minCount": 1,
						"path":     "rdfs:label",
					},
					map[string]interface{}{
						"path": "skos:member",
					},
				},
				"targetClass": "skos:Collection",
				"type":        "NodeShape",
			},
			map[string]interface{}{
				"id": "ckm:relationShape",
				"property": map[string]interface{}{
					"or": []map[string]interface{}{
						map[string]interface{}{
							"equals": "dcterms:hasPart",
						},
						map[string]interface{}{
							"equals": "dcterms:isRequiredBy",
						},
						map[string]interface{}{
							"equals": "xkos:specializes",
						},
						map[string]interface{}{
							"equals": "xkos:generalizes",
						},
						map[string]interface{}{
							"equals": "dcterms:isPartOf",
						},
						map[string]interface{}{
							"equals": "dcterms:requires",
						},
						map[string]interface{}{
							"equals": "skos:related",
						},
						map[string]interface{}{
							"equals": "skos:narrower",
						},
					},
					"path": "@id",
				},
				"targetClass": "rdf:Property",
				"type":        "NodeShape",
			},
			//以下是自己定义的数据的信息

		},
	}
}
