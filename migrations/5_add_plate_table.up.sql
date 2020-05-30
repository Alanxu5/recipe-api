CREATE TABLE `plate` (
    `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `recipe_id` INT NOT NULL,
    INDEX `fk_plate_recipe_id` (`recipe_id` ASC) VISIBLE,
    FOREIGN KEY (`recipe_id`)
        REFERENCES `recipes`.`recipe` (`id`)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);
