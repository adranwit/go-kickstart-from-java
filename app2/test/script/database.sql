DROP  TABLE IF EXISTS table1;

CREATE TABLE `table1` (
  `id`                 INT(11)   NOT NULL AUTO_INCREMENT,
  `name`               VARCHAR(255)       DEFAULT NULL,
  `category`           VARCHAR(255)       DEFAULT NULL,
  `status`             VARCHAR(10)       DEFAULT NULL,
  `groupname`         VARCHAR(10)       DEFAULT NULL,
  PRIMARY KEY (`id`)
);
