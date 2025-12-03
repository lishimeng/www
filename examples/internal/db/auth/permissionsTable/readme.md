### "角色-权限-资源"模型

#### 前缀使用: auth_


#### 相关视图
```sql
create or replace view public.user_perm_view(id, role_code, is_global, user_code, permission_code) as
SELECT DISTINCT row_number() OVER (ORDER BY t1.id, t2.id) AS id,
    t1.role_code,
    t1.is_global,
    t1.user_code,
    t2.permission_code
   FROM auth_user_roles t1
     JOIN auth_role_permissions t2 ON t1.role_code = t2.role_code;
```

```sql
create or replace view public.user_router_view
            (id, name, menu_group, route_path, component, type, icon, parent_id, route_name, keep_alive,
             sort, visible, params, perm, status, ctime, mtime, user_code, role_code)
as
SELECT t5.id,
    t5.name,
    t5.menu_group,
    t5.route_path,
    t5.component,
    t5.type,
    t5.icon,
    t5.parent_id,
    t5.route_name,
    t5.keep_alive,
    t5.sort,
    t5.visible,
    t5.params,
    t5.perm,
    t5.status,
    t5.ctime,
    t5.mtime,
    t3.user_code,
    t3.role_code
   FROM auth_user_roles t3,
    auth_role_permissions t4,
    auth_menu t5
  WHERE t3.role_code = t4.role_code AND t4.permission_code = t5.perm::text
;
```
