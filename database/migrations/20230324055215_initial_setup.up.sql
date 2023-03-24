-- create "todos" table
CREATE TABLE "public"."todos" ("id" character varying(36) NOT NULL, "name" character varying(50) NOT NULL, "description" text NULL, "status" text NULL, "createdAt" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP, "updatedAt" timestamp(3) NOT NULL, PRIMARY KEY ("id"));
