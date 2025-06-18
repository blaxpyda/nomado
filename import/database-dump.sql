DROP TABLE IF EXISTS "public"."property";

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS property_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS "public"."property" (
    "id" INT4 NOT NULL DEFAULT nextval('property_id_seq'::regclass),
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "transaction_type" VARCHAR(50) NOT NULL,
    "price" NUMERIC(10, 2) NOT NULL,
    "status" VARCHAR(50) NOT NULL,
    "agent_id" INT4 NOT NULL,
    "location" VARCHAR(255) NOT NULL,
    "property_type" VARCHAR(50) NOT NULL,
    "area" NUMERIC(10, 2) NOT NULL,
    "bedrooms" INT4 NOT NULL,
    "bathrooms" INT4 NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY ("id"),
    FOREIGN KEY ("agent_id") REFERENCES "public"."agent" ("id") ON DELETE CASCADE
);

DROP TABLE IF EXISTS "public"."agent";

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS agent_id_seq;

-- Table Definition
CREATE TABLE "public"."agent"(
    "id" INT4 NOT NULL DEFAULT nextval('agent_id_seq'::regclass),
    "first_name" VARCHAR(100) NOT NULL,
    "last_name" VARCHAR(100) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(20) NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("email")


);

DROP TABLE IF EXISTS "public"."user";

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS user_id_seq;

-- Table Definition
CREATE TABLE "public"."user"(
    "id" INT4 NOT NULL DEFAULT nextval('user_id_seq'::regclass),
    "username" VARCHAR(50) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY ("id")
);

-- Populate the property table with sample data
INSERT INTO "public"."agent" (first_name, last_name, email, password, phone)
VALUES ('Agent', 'Smith', 'agent@example.com', 'hashed_password', '0700123456');


INSERT INTO "public"."property" (
    name, description, image_url, transaction_type, price, status,
    agent_id, location, property_type, area, bedrooms, bathrooms
) VALUES
(
    'Modern Villa',
    'Spacious modern villa with pool and garden.',
    'assets/images/villa1.jpg',
    'sale',
    480000.00,
    'available',
    1,
    'Naguru, Kampala',
    'villa',
    300.0,
    4,
    3
),
(
    'City Apartment',
    '2-bedroom apartment in the city center.',
    'assets/images/villa2.jpg',
    'rent',
    1500.00,
    'available',
    1,
    'Kisementi, Kampala',
    'apartment',
    90.0,
    2,
    2
),
(
    'Country Cottage',
    'Quiet country home with lots of land.',
    'assets/images/villa3.jpg',
    'sale',
    120000.00,
    'available',
    1,
    'Entebbe, Wakiso',
    'cottage',
    150.0,
    3,
    2
),
(
    'Lake View Mansion',
    'Luxury mansion with views of Lake Victoria.',
    'assets/images/villa4.jpg',
    'sale',
    900000.00,
    'available',
    1,
    'Ggaba, Kampala',
    'mansion',
    500.0,
    6,
    5
),
(
    'Studio Apartment',
    'Affordable studio ideal for students.',
    'assets/images/villa5.jpg',
    'rent',
    600.00,
    'available',
    1,
    'Wandegeya, Kampala',
    'studio',
    45.0,
    1,
    1
),
(
    'Family Home',
    '3-bedroom family house with garage.',
    'assets/images/bung1.jpg',
    'sale',
    210000.00,
    'available',
    1,
    'Ntinda, Kampala',
    'house',
    160.0,
    3,
    2
),
(
    'Penthouse',
    'Top-floor penthouse with skyline views.',
    'assets/images/bung2.jpg',
    'rent',
    3500.00,
    'available',
    1,
    'Kololo, Kampala',
    'penthouse',
    200.0,
    3,
    3
),
(
    'Beach House',
    'House right on the beach, perfect for holidays.',
    'assets/images/bung3.jpg',
    'sale',
    320000.00,
    'available',
    1,
    'Lutembe Bay, Entebbe',
    'beach house',
    180.0,
    4,
    3
),
(
    'Suburban Duplex',
    'Modern duplex in a quiet suburb.',
    'assets/images/bung4.jpg',
    'rent',
    1800.00,
    'available',
    1,
    'Naalya, Kampala',
    'duplex',
    130.0,
    3,
    2
),
(
    'Eco Home',
    'Sustainable home built with eco-friendly materials.',
    'assets/images/bung5.jpg',
    'sale',
    275000.00,
    'available',
    1,
    'Mukono, Uganda',
    'eco-home',
    140.0,
    3,
    2
);

