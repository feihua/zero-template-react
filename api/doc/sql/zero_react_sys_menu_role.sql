create table sys_menu_role
(
    id           bigint unsigned auto_increment comment '主键'
        primary key,
    gmt_create   datetime         default CURRENT_TIMESTAMP not null comment '创建时间',
    gmt_modified datetime         default CURRENT_TIMESTAMP not null comment '修改时间',
    status_id    tinyint unsigned default 1                 not null comment '状态(1:正常，0:禁用)',
    sort         int unsigned     default 1                 not null comment '排序',
    menu_id      bigint unsigned                            not null comment '菜单ID',
    role_id      bigint unsigned                            not null comment '角色ID'
)
    comment '菜单角色关联表' charset = utf8mb4;

