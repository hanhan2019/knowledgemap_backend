package dao

import (
	"time"
)

// func GetLoginTokenRedisKey(uid string) string {
// 	return fmt.Sprintf(\"passport:logintoken:%v\", uid)
// }
// func (d *Dao) CheckIDCardInStudent(ctx context.Context, idCard string) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.STUDENT_COLLECTION_NAME)
// 	if cnt, err := col.Find(bson.M{\"idcard\": idCard}).Count(); err != nil {
// 		return err
// 	} else if cnt > 0 {
// 		return errors.New(\"errors.idcard-duplicated\")
// 	}
// 	return nil
// }

// func (d *Dao) CheckAccountInStudent(ctx context.Context, account string) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.STUDENT_COLLECTION_NAME)
// 	if cnt, err := col.Find(bson.M{\"account\": account}).Count(); err != nil {
// 		return err
// 	} else if cnt > 0 {
// 		return errors.New(\"errors.idcard-duplicated\")
// 	}
// 	return nil
// }

// func (d *Dao) CheckIDCardInTeacher(ctx context.Context, idCard string) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.TEACHER_COLLECTION_NAME)
// 	if cnt, err := col.Find(bson.M{\"idcard\": idCard}).Count(); err != nil {
// 		return err
// 	} else if cnt > 0 {
// 		return errors.New(\"errors.idcard-duplicated\")
// 	}
// 	return nil
// }
// func (d *Dao) CheckAccountInTeacher(ctx context.Context, account string) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.TEACHER_COLLECTION_NAME)
// 	if cnt, err := col.Find(bson.M{\"account\": account}).Count(); err != nil {
// 		return err
// 	} else if cnt > 0 {
// 		return errors.New(\"errors.idcard-duplicated\")
// 	}
// 	return nil
// }

// func (d *Dao) CheckIDCardInSecretary(ctx context.Context, idCard string) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.SECRETARY_COLLECTION_NAME)
// 	if cnt, err := col.Find(bson.M{\"idcard\": idCard}).Count(); err != nil {
// 		return err
// 	} else if cnt > 0 {
// 		return errors.New(\"errors.idcard-duplicated\")
// 	}
// 	return nil
// }

// func (d *Dao) CheckAccountInSecretary(ctx context.Context, account string) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	col := db.C(model.SECRETARY_COLLECTION_NAME)
// 	if cnt, err := col.Find(bson.M{\"account\": account}).Count(); err != nil {
// 		return err
// 	} else if cnt > 0 {
// 		return errors.New(\"errors.idcard-duplicated\")
// 	}
// 	return nil
// }

// func (d *Dao) newStudent(ctx context.Context, student *model.Student) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	genIdCol := mongo.NewNumberCollection(db.C(pmodel.GEN_ID_COLLECTION_NAME))
// 	//fill username
// 	studentId, _ := genIdCol.Inc(pmodel.GEN_STUDENT_ID_KEY_NAME)
// 	student.Number = fmt.Sprintf(\"%v\", studentId)
// 	col := db.C(model.STUDENT_COLLECTION_NAME)
// 	return col.Insert(student)
// }

// func (d *Dao) newTeacher(ctx context.Context, teacher *model.Teacher) error {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	genIdCol := mongo.NewNumberCollection(db.C(pmodel.GEN_ID_COLLECTION_NAME))
// 	//fill username
// 	teacherId, _ := genIdCol.Inc(pmodel.GEN_STUDENT_ID_KEY_NAME)
// 	teacher.Number = fmt.Sprintf(\"%v\", teacherId)
// 	return db.C(model.TEACHER_COLLECTION_NAME).Insert(teacher)
// }

// func createDefaultStudent(req *api.RegisterReq) *model.Student {
// 	user := new(model.Student)
// 	user.ID = bson.NewObjectId()
// 	user.Name = req.Name
// 	user.Major = req.Major
// 	user.IDCard = req.Idcard
// 	user.Account = req.Account
// 	user.Password = req.Password
// 	user.Origin = req.Origin
// 	user.Class = req.Class
// 	user.College = req.College
// 	user.AdmissionTime = req.Admissontime
// 	user.CreateTime = time.Now().Unix()
// 	return user
// }

// func createDefaultTeacher(req *api.RegisterReq) *model.Teacher {
// 	user := new(model.Teacher)
// 	user.ID = bson.NewObjectId()
// 	user.Name = req.Name
// 	user.Major = req.Major
// 	user.IDCard = req.Idcard
// 	user.Account = req.Account
// 	user.Password = req.Password
// 	user.Courses = []string{req.Course}
// 	return user
// }

