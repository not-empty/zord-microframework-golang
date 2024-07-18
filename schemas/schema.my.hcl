table "dummy" {
  schema = schema.skeleton
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

//{{newTable}}

variable "database" {
  type = string
}

schema "skeleton" {
  name = var.database
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
