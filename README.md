# knowledgemap_backend
知识图谱项目后端
/usr/local/mongodb/bin
/data/db  数据存放位置
redis地址 /usr/local/redis/
consul agent -server  -bootstrap-expect 1 -data-dir /tmp/consul -node=172.17.9.156
公47.95.145.171
cd  /usr/local/consul/
nohup ./consul agent -dev -ui -node=consul-dev -client=0.0.0.0 > consul.log 2>&1 &

所有api要加统一前缀 /api
所有登录后的操作都需要在header里添加auth-session为用户cookie

----

api:/user/register

desc:注册

method:post

param:

    - rtype 身份枚举 0学生 1老师 2教秘
    - name string 真实姓名 
    - major string 专业
    - sex string 性别
    - account string 账号
    - password string 密码
    - college string 学院

response:

    `{
    "msg": "",
    "code": 0,
    "data": {
            "user": {
                "userid": "5fb665f436d02a366cae5b0f",
                "username": "teacher1",
                "major": "computer",
                "sex": "man",
                "account": "teacher1",
                "password": "123456",
                "number": "1"
            },
            "token": "55Tfjb56Y6O58fo453M6yd40N2cas3j6I6xc6aLeN5obI0if",
            "expires": 1605875572
        }
    }`

----

api:/user/login

desc:用户登录

method:put

param:

   - account string 账号
   - password string 密码
   - ltype 身份枚举 0学生 1老师 2教秘
  
response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "user": {
                "userid": "5fb666c836d02a366cae5b10",
                "username": "teacher1",
                "usertype": 1,
                "major": "computer",
                "sex": "man",
                "account": "teacher1",
                "password": "123456",
                "number": "1"
            },
            "token": "T5vfMbR666c6EcM8M3Q6jdv0h2Qag3z646DcEaDeT51b81M0",
            "expires": 1605875825
        }
    }`

----

api:/user/changeinfo

desc:修改自己的信息

method:put

param:

   - name string 修改后的用户名
   - sex string 修改后的性别
   - college string 修改后的学院
   - major string 修改后的专业

response:

  `{
    "msg": "",
    "code": 0,
    "data": {
        "user": {
                "userid": "5fb67c5e36d02a6d18e309c9",
                "username": "teacher1_change",
                "major": "computer_change",
                "sex": "man_change",
                "account": "teacher1",
                "number": "1",
                "college": "computer_change"
            }
        }
    }`

----

api:/user/changepassword

desc:修改自己的信息

method:put

param:

   - password string 修改后的密码

response:

  `{
    "msg": "",
    "code": 0,
    "data": {}
    }`

----

api:/user/query/info

desc:查看自己的信息

method:get

param:

response:

  `{
    "msg": "",
    "code": 0,
    "data": {
            "userid": "5fb67c5e36d02a6d18e309c9",
            "username": "teacher1_change",
            "major": "computer_change",
            "sex": "man_change",
            "account": "teacher1",
            "number": "1",
            "college": "computer_change",
            "identify": 1
        }
    }`


----

api:/class/create

desc:教师或者教秘，创建班级

method:post

param:

   - classname string 课程名
   - major string 专业名
   - college string 学院
   - teachername string 教师名,教师创建必不填，教秘创建必填
   - teacherid string 教师id编号，教秘创建必填
   - introdution string 课程介绍

response:

    `{
        "msg": "",
        "code": 0,
        "data": {
            "classid": "5e97e77936d02a9dbb5ce966",
            "name": "高等数学A",
            "major": "math",
            "college": "computer",
            "teachername": "李永乐",
            "introduction": "数学课，必学的"
        }
    }`

--------

api:/class/join

desc:学生加入班级

method:put

param:

    - classid string 班级id
  
response:

    `{
        "msg": "",
        "code": 0,
        "data": {
            "classes": [
                {
                    "classid": "5e97e77936d02a9dbb5ce966",
                    "name": "高等数学A",
                    "major": "math",
                    "college": "computer",
                    "teachername": "李永乐",
                    "createTime": 1587013497,
                    "number": "10",
                     "introduction": "数学课，必学的"
                }
            ]
        }
    }`

--------

api:/class/query/myclasses

desc:学生、教师查询自己所有加入的班级

method:get
  
response:

    `{
        "msg": "",
        "code": 0,
        "data": {
            "classes": [
                {
                    "classid": "5e97e77936d02a9dbb5ce966",
                    "name": "高等数学A",
                    "major": "math",
                    "college": "computer",
                    "teachername": "李永乐",
                    "createTime": 1587013497,
                    "number": "10",
                    "introduction": "数学课，必学的"
                }
            ]
        }
    }`

--------

api:/class/query/classinfo/:classid

desc:查询班级信息,待完成

method:get

param:

    - classid string 班级id 
  
response:

    `{
        "msg": "",
        "code": 0,
        "data": {
            "classid": "5e97e77936d02a9dbb5ce966",
            "name": "高等数学A",
            "major": "math",
            "college": "computer",
            "teachername": "李永乐",
            "introduction": "数学课，必学的"

        }
    }`

--------

api:/class/query/alluserinclass/:classid

desc:查询班级所有学生信息

method:get

param:

    - classid string 班级id 
    - page string 分页
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "classid": "",
        "students": [
            {
                "userid": "5fdee86255a8d148cdc91ae2",
                "username": "张十一",
                "number": "18",
                "account": "e2",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1 //状态，1是活跃，2是封禁
            },
            {
                "userid": "5fdee86255a8d148cdc91ae1",
                "username": "张十",
                "number": "17",
                "account": "e1",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            },
            {
                "userid": "5fdee86255a8d148cdc91ad9",
                "username": "张九",
                "number": "16",
                "account": "d9",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            },
            {
                "userid": "5fdee86255a8d148cdc91ad8",
                "username": "张八",
                "number": "15",
                "account": "d8",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            },
            {
                "userid": "5fdee86255a8d148cdc91ad7",
                "username": "张七",
                "number": "14",
                "account": "d7",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            },
            {
                "userid": "5fdee86255a8d148cdc91ad6",
                "username": "张六",
                "number": "13",
                "account": "d6",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            },
            {
                "userid": "5fdee86255a8d148cdc91ad5",
                "username": "张五",
                "number": "12",
                "account": "d5",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            },
            {
                "userid": "5fdee86255a8d148cdc91ad4",
                "username": "张四",
                "number": "11",
                "account": "d4",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            },
            {
                "userid": "5fdee86255a8d148cdc91ad3",
                "username": "张三",
                "number": "10",
                "account": "d3",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1 
            },
            {
                "userid": "5fdee86255a8d148cdc91ad2",
                "username": "张二",
                "number": "8",
                "account": "d2",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            }
        ],
        "currentpage": 0,
        "totalpage": 1
    }
}`