// func (d *Dao) NewStudent(ctx context.Context, req *api.RegisterReq) error {
// 	student := createDefaultStudent(req)
// 	return d.newStudent(ctx, student)
// }

// func (d *Dao) NewTeacher(ctx context.Context, req *api.RegisterReq) error {
// 	teacher := createDefaultTeacher(req)
// 	return d.newTeacher(ctx, teacher)
// }
// func (d *Dao) FillUserByIDCardInStudent(ctx context.Context, idCard string, rsp **uapi.UserReply) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	if *rsp == nil {
// 		*rsp = &uapi.UserReply{}
// 	}
// 	err = db.C(model.STUDENT_COLLECTION_NAME).Find(bson.M{\"idcard\": idCard}).One(*rsp)
// 	if err == nil {
// 		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
// 	}
// 	return
// }

// func (d *Dao) FillUserByIDCardInTeacher(ctx context.Context, idCard string, rsp **uapi.UserReply) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	if *rsp == nil {
// 		*rsp = &uapi.UserReply{}
// 	}
// 	err = db.C(model.TEACHER_COLLECTION_NAME).Find(bson.M{\"idcard\": idCard}).One(*rsp)
// 	if err == nil {
// 		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
// 	}
// 	return
// }

// func (d *Dao) GenerateLoginToken(ctx context.Context, rsp **api.PassportUserReply) error {
// 	userid := (*rsp).User.Userid
// 	token := utils.GenSession(userid)
// 	(*rsp).Token = token
// 	expires := 24 * time.Hour
// 	(*rsp).Expires = time.Now().Add(expires).Unix()

// 	set := d.redis.Set(GetLoginTokenRedisKey(userid), token, expires)
// 	return set.Err()

// }

// func (d *Dao) CheckSessionToken(ctx context.Context, uid, token string) error {
// 	res := d.redis.Get(GetLoginTokenRedisKey(uid))
// 	if savedToken, err := res.Result(); err != nil {
// 		return err
// 	} else {
// 		if savedToken != token {
// 			return pmodel.ErrorSessionTokenNotValidate
// 		}
// 	}
// 	return nil
// }

// func (d *Dao) FillStudentByAccount(ctx context.Context, account string, rsp **uapi.UserReply) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	if *rsp == nil {
// 		*rsp = &uapi.UserReply{}
// 	}
// 	err = db.C(model.STUDENT_COLLECTION_NAME).Find(bson.M{\"account\": account}).One(*rsp)
// 	if err == nil {
// 		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
// 	}
// 	return
// }

// func (d *Dao) FillTeacherByAccount(ctx context.Context, account string, rsp **uapi.UserReply) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	if *rsp == nil {
// 		*rsp = &uapi.UserReply{}
// 	}
// 	err = db.C(model.TEACHER_COLLECTION_NAME).Find(bson.M{\"account\": account}).One(*rsp)
// 	if err == nil {
// 		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
// 	}
// 	return
// }
// func (d *Dao) FillSecretaryByAccount(ctx context.Context, account string, rsp **uapi.UserReply) (err error) {
// 	db := d.mdb.Copy()
// 	defer db.Session.Close()
// 	if *rsp == nil {
// 		*rsp = &uapi.UserReply{}
// 	}
// 	err = db.C(model.SECRETARY_COLLECTION_NAME).Find(bson.M{\"account\": account}).One(*rsp)
// 	if err == nil {
// 		(*rsp).Userid = bson.ObjectId((*rsp).Userid).Hex()
// 	}
// 	return
// }

