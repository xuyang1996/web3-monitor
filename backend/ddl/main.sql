CREATE TABLE `user`
(
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `account` VARCHAR(64) NOT NULL,
  `signature` VARCHAR(256) NOT NULL,
  `email` VARCHAR(256) NOT NULL,
  `status` VARCHAR(64) NOT NULL,
  `create_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  INDEX(`account`),
  INDEX(`signature`),
  INDEX(`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `user_task`
(
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `account` VARCHAR(64) NOT NULL,
  `chain` INT NOT NULL,
  `address` VARCHAR(64) NOT NULL,
  `type` VARCHAR(64) NOT NULL,
  `status` VARCHAR(64) NOT NULL,
  `create_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  INDEX(`account`),
  INDEX(`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `transfer_txn`
(
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `chain` INT NOT NULL,
  `blockNumber` VARCHAR(128) NOT NULL,
  `timeStamp` VARCHAR(128) NOT NULL,
  `hash` VARCHAR(128) NOT NULL,
  `blockHash` VARCHAR(128) NOT NULL,
  `from` VARCHAR(64) NOT NULL,
  `to` VARCHAR(64) NOT NULL,
  `value` VARCHAR(128) NOT NULL,
  `create_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  INDEX(`blockNumber`),
  INDEX(`hash`),
  INDEX(`from`),
  INDEX(`to`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;