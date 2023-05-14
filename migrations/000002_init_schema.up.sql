CREATE TABLE "regions" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar(128) NOT NULL
);

CREATE TABLE "users" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "first_name" varchar(30) NOT NULL,
    "last_name" varchar(30) NOT NULL,
    "second_name" varchar(30) NOT NULL,
    "profile_picture" varchar(256),
    "inn" varchar(100),
    "email" varchar(256) NOT NULL,
    "phone_number" varchar(15) NOT NULL,
    "role_id" integer NOT NULL DEFAULT 1,
    "hashed_password" varchar(100) NOT NULL,
    "status" integer NOT NULL DEFAULT 1,
    "region_id" integer NOT NULL,
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "roles" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar(30) NOT NULL
);

CREATE TABLE "permissions" (
    "id" serial PRIMARY KEY NOT NULL,
    "role_id" integer NOT NULL,
    "path_id" integer NOT NULL,
    "can_insert" boolean NOT NULL DEFAULT false,
    "can_update" boolean NOT NULL DEFAULT false,
    "can_delete" boolean NOT NULL DEFAULT false,
    "can_read" boolean NOT NULL DEFAULT false
);

CREATE TABLE "paths" (
    "id" serial PRIMARY KEY NOT NULL,
    "path" varchar(255) NOT NULL
);

CREATE TABLE "session" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "device_name" varchar(30) NOT NULL,
    "ip_address" varchar(32),
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "user_sessions" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "user_id" varchar(21) NOT NULL,
    "session_id" varchar(21) NOT NULL,
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "classificator_group" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar(30) NOT NULL
);

CREATE TABLE "classificators" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "group" integer NOT NULL DEFAULT 1,
    "question" varchar(255) NOT NULL,
    "answer" varchar(255),
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "dynamic_category" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar(30) NOT NULL
);

CREATE TABLE "dynamic_forms" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "name" varchar(255) NOT NULL,
    "category" integer NOT NULL
);

CREATE TABLE "dynamic_form_key_values" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "key" varchar(255) NOT NULL,
    "type" varchar(255) NOT NULL
);

CREATE TABLE "dynamic_form_values" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "key_id" varchar(21) NOT NULL,
    "form_id" varchar(21) NOT NULL,
    "value" varchar NOT NULL
);

CREATE TABLE "declaration" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "organization_id" varchar(21) NOT NULL,
    "danger_rate" integer NOT NULL,
    "reasons_of_danger" varchar(255),
    "converage_of_the_danger_area" varchar(255),
    "proof" varchar(255) NOT NULL,
    "location_info" varchar(255),
    "residents_info" varchar(255),
    "life_insurance" varchar(255),
    "tech_document" varchar(255),
    "status" integer NOT NULL DEFAULT 1,
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "organization" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "name" varchar(255) NOT NULL,
    "full_name" varchar(255) NOT NULL,
    "phone_number" varchar(15) NOT NULL,
    "parent_organization" varchar(36),
    "location" varchar(200),
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "payment_types" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar NOT NULL
);

CREATE TABLE "payments" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "organization_id" varchar(21) NOT NULL,
    "amount" bigint NOT NULL,
    "requisites" varchar,
    "status" integer NOT NULL DEFAULT 1,
    "type" integer NOT NULL,
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "status" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar(20) NOT NULL
);

CREATE TABLE "license_type" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar NOT NULL
);

CREATE TABLE "license" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "document_number" varchar(255),
    "granted_date" timestamp,
    "lifetime" integer,
    "organization_name" varchar(255) NOT NULL,
    "stir_number" varchar(255),
    "reestr_number" varchar(255),
    "work_category" varchar(60),
    "doc_file" varchar(255),
    "license_type" integer NOT NULL,
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "fine" (
    "id" varchar(21) PRIMARY KEY NOT NULL DEFAULT (nanoid()),
    "issued_at" timestamp NOT NULL,
    "issued_by" varchar(21) NOT NULL,
    "issued_to" varchar(21) NOT NULL,
    "status" integer NOT NULL,
    "amount" bigint NOT NULL,
    "reason" varchar(255) NOT NULL,
    "invoice_number" varchar(60) NOT NULL,
    "created_at" timestamp DEFAULT (NOW()),
    "updated_at" timestamp DEFAULT (NOW()),
    "deleted_at" timestamp DEFAULT NULL
);

CREATE UNIQUE INDEX ON "permissions" ("role_id", "path_id");

ALTER TABLE "users"
ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "users"
ADD FOREIGN KEY ("status") REFERENCES "status" ("id");

ALTER TABLE "users"
ADD FOREIGN KEY ("region_id") REFERENCES "regions" ("id");

ALTER TABLE "permissions"
ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "permissions"
ADD FOREIGN KEY ("path_id") REFERENCES "paths" ("id");

ALTER TABLE "user_sessions"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_sessions"
ADD FOREIGN KEY ("session_id") REFERENCES "session" ("id");

ALTER TABLE "classificators"
ADD FOREIGN KEY ("group") REFERENCES "classificator_group" ("id");

ALTER TABLE "dynamic_forms"
ADD FOREIGN KEY ("category") REFERENCES "dynamic_category" ("id");

ALTER TABLE "dynamic_form_values"
ADD FOREIGN KEY ("key_id") REFERENCES "dynamic_form_key_values" ("id");

ALTER TABLE "dynamic_form_values"
ADD FOREIGN KEY ("form_id") REFERENCES "dynamic_forms" ("id");

ALTER TABLE "declaration"
ADD FOREIGN KEY ("status") REFERENCES "status" ("id");

ALTER TABLE "payments"
ADD FOREIGN KEY ("organization_id") REFERENCES "organization" ("id");

ALTER TABLE "payments"
ADD FOREIGN KEY ("status") REFERENCES "status" ("id");

ALTER TABLE "payments"
ADD FOREIGN KEY ("type") REFERENCES "payment_types" ("id");

ALTER TABLE "license"
ADD FOREIGN KEY ("license_type") REFERENCES "license_type" ("id");

ALTER TABLE "fine"
ADD FOREIGN KEY ("issued_by") REFERENCES "users" ("id");

ALTER TABLE "fine"
ADD FOREIGN KEY ("issued_to") REFERENCES "users" ("id");

ALTER TABLE "fine"
ADD FOREIGN KEY ("status") REFERENCES "status" ("id");