--------

api:/class/query/classes/:college/:subject/:course/:teacher/:page

desc:根据条件检索班级

method:get

param:

    - college  学院
    - subject 学科
    - course 课程
    - teacher 教师
    - page 当前页数
  
response:

    `{
        "msg":"",
        "code":0,
        "data":
            {
                "classes":
                    [{"classid":"5fbbc8af36d02ac44f70eb67","name":"高数A","subject":"数学","course":"高数","college":"信息工程学院","teachername":"teacher7","introduction":"这是高数A"},
                    {"classid":"5fbbc88736d02ac44f70eb65","subject":"数学","course":"高数","college":"信息工程学院","teachername":"teacher7","introduction":"这是高数A"},
                    {"classid":"5fbbc80036d02ac117cb324f","subject":"数学","course":"高数","college":"信息工程学院","teachername":"teacher7","introduction":"这是高数A"}],
                    "currentpage":1
            }
    }`

--------

api:/class/query/formlist

desc:

method:get

param:
  
response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "colleges": [
            "数学科学学院",
            "中文系",
            "信息工程学院",
            "历史学院"
        ],
        "subjects": [
            "数学",
            "历史",
            "语文",
            "计算机"
        ],
        "courses": [
            "高数",
            "线代",
            "语文",
            "英语",
            "c语言"
        ]
    }
}`

--------

api:/class/query/student

desc:在班级中以用户名查询学生

method:get

param:
    - classid string 班级id
    - username string 学生用户名
    - page string 分页

  
response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "classid": "",
        "students": [
            {
                "userid": "5fdee86255a8d148cdc91ad3",
                "username": "张三",
                "number": "10",
                "account": "d3",
                "sex": "男",
                "college": "数学学院",
                "createtime": 1608444002,
                "status": 1
            }
        ],
        "currentpage": 0,
        "totalpage": 0
    }
}`


