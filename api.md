### 注册

URL：POST        /register

Header：无

Body：表单

```formdata
"username":"XXX"
"password":"XXX
```

Response：json

```json
{
    "code":0,    //自定义错误码，0表示成功
    "message":"注册成功"
}
```

```json
{
    "code":1,
    "message":"用户名已存在"
}
```

```json
{
    "code":1,
    "message":"你已登录，请先登出"
}
```

### 登录

URL：POST       /login

Header：无

Body：表单

```
"username":"XXX"
"password":"XXX"
```

Response：json

```json
{
    "code":0,    
    "message":"登录成功"
    "token":"..."    //含一串随机字符串和username(用于确认是谁)
}
```

```json
{
    "code":1,
    "message":"密码错误"
}
```

```json
{
    "code":1,
    "message":"你已登录，请先登出"
}
```

```json
{
    "code":1,
    "message":"无当前用户名"
}
```

### 登出

URL：GET       /logout

Header：req.Header.Set("Authorization", "token")

Body：无

Response：json

```json
{
    "code":0,    
    "message":"登出成功"
}
```

```json
{
    "code":1,    
    "message":"你还没登录"
}
```

### 修改密码

URL：POST       /change_password

Header：req.Header.Set("Authorization", "token")

Body：表单

```
"password":"XXX"
"changepassword":"XXX"
```

Response：json

```json
{
    "code":1,    
    "message":"你还没登录"
}
```

```json
{
    "code":0,    
    "message":"修改成功"
}
```

```json
{
    "code":1,
    "message":"密码错误"
}
```

### 发帖子

URL：POST       /sendpost

Header：req.Header.Set("Authorization", "token")

Body：表单

```
name：XXX
circle：XXX(圈子)
content：帖子文本
title：标题
image：图片
```

Response：json

```json
{
    "error":"表单错误 or 保存图片失败"（图片发了不止一张）
}
```

```json
{
    "error":"不属于这个圈子"
}
```

```json
{
     "message":"帖子创建成功"
}
```

### 看帖子

URL：POST       /readpost

Header：req.Header.Set("Authorization", "token")

Body：表单

```
title：帖子的标题
```

Response：json

```json
{
     "error":"无该帖子"
}
```

```json
{
    "name":发帖人的名字
    "content":帖子内容
    "circle":所属圈子
    "love":点赞数
    "collect":收藏量
    "title":帖子标题
    "imageurl":图片存在本地的地址
}
```

### 点赞

URL：POST       /lovepost

Header：req.Header.Set("Authorization", "token")

Body：表单

```
title：帖子的标题
```

Response：json

```json
{
     "love":点赞数
}
```

### 收藏

URL：POST       /collectpost

Header：req.Header.Set("Authorization", "token")

Body：表单

```
title：帖子的标题
```

Response：json

```json
{
     "collect":收藏数
}
```

### 分享

URL：POST       /sharepost

Header：req.Header.Set("Authorization", "token")

Body：表单

```
title：帖子的标题
```

Response：json

```json
{
    "name":发帖人的名字
    "content":帖子内容
    "circle":所属圈子
    "love":点赞数
    "collect":收藏量
    "title":帖子标题
    "imageurl":图片存在本地的地址
}
```

### 关注帖主

URL：POST       /followpost

Header：req.Header.Set("Authorization", "token")

Body：表单

```
title：帖子的标题
```

Response：json

```json
{
     "message":"关注成功"
}
```

### 发评论

URL：POST       /commentpost

Header：req.Header.Set("Authorization", "token")

Body：表单

```
title：帖子的标题
content:评论内容
name:评论人
circle:所属圈子
```

Response：json

```json
{
     "error":"不属于这个圈子"
}
```

```json
{
    "message":"评论创建成功"
}
```

### 看帖子

URL：POST       /readcomment

Header：req.Header.Set("Authorization", "token")

Body：表单

```
title：帖子的标题
```

Response：json

```json
{
    "name":评论人
    "content":评论内容
    。。。
    。。。
}
```

