## 我的服务器：112.126.68.22:8080/…

## 除了前三个都要加请求头：”Authorization“(存储token)

# 用户 /user/…

### 获取验证码 /getcode

POST

Body：表单

| email | 对应邮箱 |
| ----- | -------- |

Response：JSON

```json
{
    "success":"验证码已发送"
}
```

### 确认验证码 /checkcode

POST

Body：表单

| email | 对应邮箱 |
| ----- | -------- |
| code  | 验证码   |

Response：JSON

```json
{
    "success":"验证成功"
}
```

```json
{
    "fail":"验证码错误"
}
```

### 注册 /register

POST

Body：表单

| email    | 对应邮箱 |
| -------- | -------- |
| password | 密码     |

Response：JSON

```json
{
    "success":"注册成功"
}
```

```json
{
    "fail":"该邮箱已注册"
}
```

### 登录 /login

POST

Body：表单

| email    | 对应邮箱 |
| -------- | -------- |
| password | 密码     |

Response：JSON

```json
{
    "success":""(将会返回一个token)
}
```

```json
{
    "fail":"该邮箱未注册"
}
```

```json
{
    "fail":"密码错误"
}
```

### 登出 /logout

GET

Response：JSON

```json
{
    "success":"登出成功"
}
```

```json
{
    "fail":"token无效"
}
```

### 改密码 /changepassword （先验证码验证再这一步）

POST

Body：表单 

| newpassword | 新密码 |
| ----------- | ------ |

Response：JSON

```json
{
    "success":"密码修改成功"
}
```

```json
{
    "fail":"token无效"
}
```

### 改名 /changeusername

POST

Body：表单

| newusername | 新用户名 |
| ----------- | -------- |

Response：JSON

```json
{
    "success":”“（返回新的token，原token无效）
}
```

```json
{
    "fail":"token无效"
}
```

```json
{
    "fail":"用户名已存在"
}
```

### 上传头像 /setphoto

POST

Body：表单

| imageurl | 图片地址 |
| -------- | -------- |

Response：JSON

```json
{
    "success":"头像添加成功"
}
```

```json
{
    "fail":"token无效"
}
```

### 设置简介 /setdiscription

POST

Body：表单

| discription | 简介 |
| ----------- | ---- |

Response：JSON

```json
{
    "success":"简介修改成功"
}
```

### 通过用户id获取名字 /getname

### (因为之后很多数据以用户id存储，但显示出来要用户名)

POST

Body：表单

| id   | 用户id |
| ---- | ------ |

Response：JSON

```json
{
    "success":"用户名"
}
```

### 我组的卷 /mytest

GET

Response：JSON

```json
{
    "tests":{
        testid:卷子id
        userid:出题的用户id
        discription:卷子简介
        circle:所属圈子
        good:点赞数
        status:是否过审（之前说需要，后来又不需要，不用管）
        createtime:出卷时间
        testname:卷子名称
    }
    {
        ......
    }
}
```

### 我做过的卷 /mydotest

GET

Response：JSON

```json
{
    "tests":{
        testid:卷子id （可以通过下面的功能——用卷子id获取卷子信息）
        userid:做题的用户id
        testhistoryid:没用
    }
    {
        ......
    }
}
```

### 我出的题 /mypractice （根据需要提取数据嘻嘻）

GET

Response：JSON

```json
{
    "practices":{
        practiceid:练习id
        content:题目
        difficult:难度
        circle:所属圈子
        userid:出题人id
        answer:答案
        variety:题目类型
        imageurl:图片地址（如果有的话）
        status:是否过审（之前说需要，后来又不需要，不用管）
        explain:解析
        good:点赞数
    }
    {
        ......
    }
}
```

### 我做过的练习 /mydopractice

GET

Response：JSON

```json
{
    "practice":{
        practiceid:卷子id （可以通过下面的功能——用练习id获取练习信息）
        userid:做题的用户id
        answer:用户是否答对还是答错（true/false）
    }
    {
        ......
    }
}
```

### 用户信息 /myuser

GET

Response：JSON

```json
{
    "user":{
        id:用户id
        name:用户名
        discription:简介
        imageurl:头像地址
        password:大概没用
        email:大概没用
    }
}
```

# 练习 /practice/…

### 编练习题 /createpractice

POST

Body：表单

| variety    | 单选题/多选题/判断题         |
| ---------- | ---------------------------- |
| difficulty | 难度星数1–5                  |
| circle     | 所属圈子                     |
| imageurl   | 图片地址如果有的话           |
| content    | 题目内容                     |
| answer     | 答案（A/B/ABC/true/false..） |
| explain    | 解析                         |

