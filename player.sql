/*
 Navicat Premium Data Transfer

 Source Server         : postgre local
 Source Server Type    : PostgreSQL
 Source Server Version : 150001 (150001)
 Source Host           : localhost:5432
 Source Catalog        : player
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150001 (150001)
 File Encoding         : 65001

 Date: 13/08/2023 11:52:14
*/


-- ----------------------------
-- Table structure for players
-- ----------------------------
DROP TABLE IF EXISTS "public"."players";
CREATE TABLE "public"."players" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "bank" varchar(255) COLLATE "pg_catalog"."default",
  "no_rekening" varchar(255) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "email" varchar(255) COLLATE "pg_catalog"."default",
  "nama_rekening" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."players" OWNER TO "postgres";

-- ----------------------------
-- Records of players
-- ----------------------------
BEGIN;
INSERT INTO "public"."players" ("id", "username", "password", "bank", "no_rekening", "created_at", "updated_at", "email", "nama_rekening") VALUES ('5e8257d2-3935-11ee-a9ca-0a5c2959084c', 'admin2', '$2a$08$rKvEZAUHQkUBfQlaJNr4qet3SHeu8VHA2l1.AUhMMscvoA2nqnebK', '', '', '2023-08-13 00:26:27.671443', '2023-08-13 00:26:27.671443', 'admin2@gmail.com', '');
INSERT INTO "public"."players" ("id", "username", "password", "bank", "no_rekening", "created_at", "updated_at", "email", "nama_rekening") VALUES ('65695acc-38e6-11ee-b565-0a5c2959084c', 'aji', '$2a$08$ZjV7mcyet9AoahyEC59xcOoyUd/sBkdWrJq6QkKYBKa.hloXRBale', 'Mandiri', '1203432234', '2023-08-12 15:01:09.010092', '2023-08-13 11:39:34.210535', 'admin1@gmail.com', 'muh. tarmizi');
COMMIT;

-- ----------------------------
-- Table structure for players_deposit
-- ----------------------------
DROP TABLE IF EXISTS "public"."players_deposit";
CREATE TABLE "public"."players_deposit" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "player_id" uuid NOT NULL,
  "type" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "nominal" numeric(10,2) NOT NULL,
  "created_at" timestamp(6)
)
;
ALTER TABLE "public"."players_deposit" OWNER TO "postgres";

-- ----------------------------
-- Records of players_deposit
-- ----------------------------
BEGIN;
INSERT INTO "public"."players_deposit" ("id", "player_id", "type", "nominal", "created_at") VALUES ('009e689b-a9d8-4c5d-ad3e-b912873563c3', '65695acc-38e6-11ee-b565-0a5c2959084c', 'debit', 10000000.00, '2023-08-12 16:40:01');
INSERT INTO "public"."players_deposit" ("id", "player_id", "type", "nominal", "created_at") VALUES ('50e6556c-537a-48fe-a941-ef352e9adf39', '65695acc-38e6-11ee-b565-0a5c2959084c', 'credit', 50000.00, '2023-08-12 16:40:59');
INSERT INTO "public"."players_deposit" ("id", "player_id", "type", "nominal", "created_at") VALUES ('016fd9f2-393f-11ee-a813-0a5c2959084c', '5e8257d2-3935-11ee-a9ca-0a5c2959084c', 'debit', 1000000.00, '2023-08-13 01:35:26.489442');
COMMIT;

-- ----------------------------
-- Function structure for uuid_generate_v1
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v1"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v1mc
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1mc"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1mc"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1mc'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v1mc"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v3
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v3"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v3"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v3'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v3"("namespace" uuid, "name" text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v4
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v4"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v4"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v4'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v4"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v5
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v5"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v5"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v5'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v5"("namespace" uuid, "name" text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_nil
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_nil"();
CREATE OR REPLACE FUNCTION "public"."uuid_nil"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_nil'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_nil"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_dns
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_dns"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_dns"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_dns'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_dns"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_oid
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_oid"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_oid"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_oid'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_oid"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_url
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_url"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_url"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_url'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_url"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_x500
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_x500"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_x500"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_x500'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_x500"() OWNER TO "postgres";

-- ----------------------------
-- Primary Key structure for table players
-- ----------------------------
ALTER TABLE "public"."players" ADD CONSTRAINT "players_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table players_deposit
-- ----------------------------
ALTER TABLE "public"."players_deposit" ADD CONSTRAINT "players_deposit_pkey" PRIMARY KEY ("id");