--------

api:/class/deletestudent

desc:在班级中删除学生

method:post

param:
    - classid string 班级id
    - userid string 学生id

  
response:

    ``

--------

api:/knowledge/create

desc:创建知识点

method:post

param:
  
  name: string 知识点名称
  subject: string 学科
  course:string 课程


response:

    `{
        "msg": "",
        "code": 0,
        "data": {
            "id": "5e9b06ce36d02a7b7637eccf",
            "course": "高数"
        }
    }`

--------

api:/knowledge/query/:knowledgeId

desc:查询知识点

method:get

param:

    - id 知识点id
  
response:

   {
    "msg": "",
    "code": 0,
    "data": {
        "id": "5e9b06ce36d02a7b7637eccf",
        "course": "高数"
    }
}


--------

api:/knowledge/query/map/:subject

desc:查询学科所有知识点，待完成

method:GET

param:
       
    - subject: string 学科
  
response:

    未设计好


--------

api:/knowledge/query/my/map/:uid/:subject/:endtime

desc:分学科查询自己掌握的知识点

method:GET

param:

      - uid: string 用户id
      - subject: string 学科  
      - endtime: int64 结束时间
  
response:

    未设计好


--------

api:/question/create

desc:创建题目

method:POST

param:

    `{
        "kind": 1, int64 类型
        "content": "1+1=？", string 题目描述
        "option": [
        "1",
        "2",
        "3",
        "4"
        ],string数组 选择题的选项
        "answer": [
        "2"
        ],string数组 选择题的答案
        "subject": "math",string 学科
        "course": "小学数学",string 课程
        "knowledge": "5e9b06ce36d02a7b7637eccf" string 对应知识点的id
    }`

response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "id": "5e9b0ccf36d02a7a6979991e",
        "kind": 1,
        "content": "1+1=？",
        "option": [
            "1",
            "2",
            "3",
            "4"
        ],
        "answer": [
            "2"
        ],
        "subject": "math",
        "course": "小学数学",
        "knowledge": "5e9b06ce36d02a7b7637eccf"
    }
}`

--------

api:/question/query

desc:查询某种类型的题目

method:GET

param:

    - questiontype int64 题目类型
    - subject string 学科
    - course string 课程 
    - page int64 分页
   
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "questions": [
            {
                "id": "5fccd6bc36d02a179478822c",
                "questiontype": 2,
                "title": "多选题测试",
                "istitleimg": false,
                "options": [
                    {
                        "prefix": "A",
                        "content": "1",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "B",
                        "content": "2",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "C",
                        "content": "3",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "D",
                        "content": "4",
                        "iscontentimg": false
                    }
                ],
                "answers": [
                    {
                        "prefix": "A",
                        "content": "1",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "C",
                        "content": "3",
                        "iscontentimg": false
                    }
                ],
                "subject": "math",
                "course": "高数",
                "knowledge": "5fccd50d36d02a19df8ed447",
                "name": "多选题简单测试",
                "needcheck": false,
                "explain": "看着搞",
                "star": 2
            },
            {
                "id": "5fccd6bc36d02a179478822b",
                "questiontype": 2,
                "title": "多选题测试",
                "istitleimg": false,
                "options": [
                    {
                        "prefix": "A",
                        "content": "1",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "B",
                        "content": "2",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "C",
                        "content": "3",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "D",
                        "content": "4",
                        "iscontentimg": false
                    }
                ],
                "answers": [
                    {
                        "prefix": "A",
                        "content": "1",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "C",
                        "content": "3",
                        "iscontentimg": false
                    }
                ],
                "subject": "math",
                "course": "高数",
                "knowledge": "5fccd50d36d02a19df8ed447",
                "name": "多选题简单测试",
                "needcheck": false,
                "explain": "看着搞",
                "star": 2
            },
            {
                "id": "5fccd6bc36d02a179478822a",
                "questiontype": 2,
                "title": "多选题测试",
                "istitleimg": false,
                "options": [
                    {
                        "prefix": "A",
                        "content": "1",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "B",
                        "content": "2",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "C",
                        "content": "3",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "D",
                        "content": "4",
                        "iscontentimg": false
                    }
                ],
                "answers": [
                    {
                        "prefix": "A",
                        "content": "1",
                        "iscontentimg": false
                    },
                    {
                        "prefix": "C",
                        "content": "3",
                        "iscontentimg": false
                    }
                ],
                "subject": "math",
                "course": "高数",
                "knowledge": "5fccd50d36d02a19df8ed447",
                "name": "多选题简单测试",
                "needcheck": false,
                "explain": "看着搞",
                "star": 2
            }
        ],
        "currentpage": 0,
        "totalpage": 0
    }
}`

