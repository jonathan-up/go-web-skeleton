# TABLE USERS
DROP TABLE IF EXISTS `pre_users`;
CREATE TABLE pre_users
(
    `id`            INT AUTO_INCREMENT COMMENT '管理员ID',
    `username`      VARCHAR(255) UNIQUE                 NOT NULL COMMENT '用户名',
    `password`      VARCHAR(255)                        NOT NULL COMMENT '密码',
    `last_login_ip` CHAR(15)                            NULL comment '上次登陆IP',
    `last_login_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL comment '上次登陆时间',
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL comment '创建时间',
    PRIMARY KEY (`id`)
);
