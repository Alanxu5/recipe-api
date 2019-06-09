CREATE TABLE `recipe` (
    `id` int   NOT NULL   AUTO_INCREMENT,
    `name` varchar(50)   NOT NULL,
    `prep_time` int   NOT NULL, 
    `cook_time` int   NOT NULL,
    `servings` int   NOT NULL,
    `method` varchar(20)   NOT NULL,
    `type` varchar(20)   NOT NULL,
    `description` varchar(500)   NOT NULL,
    `directions` JSON   NOT NULL,
    CONSTRAINT `pk_recipe` PRIMARY KEY (
        `id`
     )
);

-- bake, slow cooker, rice cooker, pan
CREATE TABLE `method` (
    `id` varchar(20)   UNIQUE NOT NULL,
    `name` varchar(20)   NOT NULL
);

-- protein, vegetables, carbs
CREATE TABLE `type` (
    `id` varchar(20)   NOT NULL,
    `name` varchar(20)   NOT NULL,
    CONSTRAINT `pk_type` PRIMARY KEY (
        `id`
     )
);

CREATE TABLE `ingredient` (
    `id` int   NOT NULL   AUTO_INCREMENT,
    `food_id` int   NOT NULL,
    `recipe_id` int   NOT NULL,
    `unit` varchar(25)   NOT NULL,
    `amount` float   NOT NULL,
    CONSTRAINT `pk_ingredient` PRIMARY KEY (
        `id`
     )
);

CREATE TABLE `food` (
    `id` int   NOT NULL   AUTO_INCREMENT,
    `name` varchar(50)   NOT NULL,
    CONSTRAINT `pk_food` PRIMARY KEY (
        `id`
     )
);

CREATE TABLE `image` (
    `id` int   NOT NULL   AUTO_INCREMENT,
    `recipe_id` int   NOT NULL,
    `image_link` text   NOT NULL,
    CONSTRAINT `pk_image` PRIMARY KEY (
        `id`
     )
);

ALTER TABLE `recipe` ADD CONSTRAINT `fk_recipe_method` FOREIGN KEY(`method`)
REFERENCES `method` (`id`);

ALTER TABLE `recipe` ADD CONSTRAINT `fk_recipe_type` FOREIGN KEY(`type`)
REFERENCES `type` (`id`);

ALTER TABLE `ingredient` ADD CONSTRAINT `fk_ingredient_food_id` FOREIGN KEY(`food_id`)
REFERENCES `food` (`id`);

ALTER TABLE `ingredient` ADD CONSTRAINT `fk_ingredient_recipe_id` FOREIGN KEY(`recipe_id`)
REFERENCES `recipe` (`id`);

ALTER TABLE `image` ADD CONSTRAINT `fk_image_recipe_id` FOREIGN KEY(`recipe_id`)
REFERENCES `recipe` (`id`);