Response：JSON

```json
{
    "id":练习题的id，后面创建选项用
	"success":"等待审核",（现在没用）
}
```

### 编练习题的选项 /createoption

POST

Body：表单

| practiceid | 对应练习的id       |
| ---------- | ------------------ |
| content    | 选项内容           |
| option     | 选项（A/B/true/…） |

Response：JSON

```json
{
	"success":"等待审核",（现在没用了）
}
```

### 做练习or前面“我做过的练习“通过practiceid获取practice /getpractice

POST

Body：表单

| circle     | 所属圈子（做练习时用）                 |
| ---------- | -------------------------------------- |
| practiceid | 练习id(通过practiceid获取practice时用) |

Response：JSON

```json
{
    practiceid:练习id
    content:题目内容
    difficulty:难度
    circle:所属圈子
    userid:编题用户的id
    answer:正确答案
    variety:题目类型
    imageurl:图片地址
    status:是否过审（目前不需要）
    explain:解析
    good:点赞数
}
```

### 获取练习题的选项/getoption

POST

Body：表单

| practiceid | 练习id |
| ---------- | ------ |

Response：JSON

```json
{
    optionid:选项id
    content:选项内容
    practiceid:练习id
    option:选项（A/B/true..）
}
```

### 评论练习/commentpractice

POST

Body：表单

| practiceid | 练习id   |
| ---------- | -------- |
| content    | 评论内容 |

Response：JSON

```json
{
    ”success":评论成功
}
```

### 获取练习题的评论/getcomment

POST

Body：表单

| practiceid | 练习id |
| ---------- | ------ |

Response：JSON

```json
{
    commentid:评论id
    content:评论内容
    practiceid:对应的练习id
    userid:评论人的id
}
```

### 对答案/checkanswer

POST

Body：表单

| practiceid | 练习id                                |
| ---------- | ------------------------------------- |
| circle     | 所属圈子（前面getpractice应该有返回） |
| answer     | 用户是否答对（true/false）            |
| time       | 用时（以秒返回）                      |

Response：JSON

```json
{
    "success":成功
}
```

### 获取排名/getrank

POST

Body：表单

| circle | 圈子 |
| ------ | ---- |

Response：JSON

```json
{
    "success":用户在这个圈子的练习排名
}
```

### 获取做题总数、正确数、总时长/getuserpractice

POST

Body：表单

| circle | 圈子 |
| ------ | ---- |

Response：JSON

```json
{
    id:没用
    userid:用户的id
    practicenum:总做练习数
    correctnum:正确数（正确率自己算）
    Alltime:总时长（平均时长自己算）
    circle:对应圈子
}
```

### 点赞练习/lovepractice

POST

Body：表单

| practiceid | 练习id |
| ---------- | ------ |

Response：JSON

```json
{
    "success":点赞成功
}
```

# 卷子 /test/…

### (题目好像要归入题库，将题目往创建练习功能再发一次？)（题库选题功能在后面）

### 组卷/createtest

POST

Body：表单

| circle      | 所属圈子 |
| ----------- | -------- |
| discription | 简介     |
| testname    | 卷子名称 |

Response：JSON

```json
{
    "id":卷子id
	"success":"等待审核",（现在没用）
}
```

### 创建卷子的题目/createtest

POST

Body：表单

| testid     | 卷子id   |
| ---------- | -------- |
| content    | 内容     |
| difficulty | 难度     |
| answer     | 答案     |
| variety    | 题型     |
| imageurl   | 图片地址 |
| explain    | 解析     |

Response：JSON

```json
{
    "id":题目id
	"success":"等待审核",（现在没用）
}
```

### 创建选项/createoption

POST

Body：表单

| practiceid | 题目id   |
| ---------- | -------- |
| content    | 选项内容 |
| option     | 选项     |

Response：JSON

```json
{
    "id":选项id
	"success":"等待审核",（现在没用）
}
```

### 做卷子/gettest

POST

Body：表单

| testid | 卷子id(后面会有获取卷子id的方法) |
| ------ | -------------------------------- |

Response：JSON

```json
{
    Testid 	
	Testname 
	Userid 
	Discription 
	Circl
	Good 
	Status string
	Createtime 创建时间
}
```

### 获取卷子对应题目/getquestion

POST

Body：表单

| testid | 卷子id |
| ------ | ------ |

Response：JSON

