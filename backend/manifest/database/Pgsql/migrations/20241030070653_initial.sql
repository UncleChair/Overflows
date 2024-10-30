-- Create "casbin_rule" table
CREATE TABLE "public"."casbin_rule" ("id" serial NOT NULL, "p_type" character varying(10) NULL, "v0" character varying(25) NULL, "v1" character varying(25) NULL, "v2" character varying(25) NULL, "v3" character varying(25) NULL, "v4" character varying(25) NULL, "v5" character varying(25) NULL, "v6" character varying(25) NULL, "v7" character varying(25) NULL, PRIMARY KEY ("id"), CONSTRAINT "idx_casbin_rule" UNIQUE ("p_type", "v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7"));
-- Create "users" table
CREATE TABLE "public"."users" ("uid" character varying(10) NOT NULL, "avatar_url" character varying(45) NOT NULL, "username" character varying(25) NOT NULL, "email" character varying(45) NOT NULL, "password" character varying(60) NOT NULL, "created_at" timestamp NULL, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, "last_login" timestamp NULL, "login_attempts" integer NULL DEFAULT 0, "lock" boolean NULL DEFAULT false, "lock_at" timestamp NULL, PRIMARY KEY ("uid"), CONSTRAINT "unique_username_email" UNIQUE ("username", "email"), CONSTRAINT "users_email_key" UNIQUE ("email"), CONSTRAINT "users_username_key" UNIQUE ("username"));
-- Create index "idx_email" to table: "users"
CREATE INDEX "idx_email" ON "public"."users" ("email");
-- Create index "idx_username" to table: "users"
CREATE INDEX "idx_username" ON "public"."users" ("username");
