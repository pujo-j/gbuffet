DROP TABLE IF EXISTS `project_request`;
CREATE TABLE `project_request`
(

  `project_id`        varchar(255) COLLATE utf8_bin NOT NULL,
  `project_name`      varchar(255) COLLATE utf8_bin NOT NULL,
  `project_number`    varchar(255) COLLATE utf8_bin,
  `requester_email`   varchar(255) COLLATE utf8_bin NOT NULL,
  `requester_group`   varchar(255) COLLATE utf8_bin,
  `expected_lifetime` INTEGER                       NOT NULL DEFAULT 8,
  `request_status`    varchar(12) COLLATE utf8_bin,
  `creation`          datetime                      NOT NULL,
  `project_creation`  datetime,
  `project_deletion`  datetime,
  `requester_comment` TEXT COLLATE utf8_bin         NOT NULL,
  `admin_comment`     TEXT COLLATE utf8_bin         NOT NULL,
  `folder`            varchar(255) COLLATE utf8_bin NOT NULL,
  `tags`              JSON                          NOT NULL,
  PRIMARY KEY (`project_id`),
  FULLTEXT pr_ftx (requester_comment, admin_comment),
  INDEX pr_cdx (creation DESC),
  INDEX pr_reex (requester_email, creation DESC)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_bin;

DROP TABLE IF EXISTS `allocation`;
CREATE TABLE `allocation`
(
  `id`         BIGINT AUTO_INCREMENT         NOT NULL,
  `project_id` varchar(255) COLLATE utf8_bin NOT NULL,
  `type`       varchar(255) COLLATE utf8_bin NOT NULL,
  `region`     varchar(255) COLLATE utf8_bin,
  `zone`       varchar(255) COLLATE utf8_bin,
  `value`      BIGINT                        NOT NULL,
  PRIMARY KEY (`id`),
  INDEX a_px (`project_id`),
  CONSTRAINT `a_pidfk` FOREIGN KEY (`project_id`) REFERENCES `project_request` (`project_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_bin;


DROP TABLE IF EXISTS `user_group`;
CREATE TABLE `user_group`
(
  `user`        varchar(255) COLLATE utf8_bin NOT NULL,
  `group`       varchar(255) COLLATE utf8_bin NOT NULL,
  `last_update` datetime                      NOT NULL,
  PRIMARY KEY (`user`, `group`),
  INDEX ug_gux (`group`, `user`)
)
