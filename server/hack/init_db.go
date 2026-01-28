package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.Background()
	fmt.Println("Start initializing database...")

	// 1. 创建 hg_app_user 表
	createTableSQL := `
CREATE TABLE IF NOT EXISTS hg_app_user (
    id bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    username varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    nickname varchar(100) DEFAULT '' COMMENT '昵称',
    password_hash varchar(255) NOT NULL DEFAULT '' COMMENT '密码哈希',
    salt varchar(10) NOT NULL COMMENT '密码盐',
    mobile varchar(20) DEFAULT '' COMMENT '手机号',
    email varchar(100) DEFAULT '' COMMENT '邮箱',
    avatar varchar(500) DEFAULT '' COMMENT '头像URL',
    sex tinyint(1) DEFAULT '0' COMMENT '性别',
    birthday date DEFAULT NULL COMMENT '生日',
    last_login_at datetime DEFAULT NULL COMMENT '最后登录时间',
    last_login_ip varchar(50) DEFAULT '' COMMENT '最后登录IP',
    status tinyint(1) DEFAULT '1' COMMENT '状态',
    remark text COMMENT '备注',
    created_at datetime DEFAULT NULL COMMENT '创建时间',
    updated_at datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY uniq_username (username),
    KEY idx_mobile (mobile),
    KEY idx_email (email),
    KEY idx_status (status),
    KEY idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用_普通用户';
`
	if _, err := g.DB().Exec(ctx, createTableSQL); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
	} else {
		fmt.Println("✅ Table hg_app_user check/create successfully!")
	}

	// 2. 插入菜单和权限数据（逐条执行）
	fmt.Println("Start seeding menu data...")

	sqls := []string{
		// 插入一级菜单：C端管理
		`INSERT INTO hg_admin_menu (pid, title, name, path, icon, type, sort, status, created_at, updated_at)
		SELECT 0, 'C端管理', 'app', '/app', 'UserOutlined', 1, 90, 1, NOW(), NOW()
		WHERE NOT EXISTS (SELECT 1 FROM hg_admin_menu WHERE name = 'app')`,

		// 插入二级菜单：C端用户
		`INSERT INTO hg_admin_menu (pid, title, name, path, component, permissions, type, sort, status, created_at, updated_at)
		SELECT (SELECT id FROM hg_admin_menu WHERE name = 'app'), 'C端用户', 'app-user', 'user', '/app/user/list', '/user/list', 2, 1, 1, NOW(), NOW()
		WHERE NOT EXISTS (SELECT 1 FROM hg_admin_menu WHERE name = 'app-user')`,

		// 插入按钮权限：查看
		`INSERT INTO hg_admin_menu (pid, title, name, permissions, type, status, created_at, updated_at)
		SELECT (SELECT id FROM hg_admin_menu WHERE name = 'app-user'), '查看', 'app-user-view', '/user/view', 3, 1, NOW(), NOW()
		WHERE NOT EXISTS (SELECT 1 FROM hg_admin_menu WHERE name = 'app-user-view')`,

		// 插入按钮权限：编辑
		`INSERT INTO hg_admin_menu (pid, title, name, permissions, type, status, created_at, updated_at)
		SELECT (SELECT id FROM hg_admin_menu WHERE name = 'app-user'), '编辑', 'app-user-edit', '/user/edit', 3, 1, NOW(), NOW()
		WHERE NOT EXISTS (SELECT 1 FROM hg_admin_menu WHERE name = 'app-user-edit')`,

		// 插入按钮权限：删除
		`INSERT INTO hg_admin_menu (pid, title, name, permissions, type, status, created_at, updated_at)
		SELECT (SELECT id FROM hg_admin_menu WHERE name = 'app-user'), '删除', 'app-user-delete', '/user/delete', 3, 1, NOW(), NOW()
		WHERE NOT EXISTS (SELECT 1 FROM hg_admin_menu WHERE name = 'app-user-delete')`,
	}

	for _, sql := range sqls {
		if _, err := g.DB().Exec(ctx, sql); err != nil {
			fmt.Printf("Error executing SQL: %v\nSQL: %s\n", err, sql)
		}
	}

	fmt.Println("✅ Menu data seeded successfully!")
}
