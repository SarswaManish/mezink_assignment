Drop database If EXISTS student;

CREATE DATABASE IF NOT EXISTS student;

use student;

DROP TABLE IF EXISTS records;
CREATE TABLE records" (
                                    "id" int8,
                                    "name" varchar,
                                    "created_at" date,
                                    "marks" json
    );


INSERT INTO "records" ("id", "name", "created_at", "marks") VALUES
(1, 'Bob', '2023-05-20', '{"marks": [100,200,400,500]}');

INSERT INTO "records" ("id", "name", "created_at", "marks") VALUES
(2, 'Alice', '2023-05-21', '{"marks": [100,200,500]}');



INSERT INTO "records" ("id", "name", "created_at", "marks") VALUES
(3, 'Alena', '2023-05-21', '{"marks": [100,200,400]}');


INSERT INTO "records" ("id", "name", "created_at", "marks") VALUES
(4, 'Bella', '2023-05-21', '{"marks": [200,400]}');



INSERT INTO "records" ("id", "name", "created_at", "marks") VALUES
(5, 'Bella2', '2023-05-21', '{"marks": [200,600,50]}');



