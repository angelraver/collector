SHOW PROCEDURE STATUS WHERE db = 'collector';

DROP PROCEDURE IF EXISTS consoleAdd;

DELIMITER //

CREATE PROCEDURE consoleAdd(
  $idCompany INT(11),
  $title VARCHAR(200)
)
BEGIN
  INSERT INTO console (idCompany, title, created)
  VALUES ($idCompany, $title, NOW());
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS consoleUpdate;

DELIMITER //

CREATE PROCEDURE consoleUpdate(
  $id INT(11),
  $idCompany INT(11),
  $title VARCHAR(200)
)
BEGIN
  UPDATE game
  SET idCompany = $idCompany,
      title = $title,
      updated = NOW()
  WHERE id = $id;
END//

DELIMITER ;

DROP PROCEDURE IF EXISTS consoleDelete;

DELIMITER //

CREATE PROCEDURE consoleDelete(
  $id INT
)
BEGIN
  DELETE FROM console WHERE id = $id;
END//

DELIMITER ;

DROP PROCEDURE IF EXISTS consoleGet;

DELIMITER //

CREATE PROCEDURE consoleGet(
  $id INT
)
BEGIN
  SELECT
    id,
    idCompany,
    title
  FROM console
  WHERE id = $id OR $id = '';
END//

DELIMITER ;

