package handler

// func (s *KnowledgeMapService) GetKnowledegeMapBySubject(ctx context.Context, req *api.CRqQueryMapBySubject, rsp *api.KnowledegeMapInfo) error {
// 	if req.Subject == "" {
// 		logrus.Errorln("subject is undefine!")
// 		return errors.New("subject is undefine!")
// 	}
// 	knowledgeMap, err := gdao.QueryKnowledgeMapByCourse(req.Subject)
// 	rsp.Knowledgemap = knowledgeMap
// 	return err
// }

// func (s *KnowledgeMapService) GetMyKnowledegeMapBySubject(ctx context.Context, req *api.CRqQueryMyMapBySubject, rsp *api.KnowledegeMapInfo) error {
// 	if req.Subject == "" {
// 		logrus.Errorln("subject is undefine!")
// 		return errors.New("subject is undefine!")
// 	}
// 	mathKnowledgeMap := math.GetMathKnowledgeMap()
// 	questionInfo, err := questionSrv.GetMyQuestionInfo(ctx, &qapi.CRqQueryMyQuestionInfoBySubject{Uid: req.Uid, Subject: req.Subject, Endtime: req.Endtime})
// 	for _, v := range questionInfo.Knowledgenodes {
// 		node := &model.Node{}
// 		err := gdao.QueryNodeInfoByNodeID(ctx, bson.ObjectIdHex(v), node)
// 		if err != nil {
// 			logrus.Errorf("can't find node:%v", v)
// 			continue
// 		}
// 		nodeInfo := CreateKnowledgeMapFromNodes(ctx, node)
// 		if len(nodeInfo) != 0 {
// 			mathKnowledgeMap["@graph"] = append(mathKnowledgeMap["@graph"].([]map[string]interface{}), nodeInfo...)
// 		}
// 	}
// 	info, _ := json.Marshal(mathKnowledgeMap)
// 	rsp.Knowledgemap = string(info[:])
// 	return err
// }

// func CreateKnowledgeMapFromNodes(ctx context.Context, node *model.Node) []map[string]interface{} {
// 	info := []map[string]interface{}{}
// 	relations := new([]*model.Relation)
// 	err := gdao.QueryRelationByNodeID(ctx, node.ID, relations)
// 	if err != nil {
// 		logrus.Errorf("can't find node relations:%v", node.ID.Hex())
// 		return info
// 	}

// 	contains := []map[string]interface{}{}
// 	derives := []map[string]interface{}{}
// 	affects := []map[string]interface{}{}
// 	for _, v := range *relations {
// 		node := &model.Node{}
// 		err := gdao.QueryNodeInfoByNodeID(ctx, v.ObjectNodeID, node)
// 		if err != nil {
// 			logrus.Errorf("can't find node:%v", v)
// 			continue
// 		}
// 		switch v.Relation {
// 		case model.RelationContain:
// 			contains = append(contains, map[string]interface{}{"id": node.Kind})
// 		case model.RelationDerive:
// 			derives = append(derives, map[string]interface{}{"id": node.Kind})
// 		case model.RelationAffect:
// 			affects = append(affects, map[string]interface{}{"id": node.Kind})
// 		default:
// 		}
// 	}

// 	label := []map[string]interface{}{}
// 	for _, v := range node.Label {
// 		label = append(label, map[string]interface{}{"language": v.Language, "value": v.Value})
// 	}
// 	info = append(info, map[string]interface{}{
// 		"id":       node.Kind,
// 		"label":    label,
// 		"type":     node.Type,
// 		"contains": contains,
// 		"derives":  derives,
// 		"affects":  affects,
// 	})
// 	return info
// }

// func (s *KnowledgeMapService) CreateKnowledege(ctx context.Context, req *api.CreateKnowledegeReq, rsp *api.KnowledegeInfoReply) error {
// 	logrus.Infof("CreateKnowledege req is %v ", req)
// 	knowledge := &model.Node{bson.NewObjectId(), []model.LableInfo{model.LableInfo{"zh", req.Name}}, "default", model.NodeConcept, req.Subject, req.Course}
// 	knowledgeInfo, err := gdao.NewNode(ctx, knowledge)
// 	if err != nil {
// 		fmt.Println("create knowledge error", err)
// 		return fmt.Errorf("创建知识点失败")
// 	} else {
// 		rsp.Id = knowledgeInfo.ID.Hex()
// 		rsp.Name = knowledgeInfo.Label[0].Value
// 		rsp.Subject = knowledgeInfo.Subject
// 		rsp.Course = knowledgeInfo.Course
// 	}
// 	return nil
// }

// func (s *KnowledgeMapService) QueryKnowledegeInfo(ctx context.Context, req *api.QueryKnowledegeInfoReq, rsp *api.KnowledegeInfoReply) error {
// 	logrus.Infof("QueryKnowledegeInfo req is %v ", req)
// 	knowledege := new(model.Node)
// 	if err := gdao.FillKnowledgeByID(ctx, bson.ObjectIdHex(req.Id), knowledege); err != nil {
// 		fmt.Println("query knowledege info error", err)
// 		return fmt.Errorf("查询知识点信息失败")
// 	} else {
// 		rsp.Id = knowledege.ID.Hex()
// 		rsp.Name = knowledege.Label[0].Value
// 		rsp.Subject = knowledege.Subject
// 		rsp.Course = knowledege.Course
// 	}
// 	return nil
// }
