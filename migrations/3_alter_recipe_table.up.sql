ALTER TABLE `ingredient` RENAME COLUMN food_id TO food;
ALTER TABLE `ingredient` MODIFY COLUMN food VARCHAR(50)