--------

api:/homework/create

desc:创建试卷

method:post

param:

    `{
	"name":"家庭作业1", string 试卷名
	"classid":"5e97e77936d02a9dbb5ce966", string 班级id，即给哪些班级设置试卷
	"students":["5e919d0036d02a48ccc08225"], string 学生id，即给哪些学生设置试卷
	"questions":["5e9b0ccf36d02a7a6979991e"]，string 题目id，即给添加哪些题目
    }`
  
response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "homeworkid": "5e9b103b36d02a7a6979991f"
    }
}`

--------

api:/homework/query/:userid/:classid

desc:查询用户所需做的试卷

method:GET

param:

    - userid 用户id
    - classid 班级试卷
  
response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "homework": [
            {
                "homeworkid": "5e9b103b36d02a7a6979991f",
                "name": "家庭作业1",
                "questions": [
                    {
                        "questionid": "5e9b0ccf36d02a7a6979991e",
                        "kind": 1,
                        "content": "1+1=？",
                        "option": [
                            "1",
                            "2",
                            "3",
                            "4"
                        ]
                    }
                ]
            }
        ]
    }
}`

--------

api:/homework/do

desc:做试卷

method:put

param:

    `{
	"homeworkid":"5e9b103b36d02a7a6979991f", string 试卷id
	"userid":"5e97e77936d02a9dbb5ce966", string 用户id
	"username":"iii", string 用户名
	"questions":[
         {
             "questionid": "5e9b103b36d02a7a6979991f", string 试题id
              "answer": [
                "2"
              ],string数组 选择题的答案
         }
        ]，数组 具体每道题的答案
    }`
  
response:

    `{
    "msg":"",      失败的时候代表失败原因
    "code":0       0代表成功，1代表失败
    }`

--------

api:/homework/answerrecord/query/:homeworkid

desc:查询做试卷的记录

method:GET

param:

    - homeworkid 答题记录
  
response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "homeworkrecord": [
            {
                "questionid": "5e9b0ccf36d02a7a6979991e",
                "content": "1+1=？",
                "option": [
                    "1",
                    "2",
                    "3",
                    "4"
                ],
                "rightanswer": [
                    "2"
                ],
                "alluseranswer": [
                    {
                        "username": "学生1",
                        "userid": "5e919d0036d02a48ccc08225",
                        "answer": [
                            "2"
                        ]
                    }
                ]
            }
        ]
    }
}`


--------

api:/homework/query/info/:classid

desc:查询班级的所需试卷

method:get

param:

    - classid string 班级id

  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "homework": [
            {
                "homeworkid": "5e9b103b36d02a7a6979991f",
                "name": "家庭作业1"
            }
        ]
    }
}`



--------

api:/practice/create

desc:创建题库

method:post

param:

    - name string 题库名字
    - classid string 教师创建的话需绑定课程班级，学生创建则不需要
    - pstype 枚举 题库类型，0是普通题库，1是收藏题库，2是错题题库
    - introduction string 题库介绍

  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "practicesummaryid": "5fcccdbc36d02a1187f6528a"
    }
    }`

--------

api:/practice/query/psinfo/:psid

desc:查询题库信息

method:get

param:

    - psid string 题库id
  
response:

   `{
        "msg": "",
        "code": 0,
        "data": {
            "practicesummaryid": "5fcb8c1236d02ab338449039",
            "name": "题库一",
            "introduction": "测试创建题库",
            "pstype": 1
        }
    }`

--------

api:/practice/query/psdetailinfo/:psid/:page

desc:查询某个题库中的题目

method:get

param:

    - psid string 题库id
    - page string 当前分页
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "questions": [
            {
                "questionid": "5fccd6d236d02a17947882a4",
                "kind": 1,题目类型枚举，1选择题，2简答题
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6d236d02a17947882a3",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6d136d02a17947882a2",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6d136d02a17947882a1",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6d036d02a17947882a0",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6d036d02a179478829f",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6cf36d02a179478829e",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6ce36d02a179478829d",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6ca36d02a179478829c",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            },
            {
                "questionid": "5fccd6c636d02a179478829b",
                "kind": 1,
                "name":"鸡兔同笼",
                "knowledgename": "知识点"
            }
        ],
        "totalpage": 1
    }
}`

