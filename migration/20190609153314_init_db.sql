-- +goose Up
CREATE TABLE "public"."countries" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "code" varchar(255),
  CONSTRAINT "countries_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."cities" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "country_id" uuid,
  CONSTRAINT "cities_pkey" PRIMARY KEY ("id")
);

ALTER TABLE "public"."cities"
  ADD CONSTRAINT "fk_cities_countries" FOREIGN KEY ("country_id") REFERENCES "public"."countries" ("id");

CREATE TABLE "public"."districts" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "city_id" uuid,
  CONSTRAINT "districts_pkey" PRIMARY KEY ("id")
);

ALTER TABLE "public"."districts"
  ADD CONSTRAINT "fk_districts_countries" FOREIGN KEY ("city_id") REFERENCES "public"."cities" ("id");

CREATE TABLE "public"."addresses" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(1000),
  "address" varchar(1000),
  "district_id" uuid,
  "location" point,

  CONSTRAINT "addresses_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE "public"."addresses"
  ADD CONSTRAINT "fk_addresses_districts" FOREIGN KEY ("district_id") REFERENCES "public"."districts" ("id");

CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "email" varchar(255) NOT NULL,
  "hashed_password" varchar(255) NOT NULL,

  CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."class_categories" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "description" text,

  CONSTRAINT "class_categories_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."class_types" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "description" text,
  "class_category_id" uuid NOT NULL,

  CONSTRAINT "class_types_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE "public"."class_types"
  ADD CONSTRAINT "fk_class_types_class_categories" FOREIGN KEY ("class_category_id") REFERENCES "public"."class_categories" ("id");

CREATE TABLE "public"."levels" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "level" integer DEFAULT 0,
  "description" text,

  CONSTRAINT "levels_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."classes" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "code" varchar(255),
  "description" text,
  "class_type_id" uuid DEFAULT NULL,

  CONSTRAINT "classes_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE "public"."classes"
  ADD CONSTRAINT "fk_classes_class_types" FOREIGN KEY ("class_type_id") REFERENCES "public"."class_types" ("id");

CREATE TABLE "public"."clubs" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "code" varchar(255),
  "address_id" uuid,

  CONSTRAINT "clubs_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE "public"."clubs"
  ADD CONSTRAINT "fk_clubs_addresses" FOREIGN KEY ("address_id") REFERENCES "public"."addresses" ("id");

CREATE TABLE "public"."rooms" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "code" varchar(255),
  "club_id" uuid,

  CONSTRAINT "rooms_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE "public"."rooms"
  ADD CONSTRAINT "fk_rooms_clubs" FOREIGN KEY ("club_id") REFERENCES "public"."clubs" ("id");

CREATE TABLE "public"."instructors" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),

  CONSTRAINT "instructors_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."schedules" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" varchar(255),
  "status" integer DEFAULT 0,
  "start_at" timestamptz NOT NULL,
  "end_at" timestamptz NOT NULL,
  "club_id" uuid NOT NULL,
  "class_category_id" uuid NOT NULL,

  CONSTRAINT "schedules_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


ALTER TABLE "public"."schedules"
  ADD CONSTRAINT "fk_schedules_clubs" FOREIGN KEY ("club_id") REFERENCES "public"."clubs" ("id");

ALTER TABLE "public"."schedules"
  ADD CONSTRAINT "fk_schedules_class_categories" FOREIGN KEY ("class_category_id") REFERENCES "public"."class_categories" ("id");


CREATE TABLE "public"."schedule_details" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "status" integer DEFAULT 0,
  "schedule_id" uuid NOT NULL,
  "class_id" uuid NOT NULL,
  "start_at" timestamptz NOT NULL,
  "end_at" timestamptz NOT NULL,
  "level_id" uuid DEFAULT NULL,
  "instructor_id" uuid NOT NULL,
  "room_id" uuid NOT NULL,

  CONSTRAINT "schedule_details_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

ALTER TABLE "public"."schedule_details"
  ADD CONSTRAINT "fk_schedule_details_levels" FOREIGN KEY ("level_id") REFERENCES "public"."levels" ("id");

ALTER TABLE "public"."schedule_details"
  ADD CONSTRAINT "fk_schedule_details_schedules" FOREIGN KEY ("schedule_id") REFERENCES "public"."schedules" ("id");

ALTER TABLE "public"."schedule_details"
  ADD CONSTRAINT "fk_schedule_details_classes" FOREIGN KEY ("class_id") REFERENCES "public"."classes" ("id");

ALTER TABLE "public"."schedule_details"
  ADD CONSTRAINT "fk_schedule_details_instructors" FOREIGN KEY ("instructor_id") REFERENCES "public"."instructors" ("id");

ALTER TABLE "public"."schedule_details"
  ADD CONSTRAINT "fk_schedule_details_rooms" FOREIGN KEY ("room_id") REFERENCES "public"."rooms" ("id");

-- +goose Down
DROP TABLE "public"."schedule_details" CASCADE;
DROP TABLE "public"."schedules" CASCADE;
DROP TABLE "public"."instructors" CASCADE;
DROP TABLE "public"."rooms" CASCADE;
DROP TABLE "public"."clubs" CASCADE;
DROP TABLE "public"."classes" CASCADE;
DROP TABLE "public"."levels" CASCADE;
DROP TABLE "public"."class_categories" CASCADE;
DROP TABLE "public"."class_types" CASCADE;
DROP TABLE "public"."users" CASCADE;
DROP TABLE "public"."addresses" CASCADE;
DROP TABLE "public"."districts" CASCADE;
DROP TABLE "public"."cities" CASCADE;
DROP TABLE "public"."countries" CASCADE;