CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" varchar,
                         "password" varchar,
                         "created_at" timestamptz DEFAULT (now()),
                         "updated_at" timestamptz
);

CREATE TABLE "profile_info" (
                                "id" bigserial PRIMARY KEY,
                                "first_name" varchar,
                                "last_name" varchar,
                                "email" varchar,
                                "phone_number" varchar,
                                "profile_img" varchar,
                                "user_id" bigint,
                                "created_at" timestamptz DEFAULT (now()),
                                "updated_at" timestamptz
);

CREATE TABLE "extra_info" (
                              "id" bigserial PRIMARY KEY,
                              "university" varchar,
                              "about_me" varchar,
                              "user_id" bigint,
                              "created_at" timestamptz DEFAULT (now()),
                              "updated_at" timestamptz
);

ALTER TABLE "profile_info" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "extra_info" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "profile_info" ("email");

CREATE INDEX ON "profile_info" ("phone_number");

CREATE INDEX ON "extra_info" ("user_id");

COMMENT ON COLUMN "profile_info"."email" IS 'Must be valid email';

COMMENT ON COLUMN "profile_info"."phone_number" IS 'Must be valid phone number';

COMMENT ON COLUMN "profile_info"."profile_img" IS 'Url of the image';