--------

api:/practice/query/mypsinfo

desc:查询我的题库

method:get

param:

    - page string 当前分页
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "practicesummary": [
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6528f",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6528d",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6528c",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6528b",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6528a",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6527e",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6527d",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6527c",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6527b",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            },
            {
                "practicesummaryid": "5fcccdbc36d02a1187f6527a",
                "name": "题库二",
                "introduction": "测试创建题库",
                "pstype": 1
            }
        ],
        "totalpage": 1
    }
}`


--------

api:/practice/addquestion

desc:往某个题库添加的题目

method:post

param:

    - practicesummaryid string 题库id
    - questions []string 题目id数组
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {}
    }`


--------

api:/practice/deletequestion

desc:删除某个题库的部分题目

method:post

param:

    - practicesummaryid string 题库id
    - questions []string 题目id数组
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {}
    }`


--------

api:/paper/create

desc:教师为班级创建试卷

method:post

param:

    `{
	"name":"期中试卷试卷", //string 试卷名
	"classid":"5fbbc8af36d02ac44f70eb67", //string 班级id
	"continuingtime":3600, //in64 答题时长，单位秒
	"questions":[
		{
			"questionid":"5fccd6d236d02a17947882a4",//string 试题id
			"score":10,//int64 试题分数
			"needcheck":false //bool 是否需要人工判题
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		},
		{
			"questionid":"5fccd6d236d02a17947882a4",
			"score":10,
			"needcheck":false
		}
		]
}`
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "paperid": "5fe478b3a541ddbf7e479508" // string  试卷id
    }
    }`

--------

api:/paper/query

desc:查询班级下拥有的试卷

method:get

param:

    - classid string 班级id
    - page string 当前页数
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "paper": [
            {
                "paperid": "5fe47805a541ddbf7e4794fe",//试卷id
                "name": "期中试卷试卷",// 试卷名
                "totalscore": 110, // 试卷总分
                "continuingtime": 3600 // 答卷时长限制
            }
        ],
        "currentpage": 1,
        "totalpage": 1
    }
    }`

--------

api:/paper/query/questions

desc:根据试卷id查询试卷内容

method:Get

param:

    - paperid string 试卷id
    - paperkind string 类型,value为homework和exam两种，value为exam，则为已完成的试卷功能，value为homework，则为首页待接的查询作业接口

response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "titleitems": [
            {
                "name": "单选题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478829a",
                        "questiontype": 1,
                        "title": "测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ]
                    },
                    {
                        "id": "5fccd6bc36d02a179478829b",
                        "questiontype": 1,
                        "title": "测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ]
                    },
                    {
                        "id": "5fccd6bc36d02a179478829c",
                        "questiontype": 1,
                        "title": "测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ]
                    }
                ]
            },
            {
                "name": "多选题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478822a",
                        "questiontype": 2,
                        "title": "多选题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ]
                    },
                    {
                        "id": "5fccd6bc36d02a179478822b",
                        "questiontype": 2,
                        "title": "多选题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ]
                    },
                    {
                        "id": "5fccd6bc36d02a179478822c",
                        "questiontype": 2,
                        "title": "多选题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ]
                    }
                ]
            },
            {
                "name": "判断题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478823a",
                        "questiontype": 3,
                        "title": "判断题测试1+1=2？",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "对",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "错",
                                "iscontentimg": false
                            }
                        ]
                    },
                    {
                        "id": "5fccd6bc36d02a179478823b",
                        "questiontype": 3,
                        "title": "判断题测试5-4=1？",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "对",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "错",
                                "iscontentimg": false
                            }
                        ]
                    }
                ]
            },
            {
                "name": "填空题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478824a",
                        "questiontype": 4,
                        "title": "填空题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "窗前_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "疑是_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "举头_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "低头_",
                                "content": "",
                                "iscontentimg": false
                            }
                        ]
                    },
                    {
                        "id": "5fccd6bc36d02a179478824b",
                        "questiontype": 4,
                        "title": "填空题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "窗前_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "疑是_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "举头_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "低头_",
                                "content": "",
                                "iscontentimg": false
                            }
                        ]
                    }
                ]
            },
            {
                "name": "简答题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478825a",
                        "questiontype": 5,
                        "title": "简述笛卡尔坐标系的建立方法",
                        "istitleimg": false,
                        "options": null
                    },
                    {
                        "id": "5fccd6bc36d02a179478825b",
                        "questiontype": 5,
                        "title": "简述数学的应用方面？",
                        "istitleimg": false,
                        "options": null
                    }
                ]
            },
            {
                "name": "图片题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478826a",
                        "questiontype": 6,
                        "title": "请上传y=cosx的图像",
                        "istitleimg": false,
                        "options": null
                    },
                    {
                        "id": "5fccd6bc36d02a179478826b",
                        "questiontype": 6,
                        "title": "请上传y=sinx的图像",
                        "istitleimg": false,
                        "options": null
                    }
                ]
            },
            {
                "name": "文件题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478827a",
                        "questiontype": 7,
                        "title": "请上传第一次作业的文件",
                        "istitleimg": false,
                        "options": null
                    },
                    {
                        "id": "5fccd6bc36d02a179478827b",
                        "questiontype": 7,
                        "title": "请上传第二次作业的文件",
                        "istitleimg": false,
                        "options": null
                    }
                ]
            }
        ],
        "name": "期中试卷试卷",
        "score": 110,
        "suggesttime": 3600
    }
}`

