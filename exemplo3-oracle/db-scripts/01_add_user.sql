--declare
--userexist integer;
--begin
--  select count(*) into userexist from dba_users where username='SATO';
--  if (userexist = 0) then
--    execute immediate 'create user sato identified by abc123';
--    DEFAULT TABLESPACE users;
--    TEMPORARY TABLESPACE temp;
--    GRANT ALL PRIVILEGES TO sato;  end if;
--end;

CREATE USER sato IDENTIFIED BY abc123
DEFAULT TABLESPACE users
TEMPORARY TABLESPACE temp;
GRANT ALL PRIVILEGES TO sato;