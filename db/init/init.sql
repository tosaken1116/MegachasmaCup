CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "image_url" varchar,
  "name" varchar,
  "email" varchar NOT NULL,
  "created_at" timestamp,
  "school_id" uuid
);

CREATE TABLE "notes" (
  "id" uuid PRIMARY KEY,
  "class_id" uuid,
  "school_id" uuid,
  "title" varchar,
  "description" varchar,
  "user_id" uuid,
  "is_public" bool DEFAULT false,
  "created_at" timestamp,
  "deleted_at" timestamp,
  "update_at" timestamp
);

CREATE TABLE "schools" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "classes" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "school_id" uuid,
  "created_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "class_user" (
  "class_id" uuid,
  "user_id" uuid,
  PRIMARY KEY ("class_id", "user_id")
);

CREATE TABLE "like" (
  "note_id" uuid,
  "user_id" uuid,
  PRIMARY KEY ("note_id", "user_id")
);

CREATE TABLE "tags" (
  "id" uuid PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "Tagging" (
  "tag_id" uuid,
  "note_id" uuid,
  PRIMARY KEY ("tag_id", "note_id")
);

ALTER TABLE "users" ADD FOREIGN KEY ("school_id") REFERENCES "schools" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "classes" ADD FOREIGN KEY ("school_id") REFERENCES "schools" ("id");

ALTER TABLE "class_user" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "class_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "like" ADD FOREIGN KEY ("note_id") REFERENCES "notes" ("id");

ALTER TABLE "like" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "Tagging" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

ALTER TABLE "Tagging" ADD FOREIGN KEY ("note_id") REFERENCES "notes" ("id");