--------

api:/paper/do

desc:答卷

method:PUT

param:

    `{
	"paperid":"5fe47805a541ddbf7e4794fe",
	"paperkind":"exam",
	"answer":[
		{
			"questionid":"5fccd6bc36d02a179478822a",
			"answers":["A","B"]
		},
		{
			"questionid":"5fccd6bc36d02a179478827a",
			"answers":["http://127.0.0.1:8080/group1/default/20210102/20/40/5/transport_end.jpeg"
			]
		}
		]
}`
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {}
    }`

--------
api:/paper/answerrecord/query/list

desc:查询我所答测试试卷或家庭作业的列表

method:GET

param:

    - page int 分页
    - paperkind string exam为测试试卷,homework为家庭作业
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "recordlist": [
            {
                "paperrecordid": "5ff42fc055a8d122a8b714d4",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838528
            },
            {
                "paperrecordid": "5ff42fb955a8d122a8b714d1",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838521
            },
            {
                "paperrecordid": "5ff42fb855a8d122a8b714ce",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838520
            },
            {
                "paperrecordid": "5ff42fb755a8d122a8b714cb",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838519
            },
            {
                "paperrecordid": "5ff42fb655a8d122a8b714c8",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838518
            },
            {
                "paperrecordid": "5ff42fb655a8d122a8b714c5",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838518
            },
            {
                "paperrecordid": "5ff42fb455a8d122a8b714c2",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838516
            },
            {
                "paperrecordid": "5ff42fb355a8d122a8b714bf",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838515
            },
            {
                "paperrecordid": "5ff42ece55a8d122a8b714bc",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838286
            },
            {
                "paperrecordid": "5ff42eb855a8d122a8b714b9",
                "papername": "有所要题目类型的试卷",
                "course": "高数",
                "teachername": "teacher1",
                "status":"",
                "dotime": 1609838264
            }
        ],
        "currentpage": 0,
        "totalpage": 1
    }
}`


--------


api:/paper/answerrecord/query

desc:查询具体的测试答卷答题记录或家庭作业答题记录

method:Get

