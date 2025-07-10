table "dummy" {
  schema = schema.zord
  column "id" {
    null = false
    type = char(26)
  }
  column "name" {
    null = false
    type = char(255)
  }
  column "email" {
    null = false
    type = char(255)
  }
  primary_key {
    columns = [column.id]
  }
}
schema "zord" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}


table "societies" {
  schema = schema.zord
  column "id_society" {
    null = false
    type = char(26)
  }
  column "name" {
    null = false
    type = char(255)
  }
  column "date_deleted_at" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id_society]
  }
}