DROP PROCEDURE IF EXISTS gameAdd;

DELIMITER //

CREATE PROCEDURE gameAdd(
  $title VARCHAR(200), 
  $idConsole INT(11),
  $stars INT(11),
  $qty INT(11)
)
BEGIN
  INSERT INTO game (title, idConsole, stars, qty)
  VALUES ($title, $idConsole, $stars, $qty);
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS gameUpdate;

DELIMITER //

CREATE PROCEDURE gameUpdate(
  $id INT(11),
  $title VARCHAR(200), 
  $idConsole INT(11),
  $stars INT(11),
  $qty INT(11)
)
BEGIN
  UPDATE game
  SET title = $title,
    idConsole = $idConsole,
    stars = $stars,
    qty = $qty
  WHERE id = $id;
END//

DELIMITER ;

DROP PROCEDURE IF EXISTS gameDelete;

DELIMITER //

CREATE PROCEDURE gameDelete(
  $id INT
)
BEGIN
  DELETE FROM game WHERE id = $id;
END//

DELIMITER ;

DROP PROCEDURE IF EXISTS gameGet;

DELIMITER //

CREATE PROCEDURE gameGet(
  $id INT
)
BEGIN
  SELECT
    id,
    idConsole,
    title,
    stars,
    qty
  FROM game
  WHERE id = $id or $id IS NULL;
END//

DELIMITER ;