param:

    - paperrecordid string 试卷记录id
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "paperid": "5fe47805a541ddbf7e4794fe",
        "papername": "有所要题目类型的试卷",
        "score": 110,//总分
        "getscore": 0,//得分
        "needcheck": true,//需要人工判卷
        "paperrecord": [
            {
                "name": "单选题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478829a",
                        "questiontype": 1,
                        "title": "测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            }
                        ],//正确答案 
                        "getoptions": null,//回答的答案
                        "needcheck": false,//是否需要人工判卷
                        "explain": "简单，自己想",//解析
                        "result": 0 // 0未作答，1正确，2错误，3非最佳答案，4等待批改
                    },
                    {
                        "id": "5fccd6bc36d02a179478829b",
                        "questiontype": 1,
                        "title": "测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "简单，自己想",
                        "result": 0
                    },
                    {
                        "id": "5fccd6bc36d02a179478829c",
                        "questiontype": 1,
                        "title": "测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "简单，自己想",
                        "result": 0
                    }
                ]
            },
            {
                "name": "多选题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478822a",
                        "questiontype": 2,
                        "title": "多选题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": [
                            "A",
                            "B"
                        ],
                        "needcheck": false,
                        "explain": "看着搞",
                        "result": 2
                    },
                    {
                        "id": "5fccd6bc36d02a179478822b",
                        "questiontype": 2,
                        "title": "多选题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "看着搞",
                        "result": 0
                    },
                    {
                        "id": "5fccd6bc36d02a179478822c",
                        "questiontype": 2,
                        "title": "多选题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "2",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "D",
                                "content": "4",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "1",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "C",
                                "content": "3",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "看着搞",
                        "result": 0
                    }
                ]
            },
            {
                "name": "判断题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478823a",
                        "questiontype": 3,
                        "title": "判断题测试1+1=2？",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "对",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "错",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "对",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "看着搞",
                        "result": 0
                    },
                    {
                        "id": "5fccd6bc36d02a179478823b",
                        "questiontype": 3,
                        "title": "判断题测试5-4=1？",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "A",
                                "content": "对",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "B",
                                "content": "错",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "A",
                                "content": "对",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "看着搞",
                        "result": 0
                    }
                ]
            },
            {
                "name": "填空题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478824a",
                        "questiontype": 4,
                        "title": "填空题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "窗前_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "疑是_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "举头_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "低头_",
                                "content": "",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "窗前_",
                                "content": "明月光",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "疑是_",
                                "content": "地上霜",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "举头_",
                                "content": "望明月",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "低头_",
                                "content": "思故乡",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "看着搞",
                        "result": 0
                    },
                    {
                        "id": "5fccd6bc36d02a179478824b",
                        "questiontype": 4,
                        "title": "填空题测试",
                        "istitleimg": false,
                        "options": [
                            {
                                "prefix": "窗前_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "疑是_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "举头_",
                                "content": "",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "低头_",
                                "content": "",
                                "iscontentimg": false
                            }
                        ],
                        "rightoptions": [
                            {
                                "prefix": "窗前_",
                                "content": "明月光",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "疑是_",
                                "content": "地上霜",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "举头_",
                                "content": "望明月",
                                "iscontentimg": false
                            },
                            {
                                "prefix": "低头_",
                                "content": "思故乡",
                                "iscontentimg": false
                            }
                        ],
                        "getoptions": null,
                        "needcheck": false,
                        "explain": "看着搞",
                        "result": 0
                    }
                ]
            },
            {
                "name": "简答题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478825a",
                        "questiontype": 5,
                        "title": "简述笛卡尔坐标系的建立方法",
                        "istitleimg": false,
                        "options": null,
                        "rightoptions": null,
                        "getoptions": null,
                        "needcheck": true,
                        "explain": "看着搞",
                        "result": 0
                    },
                    {
                        "id": "5fccd6bc36d02a179478825b",
                        "questiontype": 5,
                        "title": "简述数学的应用方面？",
                        "istitleimg": false,
                        "options": null,
                        "rightoptions": null,
                        "getoptions": null,
                        "needcheck": true,
                        "explain": "看着搞",
                        "result": 0
                    }
                ]
            },
            {
                "name": "图片题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478826a",
                        "questiontype": 6,
                        "title": "请上传y=cosx的图像",
                        "istitleimg": false,
                        "options": null,
                        "rightoptions": null,
                        "getoptions": null,
                        "needcheck": true,
                        "explain": "看着搞",
                        "result": 0
                    },
                    {
                        "id": "5fccd6bc36d02a179478826b",
                        "questiontype": 6,
                        "title": "请上传y=sinx的图像",
                        "istitleimg": false,
                        "options": null,
                        "rightoptions": null,
                        "getoptions": null,
                        "needcheck": true,
                        "explain": "看着搞",
                        "result": 0
                    }
                ]
            },
            {
                "name": "文件题",
                "questionitems": [
                    {
                        "id": "5fccd6bc36d02a179478827a",
                        "questiontype": 7,
                        "title": "请上传第一次作业的文件",
                        "istitleimg": false,
                        "options": null,
                        "rightoptions": null,
                        "getoptions": [
                            "http://127.0.0.1:8080/group1/default/20210102/20/40/5/transport_end.jpeg"
                        ],
                        "needcheck": true,
                        "explain": "看着搞",
                        "result": 4
                    },
                    {
                        "id": "5fccd6bc36d02a179478827b",
                        "questiontype": 7,
                        "title": "请上传第二次作业的文件",
                        "istitleimg": false,
                        "options": null,
                        "rightoptions": null,
                        "getoptions": null,
                        "needcheck": true,
                        "explain": "看着搞",
                        "result": 0
                    }
                ]
            }
        ]
    }
}`


--------

api:/file/upload

desc:上传文件

method:PUT

param:
    - scene string 场景
    - type string 1图片，2文件
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "fileid": "6a594fdee1fa1c60eca89f39c00d8a70",
        "url": "http://127.0.0.1:8080/group1/default/20210102/20/40/5/transport_end.jpeg"
    }
    }`

