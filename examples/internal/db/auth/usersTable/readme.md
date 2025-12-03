### 系统用户

#### 只管理用户

#### 表结构
* identity: 登录标识
* * IdentityCode: 
* security: 用户安全控制表, 业务无关
* * security_code: 用户编号, 业务无关
* * UserPlatform: 用户平台, 用于saas化管理, 按需扩展为saas系统/用户系统/管理员系统.
* * Password: 密码, 密码是一种特殊credential, 单独保存. 方便加强管理


#### 描述UserPlatform
- `app_user`: 不带"org"属性的用户
- `saas_user`: 带"org"属性的用户
- `admin_user`: 系统管理类账号, 负责管理系统.

相应系统应按照规则提供UserPlatform属性,与用户类型做匹配
