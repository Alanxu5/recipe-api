ALTER TABLE `ingredient` ADD COLUMN `preparation` VARCHAR(25) NULL;

INSERT INTO `method` (name) VALUES ('Pan');
INSERT INTO `method` (name) VALUES ('Oven');
INSERT INTO `method` (name) VALUES ('Instant Pot');

INSERT INTO `type` (name) VALUES ('Protein');
INSERT INTO `type` (name) VALUES ('Carb');
INSERT INTO `type` (name) VALUES ('Vegetable');

ALTER TABLE `ingredient` DROP FOREIGN KEY `fk_ingredient_food_id`;
DROP TABLE if exists `food`;