```json
{   多个。。
    Testid 
	Questionid:题目id 	
	Content 
	Difficulty 
	Answer 
	Variety
	Imageurl 
	Explain 
}
```

### 获取选项/gettestoption

POST

Body：表单

| practiceid | 题目id   |
| ---------- | -------- |
| content    | 选项内容 |
| option     | 选项     |

Response：JSON

```json
{
    类似
}
```

### 获取成绩/getscore

POST

Body：表单

| testid     | 卷子id     |
| ---------- | ---------- |
| correctnum | 正确数     |
| time       | 用时（秒） |

Response：JSON

```json
{
    "success":成功
}
```

### 卷子排行榜/showtop

POST

Body：表单

| testid | 卷子id |
| ------ | ------ |

Response：JSON

```json
{   (返回的就是前十，已经排好序)
    topid:没用
    userid:用户id
    correctnum:正确数
    time:用时
    testid:卷子id
}
```

### 评论卷子/commenttest

POST

Body：表单

| testid  | 卷子id   |
| ------- | -------- |
| content | 评论内容 |

Response：JSON

```json
{
    ”success":评论成功
}
```

### 获取卷子的评论/gettestcomment

POST

Body：表单

| testid | id   |
| ------ | ---- |

Response：JSON

```json
{
    commentid:评论id
    content:评论内容
    testid:对应的练习id
    userid:评论人的id
}
```

###  点赞卷子/lovetest

POST

Body：表单

| testid | 卷子id |
| ------ | ------ |

Response：JSON

```json
{
    "success":点赞成功
}
```

###  推荐/recommendtest

POST

Body：表单

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{
    多条test的信息
}
```

###  最热/hottest

POST

Body：表单

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{
    多条test的信息
}
```

###  最新/newtest

POST

Body：表单

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{
    多条test的信息
}
```

###  关注的圈子的卷子/followcircletest

GET

POST

Body：表单

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{
    多条test的信息
}
```

# 圈子  /circle/..

###  创圈/createcircle

POST

Body：表单

| name        | 圈子名称 |
| ----------- | -------- |
| discription | 简介     |
| imageurl    | 图片地址 |

Response：JSON

```json
{
    "success":等待审核
}
```

###  查看待审核的圈子/pendingcircle

GET   (需要root账号登录)

Response：JSON

```json
{
    "fail":权限不足
}
```

```json
{   (一个一个返回)
    "id":圈子id
    "name":圈子名称
    "imageurl":图片
    "discription":简介
    "userid":创圈人id
    "status":pending(待审核)
}
```

###  是否过审/approvecircle

POST   (需要root账号登录)

Body：表单

| circleid | 圈子id                 |
| -------- | ---------------------- |
| decide   | 是否过审（true/false） |

Response：JSON

```json
{
    "fail":权限不足
}
```

```json
{   
    "success":审核结束
}
```

###  获取圈子/getcircle

POST   

Body：表单

| circleid | 圈子id |
| -------- | ------ |

Response：JSON

```json
{   
    圈子信息
}
```

###  发现圈子/selectcircle

GET  

Response：JSON

```json
{   
    随机过审的圈子信息，最多十条
}
```

###  关注圈子/followcircle

POST   

Body：表单

| circleid | 圈子id |
| -------- | ------ |

Response：JSON

```json
{   
   "success":关注成功
}
```

# 搜索 /search/..

###  搜索圈子/searchcircle

POST   

Body：表单

| circlekey | 圈子关键词 |
| --------- | ---------- |

Response：JSON

```json
{   
    圈子信息
}
```

###  搜索卷子/searchtest

POST   

Body：表单

| testkey | 卷子名称关键词 |
| ------- | -------------- |

Response：JSON

```json
{   
    卷子信息
}
```

###  题库选题功能/searchpractice

POST   

Body：表单

| circle | 对应圈子 |
| ------ | -------- |

Response：JSON

```json
{   
    随机练习题目
}
```

###  搜索历史/searchhistory

GET   

Response：JSON

```json
{   
    "id":没用
    "searchkey":搜索过的词
    "userid":用户id
}
```

###  清空搜索历史/deletehistory

GET   

Response：JSON

```json
{   
    "success":删除成功
}
```

七牛云token：0bNiwJGpdwmvvuVAzLDjM6gnxj9MiwmSagVpIW81:85DTubmQkSKtCyWaL5KoaucrQKU=:eyJkZWFkbGluZSI6MTczODU3NjI0MCwic2NvcGUiOiJtdXhpLW1pbmlwcm9qZWN0In0=
