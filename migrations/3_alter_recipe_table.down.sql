ALTER TABLE `ingredient` RENAME COLUMN food TO food_id;
ALTER TABLE `ingredient` MODIFY COLUMN food_id int;
