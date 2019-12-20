ALTER TABLE `ingredient` DROP COLUMN `preparation`;

DELETE FROM `method`;
DELETE FROM `type`;

CREATE TABLE `food` (
    `id` int   NOT NULL   AUTO_INCREMENT,
    `name` varchar(50)   NOT NULL,
    CONSTRAINT `pk_food` PRIMARY KEY (
        `id`
    )
);

ALTER TABLE `ingredient` ADD CONSTRAINT `fk_ingredient_food_id` FOREIGN KEY(`food_id`)
    REFERENCES `food` (`id`);
