CREATE SEQUENCE "SATO"."USUARIO_SQ" 
              INCREMENT BY 1 
              START WITH 1 
              MAXVALUE 9999999999999999999999999999 
              MINVALUE 1 
              CACHE 20;
CREATE TABLE "SATO"."USUARIO"
  (
    id NUMBER DEFAULT "SATO"."USUARIO_SQ".nextval NOT NULL,
    nome VARCHAR2(200),
    cpf VARCHAR2(20),
    datainclusao     DATE DEFAULT sysdate NOT NULL
  );