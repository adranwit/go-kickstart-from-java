DROP TABLE IF EXISTS site;

CREATE TABLE `site` (
  `id`       INT(11) NOT NULL AUTO_INCREMENT,
  `name`     VARCHAR(255)     DEFAULT NULL,
  `url`      TEXT             DEFAULT NULL,
  `category` VARCHAR(255)     DEFAULT NULL,
  `status`   VARCHAR(10)      DEFAULT NULL,
  PRIMARY KEY (`id`)
);
