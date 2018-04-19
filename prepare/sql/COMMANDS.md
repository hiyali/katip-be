# artisan
```shell
ALTER TABLE talbe_name MODIFY COLUMN column_name varchar(30); // modify column type

ALTER TABLE tbl_Country DROP COLUMN IsDeleted1, DROP COLUMN IsDeleted2; // delete column

DELETE FROM tbl_Country WHERE id > 10; // delete row

SELECT COLUMN_NAME, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH
			 FROM information_schema.columns
			 WHERE TABLE_NAME = 'YOUR_TABLE_NAME'; // show data type for column
```
