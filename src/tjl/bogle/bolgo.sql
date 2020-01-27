create table article(
    id bigint(20) not null auto_increment comment '文章id',
    category_id bigint(20) unsigned not null  comment '分类id',
    content longtext character set utf8mb4 collate utf8mb4_general_ci not null comment '文章内容',
    title varchar(24) character set utf8mb4 collate utf8mb4_general_ci not null comment '文章标题',
    view_count int(255) unsigned not null comment '阅读次数',
    comment_count int(255) unsigned not null comment '评论次数',
    username varchar(128) character set utf8mb4 collate utf8mb4_general_ci not null comment '作者',
    status int(10) unsigned not null comment '状态',
    summary varchar(256) character set  utf8mb4 collate utf8mb4_general_ci not null comment '文章摘要',
    create_time timestamp(0) not null  default current_timestamp comment '发布时间',
    update_time timestamp(0) null default current_timestamp on update  current_timestamp(0) comment '更新时间',
    primary key ('id') using btree ,
    index 'idx_view_count'('view_count') using btree comment '阅读次数索引',
    index 'idx_comment_count'('comment_count') using btree comment '评论数索引',
    index 'idx_category_id'('category_id') using btree comment '分类索引'
)engine = InnoDB character set = utf8mb4 collate = utf8mb4_general_ci row_format  = dynamic ;

drop table if exists cateogry;
create table cateogry(
    id bigint(20) not null  auto_increment,
    category_name varchar(255) character set utf8mb4 collate utf8mb4_general_ci not null comment '分类名称',
    category_no int(10) unsigned not null comment '分类排序',
    create_time timestamp(0) null default current_timestamp,
    update_time timestamp(0) null default current_timestamp on update current_timestamp(),
    primary key ('id') using btree
)engine = innodb character set  = utf8mb4 collate = utf8mb4_general_ci row_format = dynamic ;

drop table if exists comment;
create table comment (
    id bigint(20) unsigned not null  comment '评论id',
    content text character set utf8mb4 collate utf8mb4_general_ci not null comment '评论内容',
    username varchar(64) character set utf8mb4 collate utf8mb4_general_ci not null comment '评论作者',
    create_time timestamp(0) not null on update current_timestamp(0) comment '评论发布时间',
    status int(255) unsigned not null comment '评论状态:0 删除,1正常',
    article_id bigint(20) unsigned null default null,
    primary key ('id') using btree
)engine = innodb character set  = utf8mb4 collate = utf8mb4_general_ci row_format = dynamic ;

drop table if exists leave;
create  table leave (
    id bigint(20) not null auto_increment,
    username varchar(255) character set utf8mb4 collate utf8mb4_general_ci not null ,
    content text character set utf8mb4 collate utf8mb4_general_ci not null ,
    create_time timestamp(0) not null default current_timestamp,
    update_time timestamp(0) not null default current_timestamp on update current_timestamp(0),
    primary key ('id') using btree
)engine = innodb character set  = utf8mb4 collate = utf8mb4_general_ci row_format = dynamic ;