--------

api:/paper/query/recommend

desc:首页查看推荐试卷和待做作业

method:GET

param:

response:

`{
    "msg": "",
    "code": 0,
    "data": {
        "homework": [
            {
                "paperid": "5fe47805a541ddbf7e479ad1",
                "name": "家庭作业1",
                "score": 0,
                "suggesttime": 0,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e479ad2",
                "name": "家庭作业2",
                "score": 0,
                "suggesttime": 0,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e479ad3",
                "name": "家庭作业3",
                "score": 0,
                "suggesttime": 0,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e479ad4",
                "name": "家庭作业4",
                "score": 0,
                "suggesttime": 0,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e479ad5",
                "name": "家庭作业5",
                "score": 0,
                "suggesttime": 0,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e479ad6",
                "name": "家庭作业6",
                "score": 0,
                "suggesttime": 0,
                "origin": "高数Z"
            }
        ],
        "exam": [
            {
                "paperid": "5fe47805a541ddbf7e4794fe",
                "name": "有所要题目类型的试卷",
                "score": 110,
                "suggesttime": 3600,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e4794d3",
                "name": "期中试卷试卷",
                "score": 110,
                "suggesttime": 3600,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e4794d2",
                "name": "期中试卷试卷",
                "score": 110,
                "suggesttime": 3600,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e4794d1",
                "name": "期中试卷试卷",
                "score": 110,
                "suggesttime": 3600,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e4794c9",
                "name": "期中试卷试卷",
                "score": 110,
                "suggesttime": 3600,
                "origin": "高数Z"
            },
            {
                "paperid": "5fe47805a541ddbf7e4794c8",
                "name": "期中试卷试卷",
                "score": 110,
                "suggesttime": 3600,
                "origin": "高数Z"
            }
        ]
    }
}`


--------

api:/question/do

desc:做单个题目

method:PUT

param:

   `{
			"questionid":"5fccd6bc36d02a179478821c",
			"answers":["A","B"]	
}`
  
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "id": "5fccd6bc36d02a179478821c",
        "questiontype": 2,
        "title": "多选题测试",
        "istitleimg": false,
        "options": [
            {
                "prefix": "A",
                "content": "1",
                "iscontentimg": false
            },
            {
                "prefix": "B",
                "content": "2",
                "iscontentimg": false
            },
            {
                "prefix": "C",
                "content": "3",
                "iscontentimg": false
            },
            {
                "prefix": "D",
                "content": "4",
                "iscontentimg": false
            }
        ],
        "rightoptions": [
            {
                "prefix": "A",
                "content": "1",
                "iscontentimg": false
            },
            {
                "prefix": "C",
                "content": "3",
                "iscontentimg": false
            }
        ],
        "getoptions": [
            "A",
            "B"
        ],
        "needcheck": false,
        "explain": "看着搞",
        "result": 2
    }
}`