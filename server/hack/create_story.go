package main

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	db := g.DB()
	ctx := context.Background()
	_, err := db.Exec(ctx, `
CREATE TABLE IF NOT EXISTS hg_story (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  ip_id bigint(20) NOT NULL COMMENT '关联的IP ID',
  chat_id varchar(36) NOT NULL DEFAULT '' COMMENT '关联的会话ID',
  title varchar(255) NOT NULL COMMENT '故事标题',
  content text COMMENT '故事内容',
  created_at datetime DEFAULT NULL COMMENT '创建时间',
  updated_at datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='故事表';
`)
	if err != nil {
		glog.Fatal(ctx, err)
	}
	glog.Info(ctx, "table hg_story created successfully")
}
