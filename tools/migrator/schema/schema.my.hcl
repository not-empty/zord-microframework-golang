table "dummy" {
  schema = schema.skeleton
  column "id" {
    type = tinytext
    null = false
  }
  column "name" {
    type = tinytext
    null = false
  }
  column "email" {
    type = tinytext
    null = false
  }
}
schema "skeleton" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
