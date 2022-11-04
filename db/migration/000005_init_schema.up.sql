ALTER TABLE "accounts" ALTER COLUMN "owner" SET NOT NULL;

ALTER TABLE "accounts" ALTER COLUMN "balance" SET NOT NULL;

ALTER TABLE "accounts" ALTER COLUMN "currency" SET NOT NULL;

ALTER TABLE "accounts" ALTER COLUMN "created_at" SET NOT NULL; 
ALTER TABLE "accounts" ALTER COLUMN "created_at" SET DEFAULT (now());

ALTER TABLE "entries" ALTER COLUMN "account_id" SET NOT NULL;

ALTER TABLE "entries" ALTER COLUMN "amount" SET NOT NULL;

ALTER TABLE "entries" ALTER COLUMN "created_at" SET NOT NULL;
ALTER TABLE "entries" ALTER COLUMN "created_at" SET DEFAULT (now());