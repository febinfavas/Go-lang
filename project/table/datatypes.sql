CREATE DATABASE test;

CREATE TABLE datatypes(ID smallserial,NAME varchar(50),STORAGE_SIZE varchar(10),DISCRIPTION text);

INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('smallint','2-bytes','small-range integer');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('int/integer','4-bytes','typical-range integer');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('bigint','8-bytes','large-range integer');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('decimal','variable','user-specified precision');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('numeric','variable','user-specified precision');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('real','4-bytes','variable precision');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('double precision','8-bytes','variable precision');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('smallserial','2-bytes','small autoincrementing integer');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('serial','4-bytes','typical autoincrementing integer');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('bigserial','8-bytes','large autoincrementing integer');

SELECT * FROM datatypes;

INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('money','8-bytes','currency amount');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('varchar','variable','variable-length with limit');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('varchar','variable','variable-length with limit');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('test','variable','varaible unlimited length');

ALTER TABLE datatypes ALTER COLUMN storage_size TYPE VARCHAR(20);

INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('bytea','1 or 4 bytes','varaible length binary string');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('timestamp','8-bytes','both time and date ,no time zone');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('timestamp','8-bytes','both time and date with time zone');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('date','4-bytes','only date');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('time','8-bytes','only time');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('time','12-bytes','only time with time zone');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('interval','10-bytes','time interval');
INSERT INTO datatypes(NAME,STORAGE_SIZE,DISCRIPTION) VALUES ('iboolean','1-bytes','state of true or false');

UPDATE datatypes SET name='boolean' WHERE id = 22;