func (d *Dao) InitAllCourseKnowledgeMap() error {
	courseName_object_expires := 0 * time.Hour
	courseName_object := "object_course"
	objectCourse := "{\"@context\":{\"owl\":\"http://www.w3.org/2002/07/owl#\",\"Class\":\"owl:Class\",\"Concept\":\"skos:Concept\",\"Collection\":\"skos:Collection\",\"Relation\":\"rdf:Property\",\"affects\":\"dcterms:isRequiredBy\",\"belongsTo\":\"dcterms:isPartOf\",\"ckm\":\"https://github.com/Atuno/CKM-Toolkit/blob/master/jsonld/\",\"contains\":\"dcterms:hasPart\",\"dcterms\":\"http://purl.org/dc/terms/\",\"derives\":\"xkos:specializes\",\"id\":\"@id\",\"inherits\":\"xkos:generalizes\",\"isDependentOn\":\"dcterms:requires\",\"label\":\"rdfs:label\",\"language\":\"@language\",\"rdf\":\"http://www.w3.org/1999/02/22-rdf-syntax-ns#\",\"rdfs\":\"http://www.w3.org/2000/01/rdf-schema#\",\"skos\":\"http://www.w3.org/2004/02/skos/core#\",\"type\":\"@type\",\"value\":\"@value\",\"xkos\":\"http://rdf-vocabulary.ddialliance.org/xkos#\",\"related\":\"skos:related\",\"broader\":\"skos:broader\",\"narrower\":\"skos:narrower\",\"member\":\"skos:member\",\"sh\":\"http://www.w3.org/ns/shacl#\",\"NodeShape\":\"sh:NodeShape\",\"targetClass\":\"sh:targetClass\",\"property\":\"sh:property\",\"path\":\"sh:path\",\"maxCount\":\"sh:maxCount\",\"minCount\":\"sh:minCount\",\"datatype\":\"sh:datatype\",\"or\":\"sh:or\",\"equals\":\"sh:equals\",\"xsd\":\"http://www.w3.org/2001/XMLSchema#\",\"isSubjectOf\":\"skos:isSubjectOf\",\"weight\":\"rdf:value\"},\"@graph\":[{\"id\":\"dcterms:hasPart\",\"label\":{\"language\":\"zh\",\"value\":\"包含\"},\"type\":\"Relation\"},{\"id\":\"dcterms:isRequiredBy\",\"label\":{\"language\":\"zh\",\"value\":\"影响\"},\"type\":\"Relation\"},{\"id\":\"xkos:specializes\",\"label\":{\"language\":\"zh\",\"value\":\"派生\"},\"type\":\"Relation\"},{\"derives\":[{\"id\":\"ckm:privateMemberAccess\"},{\"id\":\"ckm:protectedMemberAccess\"},{\"id\":\"ckm:publicMemberAccess\"}],\"id\":\"ckm:accessSpecifier\",\"label\":[{\"language\":\"en\",\"value\":\"access specifier\"},{\"language\":\"zh\",\"value\":\"访问控制属性\"}],\"type\":\"Concept\"},{\"contains\":[{\"id\":\"ckm:member\"},{\"id\":\"ckm:composition\"},{\"id\":\"ckm:object\"},{\"id\":\"ckm:friend\"}],\"id\":\"ckm:class\",\"label\":[{\"language\":\"en\",\"value\":\"Class\"},{\"language\":\"zh\",\"value\":\"类\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:initializerList\"},\"id\":\"ckm:composition\",\"label\":[{\"language\":\"en\",\"value\":\"Composition\"},{\"language\":\"zh\",\"value\":\"类的组合\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:initializerList\"},\"derives\":[{\"id\":\"ckm:defaultConstructor\"},{\"id\":\"ckm:copyConstructor\"}],\"id\":\"ckm:constructor\",\"label\":[{\"language\":\"en\",\"value\":\"constructor\"},{\"language\":\"zh\",\"value\":\"构造函数\"}],\"type\":\"Concept\"},{\"contains\":[{\"id\":\"ckm:deepCopy\"},{\"id\":\"ckm:shallowCopy\"}],\"id\":\"ckm:copyConstructor\",\"label\":[{\"language\":\"en\",\"value\":\"copy constructor\"},{\"language\":\"zh\",\"value\":\"拷贝构造函数\"}],\"type\":\"Concept\"},{\"derives\":{\"id\":\"ckm:staticDataMember\"},\"id\":\"ckm:dataMember\",\"label\":[{\"language\":\"en\",\"value\":\"data member\"},{\"language\":\"zh\",\"value\":\"数据成员\"}],\"type\":\"Concept\"},{\"id\":\"ckm:deepCopy\",\"label\":[{\"language\":\"en\",\"value\":\"deep copy\"},{\"language\":\"zh\",\"value\":\"深拷贝\"}],\"type\":\"Concept\"},{\"id\":\"ckm:defaultConstructor\",\"label\":[{\"language\":\"en\",\"value\":\"default constructor\"},{\"language\":\"zh\",\"value\":\"缺省构造函数\"}],\"type\":\"Concept\"},{\"id\":\"ckm:destructor\",\"label\":[{\"language\":\"en\",\"value\":\"destructor\"},{\"language\":\"zh\",\"value\":\"析构函数\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:friendFunction\"},\"id\":\"ckm:friend\",\"label\":[{\"language\":\"en\",\"value\":\"friend\"},{\"language\":\"zh\",\"value\":\"友元\"}],\"type\":\"Concept\"},{\"id\":\"ckm:friendFunction\",\"label\":[{\"language\":\"en\",\"value\":\"friend function\"},{\"language\":\"zh\",\"value\":\"友元函数\"}],\"type\":\"Concept\"},{\"id\":\"ckm:initializerList\",\"label\":[{\"language\":\"en\",\"value\":\"initializer list\"},{\"language\":\"zh\",\"value\":\"初始化列表\"}],\"type\":\"Concept\"},{\"contains\":{\"id\":\"ckm:accessSpecifier\"},\"derives\":[{\"id\":\"ckm:memberFunction\"},{\"id\":\"ckm:dataMember\"}],\"id\":\"ckm:member\",\"label\":[{\"language\":\"en\",\"value\":\"member\"},{\"language\":\"zh\",\"value\":\"成员\"}],\"type\":\"Concept\"},{\"derives\":[{\"id\":\"ckm:constructor\"},{\"id\":\"ckm:destructor\"},{\"id\":\"ckm:staticMemberFunction\"}],\"id\":\"ckm:memberFunction\",\"label\":[{\"language\":\"en\",\"value\":\"member function\"},{\"language\":\"zh\",\"value\":\"成员函数\"}],\"type\":\"Concept\"},{\"affects\":{\"id\":\"ckm:deepCopy\"},\"id\":\"ckm:memoryLeak\",\"label\":[{\"language\":\"en\",\"value\":\"memory leak\"},{\"language\":\"zh\",\"value\":\"内存泄漏\"}],\"type\":\"Concept\"},{\"contains\":[{\"id\":\"ckm:objectArray\"},{\"id\":\"ckm:thisPointer\"}],\"id\":\"ckm:object\",\"label\":[{\"language\":\"en\",\"value\":\"Object\"},{\"language\":\"zh\",\"value\":\"对象\"}],\"type\":\"Concept\"},{\"id\":\"ckm:objectArray\",\"label\":[{\"language\":\"en\",\"value\":\"object array\"},{\"language\":\"zh\",\"value\":\"对象数组\"}],\"type\":\"Concept\"},{\"id\":\"ckm:privateMemberAccess\",\"label\":[{\"language\":\"en\",\"value\":\"private member access\"},{\"language\":\"zh\",\"value\":\"私有访问\"}],\"type\":\"Concept\"},{\"id\":\"ckm:protectedMemberAccess\",\"label\":[{\"language\":\"en\",\"value\":\"protected member access\"},{\"language\":\"zh\",\"value\":\"保护访问\"}],\"type\":\"Concept\"},{\"id\":\"ckm:publicMemberAccess\",\"label\":[{\"language\":\"en\",\"value\":\"public member access\"},{\"language\":\"zh\",\"value\":\"公有访问\"}],\"type\":\"Concept\"},{\"id\":\"ckm:shallowCopy\",\"label\":[{\"language\":\"en\",\"value\":\"shallow copy\"},{\"language\":\"zh\",\"value\":\"浅拷贝\"}],\"type\":\"Concept\"},{\"id\":\"ckm:staticDataMember\",\"label\":[{\"language\":\"en\",\"value\":\"static member variable\"},{\"language\":\"zh\",\"value\":\"静态数据成员\"}],\"type\":\"Concept\"},{\"id\":\"ckm:staticMemberFunction\",\"label\":[{\"language\":\"en\",\"value\":\"static member function\"},{\"language\":\"zh\",\"value\":\"静态成员函数\"}],\"type\":\"Concept\"},{\"id\":\"ckm:thisPointer\",\"label\":[{\"language\":\"en\",\"value\":\"this pointer\"},{\"language\":\"zh\",\"value\":\"this指针\"}],\"type\":\"Concept\"},{\"id\":\"ckm:summaryOfOOP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"面向对象程序设计概述\"}]},{\"id\":\"ckm:summaryOfCPP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"C++概述\"}],\"member\":[{\"id\":\"ckm:memoryLeak\",\"type\":\"Concept\"}]},{\"id\":\"ckm:classAndObject\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"类和对象\"}],\"member\":[{\"id\":\"ckm:accessSpecifier\",\"type\":\"Concept\"},{\"id\":\"ckm:class\",\"type\":\"Concept\"},{\"id\":\"ckm:composition\",\"type\":\"Concept\"},{\"id\":\"ckm:constructor\",\"type\":\"Concept\"},{\"id\":\"ckm:copyConstructor\",\"type\":\"Concept\"},{\"id\":\"ckm:dataMember\",\"type\":\"Concept\"},{\"id\":\"ckm:deepCopy\",\"type\":\"Concept\"},{\"id\":\"ckm:defaultConstructor\",\"type\":\"Concept\"},{\"id\":\"ckm:destructor\",\"type\":\"Concept\"},{\"id\":\"ckm:friend\",\"type\":\"Concept\"},{\"id\":\"ckm:friendFunction\",\"type\":\"Concept\"},{\"id\":\"ckm:initializerList\",\"type\":\"Concept\"},{\"id\":\"ckm:member\",\"type\":\"Concept\"},{\"id\":\"ckm:memberFunction\",\"type\":\"Concept\"},{\"id\":\"ckm:object\",\"type\":\"Concept\"},{\"id\":\"ckm:objectArray\",\"type\":\"Concept\"},{\"id\":\"ckm:privateMemberAccess\",\"type\":\"Concept\"},{\"id\":\"ckm:protectedMemberAccess\",\"type\":\"Concept\"},{\"id\":\"ckm:publicMemberAccess\",\"type\":\"Concept\"},{\"id\":\"ckm:shallowCopy\",\"type\":\"Concept\"},{\"id\":\"ckm:staticDataMember\",\"type\":\"Concept\"},{\"id\":\"ckm:staticMemberFunction\",\"type\":\"Concept\"},{\"id\":\"ckm:thisPointer\",\"type\":\"Concept\"}]},{\"id\":\"ckm:derivedClassAndInheritance\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"派生类和继承\"}]},{\"id\":\"ckm:polymorphism\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"多态性\"}]},{\"id\":\"ckm:templateAndException\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"模板与异常处理\"}]},{\"id\":\"ckm:streamAndIO\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"C++的流类库与输入输出\"}]},{\"id\":\"ckm:designOfOOP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"面向对象程序设计方法与实例\"}]},{\"id\":\"ckm:OOP_CPP\",\"type\":\"Collection\",\"label\":[{\"language\":\"zh\",\"value\":\"面向对象程序设计（C++）\"}],\"member\":[{\"id\":\"ckm:summaryOfOOP\",\"type\":\"Collection\"},{\"id\":\"ckm:summaryOfCPP\",\"type\":\"Collection\"},{\"id\":\"ckm:classAndObject\",\"type\":\"Collection\"},{\"id\":\"ckm:derivedClassAndInheritance\",\"type\":\"Collection\"},{\"id\":\"ckm:polymorphism\",\"type\":\"Collection\"},{\"id\":\"ckm:templateAndException\",\"type\":\"Collection\"},{\"id\":\"ckm:streamAndIO\",\"type\":\"Collection\"},{\"id\":\"ckm:designOfOOP\",\"type\":\"Collection\"}]},{\"type\":\"NodeShape\",\"id\":\"ckm:conceptShape\",\"targetClass\":\"skos:Concept\",\"property\":[{\"path\":\"@id\",\"minCount\":1},{\"path\":\"@type\",\"minCount\":1},{\"path\":\"rdfs:label\",\"label\":\"名称\",\"datatype\":\"xsd:string\",\"minCount\":1},{\"path\":\"skos:isSubjectOf\",\"label\":\"\",\"datatype\":\"xsd:string\",\"minCount\":0},{\"path\":\"rdf:value\",\"label\":\"\",\"datatype\":\"xsd:integer\",\"minCount\":0}]},{\"type\":\"NodeShape\",\"id\":\"ckm:collectionShape\",\"targetClass\":\"skos:Collection\",\"property\":[{\"path\":\"@id\",\"minCount\":1},{\"path\":\"@type\",\"minCount\":1},{\"path\":\"rdfs:label\",\"label\":\"名称\",\"minCount\":1,\"datatype\":\"xsd:string\"},{\"path\":\"skos:member\"}]},{\"type\":\"NodeShape\",\"id\":\"ckm:relationShape\",\"targetClass\":\"rdf:Property\",\"property\":[{\"path\":\"@id\",\"or\":[{\"equals\":\"dcterms:hasPart\"},{\"equals\":\"dcterms:isRequiredBy\"},{\"equals\":\"xkos:specializes\"},{\"equals\":\"xkos:generalizes\"},{\"equals\":\"dcterms:isPartOf\"},{\"equals\":\"dcterms:requires\"},{\"equals\":\"skos:related\"},{\"equals\":\"skos:narrower\"}]}]}]}"
	set := d.redis.Set(courseName_object, objectCourse, courseName_object_expires)
	return set.Err()
}

func (d *Dao) QueryKnowledgeMapByCourse(course string) (knowledgeMap string, err error) {
	res := d.redis.Get(course)
	if knowledgeMap, err = res.Result(); err != nil {
		return "", err
	} else {
		return knowledgeMap, nil
	}
}
