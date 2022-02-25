DROP PROCEDURE IF EXISTS companyAdd;

DELIMITER //

CREATE PROCEDURE companyAdd(
  $title VARCHAR(200)
)
BEGIN
  INSERT INTO company (title, created)
  VALUES ($title, NOW());
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS companyUpdate;

DELIMITER //

CREATE PROCEDURE companyUpdate(
  $id INT(11),
  $title VARCHAR(200)
)
BEGIN
  UPDATE company
  SET title = $title,
      updated = NOW()
  WHERE id = $id;
END//

DELIMITER ;

DROP PROCEDURE IF EXISTS companyDelete;

DELIMITER //

CREATE PROCEDURE companyDelete(
  $id INT
)
BEGIN
  DELETE FROM company WHERE id = $id;
END//

DELIMITER ;

DROP PROCEDURE IF EXISTS companyGet;

DELIMITER //

CREATE PROCEDURE companyGet(
  $id INT
)
BEGIN
  SELECT
    id,
    title
  FROM company
  WHERE id = $id OR $id = 0;
END//

DELIMITER ;

