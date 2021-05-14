DROP DATABASE IF EXISTS codingOrganizer;
CREATE DATABASE codingOrganizer;
USE codingOrganizer;

CREATE TABLE entries (

  entry_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) DEFAULT "",
  url VARCHAR(2083) DEFAULT "",
  codeblock TEXT,
  notes TEXT

);

CREATE TABLE `tags` (

  `tag_id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  PRIMARY KEY (`tag_id`)

);

-- insert into tags (name)
-- values 
--   ("PHP"),
--   ("Ruby"),
--   ("Golang"),
--   ("SQL"),
--   ("Command Line"),
--   ("Linux"),
--   ("JS");

create table entry_has_tag (
  entry_id INT UNSIGNED NOT NULL,
  `tag_id` int unsigned NOT NULL,
  PRIMARY KEY (`entry_id`, `tag_id`)
);

alter table entry_has_tag
add constraint foreign key (entry_id) references entries(entry_id)
on delete cascade
on update cascade;

-- insert into entry_has_tag (entry_id, tag_id)
-- values
-- (3, 2),
-- (3, 7),
-- (5, 1),
-- (5, 2);

-- SELECT tags.*
-- FROM tags
-- INNER JOIN entry_has_tag ON (entry_has_tag.tag_id=tags.tag_id)
-- WHERE
--   entry_has_tag.entry_id=3;