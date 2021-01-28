CREATE DATABASE IF NOT EXISTS `work-meow` DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

USE `work-meow`;

/*******************************/
/*   DatabaseName = work-meow  */
/*   TableName = cat_profile   */
/*******************************/
CREATE TABLE `cat_profile` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `host_id` varchar(32) NOT NULL COMMENT 'hostID',
  `nickname` varchar(64) NOT NULL COMMENT '昵称',
  `avatar_type` tinyint(3) NOT NULL COMMENT '头像类型',
  `meow_id` varchar(32) NOT NULL COMMENT '喵号',
  `level` tinyint(3) NOT NULL COMMENT '级别',
  `money` double NOT NULL COMMENT '余额',
  `is_worked` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否有打工经历',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'table''s create time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'table''s modify time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_host_meow` (`host_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='喵星人信息';


CREATE TABLE `meow_tags` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键自增ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名称',
  `host_id` varchar(32) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='论坛标签';

CREATE TABLE `meow_tiezis` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `title` longtext NOT NULL COMMENT '帖子标题',
  `content` longtext NOT NULL COMMENT '帖子内容',
  `host_id` varchar(32) NOT NULL DEFAULT '' COMMENT '发布者',
  `read_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '浏览次数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `weight` int(11) NOT NULL COMMENT '权值',
  `tag_id` int(11) NOT NULL COMMENT '标签ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='论坛帖子';

CREATE TABLE `meow_comments` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `content` longtext NOT NULL COMMENT '评论内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `host_id` varchar(32) NOT NULL COMMENT 'HostID',
  `tiezi_id` int(11) NOT NULL COMMENT '帖子ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子评论';