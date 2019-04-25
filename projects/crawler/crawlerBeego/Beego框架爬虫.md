# Beego框架爬虫
> 以爬取豆瓣网一个静态页面为例
>
> https://movie.douban.com/subject/25827935/

**项目构建**

```
bee new crawlerBeego
```

## 数据存储-创建表
```
CREATE TABLE `movie_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(11) unsigned NOT NULL COMMENT '电影id',
  `movie_name` varchar(100) COMMENT '电影名称',
  `movie_pic` varchar(200) COMMENT '电影图片',
  `movie_director` varchar(50) COMMENT '电影导演',
  `movie_writer` varchar(50) COMMENT '电影编剧',
  `movie_country` varchar(50) COMMENT '电影产地',
  `movie_language` varchar(50) COMMENT '电影语言',
  `movie_main_character` varchar(300) COMMENT '电影主演',
  `movie_type` varchar(50) COMMENT '电影类型',
  `movie_on_time` timestamp DEFAULT '2018-12-12 00:00:00' COMMENT '电影上映时间',
  `movie_span` varchar(20) COMMENT '电影时长',
  `movie_grade` varchar(5) COMMENT '电影评分',
  `remark` varchar(500) DEFAULT '' COMMENT '备注',
  `_create_time` timestamp NOT NULL DEFAULT '2018-12-12 00:00:00' COMMENT '创建时间',
  `_modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `_status` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `idx_movie_id` (`movie_id`),
  KEY `idx_create_time` (`_create_time`),
  KEY `idx_modify_time` (`_modify_time`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COMMENT='电影信息表';
```