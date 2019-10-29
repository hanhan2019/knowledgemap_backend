package main

import (
	"encoding/json"
	"fmt"
)

type Model struct {
	Content Content `json:"@content"`
}

type Content struct {
	Class         string
	Collection    string
	Concept       string
	NodeShape     string
	Relation      string
	Affects       string
	BelongsTo     string
	Broader       string
	Ckm           string
	Contains      string
	Datatype      string
	Dcterms       string
	Derives       string
	Equals        string
	Id            string
	Inherits      string
	IsDependentOn string
	IsSubjectOf   string
	Label         string
	Language      string
	MaxCount      string
	Member        string
	MinCount      string
	Narrower      string
	Or            string
	Owl           string
	Path          string
	Property      string
	Rdf           string
	Rdfs          string
	Related       string
	Sh            string
	Skos          string
	TargetClass   string
	Type          string
	Value         string
	Weight        string
	Xkos          string
	Xsd           string
}

func main() {
	knowledgeMap := map[string]interface{}{
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

	graph := []map[string]interface{}{
		map[string]interface{}{
			"@id":         "local:1",
			"isSubjectOf": "数组",
			"label":       "数据结构",
			"type":        "Concept",
			"weight":      "1",
			"affects": []map[string]interface{}{
				map[string]interface{}{
					"id": "local:1",
				},
			},
		},
		map[string]interface{}{
			"id":          "local:1",
			"isSubjectOf": "数组1",
			"label":       "数据结构1",
			"type":        "Concept",
			"weight":      "2",
		},
	}
	knowledgeMap["@graph"] = append(knowledgeMap["@graph"].([]map[string]interface{}), graph...)
	// a := "{\"@context\":{\"owl\":\"http://www.w3.org/2002/07/owl#\",\"Class\":\"owl:Class\",\"Concept\":\"skos:Concept\",\"Collection\":\"skos:Collection\",\"Relation\":\"rdf:Property\",\"affects\":\"dcterms:isRequiredBy\",\"belongsTo\":\"dcterms:isPartOf\",\"ckm\":\"https://github.com/Atuno/CKM-Toolkit/blob/master/jsonld/\",\"contains\":\"dcterms:hasPart\",\"dcterms\":\"http://purl.org/dc/terms/\",\"derives\":\"xkos:specializes\",\"id\":\"@id\",\"inherits\":\"xkos:generalizes\",\"isDependentOn\":\"dcterms:requires\",\"label\":\"rdfs:label\",\"language\":\"@language\",\"rdf\":\"http://www.w3.org/1999/02/22-rdf-syntax-ns#\",\"rdfs\":\"http://www.w3.org/2000/01/rdf-schema#\",\"skos\":\"http://www.w3.org/2004/02/skos/core#\",\"type\":\"@type\",\"value\":\"@value\",\"xkos\":\"http://rdf-vocabulary.ddialliance.org/xkos#\",\"related\":\"skos:related\",\"broader\":\"skos:broader\",\"narrower\":\"skos:narrower\",\"member\":\"skos:member\",\"sh\":\"http://www.w3.org/ns/shacl#\",\"NodeShape\":\"sh:NodeShape\",\"targetClass\":\"sh:targetClass\",\"property\":\"sh:property\",\"path\":\"sh:path\",\"maxCount\":\"sh:maxCount\",\"minCount\":\"sh:minCount\",\"datatype\":\"sh:datatype\",\"or\":\"sh:or\",\"equals\":\"sh:equals\",\"xsd\":\"http://www.w3.org/2001/XMLSchema#\",\"isSubjectOf\":\"skos:isSubjectOf\",\"weight\":\"rdf:value\"},\"@graph\":[{\"id\":\"dcterms:hasPart\",\"label\":{\"language\":\"zh\",\"value\":\"包含\"},\"type\":\"Relation\"},{\"id\":\"dcterms:isRequiredBy\",\"label\":{\"language\":\"zh\",\"value\":\"影响\"},\"type\":\"Relation\"},{\"id\":\"xkos:specializes\",\"label\":{\"language\":\"zh\",\"value\":\"派生\"},\"type\":\"Relation\"},{\"derives\":[{\"id\":\"ckm:privateMemberAccess\"},{\"id\":\"ckm:protectedMemberAccess\"},{\"id\":\"ckm:publicMemberAccess\"}],\"id\":\"ckm:accessSpecifier\",\"label\":[{\"language\":\"en\",\"value\":\"access specifier\"},{\"language\":\"zh\",\"value\":\"访问控制属性\"}],\"type\":\"Concept\"},{\"contains\":[{\"id\":\"ckm:member\"},{\"id\":\"ckm:composition\"},{\"id\":\"ckm:object\"},{\"id\":\"ckm:friend\"}],\"id\":\"ckm:class\",\"label\":[{\"language\":\"en\",\"value\":\"Class\"},{\"language\":\"zh\",\"value\":\"类\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:initializerList\"},\"id\":\"ckm:composition\",\"label\":[{\"language\":\"en\",\"value\":\"Composition\"},{\"language\":\"zh\",\"value\":\"类的组合\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:initializerList\"},\"derives\":[{\"id\":\"ckm:defaultConstructor\"},{\"id\":\"ckm:copyConstructor\"}],\"id\":\"ckm:constructor\",\"label\":[{\"language\":\"en\",\"value\":\"constructor\"},{\"language\":\"zh\",\"value\":\"构造函数\"}],\"type\":\"Concept\"},{\"contains\":[{\"id\":\"ckm:deepCopy\"},{\"id\":\"ckm:shallowCopy\"}],\"id\":\"ckm:copyConstructor\",\"label\":[{\"language\":\"en\",\"value\":\"copy constructor\"},{\"language\":\"zh\",\"value\":\"拷贝构造函数\"}],\"type\":\"Concept\"},{\"derives\":{\"id\":\"ckm:staticDataMember\"},\"id\":\"ckm:dataMember\",\"label\":[{\"language\":\"en\",\"value\":\"data member\"},{\"language\":\"zh\",\"value\":\"数据成员\"}],\"type\":\"Concept\"},{\"id\":\"ckm:deepCopy\",\"label\":[{\"language\":\"en\",\"value\":\"deep copy\"},{\"language\":\"zh\",\"value\":\"深拷贝\"}],\"type\":\"Concept\"},{\"id\":\"ckm:defaultConstructor\",\"label\":[{\"language\":\"en\",\"value\":\"default constructor\"},{\"language\":\"zh\",\"value\":\"缺省构造函数\"}],\"type\":\"Concept\"},{\"id\":\"ckm:destructor\",\"label\":[{\"language\":\"en\",\"value\":\"destructor\"},{\"language\":\"zh\",\"value\":\"析构函数\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:friendFunction\"},\"id\":\"ckm:friend\",\"label\":[{\"language\":\"en\",\"value\":\"friend\"},{\"language\":\"zh\",\"value\":\"友元\"}],\"type\":\"Concept\"},{\"id\":\"ckm:friendFunction\",\"label\":[{\"language\":\"en\",\"value\":\"friend function\"},{\"language\":\"zh\",\"value\":\"友元函数\"}],\"type\":\"Concept\"},{\"id\":\"ckm:initializerList\",\"label\":[{\"language\":\"en\",\"value\":\"initializer list\"},{\"language\":\"zh\",\"value\":\"初始化列表\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:accessSpecifier\"},\"derives\":[{\"id\":\"ckm:memberFunction\"},{\"id\":\"ckm:dataMember\"}],\"id\":\"ckm:member\",\"label\":[{\"language\":\"en\",\"value\":\"member\"},{\"language\":\"zh\",\"value\":\"成员\"}],\"type\":\"Concept\"},{\"derives\":[{\"id\":\"ckm:constructor\"},{\"id\":\"ckm:destructor\"},{\"id\":\"ckm:staticMemberFunction\"}],\"id\":\"ckm:memberFunction\",\"label\":[{\"language\":\"en\",\"value\":\"member function\"},{\"language\":\"zh\",\"value\":\"成员函数\"}],\"type\":\"Concept\"},{\"affects\":{\"id\":\"ckm:deepCopy\"},\"id\":\"ckm:memoryLeak\",\"label\":[{\"language\":\"en\",\"value\":\"memory leak\"},{\"language\":\"zh\",\"value\":\"内存泄漏\"}],\"type\":\"Concept\"},{\"contains\":[{\"id\":\"ckm:objectArray\"},{\"id\":\"ckm:thisPointer\"}],\"id\":\"ckm:object\",\"label\":[{\"language\":\"en\",\"value\":\"Object\"},{\"language\":\"zh\",\"value\":\"对象\"}],\"type\":\"Concept\"},{\"id\":\"ckm:objectArray\",\"label\":[{\"language\":\"en\",\"value\":\"object array\"},{\"language\":\"zh\",\"value\":\"对象数组\"}],\"type\":\"Concept\"},{\"id\":\"ckm:privateMemberAccess\",\"label\":[{\"language\":\"en\",\"value\":\"private member access\"},{\"language\":\"zh\",\"value\":\"私有访问\"}],\"type\":\"Concept\"},{\"id\":\"ckm:protectedMemberAccess\",\"label\":[{\"language\":\"en\",\"value\":\"protected member access\"},{\"language\":\"zh\",\"value\":\"保护访问\"}],\"type\":\"Concept\"},{\"id\":\"ckm:publicMemberAccess\",\"label\":[{\"language\":\"en\",\"value\":\"public member access\"},{\"language\":\"zh\",\"value\":\"公有访问\"}],\"type\":\"Concept\"},{\"id\":\"ckm:shallowCopy\",\"label\":[{\"language\":\"en\",\"value\":\"shallow copy\"},{\"language\":\"zh\",\"value\":\"浅拷贝\"}],\"type\":\"Concept\"},{\"id\":\"ckm:staticDataMember\",\"label\":[{\"language\":\"en\",\"value\":\"static member variable\"},{\"language\":\"zh\",\"value\":\"静态数据成员\"}],\"type\":\"Concept\"},{\"id\":\"ckm:staticMemberFunction\",\"label\":[{\"language\":\"en\",\"value\":\"static member function\"},{\"language\":\"zh\",\"value\":\"静态成员函数\"}],\"type\":\"Concept\"},{\"id\":\"ckm:thisPointer\",\"label\":[{\"language\":\"en\",\"value\":\"this pointer\"},{\"language\":\"zh\",\"value\":\"this指针\"}],\"type\":\"Concept\"},{\"id\":\"ckm:summaryOfOOP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"面向对象程序设计概述\"}]},{\"id\":\"ckm:summaryOfCPP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"C++概述\"}],\"member\":[{\"id\":\"ckm:memoryLeak\",\"type\":\"Concept\"}]},{\"id\":\"ckm:classAndObject\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"类和对象\"}],\"member\":[{\"id\":\"ckm:accessSpecifier\",\"type\":\"Concept\"},{\"id\":\"ckm:class\",\"type\":\"Concept\"},{\"id\":\"ckm:composition\",\"type\":\"Concept\"},{\"id\":\"ckm:constructor\",\"type\":\"Concept\"},{\"id\":\"ckm:copyConstructor\",\"type\":\"Concept\"},{\"id\":\"ckm:dataMember\",\"type\":\"Concept\"},{\"id\":\"ckm:deepCopy\",\"type\":\"Concept\"},{\"id\":\"ckm:defaultConstructor\",\"type\":\"Concept\"},{\"id\":\"ckm:destructor\",\"type\":\"Concept\"},{\"id\":\"ckm:friend\",\"type\":\"Concept\"},{\"id\":\"ckm:friendFunction\",\"type\":\"Concept\"},{\"id\":\"ckm:initializerList\",\"type\":\"Concept\"},{\"id\":\"ckm:member\",\"type\":\"Concept\"},{\"id\":\"ckm:memberFunction\",\"type\":\"Concept\"},{\"id\":\"ckm:object\",\"type\":\"Concept\"},{\"id\":\"ckm:objectArray\",\"type\":\"Concept\"},{\"id\":\"ckm:privateMemberAccess\",\"type\":\"Concept\"},{\"id\":\"ckm:protectedMemberAccess\",\"type\":\"Concept\"},{\"id\":\"ckm:publicMemberAccess\",\"type\":\"Concept\"},{\"id\":\"ckm:shallowCopy\",\"type\":\"Concept\"},{\"id\":\"ckm:staticDataMember\",\"type\":\"Concept\"},{\"id\":\"ckm:staticMemberFunction\",\"type\":\"Concept\"},{\"id\":\"ckm:thisPointer\",\"type\":\"Concept\"}]},{\"id\":\"ckm:derivedClassAndInheritance\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"派生类和继承\"}]},{\"id\":\"ckm:polymorphism\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"多态性\"}]},{\"id\":\"ckm:templateAndException\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"模板与异常处理\"}]},{\"id\":\"ckm:streamAndIO\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"C++的流类库与输入输出\"}]},{\"id\":\"ckm:designOfOOP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"面向对象程序设计方法与实例\"}]},{\"id\":\"ckm:OOP_CPP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"面向对象程序设计（C++）\"}],\"member\":[{\"id\":\"ckm:summaryOfOOP\",\"type\":\"Collection\"},{\"id\":\"ckm:summaryOfCPP\",\"type\":\"Collection\"},{\"id\":\"ckm:classAndObject\",\"type\":\"Collection\"},{\"id\":\"ckm:derivedClassAndInheritance\",\"type\":\"Collection\"},{\"id\":\"ckm:polymorphism\",\"type\":\"Collection\"},{\"id\":\"ckm:templateAndException\",\"type\":\"Collection\"},{\"id\":\"ckm:streamAndIO\",\"type\":\"Collection\"},{\"id\":\"ckm:designOfOOP\",\"type\":\"Collection\"}]},{\"type\":\"NodeShape\",\"id\":\"ckm:conceptShape\",\"targetClass\":\"skos:Concept\",\"property\":[{\"path\":\"@id\",\"minCount\":1},{\"path\":\"@type\",\"minCount\":1},{\"path\":\"rdfs:label\",\"label\":\"名称\",\"datatype\":\"xsd:string\",\"minCount\":1},{\"path\":\"skos:isSubjectOf\",\"label\":\"\",\"datatype\":\"xsd:string\",\"minCount\":0},{\"path\":\"rdf:value\",\"label\":\"\",\"datatype\":\"xsd:integer\",\"minCount\":0}]},{\"type\":\"NodeShape\",\"id\":\"ckm:collectionShape\",\"targetClass\":\"skos:Collection\",\"property\":[{\"path\":\"@id\",\"minCount\":1},{\"path\":\"@type\",\"minCount\":1},{\"path\":\"rdfs:label\",\"label\":\"名称\",\"minCount\":1,\"datatype\":\"xsd:string\"},{\"path\":\"skos:member\"}]},{\"type\":\"NodeShape\",\"id\":\"ckm:relationShape\",\"targetClass\":\"rdf:Property\",\"property\":[{\"path\":\"@id\",\"or\":[{\"equals\":\"dcterms:hasPart\"},{\"equals\":\"dcterms:isRequiredBy\"},{\"equals\":\"xkos:specializes\"},{\"equals\":\"xkos:generalizes\"},{\"equals\":\"dcterms:isPartOf\"},{\"equals\":\"dcterms:requires\"},{\"equals\":\"skos:related\"},{\"equals\":\"skos:narrower\"}]}]}]}"
	// var mapResult map[string]interface{}
	// err := json.Unmarshal([]byte(a), &mapResult)
	// if err != nil {
	// 	fmt.Println("JsonToMapDemo err: ", err)
	// }
	// fmt.Println(mapResult)
	// fmt.Println(mapResult["@graph"][])
	// fmt.Println(knowledgeMap["@graph"])
	jsonStr, err := json.Marshal(knowledgeMap)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}

// map[@context:
// 	map[
// 		NodeShape:sh:NodeShape
// 		or:sh:or
// 		Collection:skos:Collection
// 		xkos:http://rdf-vocabulary.ddialliance.org/xkos#
// 		narrower:skos:narrower
// 		inherits:xkos:generalizes
// 		equals:sh:equals
// 		xsd:http://www.w3.org/2001/XMLSchema#
// 		skos:http://www.w3.org/2004/02/skos/core#
// 		broader:skos:broader
// 		Class:owl:Class dcterms:http://purl.org/dc/terms/
// 		label:rdfs:label
// 		value:@value
// 		targetClass:sh:targetClass
// 		isSubjectOf:skos:isSubjectOf
// 		ckm:https://github.com/Atuno/CKM-Toolkit/blob/master/jsonld/
// 		derives:xkos:specializes
// 		isDependentOn:dcterms:requires
// 		rdfs:http://www.w3.org/2000/01/rdf-schema#
// 		member:skos:member
// 		path:sh:path
// 		minCount:sh:minCount
// 		Relation:rdf:Property
// 		affects:dcterms:isRequiredBy
// 		contains:dcterms:hasPart
// 		related:skos:related
// 		datatype:sh:datatype
// 		id:@id
// 		language:@language
// 		rdf:http://www.w3.org/1999/02/22-rdf-syntax-ns#
// 		owl:http://www.w3.org/2002/07/owl#
// 		Concept:skos:Concept
// 		belongsTo:dcterms:isPartOf
// 		maxCount:sh:maxCount
// 		weight:rdf:value
// 		type:@type
// 		sh:http://www.w3.org/ns/shacl#
// 		property:sh:property
// 		]
//   @graph:[
// 	map[
// 		label:map[language:zh value:包含] type:Relation id:dcterms:hasPart] map[label:map[language:zh value:影响] type:Relation id:dcterms:isRequiredBy] map[id:xkos:specializes label:map[language:zh value:派生] type:Relation] map[derives:[map[id:ckm:privateMemberAccess] map[id:ckm:protectedMemberAccess] map[id:ckm:publicMemberAccess]] id:ckm:accessSpecifier label:[map[value:access specifier language:en] map[value:访问控制属性 language:zh]] type:Concept] map[contains:[map[id:ckm:member] map[id:ckm:composition] map[id:ckm:object] map[id:ckm:friend]] id:ckm:class label:[map[language:en value:Class] map[language:zh value:类]] type:Concept] map[contains:map[id:ckm:initializerList] id:ckm:composition label:[map[language:en value:Composition] map[language:zh value:类的组合]] type:Concept] map[derives:[map[id:ckm:defaultConstructor] map[id:ckm:copyConstructor]] id:ckm:constructor label:[map[language:en value:constructor] map[language:zh value:构造函数]] type:Concept contains:map[id:ckm:initializerList]] map[label:[map[language:en value:copy constructor] map[language:zh value:拷贝构造函数]] type:Concept contains:[map[id:ckm:deepCopy] map[id:ckm:shallowCopy]] id:ckm:copyConstructor] map[derives:map[id:ckm:staticDataMember] id:ckm:dataMember label:[map[language:en value:data member] map[language:zh value:数据成员]] type:Concept] map[id:ckm:deepCopy label:[map[language:en value:deep copy] map[language:zh value:深拷贝]] type:Concept] map[label:[map[language:en value:default constructor] map[language:zh value:缺省构造函数]] type:Concept id:ckm:defaultConstructor] map[type:Concept id:ckm:destructor label:[map[language:en value:destructor] map[language:zh value:析构函数]]] map[type:Concept contains:map[id:ckm:friendFunction] id:ckm:friend label:[map[language:en value:friend] map[language:zh value:友元]]] map[id:ckm:friendFunction label:[map[language:en value:friend function] map[language:zh value:友元函数]] type:Concept] map[id:ckm:initializerList label:[map[value:initializer list language:en] map[language:zh value:初始化列表]] type:Concept] map[contains:map[id:ckm:accessSpecifier] derives:[map[id:ckm:memberFunction] map[id:ckm:dataMember]] id:ckm:member label:[map[language:en value:member] map[language:zh value:成员]] type:Concept] map[derives:[map[id:ckm:constructor] map[id:ckm:destructor] map[id:ckm:staticMemberFunction]] id:ckm:memberFunction label:[map[language:en value:member function] map[language:zh value:成员函数]] type:Concept] map[id:ckm:memoryLeak label:[map[language:en value:memory leak] map[language:zh value:内存泄漏]] type:Concept affects:map[id:ckm:deepCopy]] map[label:[map[language:en value:Object] map[language:zh value:对象]] type:Concept contains:[map[id:ckm:objectArray] map[id:ckm:thisPointer]] id:ckm:object] map[id:ckm:objectArray label:[map[language:en value:object array] map[language:zh value:对象数组]] type:Concept] map[id:ckm:privateMemberAccess label:[map[language:en value:private member access] map[language:zh value:私有访问]] type:Concept] map[id:ckm:protectedMemberAccess label:[map[language:en value:protected member access] map[language:zh value:保护访问]] type:Concept] map[label:[map[value:public member access language:en] map[language:zh value:公有访问]] type:Concept id:ckm:publicMemberAccess] map[id:ckm:shallowCopy label:[map[language:en value:shallow copy] map[language:zh value:浅拷贝]] type:Concept] map[id:ckm:staticDataMember label:[map[language:en value:static member variable] map[value:静态数据成员 language:zh]] type:Concept] map[type:Concept id:ckm:staticMemberFunction label:[map[language:en value:static member function] map[language:zh value:静态成员函数]]] map[label:[map[value:this pointer language:en] map[language:zh value:this指针]] type:Concept id:ckm:thisPointer] map[id:ckm:summaryOfOOP type:Collection label:[map[language:zh value:面向对象程序设计概述]]] map[label:[map[language:zh value:C++概述]] member:[map[id:ckm:memoryLeak type:Concept]] id:ckm:summaryOfCPP type:Collection] map[member:[map[id:ckm:accessSpecifier type:Concept] map[id:ckm:class type:Concept] map[id:ckm:composition type:Concept] map[type:Concept id:ckm:constructor] map[id:ckm:copyConstructor type:Concept] map[type:Concept id:ckm:dataMember] map[id:ckm:deepCopy type:Concept] map[id:ckm:defaultConstructor type:Concept] map[id:ckm:destructor type:Concept] map[id:ckm:friend type:Concept] map[id:ckm:friendFunction type:Concept] map[id:ckm:initializerList type:Concept] map[type:Concept id:ckm:member] map[id:ckm:memberFunction type:Concept] map[id:ckm:object type:Concept] map[id:ckm:objectArray type:Concept] map[id:ckm:privateMemberAccess type:Concept] map[id:ckm:protectedMemberAccess type:Concept] map[id:ckm:publicMemberAccess type:Concept] map[id:ckm:shallowCopy type:Concept] map[id:ckm:staticDataMember type:Concept] map[id:ckm:staticMemberFunction type:Concept] map[type:Concept id:ckm:thisPointer]] id:ckm:classAndObject type:Collection label:[map[language:zh value:类和对象]]] map[label:[map[language:zh value:派生类和继承]] id:ckm:derivedClassAndInheritance type:Collection] map[id:ckm:polymorphism type:Collection label:[map[language:zh value:多态性]]] map[id:ckm:templateAndException type:Collection label:[map[language:zh value:模板与异常处理]]] map[id:ckm:streamAndIO type:Collection label:[map[language:zh value:C++的流类库与输入输出]]] map[id:ckm:designOfOOP type:Collection label:[map[language:zh value:面向对象程序设计方法与实例]]] map[id:ckm:OOP_CPP type:Collection label:[map[language:zh value:面向对象程序设计（C++）]] member:[map[id:ckm:summaryOfOOP type:Collection] map[id:ckm:summaryOfCPP type:Collection] map[id:ckm:classAndObject type:Collection] map[type:Collection id:ckm:derivedClassAndInheritance] map[id:ckm:polymorphism type:Collection] map[id:ckm:templateAndException type:Collection] map[type:Collection id:ckm:streamAndIO] map[type:Collection id:ckm:designOfOOP]]] map[type:NodeShape id:ckm:conceptShape targetClass:skos:Concept property:[map[path:@id minCount:1] map[path:@type minCount:1] map[label:名称 datatype:xsd:string minCount:1 path:rdfs:label] map[datatype:xsd:string minCount:0 path:skos:isSubjectOf label:] map[datatype:xsd:integer minCount:0 path:rdf:value label:]]] map[type:NodeShape id:ckm:collectionShape targetClass:skos:Collection property:[map[path:@id minCount:1] map[minCount:1 path:@type] map[path:rdfs:label label:名称 minCount:1 datatype:xsd:string] map[path:skos:member]]] map[type:NodeShape id:ckm:relationShape targetClass:rdf:Property property:[map[path:@id or:[map[equals:dcterms:hasPart] map[equals:dcterms:isRequiredBy] map[equals:xkos:specializes] map[equals:xkos:generalizes] map[equals:dcterms:isPartOf] map[equals:dcterms:requires] map[equals:skos:related] map[equals:skos:narrower]]]]]]]
