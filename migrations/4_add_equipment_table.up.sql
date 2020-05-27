CREATE TABLE equipment (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(45) NULL,
  `equipment` VARCHAR(45) NOT NULL,
  `affiliate_link` VARCHAR(250) NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE recipe_equipment (
  `id` INT NOT NULL AUTO_INCREMENT,
  `recipe_id` INT NULL,
  `equipment_id` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_equipment_idx` (`equipment_id` ASC) VISIBLE,
  FOREIGN KEY (`equipment_id`)
      REFERENCES `recipes`.`equipment` (`id`)
      ON DELETE CASCADE
      ON UPDATE NO ACTION,
  FOREIGN KEY (`recipe_id`)
      REFERENCES `recipes`.`recipe` (`id`)
      ON DELETE CASCADE
      ON UPDATE NO ACTION
);



