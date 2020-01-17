ALTER TABLE `ingredient` RENAME COLUMN food_id TO food;
ALTER TABLE `ingredient` MODIFY COLUMN food VARCHAR(50);
ALTER TABLE `ingredient` DROP INDEX `fk_ingredient_food_id`