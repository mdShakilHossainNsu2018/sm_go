-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema sm_db
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `sm_db` ;

-- -----------------------------------------------------
-- Schema sm_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `sm_db` ;
USE `sm_db` ;

-- -----------------------------------------------------
-- Table `sm_db`.`album`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sm_db`.`album` ;

CREATE TABLE IF NOT EXISTS `sm_db`.`album` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(128) NOT NULL,
  `artist` VARCHAR(255) NOT NULL,
  `price` DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 5
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sm_db`.`timestamps`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sm_db`.`timestamps` ;

CREATE TABLE IF NOT EXISTS `sm_db`.`timestamps` (
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sm_db`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sm_db`.`users` ;

CREATE TABLE IF NOT EXISTS `sm_db`.`users` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(16) NOT NULL,
  `email` VARCHAR(255) NULL,
  `password` TEXT NOT NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sm_db`.`posts`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sm_db`.`posts` ;

CREATE TABLE IF NOT EXISTS `sm_db`.`posts` (
  `post_id` INT NOT NULL AUTO_INCREMENT,
  `users_id` INT NOT NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL,
  `content` TEXT NULL,
  PRIMARY KEY (`post_id`),
  UNIQUE INDEX `id_UNIQUE` (`post_id` ASC) VISIBLE,
  INDEX `fk_posts_users_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_posts_users`
    FOREIGN KEY (`users_id`)
    REFERENCES `sm_db`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sm_db`.`files`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sm_db`.`files` ;

CREATE TABLE IF NOT EXISTS `sm_db`.`files` (
  `file_id` INT NOT NULL AUTO_INCREMENT,
  `url` TEXT NULL,
  `alt` VARCHAR(256) NULL,
  `posts_post_id` INT NOT NULL,
  `type` VARCHAR(45) NULL,
  PRIMARY KEY (`file_id`),
  INDEX `fk_Images_posts1_idx` (`posts_post_id` ASC) VISIBLE,
  CONSTRAINT `fk_Images_posts1`
    FOREIGN KEY (`posts_post_id`)
    REFERENCES `sm_db`.`posts` (`post_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sm_db`.`profile`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sm_db`.`profile` ;

CREATE TABLE IF NOT EXISTS `sm_db`.`profile` (
  `profile_id` INT NOT NULL AUTO_INCREMENT,
  `image` TEXT NULL,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `phone` VARCHAR(45) NULL,
  `users_user_id` INT NOT NULL,
  `about_me` VARCHAR(1200) NULL,
  PRIMARY KEY (`profile_id`),
  INDEX `fk_profile_users1_idx` (`users_user_id` ASC) VISIBLE,
  CONSTRAINT `fk_profile_users1`
    FOREIGN KEY (`users_user_id`)
    REFERENCES `sm_db`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
