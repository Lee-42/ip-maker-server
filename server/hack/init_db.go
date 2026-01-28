package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	fmt.Println("Start creating table hg_app_user...")

	// 尝试直接使用 g.DB()，如果环境配置正确（配置了 manifest/config/config.yaml），它应该能连上。
	// 但这是一个独立的脚本，可能没加载配置。为了保险，我们手动配置一下默认的连接，
	// 并允许失败（如果连不上，打印错误提示用户检查）。

	// 注意：这里硬编码了密码，如果不对，用户需要在运行前修改此文件
	g.Cfg().GetAdapter().(*g.AdapterFile).SetFileName("manifest/config/config.yaml")

	sql := `
CREATE TABLE IF NOT EXISTS hg_app_user (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL DEFAULT '',
    nickname VARCHAR(100) DEFAULT '',
    password_hash VARCHAR(255) NOT NULL DEFAULT '',
    salt VARCHAR(10) NOT NULL,
    mobile VARCHAR(20) DEFAULT '',
    email VARCHAR(100) DEFAULT '',
    avatar VARCHAR(500) DEFAULT '',
    sex SMALLINT DEFAULT 0,
    birthday DATE,
    last_login_at TIMESTAMP,
    last_login_ip VARCHAR(50) DEFAULT '',
    status SMALLINT DEFAULT 1,
    remark TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_app_user_username ON hg_app_user(username);
CREATE INDEX IF NOT EXISTS idx_app_user_mobile ON hg_app_user(mobile);
CREATE INDEX IF NOT EXISTS idx_app_user_email ON hg_app_user(email);
CREATE INDEX IF NOT EXISTS idx_app_user_status ON hg_app_user(status);
CREATE INDEX IF NOT EXISTS idx_app_user_created_at ON hg_app_user(created_at);
`
	if _, err := g.DB().Exec(context.Background(), sql); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		// 尝试 fallback 连接串
		fmt.Println("Trying fallback connection string...")
		g.DB().GetConfig().AddDefaultGroupNode(g.DBConfigNode{
			Type: "pgsql",
			Link: "pgsql:postgres:hg123456@tcp(127.0.0.1:5432)/hotgo",
		})
		if _, err := g.DB().Exec(context.Background(), sql); err != nil {
			panic(fmt.Sprintf("Still failed: %v", err))
		}
	}

	fmt.Println("✅ Table hg_app_user created successfully!")
}
