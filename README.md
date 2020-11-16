# knowledgemap_backend
知识图谱项目后端
/usr/local/mongodb/bin
/data/db  数据存放位置
redis地址 /usr/local/redis/
consul agent -server  -bootstrap-expect 1 -data-dir /tmp/consul -node=172.17.9.156
公47.95.145.171

所有api要加统一前缀 /api

----

api:/user/register

desc:注册

method:post

param:

    - rtype 身份枚举 0学生 1老师 2教秘
    - name string 真实姓名 
    - major string 专业
    - idcard string 身份证号 唯一id
    - origin string 籍贯
    - class string 所属班级
    - admissiontime string 入学时间
    - account string 账号
    - password string 密码
    - college string 学院
    - course string 课程

response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "user": {
            "userid": "5e97d27e36d02a8f32563289", 系统内部自建用户id
            "username": "teacher", 用户名 
            "usertype": 1, 用户类型枚举
            "major": "computer", 专业
            "idcard": "2123456789051", 身份证号
            "courses": [
                ""  
            ], 课程
            "password": "123456", 密码
            "number": "5" 学号，教师号，教秘编号
            },
        "token": "M5jez947zdM227ceu3Q6IdM0j2IaT8EfO3j225z6T3N2v8j9", 登录密令
        "expires": 1587304229 过期时间
        }
    }`

----

api:/user/login

desc:登陆

method:put

param:

    - ltype 身份枚举 0学生 1老师 2教秘
    - account string 账号
    - password string 密码

response:

    `{
    "msg": "",
    "code": 0,
    "data": {
        "user": {
            "userid": "5e97d27e36d02a8f32563289", 系统内部自建用户id
            "username": "teacher", 用户名 
            "usertype": 1, 用户类型枚举
            "major": "computer", 专业
            "idcard": "2123456789051", 身份证号
            "courses": [
                ""  
            ], 课程
            "password": "123456", 密码
            "number": "5" 学号，教师号，教秘编号
            },
        "token": "M5jez947zdM227ceu3Q6IdM0j2IaT8EfO3j225z6T3N2v8j9", 登录密令
        "expires": 1587304229 过期时间
        }
    }`

----

api:/user/changepassword

desc:修改用户信息,应该修改为修改用户信息，目前暂时为修改密码，待修改完成后再对接

method:put

param:

   - account string 账号
   - password string 密码
  
response:

    `{
    "msg":"",      失败的时候代表失败原因
    "code":0       0代表成功，1代表失败
    }`

----

api:/user/querymyinfo

desc:查询自己的信息 这个接口需要登陆,待完成

method:get

response:

  `{
    "msg": "",
    "code": 0,
    "data": {
        "user": {
            "userid": "5e97d27e36d02a8f32563289", 系统内部自建用户id
            "username": "teacher", 用户名 
            "usertype": 1, 用户类型枚举
            "major": "computer", 专业
            "idcard": "2123456789051", 身份证号
            "courses": [
                ""  
            ], 课程
            "password": "123456", 密码
            "number": "5" 学号，教师号，教秘编号
            },
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
   - teachername string 教师名
   - teacherid string 教师id编号

response:

    `{
        "msg": "",
        "code": 0,
        "data": {
            "classid": "5e97e77936d02a9dbb5ce966",
            "name": "高等数学A",
            "major": "math",
            "college": "computer",
            "teachername": "李永乐"
        }
    }`

--------

api:/class/join

desc:学生加入班级

method:put

param:

    - userid string 学生id
    - classid string 班级id
    - username string 学生名
    - status string 状态
  
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
                    "number": "10"
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
                    "number": "10"
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
            "teachername": "李永乐"
        }
    }`

--------

api:/class/query/alluserinclass/:classid

desc:查询班级所有学生信息

method:get

param:

    - classid string 班级id 
  
response:

   `{
        "msg": "",
        "code": 0,
        "data": {
            "students": [
                {
                    "userid": "5e919d0036d02a48ccc08225",
                    "username": "黄"
                }
            ]
        }
    }`

--------

api:/class/invitation/create

desc:创建班级邀请码

method:put

param:

    - classid  班级id
    - invitationcode 邀请码
    - userid 用户id
  
response:

    `{
    "msg":"",      失败的时候代表失败原因
    "code":0       0代表成功，1代表失败
    }`

--------

api:/class/invitation/drop

desc:删除班级邀请码

method:put

param:

    - classid  班级id
    - invitationcode 邀请码
    - userid 用户id
  
response:

    `{
    "msg":"",      失败的时候代表失败原因
    "code":0       0代表成功，1代表失败
    }`

--------

api:/class/invitation/query/:invitationcode

desc:通过邀请码获得班级信息

method:post

param:

    - classid  班级id
    - invitationcode 邀请码
    - userid 用户id
  
response:

    `{
        "msg": "",
        "code": 0,
        "data": {
            "classid": "5e97e77936d02a9dbb5ce966",
            "name": "高等数学A",
            "major": "math",
            "college": "computer",
            "teachername": "李永乐"
        }
    }`

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

api:/question/query/:kind/:course/:subject/:knowledge

desc:根据知识点查询某种类型的题目

method:GET

param:

    - kind int64 题目类型
    - subject string 学科
    - course string 课程 
    - knowledge string 知识点id
   
response:

   `{
    "msg": "",
    "code": 0,
    "data": {
        "questions": [
            {
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
                "course": "xx",
                "knowledge": "5e9b06ce36d02a7b7637eccf"
            }
        ]
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